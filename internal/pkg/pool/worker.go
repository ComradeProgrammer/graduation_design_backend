package pool

import (
	"fmt"
	"graduation_design/internal/pkg/logs"
)
type Worker struct{
	taskChan chan Task
	finishChan chan string
	abortChan chan struct{}
	//Running bool
	ID int
}

func NewWorker(taskChan chan Task,finishChan chan string,id int)(*Worker){
	var res=Worker{
		taskChan: taskChan,
		finishChan: finishChan,
		abortChan:make(chan struct{},1),
		//Running: false,
		ID:id,
	}
	return &res
}

func (w*Worker)Abort(){
	logs.Info("worker %d, abort is called",w.ID)
	w.abortChan<-struct{}{}
}

func (w *Worker)Run(){
	go w.run()
}

func (w *Worker)run(){
	//w.Running=true
	//defer func(){w.Running=false}()
	defer logs.Info("worker %d quited",w.ID)
	for{
		select{
		case <-w.abortChan:
			break
		case f:=<-w.taskChan:
			logs.Info("worker %d fetch task",w.ID)
			f()
			w.finishChan<-fmt.Sprintf("worker %d finished task",w.ID)
		}
	}
}