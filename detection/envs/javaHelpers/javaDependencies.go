package javaHelpers

type JavaDependencies struct {
	Dependencies []string
}

type JavaEnvWithDependencies interface {
	GetDependencies() JavaDependencies
}
