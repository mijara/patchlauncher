package port

type OpenerService interface {
	Open(file string) error
}
