package router

import (
	"net/http"
	"testing"
)

func TestCheckMethod(t *testing.T) {
	got := CheckMethod("", nil)
	want := "method can't be empty"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := CheckMethod(http.MethodGet, nil)
	want2 := "allow methods can't be nil"
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := CheckMethod(http.MethodGet, []string{})
	want3 := "allow methods can't be empty"
	if got3 == nil || got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := CheckMethod("GET", []string{"POST"})
	want4 := `unsupported method "GET" in ["POST"]`
	if got4 == nil || got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := CheckMethod("GET", []string{"Get", "POST"})
	want5 := `unsupported method "GET" in ["Get" "POST"]`
	if got5 == nil || got5.Error() != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := CheckMethod("GET", []string{"GET"})
	if got6 != nil {
		t.Errorf("got %v; want %v", got6, nil)
	}

	got7 := CheckMethod("POST", []string{"GET", "POST"})
	if got7 != nil {
		t.Errorf("got %v; want %v", got7, nil)
	}
}
