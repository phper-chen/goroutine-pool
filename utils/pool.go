package utils

import (
	"fmt"
	"sync"
)

//定义任务Task类型,每一个任务Task都可以抽象成一个函数
type Task struct {
	f func() error
}

//通过NewTask来创建一个Task
func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}

	return &t
}

//执行Task任务的方法
func (t *Task) Execute() {
	t.f() //调用任务所绑定的函数
}

//定义一个可控的池类型
type Pool struct {
	sync.WaitGroup
	//协程池最大worker数量,即并发数,限定Goroutine的个数
	workerNum int

	//协程池内部的任务就绪队列
	JobsChannel chan *Task
}

//协程池创建一个worker并且开始工作
func (p *Pool) worker(workID int) {
	defer p.Done()
	//worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChannel {
		//如果拿到任务,则执行task任务
		task.Execute()
		fmt.Println("worker ID ", workID, " 抓取任务完成")
	}
}

//让协程池Pool开始工作
func (p *Pool) Run() {
	//根据协程池的worker数量限定,开启固定数量的Worker,
	//每一个Worker用一个Goroutine承载
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
	//执行完毕需要关闭JobsChannel
	close(p.JobsChannel)
	p.Wait()

}

//创建一个新的可控的协程池
func NewPool(total, cap int) *Pool {
	p := Pool{
		workerNum:   cap,
		JobsChannel: make(chan *Task, total),
	}
	p.Add(cap)
	return &p
}
