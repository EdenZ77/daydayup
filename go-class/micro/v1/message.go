package v1

type Request struct {
	ServiceName string
	Method      string
	Data        []byte
}

type Response struct {
	Data  []byte
	Error string
	Meta  map[string]string
}
