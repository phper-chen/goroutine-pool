package main

import (
	"fmt"
	handlers "goroutine-pool/handlers"
	utils "goroutine-pool/utils"
	"os"
	"time"
)

func main() {
	totalRequestNum, concurrentNum, err := handlers.CheckAndGetCmdParams(os.Args)
	if err != nil {
		fmt.Println(err)
	}
	//定义抓取任务
	t := utils.NewTask(handlers.Crawl)

	//创建一个协程池,最大开启concurrentNum个协程worker
	p := utils.NewPool(totalRequestNum, concurrentNum)

	//创建任务通道
	for i := 0; i < totalRequestNum; i++ {
		p.JobsChannel <- t
	}
	start := time.Now()

	//启动
	p.Run()
	end := time.Now()

	fmt.Println(end.Sub(start).Seconds())

}
