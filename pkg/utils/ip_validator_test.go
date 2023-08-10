package utils

import "testing"

func TestValidIPAddr(t *testing.T) {
	ip_addr := "255.255.255.255"
	if ValidIPAddress(ip_addr) == false {
		t.Errorf(ip_addr + " failed validation")
	}
	ip_addr = "192.168.0.1"
	if ValidIPAddress(ip_addr) == false {
		t.Errorf(ip_addr + " failed validation")
	}
	ip_addr = "10.10.10.10"
	if ValidIPAddress(ip_addr) == false {
		t.Errorf(ip_addr + " failed validation")
	}
	ip_addr = "0.0.0.0"
	if ValidIPAddress(ip_addr) == false {
		t.Errorf(ip_addr + " failed validation")
	}
}

func TestInvalidIPAddr(t *testing.T) {
	ip_addr := "255.255.255"
	if ValidIPAddress(ip_addr) != false {
		t.Errorf(ip_addr + " passed validation when it should've failed")
	}
	ip_addr = "1924.168.0.1"
	if ValidIPAddress(ip_addr) != false {
		t.Errorf(ip_addr + " passed validation when it should've failed")
	}
	ip_addr = "10.10"
	if ValidIPAddress(ip_addr) != false {
		t.Errorf(ip_addr + " passed validation when it should've failed")
	}
	ip_addr = "1"
	if ValidIPAddress(ip_addr) != false {
		t.Errorf(ip_addr + " passed validation when it should've failed")
	}
	ip_addr = "...."
	if ValidIPAddress(ip_addr) != false {
		t.Errorf(ip_addr + " passed validation when it should've failed")
	}
}
