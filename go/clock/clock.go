package clock

import "fmt"

type Clock int

const DAY_IN_MINUTES = 1440

func New(hour int, minute int) Clock {
	c := Clock((60*hour + minute) % DAY_IN_MINUTES)

	if c < 0 {
		c += DAY_IN_MINUTES
	}

	return c
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", (c / 60), (c % 60))
}
