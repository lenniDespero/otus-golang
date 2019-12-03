package hw1

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const server = "0.pool.ntp.org"

func CurrentTime() {
	time, err := ntp.Time(server)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time is: %s\n", time.Local())
}
