package router

import "testing"

func TestRouteModule(t *testing.T) {
	route := NewRoute()

	_, got := route.GetModule("")
	want := "module's name can't be empty"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := route.GetModule("index")
	want2 := "no module \"index\""
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := route.PutModule(nil)
	want3 := "module can't be nil"
	if got3 == nil || got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}

func TestRouteInterceptor(t *testing.T) {
	task := NewRoute()

	got := task.AddInterceptor(nil)
	want := "interceptor can't be nil"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
