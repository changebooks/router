package controller_sample

import (
	"fmt"
	"github.com/changebooks/http"
	"github.com/changebooks/kernel"
	"github.com/changebooks/router"
)

type Test struct {
	router.Action
}

func NewTest() *Test {
	return &Test{}
}

func (x *Test) GetName() string {
	return "test"
}

func (x *Test) Before(holder *router.Holder) error {
	fmt.Println("Test::Before()")
	return nil
}

func (x *Test) Run(holder *router.Holder) *kernel.Result {
	fmt.Println("Test::Run()")
	return kernel.NewSuccess(nil)
}

func (x *Test) After(holder *router.Holder) {
	fmt.Println("Test::After()")
}

func (x *Test) AllowMethods() []string {
	return []string{http.MethodPost}
}
