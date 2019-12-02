package hw1

import (
	"fmt"
	"github.com/beevik/ntp"
)

func CurrentTime() error {
	const server = "0.pool.ntp.org"
	time, err := ntp.Time(server)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Current time is: %s", time)
	return nil
}