package ports

type PingPort interface {
	Ping() (string, error)
}
