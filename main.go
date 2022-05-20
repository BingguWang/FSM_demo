package main

import (
	"fmt"
	"github.com/wbing441282413/FSM_demo/event"
	"github.com/wbing441282413/FSM_demo/fsm"
	"github.com/wbing441282413/FSM_demo/status"
)

func main() {
	// 假设一个状态和事件
	modelStatus := status.ModelStatusToDeliver
	e := event.DeliverOrder
	// 获取状态机
	f := fsm.NewFSM()
	newStatus := f.Call(modelStatus, e)
	fmt.Println("新状态是：", newStatus.FSMStatusDesc())
}
