package main

import (
	"flag"
	"fmt"
	"github.com/cloudapp3/host/info"
)

var token = flag.String("t", "", "host tg token")
func main()  {
   flag.Parse()
   if *token == ""{
	   panic("please enter -t token")
   }
   fmt.Println("host running")
   info.SetToken(*token)
   info.HostIP()
   info.StartActions()
}
