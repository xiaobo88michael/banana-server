package utils

import (
	"github.com/google/uuid"
	"runtime"
	"strings"
)

func GetSelfFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return cleanUpFuncName(runtime.FuncForPC(pc).Name())
}
func cleanUpFuncName(funcName string) string {
	end := strings.LastIndex(funcName, ".")
	if end == -1 {
		return ""
	}
	return funcName[end+1:]
}

func NewUUID() string {
	return uuid.NewString()
}
