package networking

import (
	"fmt"
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

func GetLocalNetworkAddresses() ([]string, error) {
	baseAddress := "192.168.0."
	number := 2
	pinger := fastping.NewPinger()
	defer pinger.Stop()

	for {
		if number > 255 {
			break
		}
		pingAddress := fmt.Sprintf("%s%d", baseAddress, number)
		remoteAddress, err := net.ResolveIPAddr("ip4:icmp", pingAddress)
		if err != nil {
			continue
		}
		pinger.AddIPAddr(remoteAddress)
		number++
	}
	actualAddresses := make([]string, 0)
	pinger.OnRecv = func(i *net.IPAddr, d time.Duration) {
		actualAddresses = append(actualAddresses, i.IP.String())
	}
	err := pinger.Run()

	return actualAddresses, err
}
