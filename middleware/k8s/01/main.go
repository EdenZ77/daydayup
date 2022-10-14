package main

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"net/http"
	"os"
)

type K8sProxyPram struct {
	PopName    string
	KubeConfig string
}

type K8sProxy struct {
	popName    string
	KubeConfig string
	restConfig *rest.Config
	*kubernetes.Clientset
	RestClient       *rest.RESTClient
	DynamicClient    dynamic.Interface
	metricsClientset *metrics.Clientset
}

var kubeconfig192 = "apiVersion: v1\nclusters:\n- cluster:           \n    server: https://172.30.3.192:59100\n    insecure-skip-tls-verify: true\n  name: cluster.local\ncontexts:\n- context:\n    cluster: cluster.local\n    user: kubernetes-admin\n  name: kubernetes-admin@cluster.local\ncurrent-context: kubernetes-admin@cluster.local\nkind: Config\npreferences: {}\nusers:\n- name: kubernetes-admin\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lJVEZQNlJJV0Q5OTR3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWdGdzB5TWpBNE1qWXdOekUzTXpGYUdBOHlNRFV5TURneE9EQTNNVGN6TmxvdwpOREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhHVEFYQmdOVkJBTVRFR3QxWW1WeWJtVjBaWE10CllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRREpEaWJBS2NNSFVnelEKZVdZSnBJSSt0N2RNTGZWR2hiMDdyZFpCdlZNNEUzbkczbS9TaUFlNUxiNGgvdHBWcU00ZDR1dHYxeXB5R1VZVQo0ZHZkV24vUjR6SWJaTFE0aUgvcW50a21jRlBOdXFBZUZNdEJicVQ2cVlxQVh5MnJZUXlDMklVYXlTYnZlUmIwClhOZ05CVkZ5Z3dINUJyaFVxR1lpU0FHMzlQcGwzZVBpREJYbG05WWpmTFRQTVBQWUcwRUlUV0FYVVhpeUhRTVEKb0hENUpCSjFNalFYTERuQVYwbU82UHFvNXZacGIzUXRPT1NHQ0ZmMlZPcGVBZzI4K04zZnltNXFBUjZrK3lNdApFL3RBbUx2REdtUkdmbW9HZmlvMm8xcFYzODNLT2xmUWliTm4wN3Nlc1Y3UlJ1M0RhYkFSVEZhS3QxZ00rZnpNCkVTb2NWUFUxQWdNQkFBR2pTREJHTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQWZCZ05WSFNNRUdEQVdnQlRpSnFZUUE2M055SWM5MWE0Z01uamJWeDV6ZVRBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFucUZqb0xrZVdWZjhWTVhiV1QyYXp5UHhlTkVHVVlnWkE1aDZIL1dnU1lyTkJMT3lHQUJHCkdTb3dSYWw4UW8zTUdKeDRXVlI2ME5tcEI1Z0E5OTUvcDg3M09tNUY1V0FCblVRaUxwMlhGUW1lUVlKblV4NksKRThPV2hJb09IZG05ODBVWmdJYnJGeGRBRmV6V095aUE2Tnh4THhsaCtidGVvY0hpbzV3aDlYQWw4Rmt1ZDF1Sgpic1d0cFJCK2lWWUNVdW8vakhTaGNya1JQZk9LTXZGSTVJM0hNckFmbkYzZFNobmpxRDB4Vm5IRHIzbUpCSmRrCnYrTktKVm5jWGI0T3R5U0k0c2ZxVVJ4UHRlM1lTdU5NbFlIbFFQNWVLeDUwQ3VkclhuU0FrVjBxdnUwYUd2d1IKdUpSOEV2UUJZbXhRN01UNEtvZVlWVy90VWgvcm42UndVQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K                               \n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBeVE0bXdDbkRCMUlNMEhsbUNhU0NQcmUzVEMzMVJvVzlPNjNXUWIxVE9CTjV4dDV2CjBvZ0h1UzIrSWY3YVZhak9IZUxyYjljcWNobEdGT0hiM1ZwLzBlTXlHMlMwT0loLzZwN1pKbkJUemJxZ0hoVEwKUVc2aytxbUtnRjh0cTJFTWd0aUZHc2ttNzNrVzlGellEUVZSY29NQitRYTRWS2htSWtnQnQvVDZaZDNqNGd3Vgo1WnZXSTN5MHp6RHoyQnRCQ0UxZ0YxRjRzaDBERUtCdytTUVNkVEkwRnl3NXdGZEpqdWo2cU9iMmFXOTBMVGprCmhnaFg5bFRxWGdJTnZQamQzOHB1YWdFZXBQc2pMUlA3UUppN3d4cGtSbjVxQm40cU5xTmFWZC9OeWpwWDBJbXoKWjlPN0hyRmUwVWJ0dzJtd0VVeFdpcmRZRFBuOHpCRXFIRlQxTlFJREFRQUJBb0lCQUdDV1I4OXhRSnczc3FoRQphZHlnQjJJUjZDRFV3MHhKMjRyY0RGNHdrZFlTcFNJQW9qL0YwZEFJUlpzWFQ5UnU3L1l6bVY5MVFwTGx5V0VtCnovVWJFT1RIL0w1a05xQWlFekduZGpLZEsrVmRqcVprM3ZCa015V29aVDBlZkZZa25Wb09vb01udDJpOEIyY0YKWTFWK0JJNjZtU3dGS25DaEpjKzZQL2tiMjE4cnl1VmovMEJDZHhyeUU4ZDRicGdTdElnNXgranJ1UU5SOXlrSgo3MTh2SFVScVg4a1NQM3lLTEczKzNUQjA0WnNxdDJIUCtZdDF0ZjlvK1Zid2pvSHJBZWhtb0M4Zld3aVFWNSt2Cm05SUl5M0l2MDJRdHpOTWFnR1VOMlpTTHJNUldNNFlBK3NVRGZESk5BdDcwOTFrSlB1Q3R4NXV3SFlkV0ZSTmUKRmJLd0l0RUNnWUVBK1dEMnBQNXpTRGg5bWdRZUdDYXpuWDNPclZkbmd1UGpKK2ZUbkVHVDAxM05WbEZjKzFiWAphRXpuNWxJdnlNQ1YyOHBTSGRwRU8yMU1XNFBXSFI0Yzc3UU5nWk9KbXlvMy9VMWdmckdGQTRxVHM1a3hXTzNsCnZuYXl5R1NBSjNGUzR2ZkpxSXFHaDFKT0RXNXBYSWNKUmNQUGRBYXJBZUdveDlZZVVxUm1wMGNDZ1lFQXptUzcKUHRpQ0toYkR5bGtDT0swaU1iYk9rOE93a2tHSWs5KzFjVWZhc1BhY3ZkN200bkZiWVN6V0oxNytFSjFtRkljagprY3B4VXlyUmhVSWF1UVFnd09iUWQ3UDJUeFZTMXVkN0Q0WUh2Uk5DVDVkbkExMElSczQ3YzlhWWhqN2NST2RxCkk1bHY3Ym5uZk44a2pLcUlPZE5hYUJUeVo1dFZxTTBTZXJQVGRhTUNnWUVBNm1qRmp2UUxJeERPcDQ1RlI0aGgKZjZHNU8yRVVVSW1yaFdBNW5nQmFWdTB1VFh2dmZlWDBWdnNyWkdsT3QxS255dERUL1hHa2Y0UE9xWnMwRVd5ego0SEdMM0lmMWFoLzJQeWlUa3FPRkYzNFVObGJDZHdndjA2ZTVoL3BJS0VzeWtWdy9keWkzS2M0b3hpRko4b3FRCkliN2Nhd0Mxai9BdytaOEFJOGliSVpNQ2dZRUF5YzFHSStHd0M1VXNwTm00eVUvSGdsSmEwN0hnSUhFQktJencKcksxMEQ5bGhVbWp5MlcrNnlGMzltb3RQNFZEMDhaZGMyUHpYSjFsVGVYYzBCN2tZaVdSbGF0VkVQUGo5Z1hEZQpLMFNDcG9XQkxhODhvdFpBOUhKTFFTME8veHZSWlhIYm5xazAvbnpwOFhlQkZpVGJnNmE2MjgrM1lFUktVZjBKClYzNGlnUDhDZ1lFQXFGTzNBZ1ZCWWxvalV2NkJWQWJYNFF0cmQrcDF1bERNa256KzV4RnlDNm5VSnF6NmJRUXoKbGp2UDhvL2QwRmZDV3Y3bHI5UVd1MGpTR05ZV3VnRDlQWWUrMm9EWGZ6QktLYXZGNTZWdWhzMGVUSmtXZnl3UQpZUHdnZmhxRUFyOGVhZUpFcEcxcEdEQ21DT0puUXNDcnBxanQ1MmZ4dE1IVjBkY243V2VLOUNjPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
var kubeconfig229 = "apiVersion: v1\nkind: Config\ncurrent-context: kubernetes-admin@cluster.local\ncontexts:\n- name: kubernetes-admin@cluster.local\n  context:\n    cluster: kubernetes-admin@cluster.local\n    user: devops-cluster-admin\nclusters:\n- name: kubernetes-admin@cluster.local\n  cluster:\n    #certificate-authority-data: \"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2VENDQWRHZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRJeU1Ea3dNVEEzTlRjek5Gb1lEekl3TlRJd09ESTBNRGMxTnpNMFdqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCjRaejBxQVluYWRxTFkwR05ocHlveGdhKytOTTlLaHZIWG9hL0ZuZmRoRjBVMkJIdGM1VFFOWHRCaFVXZUNlVEIKcEZDdHhPRFdnSWpsWXlkeW5FUlBwQWpUcG11NVk1bk9VVWZGVmhQdFB4QWlmTkUzOS9VT1BIUU5lZkZHR0RNMQpTTEtLVWYwYXJVenFESDh0Yy9SVVROS1lTTnBEWXBNZDREcjUrWnRtcmhXZGpRTHczT2V2M0RiTDNaK3dndGw0CnN5NVlIYjF2N3pUQ2t3c0czV0VhdkNmejZ6TmNGcG9kbmFweHJ5RkhZYUFPNmJYUHdyRWYvL0JiR2x1YWxBaUIKaVA5Yk9TK3ozd1J2SzYyc3A1elVqOWRXOEswQVZLNmVYWWZISHRXVXdBY0lVSWFHczN1cFhZa0NnVFFQV2poeApNNFQrckZJT2w0ZS93N1lscVRWNUZ3SURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVVS0JpQ1BxejQ4MkN3eFQ5d1pHbTl6NSs4TzB3RFFZSktvWkkKaHZjTkFRRUxCUUFEZ2dFQkFLNm5aMzM1clgzWUp6TlF1SmFvUERCd3FKeTRPQ2s3YkNOdTU0TzAwWWRMK1pGWgpoQS9OWDN4TWNRUDREYk9EN0lDSjAxVTROTjBZMmEvMURGT2dPN1JKai90VmNVWUlhQU9mUFRLbVNrS1p0UnNuClVXZmkxLzVJSGF6ZnAwdDdZUXQwZmFDT1pPdkRyTm45TC9jeEtBS2FpZ1hIUmpLQndsZ2wvKyt2dDdBa0Z4R0IKT3J1aTN3Y29MU1JCUnpXQ1pEclpzejFnMkVwNW02V2ZOaTA1WndHMUNYRmhHdFlUUFZrMm9JWmdBYWwzYTkrVAprS1BNNXVRRWRJU2VwUzQ1b0Q1Qzk1ZS93OG9hMXkzTjdXMXFLOVJIbW9kT1VpTklQUE44aVNtSTFOZ0NXYTNXCmtOeEFsc2JIUG5SSjJXSy9zUUhuME5QN0FyU3hkaWNaWGtUQjZVbz0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"\n    #server: https://kubernetes.skyguard.cluster.local:6443\n    server: https://172.30.3.229:59100\n    insecure-skip-tls-verify: true\nusers:\n- name: devops-cluster-admin\n  user:\n    token: eyJhbGciOiJSUzI1NiIsImtpZCI6Ijh3eW1qX0RtcXBUMXZKMkR1UVhEZ1R2ZEt2WjVvV0pjalg4NWkyMWd2TE0ifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXZvcHMtY2x1c3Rlci1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZXZvcHMtY2x1c3Rlci1hZG1pbi10b2tlbi02NG1sNSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJkZXZvcHMtY2x1c3Rlci1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjVhNjZhYmUzLTgxMmEtNDg0MS1hMmZlLTdkOThhZWQzN2YyMyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXZvcHMtY2x1c3Rlci1hZG1pbjpkZXZvcHMtY2x1c3Rlci1hZG1pbiJ9.ekEfzxxOOMdaKYOIhfBH_bcdwS9yEmiKbr2cjKDXQwnluxlXv5xdDSchoOSb26m9grg1acY6daoq8605zWr72LcHI-ZgLk6sDjQwQI-2bMFksyZAmDX6TdW7uNuIceu3LDJpikHWe4wiuhBzqwGPJ6CpBa_t_v_c6cNKEa97mhEAghbIPv-D4Z91kGgmBlTdyEkJBoqMiH-uRycaTsjAWsvALLPKGuE4RJBjyNHY-3eL_3_m6WCeL5x8kXM9xPvp-WdUMqjTKDM8kiVFqSfhF-BV8PnNmrabkmdamcSR1XjK3-PSr--NTSuay_kW_LlvISENAkpVN1mxrYMizE4zyQ"

func main() {

	pops := []*K8sProxyPram{
		{
			PopName:    "cd-ops-192",
			KubeConfig: kubeconfig192,
		},
		{
			PopName:    "cd-pop-229",
			KubeConfig: kubeconfig229,
		},
	}
	for _, pop := range pops {
		k8sProxy, err := NewK8sProxy(pop)
		if err != nil {
			return
		}
		err = k8sProxy.CheckConnection(context.Background())
		if err != nil {
			return
		}
	}
}

func NewK8sProxy(k8sProxyPram *K8sProxyPram) (*K8sProxy, error) {
	proxy := &K8sProxy{
		popName:    k8sProxyPram.PopName,
		KubeConfig: k8sProxyPram.KubeConfig,
	}

	configFile := "kubeConfig." + proxy.popName + ".yaml"
	if err := os.WriteFile(configFile, []byte(proxy.KubeConfig), 0o666); err != nil {
		return nil, errors.Wrap(err, "create kubeConfig temp file error")
	}
	defer os.Remove(configFile)
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		return nil, errors.Wrap(err, "build rest config error")
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "create clientSet error")
	}

	restClient := clientSet.CoreV1().RESTClient().(*rest.RESTClient)

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "create dynamicClient error")
	}

	metricsClientset, err := metrics.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "create metricsClientset error")
	}

	proxy.restConfig = config
	proxy.Clientset = clientSet
	proxy.DynamicClient = dynamicClient
	proxy.RestClient = restClient
	proxy.metricsClientset = metricsClientset
	return proxy, nil
}

func (p *K8sProxy) CheckConnection(ctx context.Context) error {
	version, err := p.getVersion(ctx)
	if err != nil {
		return errors.Wrapf(err, "get pop [%s] version error", p.popName)
	}
	fmt.Println("success connection", version)
	return nil
}

type K8sVersionResponse struct {
	GitVersion string
}

func (p *K8sProxy) getVersion(ctx context.Context) (string, error) {
	getVersionURL := p.restConfig.Host + "/version"
	bytes, err := p.get(ctx, getVersionURL)
	if err != nil {
		return "", err
	}
	k8sVersionResponse := K8sVersionResponse{}
	if err = jsoniter.Unmarshal(bytes, &k8sVersionResponse); err != nil {
		return "", errors.Wrap(err, "parse k8s version response json error")
	}
	return k8sVersionResponse.GitVersion, err
}

// RestClient的Get请求
func (p *K8sProxy) get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "create request error")
	}
	req = req.WithContext(ctx)
	resp, err := p.RestClient.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "get request error")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "copy response body error")
	}
	return data, nil
}
