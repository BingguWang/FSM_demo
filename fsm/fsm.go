package fsm

import (
    "github.com/wbing441282413/FSM_demo/event"
    "github.com/wbing441282413/FSM_demo/status"
    "log"
    "sync"
)

type FSMStatus interface {
    FSMStatusID() int
    FSMStatusDesc() string
}

type FSMEvent interface {
    FSMEventID() int
    FSMEventDesc() string
}

type FSMHandler func(status FSMStatus, event FSMEvent) FSMStatus // 传入状态和事件返回一个新的事件

type FSM interface {
    AddHandler(status FSMStatus, event FSMEvent, handler FSMHandler) FSM
    Call(status FSMStatus, event FSMEvent) FSMStatus
}

type fsm struct {
    mu       sync.Mutex                 // 排它锁
    State    FSMStatus                  // 当前的状态(调用后)
    Handlers map[int]map[int]FSMHandler // 状态-事件-处理器map
}

func (f *fsm) AddHandler(status FSMStatus, event FSMEvent, handler FSMHandler) FSM {
    if _, ok := f.Handlers[status.FSMStatusID()]; !ok {
        f.Handlers[status.FSMStatusID()] = make(map[int]FSMHandler)
    }
    if _, ok := f.Handlers[status.FSMStatusID()][event.FSMEventID()]; !ok {
        f.Handlers[status.FSMStatusID()][event.FSMEventID()] = handler
    }
    return f
}

func (f *fsm) Call(status FSMStatus, event FSMEvent) FSMStatus {
    if _, ok := f.Handlers[status.FSMStatusID()]; !ok {
        log.Fatal("没有状态的映射关系")
    }
    if _, ok := f.Handlers[status.FSMStatusID()][event.FSMEventID()]; !ok {
        log.Fatal("没有事件的映射关系")
    }
    handler := f.Handlers[status.FSMStatusID()][event.FSMEventID()]
    f.State = handler(status, event)
    return f.State
}

// 定义事件，定制handler
func NewFSM() FSM {
    f := &fsm{
        State:    nil,
        Handlers: make(map[int]map[int]FSMHandler),
    }
    f.AddHandler(
        status.ModelStatusToPay,
        event.PayOrder,
        func(s FSMStatus, e FSMEvent) FSMStatus {
            return status.ModelStatusToDeliver
        },
    ).AddHandler(
        status.ModelStatusToDeliver,
        event.DeliverOrder,
        func(s FSMStatus, e FSMEvent) FSMStatus {
            return status.ModelStatusDelivered
        },
    ).AddHandler(
        status.ModelStatusDelivered,
        event.ReceiveOrder,
        func(s FSMStatus, e FSMEvent) FSMStatus {
            return status.ModelStatusReceivered
        },
    )
    return f
}
