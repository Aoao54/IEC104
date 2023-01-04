package main

import (
	"fmt"
	"github.com/9d77v/iec104"
	_ "github.com/9d77v/iec104/example/client/worker"
	"github.com/sirupsen/logrus"
	"time"
)

type myClient struct{}
func main() {
	address := "127.0.0.1:6666"
	//subAddress := ""
	//if config.SubServerHost != "" && config.SubServerPort != 0 {
	//	subAddress = fmt.Sprintf("%s:%d", config.SubServerHost, config.ServerPort)
	//}
	mycli := &myClient{}
	logger := logrus.New()
	//logger.SetLevel(logrus.DebugLevel)
	//logger.Hooks.Add(utils.NewContextHook())
	client := iec104.NewClient( mycli ,address, logger)
	client.Run()
	client.SendTotalCall()
	time.Sleep(2*time.Second)

}
func (c *myClient) Datahandler (data *iec104.APDU) error  {
	if data.Signals[0].TypeID == 21{
		i := int64(data.Signals[0].Ts)
		tm := time.Unix(i, 0)
		fmt.Println(tm)
	}
	println("do task")
	return nil
}
