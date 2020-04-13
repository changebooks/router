package route_sample

import (
	"github.com/changebooks/router"
	"github.com/changebooks/router/route_sample/module_sample"
)

func NewRoute() (*router.Route, error) {
	route := router.NewRoute()

	if err := route.AddInterceptor(NewInterceptor()); err != nil {
		return nil, err
	}

	module, err := module_sample.NewModule()
	if err != nil {
		return nil, err
	}

	if err := route.PutModule(module); err != nil {
		return nil, err
	}

	return route, nil
}
