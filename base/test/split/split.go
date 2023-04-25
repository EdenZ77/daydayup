package split

import "strings"

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		// 当分隔符出现在第一个位置时，i=0
		if i != 0 {
			result = append(result, s[:i])
		}
		// 分隔符可能是多个字符组成
		// 当分隔符出现在最后一个位置时，i+len(sep)将越界
		if i+len(sep) < len(s) {
			s = s[i+len(sep):]
		} else {
			s = ""
			break
		}
		i = strings.Index(s, sep)
	}
	if len(s) > 0 {
		result = append(result, s)
	}
	return
}
