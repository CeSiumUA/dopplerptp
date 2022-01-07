package utils

import (
	"testing"
)

func TestPublicIpGen(t *testing.T) {
	publicIps := GeneratePublicIPs(10)
	publicIPsCount := len(publicIps)
	t.Logf("Generated %d public IPs", publicIPsCount)
	if publicIPsCount != 10 {
		t.Errorf("got %d public ips, expected %d", publicIPsCount, 10)
	}
	t.Logf("%v", publicIps)
}
