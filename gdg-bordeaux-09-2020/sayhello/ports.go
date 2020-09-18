package sayhello

//go:generate mockgen -source ports.go -destination portsmock_test.go -package sayhello_test

type IPStack interface {
	GetCountryCode(IP string) (string, error)
}

type DB interface {
	Hit(IP string) error
}

type Clock interface {
	Now() string
}
