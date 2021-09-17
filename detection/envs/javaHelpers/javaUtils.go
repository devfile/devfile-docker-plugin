package javaHelpers

import (
	"errors"
	"strings"
)

func ConvertJavaVersion(ver string) (string, error) {
	verSplit := strings.Split(ver, ".")
	if len(verSplit) == 0 {
		return "", errors.New("not enough version components")
	}
	if verSplit[0] == "1" && len(verSplit) > 1 {
		verSplit = verSplit[1:]
	}
	return strings.Join(verSplit, "."), nil
}
