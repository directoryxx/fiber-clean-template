package interfaces

type Logger interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
