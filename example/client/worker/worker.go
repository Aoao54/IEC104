package worker

import (
	"github.com/9d77v/iec104"
)

//Task 数据处理任务
func Task(data *iec104.APDU) (error){
	//TODO 自定义数据处理
	println("do task")
	return nil
}

func Datahandler (data *iec104.APDU) error  {
	println("do task")
	return nil
}