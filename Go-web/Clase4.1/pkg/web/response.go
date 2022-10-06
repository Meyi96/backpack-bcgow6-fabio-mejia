package web

import "strconv"

type Responese struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) (int, Responese) {
	if code < 300 {
		return code, Responese{strconv.FormatInt(int64(code), 10), data, ""}
	}
	return code, Responese{strconv.FormatInt(int64(code), 10), nil, err}
}
