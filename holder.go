package router

import (
	"errors"
	"github.com/changebooks/kernel"
	"github.com/changebooks/log"
	"net/http"
	"sync"
)

type Holder struct {
	mu             sync.Mutex // ensures atomic writes; protects the following fields
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Path           *Path
	IdRegister     *log.IdRegister
	Result         *kernel.Result
	attributes     map[string]interface{}
}

func NewHolder(w http.ResponseWriter, req *http.Request) (*Holder, error) {
	if w == nil {
		return nil, errors.New("response writer can't be nil")
	}

	if req == nil {
		return nil, errors.New("request can't be nil")
	}

	idRegister, _ := NewIdRegister(req)
	path := NewPath(req.RequestURI)

	return &Holder{
		ResponseWriter: w,
		Request:        req,
		Path:           path,
		IdRegister:     idRegister,
		Result:         nil,
		attributes:     make(map[string]interface{}),
	}, nil
}

func (x *Holder) GetAttribute(key string) interface{} {
	return x.attributes[key]
}

func (x *Holder) SetAttribute(key string, value interface{}) {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.attributes[key] = value
}

func (x *Holder) GetAttributes() map[string]interface{} {
	return x.attributes
}
