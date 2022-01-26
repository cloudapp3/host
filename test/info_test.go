package test

import (
	"fmt"
	"github.com/cloudapp3/host/info"
	"testing"
)

func TestGetHostInfo(t *testing.T) {
	t.Logf("test get host info")
	i := info.GetLinuxInfo()
	if i.HostName == "" {
		t.Errorf("get host name fail")
	}
	fmt.Println(fmt.Sprintf("info %+v", i))
}
