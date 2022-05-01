package ports

type APIPort interface {
	Ping() (string, error)
}
