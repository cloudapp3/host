package info

import (
	"context"
	"fmt"
	"github.com/cloudapp3/host/model"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"time"
)

const (
	url = "host.bestcheapvps.org:443"
)

var (
	c     *client.Client
	token string
)

func SetToken(t string) {
	token = t
}

func StartActions() {
	connectRPCClient()
	i := GetLinuxInfo()
	reportHostInfo(&i)
	t := time.NewTimer(time.Minute * 1)
	for _ = range t.C {
		i := GetLinuxInfo()
		reportHostInfo(&i)
	}
}

func connectRPCClient() {
	c = client.NewClient(client.DefaultOption)
	err := c.Connect("tcp", url)
	if err != nil {
		fmt.Println("connect rpc client fail" + err.Error())
		return
	}
}
func reportHostInfo(i *model.HostInfo) {
	var reply string
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"ip": clientIP,
		"token": token})
	err := c.Call(ctx, "host", "reportHostInfo", i, &reply)
	if err != nil {
		connectRPCClient()
		fmt.Println("report host info client error " + err.Error())
	}
	if reply != "ok" {
		fmt.Println("report host info reply not success " + reply)
	}else{
		//fmt.Println("ok")
	}
}
