package router

type Path struct {
	uri        string   // RequestURI，无"?query"
	module     string   // 模块名：全小写，"_"拼接
	controller string   // 控制器名：全小写，"_"拼接
	action     string   // 方法名：全小写，"_"拼接
	parameters []string // 参数
}

// s is a http.Request.RequestURI (eg. /a/b/c?a=1)
func NewPath(s string) *Path {
	uri := InterceptUri(s)

	module, controller, action, parameters := AnalyzeUri(uri)

	if module = PathFormat(module); module == "" {
		module = DefaultModule
	}

	if controller = PathFormat(controller); controller == "" {
		controller = DefaultController
	}

	if action = PathFormat(action); action == "" {
		action = DefaultAction
	}

	return &Path{
		uri:        uri,
		module:     module,
		controller: controller,
		action:     action,
		parameters: parameters,
	}
}

func (x *Path) GetUri() string {
	return x.uri
}

func (x *Path) GetModule() string {
	return x.module
}

func (x *Path) GetController() string {
	return x.controller
}

func (x *Path) GetAction() string {
	return x.action
}

func (x *Path) GetParameters() []string {
	return x.parameters
}
