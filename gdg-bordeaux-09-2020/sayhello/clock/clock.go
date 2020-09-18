package clock

import "time"

type Clock struct{}

func New() *Clock {
	return &Clock{}
}

func (*Clock) Now() string {
	return time.Now().UTC().Format(time.Kitchen)
}
