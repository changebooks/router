package router

import "testing"

func TestInterceptUri(t *testing.T) {
	got := InterceptUri("")
	if got != "" {
		t.Errorf("got %q; want %q", got, "")
	}

	got2 := InterceptUri("?")
	if got2 != "" {
		t.Errorf("got %q; want %q", got2, "")
	}

	got3 := InterceptUri(" ?")
	want3 := " "
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := InterceptUri(" a? ")
	want4 := " a"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}
}

func TestAnalyzeUri0(t *testing.T) {
	module, controller, action, parameters := AnalyzeUri("")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "" {
		t.Errorf("got %q; want %q", controller, "")
	}
	if action != "" {
		t.Errorf("got %q; want %q", action, "")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}

	module, controller, action, parameters = AnalyzeUri("  ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "" {
		t.Errorf("got %q; want %q", controller, "")
	}
	if action != "" {
		t.Errorf("got %q; want %q", action, "")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}

	module, controller, action, parameters = AnalyzeUri(" / ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "" {
		t.Errorf("got %q; want %q", controller, "")
	}
	if action != "" {
		t.Errorf("got %q; want %q", action, "")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}
}

func TestAnalyzeUri1(t *testing.T) {
	module, controller, action, parameters := AnalyzeUri(" a ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "" {
		t.Errorf("got %q; want %q", controller, "")
	}
	if action != "a" {
		t.Errorf("got %q; want %q", action, "a")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}

	module, controller, action, parameters = AnalyzeUri(" /a ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "" {
		t.Errorf("got %q; want %q", controller, "")
	}
	if action != "a" {
		t.Errorf("got %q; want %q", action, "a")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}

	module, controller, action, parameters = AnalyzeUri(" /a/ ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "" {
		t.Errorf("got %q; want %q", controller, "")
	}
	if action != "a" {
		t.Errorf("got %q; want %q", action, "a")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}
}

func TestAnalyzeUri2(t *testing.T) {
	module, controller, action, parameters := AnalyzeUri(" a/b ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "a" {
		t.Errorf("got %q; want %q", controller, "a")
	}
	if action != "b" {
		t.Errorf("got %q; want %q", action, "b")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}

	module, controller, action, parameters = AnalyzeUri(" /a/b/ ")
	if module != "" {
		t.Errorf("got %q; want %q", module, "")
	}
	if controller != "a" {
		t.Errorf("got %q; want %q", controller, "a")
	}
	if action != "b" {
		t.Errorf("got %q; want %q", action, "b")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}
}

func TestAnalyzeUri3(t *testing.T) {
	module, controller, action, parameters := AnalyzeUri(" a/b/c ")
	if module != "a" {
		t.Errorf("got %q; want %q", module, "a")
	}
	if controller != "b" {
		t.Errorf("got %q; want %q", controller, "b")
	}
	if action != "c" {
		t.Errorf("got %q; want %q", action, "c")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}

	module, controller, action, parameters = AnalyzeUri(" /a/b/c/ ")
	if module != "a" {
		t.Errorf("got %q; want %q", module, "a")
	}
	if controller != "b" {
		t.Errorf("got %q; want %q", controller, "b")
	}
	if action != "c" {
		t.Errorf("got %q; want %q", action, "c")
	}
	if parameters != nil {
		t.Errorf("got %q; want %v", parameters, nil)
	}
}

func TestAnalyzeUri4(t *testing.T) {
	module, controller, action, parameters := AnalyzeUri(" a/b/c/d ")
	if module != "a" {
		t.Errorf("got %q; want %q", module, "a")
	}
	if controller != "b" {
		t.Errorf("got %q; want %q", controller, "b")
	}
	if action != "c" {
		t.Errorf("got %q; want %q", action, "c")
	}
	if len(parameters) != 1 {
		t.Errorf("got %v; want %v", parameters, []string{"d"})
	}

	module, controller, action, parameters = AnalyzeUri(" /a/b/c/d/ ")
	if module != "a" {
		t.Errorf("got %q; want %q", module, "a")
	}
	if controller != "b" {
		t.Errorf("got %q; want %q", controller, "b")
	}
	if action != "c" {
		t.Errorf("got %q; want %q", action, "c")
	}
	if len(parameters) != 1 {
		t.Errorf("got %v; want %v", parameters, []string{"d"})
	}
}

func TestAnalyzeUri5(t *testing.T) {
	module, controller, action, parameters := AnalyzeUri(" a/b/c/d/e ")
	if module != "a" {
		t.Errorf("got %q; want %q", module, "a")
	}
	if controller != "b" {
		t.Errorf("got %q; want %q", controller, "b")
	}
	if action != "c" {
		t.Errorf("got %q; want %q", action, "c")
	}
	if len(parameters) != 2 {
		t.Errorf("got %v; want %v", parameters, []string{"d", "e"})
	}

	module, controller, action, parameters = AnalyzeUri(" /a/b/c/d/e/ ")
	if module != "a" {
		t.Errorf("got %q; want %q", module, "a")
	}
	if controller != "b" {
		t.Errorf("got %q; want %q", controller, "b")
	}
	if action != "c" {
		t.Errorf("got %q; want %q", action, "c")
	}
	if len(parameters) != 2 {
		t.Errorf("got %v; want %v", parameters, []string{"d", "e"})
	}
}
