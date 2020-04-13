package router

import (
	"errors"
	"fmt"
	"sync"
)

type Module struct {
	mu           sync.Mutex             // ensures atomic writes; protects the following fields
	name         string                 // module's name，全小写，"_"拼接
	controllers  map[string]*Controller // [controller's name => *Controller]
	interceptors []Interceptor          // 拦截
}

func NewDefaultModule() *Module {
	r, _ := NewModule(DefaultModule)
	return r
}

func NewModule(name string) (*Module, error) {
	if name = PathFormat(name); name == "" {
		return nil, errors.New("module's name can't be empty")
	}

	return &Module{
		name:        name,
		controllers: make(map[string]*Controller),
	}, nil
}

func (x *Module) GetName() string {
	return x.name
}

func (x *Module) GetController(name string) (*Controller, error) {
	if name == "" {
		return nil, errors.New("controller's name can't be empty")
	}

	if r, ok := x.controllers[name]; ok {
		if r != nil {
			return r, nil
		} else {
			return nil, fmt.Errorf("controller %q can't be nil", name)
		}
	} else {
		return nil, fmt.Errorf("no controller %q", name)
	}
}

func (x *Module) PutController(c *Controller) error {
	if c == nil {
		return errors.New("controller can't be nil")
	}

	name := c.GetName()
	if name == "" {
		return errors.New("controller's name can't be empty")
	}

	if _, ok := x.controllers[name]; ok {
		return fmt.Errorf("duplicated controller %q", name)
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if _, ok := x.controllers[name]; ok {
		return fmt.Errorf("duplicated controller %q", name)
	}

	x.controllers[name] = c
	return nil
}

func (x *Module) GetControllers() map[string]*Controller {
	return x.controllers
}

func (x *Module) Before(holder *Holder) error {
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

func (x *Module) After(holder *Holder) {
	interceptors := x.GetInterceptors()
	if len(interceptors) > 0 {
		for _, i := range interceptors {
			i.After(holder)
		}
	}
}

func (x *Module) GetInterceptors() []Interceptor {
	return x.interceptors
}

func (x *Module) AddInterceptor(i Interceptor) error {
	if i == nil {
		return errors.New("interceptor can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.interceptors = append(x.interceptors, i)
	return nil
}
