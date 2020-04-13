package controller_sample

import "github.com/changebooks/router"

func NewController() (*router.Controller, error) {
	controller, err := router.NewController("controller")
	if err != nil {
		return nil, err
	}

	if err := controller.AddInterceptor(NewInterceptor()); err != nil {
		return nil, err
	}

	if err := controller.PutAction(NewTest()); err != nil {
		return nil, err
	}

	return controller, nil
}
