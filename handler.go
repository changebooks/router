package router

import (
	"errors"
	"github.com/changebooks/kernel"
	"github.com/changebooks/log"
	"net/http"
	"time"
)

type Handler struct {
	route        *Route
	logger       *log.Logger
	systemRunErr []byte
}

func NewHandler(route *Route, logger *log.Logger) (*Handler, error) {
	if route == nil {
		return nil, errors.New("route can't be nil")
	}

	if logger == nil {
		return nil, errors.New("logger can't be nil")
	}

	systemRunErr, _ := kernel.ResultFormat(kernel.NewSystemRunErr(nil))

	return &Handler{
		route:        route,
		logger:       logger,
		systemRunErr: systemRunErr,
	}, nil
}

func (x *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tag := "ServeHTTP"

	holder, remark, err := x.Run(w, req)

	var idRegister *log.IdRegister
	if holder != nil {
		idRegister = holder.IdRegister
	}

	if remark == nil {
		remark = NewRemark(req, nil, time.Now(), time.Now())
	}

	if err == nil {
		x.logger.N(tag, Success, remark, idRegister)
	} else {
		x.logger.C(tag, RunFailure, remark, err, "", idRegister)
	}

	if n, err := x.Write(holder); err != nil {
		x.logger.M(tag, WriteFailure, remark, err, "", idRegister)

		if n == 0 {
			if _, err = x.WriteSystemRunErr(w); err != nil {
				x.logger.M(tag, WriteFailure, remark, err, "", idRegister)
			}
		}
	}
}

func (x *Handler) Run(w http.ResponseWriter, req *http.Request) (*Holder, *Remark, error) {
	holder, err := NewHolder(w, req)
	if err != nil {
		return nil, nil, err
	}

	start := time.Now()

	err = x.route.Run(holder)

	done := time.Now()

	if holder == nil {
		return nil, nil, errors.New("ran holder can't be nil")
	}

	remark := NewRemark(holder.Request, holder.Result, start, done)
	return holder, remark, err
}

func (x *Handler) Write(holder *Holder) (int, error) {
	if holder == nil {
		return 0, errors.New("holder can't be nil")
	}

	if holder.ResponseWriter == nil {
		return 0, errors.New("response writer can't be nil")
	}

	if holder.Result == nil {
		return 0, errors.New("result can't be nil")
	}

	json, err := kernel.ResultFormat(holder.Result)
	if err != nil {
		return 0, err
	}

	return holder.ResponseWriter.Write(json)
}

func (x *Handler) WriteSystemRunErr(w http.ResponseWriter) (int, error) {
	if w == nil {
		return 0, errors.New("response writer can't be nil")
	}

	return w.Write(x.systemRunErr)
}

func (x *Handler) GetRoute() *Route {
	return x.route
}

func (x *Handler) GetLogger() *log.Logger {
	return x.logger
}
