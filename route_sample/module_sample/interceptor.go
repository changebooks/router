package module_sample

import (
	"fmt"
	"github.com/changebooks/router"
)

type Interceptor struct {
	router.Interceptor
}

func NewInterceptor() *Interceptor {
	return &Interceptor{}
}

func (x *Interceptor) Before(holder *router.Holder) error {
	fmt.Println("module_sample's Interceptor::Before()")
	return nil
}

func (x *Interceptor) After(holder *router.Holder) {
	fmt.Println("module_sample's Interceptor::After()")
}
