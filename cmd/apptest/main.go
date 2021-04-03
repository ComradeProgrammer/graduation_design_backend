package main

import (
	"fmt"
	"graduation_design/internal/pkg/pool"
	"time"
)


func main(){
	var p=pool.NewPool(3,6)
	for i:=0;i<6;i++{
		var i2=i
		p.AddTask(func(){
			time.Sleep(1*time.Second)
			fmt.Println(i2)
		})
	}
	p.Run()
	p.Wait()

}