package geojs

import (
	"testing"

	"github.com/akwong189/cliweather/pkg/utils"
)

func TestGetIpAddress(t *testing.T) {
	addr := getPublicIPAddress()
	if !utils.ValidIPAddress(addr) {
		t.Errorf("address, " + addr + ", retrieved is not a valid IP address")
	}
}
