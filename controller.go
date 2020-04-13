package router

import (
	"errors"
	"fmt"
	"sync"
)

type Controller struct {
	mu           sync.Mutex        // ensures atomic writes; protects the following fields
	name         string            // controller's name，全小写，"_"拼接
	actions      map[string]Action // [action's name => Action]
	interceptors []Interceptor     // 拦截
}

func NewDefaultController() *Controller {
	r, _ := NewController(DefaultController)
	return r
}

func NewController(name string) (*Controller, error) {
	if name = PathFormat(name); name == "" {
		return nil, errors.New("controller's name can't be empty")
	}

	return &Controller{
		name:    name,
		actions: make(map[string]Action),
	}, nil
}

func (x *Controller) GetName() string {
	return x.name
}

func (x *Controller) GetAction(name string) (Action, error) {
	if name == "" {
		return nil, errors.New("action's name can't be empty")
	}

	if r, ok := x.actions[name]; ok {
		if r != nil {
			return r, nil
		} else {
			return nil, fmt.Errorf("action %q can't be nil", name)
		}
	} else {
		return nil, fmt.Errorf("no action %q", name)
	}
}

func (x *Controller) PutAction(a Action) error {
	if a == nil {
		return errors.New("action can't be nil")
	}

	name := a.GetName()
	if name == "" {
		return errors.New("action's name can't be empty")
	}

	if _, ok := x.actions[name]; ok {
		return fmt.Errorf("duplicated action %q", name)
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if _, ok := x.actions[name]; ok {
		return fmt.Errorf("duplicated action %q", name)
	}

	x.actions[name] = a
	return nil
}

func (x *Controller) GetActions() map[string]Action {
	return x.actions
}

func (x *Controller) Before(holder *Holder) error {
	interceptors := x.GetInterceptors()
	if len(interceptors) > 0 {
		for _, i := range interceptors {
			if err := i.Before(holder); err != nil {
				return err
			}
		}
	}

	return nil
}

func (x *Controller) After(holder *Holder) {
	interceptors := x.GetInterceptors()
	if len(interceptors) > 0 {
		for _, i := range interceptors {
			i.After(holder)
		}
	}
}

func (x *Controller) GetInterceptors() []Interceptor {
	return x.interceptors
}

func (x *Controller) AddInterceptor(i Interceptor) error {
	if i == nil {
		return errors.New("interceptor can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.interceptors = append(x.interceptors, i)
	return nil
}
