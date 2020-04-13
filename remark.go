package router

import (
	"github.com/changebooks/kernel"
	"net/http"
	"net/url"
	"time"
)

type Remark struct {
	Uri      string        `json:"uri"`
	Method   string        `json:"method"`
	Header   http.Header   `json:"header"`
	Form     url.Values    `json:"form"`
	PostForm url.Values    `json:"postForm"`
	Code     uint          `json:"code"`
	Message  string        `json:"message"`
	Data     interface{}   `json:"data"`
	Total    time.Duration `json:"total"`
	Start    time.Time     `json:"start"`
	Done     time.Time     `json:"done"`
}

func NewRemark(req *http.Request, result *kernel.Result, start time.Time, done time.Time) *Remark {
	total := done.Sub(start)

	var uri string
	var method string
	var header http.Header
	var form url.Values
	var postForm url.Values

	if req != nil {
		uri = req.RequestURI
		method = req.Method
		header = req.Header
		form = req.Form
		postForm = req.PostForm
	}

	var code uint
	var message string
	var data interface{}

	if result != nil {
		code = result.GetCode()
		message = result.GetMessage()
		data = result.GetData()
	}

	return &Remark{
		Uri:      uri,
		Method:   method,
		Header:   header,
		Form:     form,
		PostForm: postForm,
		Code:     code,
		Message:  message,
		Data:     data,
		Total:    total,
		Start:    start,
		Done:     done,
	}
}
