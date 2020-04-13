package router

import "strings"

// 分析RequestURI，去掉"?query"
func InterceptUri(s string) string {
	i := strings.IndexRune(s, UriSeparator)
	if i > -1 {
		return s[0:i]
	} else {
		return s
	}
}

// 分析RequestURI，获取"模块名"、"控制器名"、"方法名"、"参数列表"
func AnalyzeUri(s string) (module string, controller string, action string, parameters []string) {
	segments := strings.Split(strings.Trim(strings.TrimSpace(s), PathSeparator), PathSeparator)

	size := len(segments)
	switch size {
	case PathSegment0:
		break
	case PathSegment1: // action
		action = segments[0]
		break
	case PathSegment2: // controller/action
		controller = segments[0]
		action = segments[1]
		break
	case PathSegment3: // module/controller/action
		module = segments[0]
		controller = segments[1]
		action = segments[2]
		break
	default:
		// module/controller/action/parameter1
		// module/controller/action/parameter1/parameter2
		module = segments[0]
		controller = segments[1]
		action = segments[2]
		parameters = make([]string, size-PathSegment3)
		for i := PathSegment3; i < size; i++ {
			parameters[i-PathSegment3] = segments[i]
		}
		break
	}

	return
}
