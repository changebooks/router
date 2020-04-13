package router

import (
	"errors"
	"fmt"
	"sync"
)

type Route struct {
	mu           sync.Mutex         // ensures atomic writes; protects the following fields
	modules      map[string]*Module // [module's name => *Module]
	interceptors []Interceptor      // 拦截
}

func NewRoute() *Route {
	return &Route{
		modules: make(map[string]*Module),
	}
}

func (x *Route) GetModule(name string) (*Module, error) {
	if name == "" {
		return nil, errors.New("module's name can't be empty")
	}

	if r, ok := x.modules[name]; ok {
		if r != nil {
			return r, nil
		} else {
			return nil, fmt.Errorf("module %q can't be nil", name)
		}
	} else {
		return nil, fmt.Errorf("no module %q", name)
	}
}

func (x *Route) PutModule(m *Module) error {
	if m == nil {
		return errors.New("module can't be nil")
	}

	name := m.GetName()
	if name == "" {
		return errors.New("module's name can't be empty")
	}

	if _, ok := x.modules[name]; ok {
		return fmt.Errorf("duplicated module %q", name)
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if _, ok := x.modules[name]; ok {
		return fmt.Errorf("duplicated module %q", name)
	}

	x.modules[name] = m
	return nil
}

func (x *Route) GetModules() map[string]*Module {
	return x.modules
}

func (x *Route) Before(holder *Holder) error {
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

func (x *Route) After(holder *Holder) {
	interceptors := x.GetInterceptors()
	if len(interceptors) > 0 {
		for _, i := range interceptors {
			i.After(holder)
		}
	}
}

func (x *Route) GetInterceptors() []Interceptor {
	return x.interceptors
}

func (x *Route) AddInterceptor(i Interceptor) error {
	if i == nil {
		return errors.New("interceptor can't be nil")
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.interceptors = append(x.interceptors, i)
	return nil
}
