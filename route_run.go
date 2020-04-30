package router

import (
	"errors"
	"github.com/changebooks/kernel"
)

func (x *Route) Run(holder *Holder) error {
	if holder == nil {
		return errors.New("holder can't be nil")
	}

	module, controller, action, err := x.AnalyzePath(holder.Path)
	if err != nil {
		holder.Result = kernel.NewPageNotFound(nil)
		return err
	}

	if holder.Request == nil {
		return errors.New("request can't be nil")
	}

	if err = CheckMethod(holder.Request.Method, action.AllowMethods()); err != nil {
		holder.Result = kernel.NewMethodNotAllowed(nil)
		return err
	}

	err = x.Before(holder)
	defer x.After(holder)
	if err != nil {
		return err
	}

	err = module.Before(holder)
	defer module.After(holder)
	if err != nil {
		return err
	}

	err = controller.Before(holder)
	defer controller.After(holder)
	if err != nil {
		return err
	}

	err = action.Before(holder)
	defer action.After(holder)
	if err != nil {
		return err
	}

	holder.Result = action.Run(holder)
	return nil
}

func (x *Route) AnalyzePath(path *Path) (module *Module, controller *Controller, action Action, err error) {
	if path == nil {
		err = errors.New("path can't be nil")
		return
	}

	if module, err = x.GetModule(path.GetModule()); err != nil {
		return
	}

	if controller, err = module.GetController(path.GetController()); err != nil {
		return
	}

	if action, err = controller.GetAction(path.GetAction()); err != nil {
		return
	}

	return
}
