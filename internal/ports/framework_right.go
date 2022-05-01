package ports

type DbPort interface {
	Ping() (string, error)
}

