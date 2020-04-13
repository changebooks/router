package router

import (
	"errors"
	"fmt"
)

// 支持的请求方式？
// method 和 allowMethods 全大写
func CheckMethod(method string, allowMethods []string) error {
	if method == "" {
		return errors.New("method can't be empty")
	}

	if allowMethods == nil {
		return errors.New("allow methods can't be nil")
	}

	if len(allowMethods) == 0 {
		return errors.New("allow methods can't be empty")
	}

	for _, allow := range allowMethods {
		if method == allow {
			return nil
		}
	}

	return fmt.Errorf("unsupported method %q in %q", method, allowMethods)
}
