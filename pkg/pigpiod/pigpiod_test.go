package pigpiod

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCommands(*testing.T) {

	piaddr := "192.168.15.191:8888"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := Connect(ctx, piaddr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	busInst, err := c.Read(23)
	fmt.Println(busInst, err)
	//d, err := c.ReadI2c(int(busInst), 0xF, 16)
	//ins := &inputs.Inputs{}
	//data := ins.DecodeData(d)
	//fmt.Println(data.UI1.Temp)
	//
	//fmt.Println(d, err)

}
