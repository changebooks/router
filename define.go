package router

const (
	ReqTraceId        = "trace_id" // Header/Request取日志链路id的Key
	ReqLogId          = "log_id"   // Header/Request取日志id的Key
	UriSeparator      = '?'        // module/controller/action/parameter1/parameter2?query
	PathSeparator     = "/"        // module/controller/action/parameter1/parameter2
	DefaultModule     = "default"
	DefaultController = "default"
	DefaultAction     = "default"
	PathSegment0      = 0 // 模块名（默认）、控制器名（默认）、方法名（默认）(eg. default/default/default)
	PathSegment1      = 1 // 模块名（默认）、控制器名（默认）、方法名（指定）(eg. default/default/action)
	PathSegment2      = 2 // 模块名（默认）、控制器名（指定）、方法名（指定）(eg. default/controller/action)
	PathSegment3      = 3 // 模块名（指定）、控制器名（指定）、方法名（指定）(eg. module/controller/action)
	Success           = "successful"
	RunFailure        = "run failed"
	WriteFailure      = "write failed"
)
