package app

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Handler 处理好封装的Handler，直接传需要返回的东西
type Handler struct {
	w http.ResponseWriter
	r *http.Request
}

// DecodePath 根据url获得想要的参数
func (h *Handler) DecodePath(index int) string {
	urls := strings.Split(h.r.URL.Path, "/")
	return urls[index]
}

// DecodeQuery 读取需要请求的数据
func (h *Handler) DecodeQuery(key string) string {
	return h.r.URL.Query().Get(key)
}

// DecodePost 读取POST数据
func (h *Handler) DecodePost(body interface{}) error {
	err := h.r.ParseForm()
	if err != nil {
		return err
	}
	return json.NewDecoder(h.r.Body).Decode(&body)
}

// Send 发送数据
func (h *Handler) Send(body interface{}, state bool) (err error) {
	h.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	h.w.WriteHeader(http.StatusOK)

	res := &Response{
		State:    state,
		Response: body}

	msg, _ := json.Marshal(res)
	h.w.Write(msg)

	return nil
}

// CreateHandler 封装处理handler
func CreateHandler(w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{w, r}
}
