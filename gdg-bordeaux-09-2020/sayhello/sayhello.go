package sayhello

import (
	"fmt"
)

var hellos = map[string]string{
	"FR": "Salut, il est", // 2.0.0.0
	"GB": "Hello, it is",  // 3.8.16.0
	"ES": "Hola, son las", // 2.136.0.0
}

type SayHello struct {
	db      DB
	ipStack IPStack
	clock   Clock
}

func New(db DB, ipStack IPStack, clock Clock) *SayHello {
	return &SayHello{
		db:      db,
		ipStack: ipStack,
		clock:   clock,
	}
}

func (sh *SayHello) SayHello(ip string) (string, error) {

	// 1. Call ipstack API to fetch IP location
	cc, err := sh.ipStack.GetCountryCode(ip)
	if err != nil {
		return "", err
	}

	// 2. Increase IP hits for tracability/ressource threshold
	sh.db.Hit(ip)
	if err != nil {
		return "", err
	}

	// 3. Get translation
	hello, ok := hellos[cc]
	if !ok {
		hello = hellos["GB"]
	}

	// 5. Print time in writer
	return fmt.Sprintf("%s %s\n", hello, sh.clock.Now()), nil
}
