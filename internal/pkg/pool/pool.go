package pool

import (
	"graduation_design/internal/pkg/logs"
	"time"
)

type Task func()
type Pool struct{
	taskNum int
	allFinishChan chan struct{}
	finishChan chan string
	taskChan chan Task
	workers []*Worker
}

func NewPool(workerNum int,taskNum int)(*Pool){
	logs.Info("Pool is created,workerNum %d,taskBufferNum %d",workerNum,taskNum)
	var ret =Pool{
		taskNum :taskNum,
		taskChan:make(chan Task,taskNum),
		finishChan: make(chan string,taskNum),
		workers:make([]*Worker,workerNum) ,
		allFinishChan: make(chan struct{}),
	}
	for i:=0;i<workerNum;i++{
		ret.workers[i]=NewWorker(ret.taskChan,ret.finishChan,i)
	}
	return &ret
}

func (p*Pool)AddTask(task Task){
	p.taskChan<-task
}



func (p *Pool)Run(){
	go p.monitor()
	for _,w:=range p.workers{
		w.Run()
	}
}

func (p *Pool)WaitWithTimeOut(t time.Duration){
	select{
	case <-p.allFinishChan:
	case <-time.After(t):
	}
	
	p.destroy()
}

func (p *Pool)Wait(){

	_=<-p.allFinishChan
	p.destroy()
}

func (p*Pool)destroy(){
	for _,w:=range p.workers{
		w.Abort()
	}
}

func (p* Pool)monitor(){
	for i:=0;i<p.taskNum;i++{
		logs.Info("%s",<-p.finishChan)
	}
	p.allFinishChan<-struct{}{}
}


