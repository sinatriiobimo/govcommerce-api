package middleware

import (
	"encoding/json"
	"net/http"
	"tlkm-api/pkg"
)

func setError(w http.ResponseWriter, code int, err error) {
	resp := pkg.Response{}
	resp.Header.Message = err.Error()
	resp.Header.Status = http.StatusUnauthorized
	errMsg, _ := JsonMarshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(errMsg)
}

func JsonMarshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	return data, err
}
