package main

import (
	"fmt"
	"rest_servers/dao"
)

func main() {
	ts := dao.TaskStore{}
	dao.InitConn()
	task, _ := ts.GetTask(2)
	fmt.Println(task)
}
