package app

import "net/http"

// URLAddOnetime 根据请求往数据库中添加一个一次性事件
func URLAddOnetime(w http.ResponseWriter, r *http.Request) {
	handler := CreateHandler(w, r)
	body := OnetimeBody{}
	err := handler.DecodePost(&body)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	newOnetime := Service.Onetime.CreateOnetime(body.OwnID, body.DDL, body.Content)
	err = Service.Onetime.PostOnetime(newOnetime)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

// URLDeleteOnetime 根据URL的请求从数据库中删除一个一次性事件
func URLDeleteOnetime(w http.ResponseWriter, r *http.Request) {
	handler := CreateHandler(w, r)
	id := handler.DecodePath(2)
	err := Service.Onetime.DeleteOnetime(id)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}
