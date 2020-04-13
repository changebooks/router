package router

import "testing"

func TestPathFormat(t *testing.T) {
	got := PathFormat("")
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := PathFormat(" Abc-Def  ")
	want2 := "abc_def"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := PathFormat(" -A-b_c-  ")
	want3 := "_a_b_c_"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}
