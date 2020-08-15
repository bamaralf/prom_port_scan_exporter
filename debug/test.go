package test

import (
	"fmt"
	"time"
	"github.com/anvie/port-scanner"
)

func portScan(pod string) []int{ 
	  ps := portscanner.NewPortScanner(pod, 2*time.Second, 5)
	  openedPorts := ps.GetOpenedPort(1, 65535)
	return openedPorts
}

main() {
	ps := portScan("10.7.63.167")
	fmt.Print(ps)
}