package event

type Event struct {
	EventID   int
	EventDesc string
}

// 事件枚举
const (
	EVENT_PAY_ORDER       = 0
	EVENT_DELIVER_ORDER   = 1
	EVENT_RECEIVEED_ORDER = 2
)

const (
	EVENT_PAY_ORDER_DESC       = "事件：支付订单"
	EVENT_DELIVER_ORDER_DESC   = "事件：订单发货"
	EVENT_RECEIVEED_ORDER_DESC = "事件：订单收货"
)

var ModelEventMap = map[int]string{
	EVENT_PAY_ORDER:       EVENT_PAY_ORDER_DESC,
	EVENT_DELIVER_ORDER:   EVENT_DELIVER_ORDER_DESC,
	EVENT_RECEIVEED_ORDER: EVENT_RECEIVEED_ORDER_DESC,
}

// 事件列举
var (
	PayOrder = Event{
		EVENT_PAY_ORDER,
		ModelEventMap[EVENT_PAY_ORDER],
	}

	DeliverOrder = Event{
		EVENT_DELIVER_ORDER,
		ModelEventMap[EVENT_DELIVER_ORDER],
	}

	ReceiveOrder = Event{
		EVENT_RECEIVEED_ORDER,
		ModelEventMap[EVENT_RECEIVEED_ORDER],
	}
)

// 因为会用到FSM接口的方法，所以要实现FSMEvent接口才行
func (e Event) FSMEventID() int {
	return e.EventID
}

func (e Event) FSMEventDesc() string {
	return e.EventDesc
}
