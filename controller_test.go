package router

import "testing"

func TestControllerName(t *testing.T) {
	_, got := NewController("")
	want := "controller's name can't be empty"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := NewDefaultController()
	want2 := "default"
	if got2 == nil {
		t.Errorf("got nil; want &Controller{}")
	} else if got2.GetName() != want2 {
		t.Errorf("got %q; want %q", got2.GetName(), want2)
	}
}

func TestControllerAction(t *testing.T) {
	controller := NewDefaultController()

	_, got := controller.GetAction("")
	want := "action's name can't be empty"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := controller.GetAction("index")
	want2 := "no action \"index\""
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := controller.PutAction(nil)
	want3 := "action can't be nil"
	if got3 == nil || got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}

func TestControllerInterceptor(t *testing.T) {
	controller := NewDefaultController()

	got := controller.AddInterceptor(nil)
	want := "interceptor can't be nil"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
