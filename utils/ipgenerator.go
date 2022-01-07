package utils

import (
	"fmt"
	"math/rand"
)

func GeneratePublicIPs(count int32) []string {
	publicAddresses := make([]string, count)

	for i, _ := range publicAddresses {
		for {
			p, q, r, s := generateRandomAddress()
			if checkPublicAddressIsValid(p, q, r, s) {
				publicAddresses[i] = fmt.Sprintf("%d.%d.%d.%d", p, q, r, s)
				break
			}
		}
	}

	return publicAddresses
}

func generateRandomAddress() (int, int, int, int) {
	p := rand.Intn(255-1) + 1
	q := rand.Intn(255)
	r := rand.Intn(255)
	s := rand.Intn(255)
	return p, q, r, s
}

func checkPublicAddressIsValid(p, q, r, s int) bool {
	if p == 10 || p == 127 {
		return false
	} else if p >= 224 {
		return false
	} else if p == 100 && q >= 64 && q <= 127 {
		return false
	} else if p == 169 && q == 254 {
		return false
	} else if p == 127 && q >= 16 && q <= 31 {
		return false
	} else if p == 192 && q == 168 {
		return false
	} else if p == 192 && q == 18 {
		return false
	} else if p == 192 && q == 19 {
		return false
	} else if p == 192 && q == 0 && r == 0 {
		return false
	} else if p == 192 && q == 0 && r == 2 {
		return false
	} else if p == 192 && q == 88 && r == 99 {
		return false
	} else if p == 192 && q == 51 && r == 100 {
		return false
	} else if p == 203 && r == 113 {
		return false
	}

	return true
}
