package net_listener

import (
	"net"
	"testing"
)

func Test_Net_LookupHost(t *testing.T) {
	host := "www.baidu.com"
	addrs, err := net.LookupHost(host)
	if err != nil {
		t.Fatal(err)
	}
	for _, addr := range addrs {
		t.Logf("%s \n", addr)
	}
}

func Test_Net_LookupPort(t *testing.T) {
	service := "ntp" // Network timing protocol
	// KEEP IN MIND : only transport protocol has port
	port, err := net.LookupPort("tcp", service)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s port is : %d\n", service, port)
}
