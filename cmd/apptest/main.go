package main

import (
	"fmt"
	"graduation_design/internal/app/model"
)

// func main() {
// 	var p = pool.NewPool(3, 6)
// 	for i := 0; i < 6; i++ {
// 		var i2 = i
// 		p.AddTask(func() {
// 			time.Sleep(1 * time.Second)
// 			fmt.Println(i2)
// 		})
// 	}
// 	p.Run()
// 	p.Wait(10 * time.Second)

// }

func main() {
	var token = "2852b707ac3a644f68688d19ce84cc406a20be0d144a0e0cb05e509b55987e78"
	res, _ := model.GetAllIssueNotes(token, 3)
	fmt.Println(len(res))
}
