package port

type Logger interface {
	Info(string)
	Debug(string)
}
