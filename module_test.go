package router

import "testing"

func TestModuleName(t *testing.T) {
	_, got := NewModule("")
	want := "module's name can't be empty"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := NewDefaultModule()
	want2 := "default"
	if got2 == nil {
		t.Errorf("got nil; want &Module{}")
	} else if got2.GetName() != want2 {
		t.Errorf("got %q; want %q", got2.GetName(), want2)
	}
}

func TestModuleController(t *testing.T) {
	module := NewDefaultModule()

	_, got := module.GetController("")
	want := "controller's name can't be empty"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := module.GetController("index")
	want2 := "no controller \"index\""
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := module.PutController(nil)
	want3 := "controller can't be nil"
	if got3 == nil || got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}

func TestModuleInterceptor(t *testing.T) {
	task := NewDefaultModule()

	got := task.AddInterceptor(nil)
	want := "interceptor can't be nil"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
