package info

import (
	"bytes"
	"fmt"
	"github.com/cloudapp3/host/model"
	"github.com/smallnest/rpcx/client"
	"github.com/vmihailenco/msgpack/v5"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	url = "https://host.bestcheapvps.org/api/host/info"
)

var (
	c     *client.Client
	token string
)

func SetToken(t string) {
	token = t
}

func StartActions() {
	//connectRPCClient()
	i := GetLinuxInfo()
	reportHostInfo(&i)
	t := time.NewTicker(time.Minute * 1)
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
	httpClient := &http.Client{}
	b, _ := msgpack.Marshal(i)
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("report http new request error " + err.Error())
		return
	}
	req.Header.Set("token", token)
	req.Header.Set("ip", getClientIP())
	rsp, rErr := httpClient.Do(req)
	if rErr != nil {
		fmt.Println("report request server error " + rErr.Error())
		return
	}
	body, _ := ioutil.ReadAll(rsp.Body)
	_ = rsp.Body.Close()
	s := string(body)
	if s != "ok" {
		fmt.Println("report server info not success " + s)
	}

}

//func reportHostInfo(i *model.HostInfo) {
//	var reply string
//	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"ip": clientIP,
//		"token": token})
//	err := c.Call(ctx, "host", "reportHostInfo", i, &reply)
//	if err != nil {
//		connectRPCClient()
//		fmt.Println("report host info client error " + err.Error())
//	}
//	if reply != "ok" {
//		fmt.Println("report host info reply not success " + reply)
//	}else{
//		fmt.Println("ok")
//	}
//}
