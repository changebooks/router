package module_sample

import (
	"github.com/changebooks/router"
	"github.com/changebooks/router/route_sample/module_sample/controller_sample"
)

func NewModule() (*router.Module, error) {
	module, err := router.NewModule("module")
	if err != nil {
		return nil, err
	}

	if err := module.AddInterceptor(NewInterceptor()); err != nil {
		return nil, err
	}

	controller, err := controller_sample.NewController()
	if err != nil {
		return nil, err
	}

	if err := module.PutController(controller); err != nil {
		return nil, err
	}

	return module, nil
}
