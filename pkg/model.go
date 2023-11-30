package pkg

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		Header     Header         `json:"header"`
		Meta       Meta           `json:"meta"`
		Data       interface{}    `json:"data"`
		Pagination Pagination     `json:"pagination"`
		Error      ErrorAttribute `json:"error"`
	}
	Header struct {
		Message     string  `json:"message"`
		ProcessTime float64 `json:"process_time"`
		Status      int     `json:"-"`
	}
	Pagination struct {
		Sort       []Sort `json:"sort"`
		LastCursor int    `json:"last_cursor"`
		Size       int    `json:"size"`
		TotalPage  int    `json:"total_page"`
		TotalData  int    `json:"total_data"`
	}
	Sort struct {
		Type  string `json:"type"`
		Field string `json:"field"`
	}
	Meta struct {
		Behavior            int  `json:"behavior"`
		CategoryRedirection bool `json:"category_redirection"`
	}
	ErrorAttribute struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Code    int64  `json:"code"`
	}
)

func (resp Response) Render(w http.ResponseWriter, r *http.Request) {
	if resp.Header.Message == "" {
		if resp.Header.Status == 200 {
			resp.Header.Message = "success"
		} else {
			resp.Header.Message = "fail"
		}
	}
	if len(resp.Pagination.Sort) == 0 {
		resp.Pagination.Sort = make([]Sort, 0)
	}
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(resp.Header.Status)
	_, _ = w.Write(body)
}
