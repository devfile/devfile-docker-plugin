package javaHelpers

type JavaDependenciesPair struct {
	//Path string
	Name string
}

type JavaDependencies struct {
	Dependencies []JavaDependenciesPair
}

type JavaEnvWithDependencies interface {
	GetDependencies() JavaDependencies
}
