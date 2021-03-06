package pigpiod

import (
	"context"
	"net"
)

func Connect(ctx context.Context, address string) (*Conn, error) {
	var dialer net.Dialer
	tcp, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	c := Conn{
		Tcp:             tcp,
		DutyCycleRanges: make(map[int]uint32),
	}
	return &c, nil
}

type Conn struct {
	Tcp             net.Conn
	DutyCycleRanges map[int]uint32
}

func New(c *Conn) *Conn {
	return c
}

func (c *Conn) Close() error {
	return c.Tcp.Close()
}

type GpioMode uint32

const (
	ModeInput  GpioMode = 0
	ModeOutput GpioMode = 1
	ModeAlt0   GpioMode = 4
	ModeAlt1   GpioMode = 5
	ModeAlt2   GpioMode = 6
	ModeAlt3   GpioMode = 7
	ModeAlt4   GpioMode = 3
	ModeAlt5   GpioMode = 2
)

type Level uint32

const (
	LevelLow  Level = 0
	LevelHigh Level = 1
)
