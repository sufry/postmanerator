package themes

import (
	"strings"
)

func helperMethod(method string) string {
	switch strings.ToLower(method) {
	case "get":
		return "primary"
	case "post":
		return "success"
	case "put":
		return "info"
	case "patch":
		return "warning"
	case "delete":
		return "danger"
	default:
		return "default"
	}
}
