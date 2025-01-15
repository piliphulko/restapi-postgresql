package f16

import "strings"

func ignoringHealthRequest(infoFullMethod string) (methodName string, ok bool) {
	parts := strings.Split(infoFullMethod, "/")
	methodName = parts[len(parts)-1]

	if methodName == "Check" {
		return "", true
	} else {
		return methodName, false
	}
}
