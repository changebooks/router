package router

import (
	"errors"
	"github.com/changebooks/log"
	"net/http"
)

func NewIdRegister(req *http.Request) (*log.IdRegister, error) {
	if req == nil {
		return nil, errors.New("req can't be nil")
	}

	traceId := ""
	logId := ""

	header := req.Header
	if header != nil {
		traceId = header.Get(ReqTraceId)
		logId = header.Get(ReqLogId)
	}

	if traceId == "" {
		traceId = log.IdGenerator()
	}

	r := &log.IdRegister{}
	r.SetTraceId(traceId)
	_ = r.SetId(logId)

	return r, nil
}
