package status

type ModelStatus struct {
	TotalStatus   Status
	OrderStatus   Status
	DeliverStatus Status
}

type Status struct {
	StatusID   int
	StatusDesc string
}

// 状态枚举
var (
	// 未支付
	ModelStatusToPay = ModelStatus{
		TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_PAY_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_PAY_FLAG]},
		OrderStatus:   Status{StatusID: ORDER_STATUS_TO_PAY_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_PAY_FLAG]},
		DeliverStatus: Status{StatusID: DELIVER_STATUS_DONT_DELIVER_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_DONT_DELIVER_FLAG]},
	}

	// 待发货
	ModelStatusToDeliver = ModelStatus{
		TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_DELIVER_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_DELIVER_FLAG]},
		OrderStatus:   Status{StatusID: ORDER_STATUS_TO_DELIVER_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_DELIVER_FLAG]},
		DeliverStatus: Status{StatusID: DELIVER_STATUS_TO_DELIVER_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_TO_DELIVER_FLAG]},
	}

	// 已发货
	ModelStatusDelivered = ModelStatus{
		TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_RECEIVE_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_RECEIVE_FLAG]},
		OrderStatus:   Status{StatusID: ORDER_STATUS_TO_RECEIVE_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_RECEIVE_FLAG]},
		DeliverStatus: Status{StatusID: DELIVER_STATUS_DELIVERED_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_DELIVERED_FLAG]},
	}

	// 已签收
	ModelStatusReceivered = ModelStatus{
		TotalStatus:   Status{StatusID: TOTAL_STATUS_TO_COMMENT_FLAG, StatusDesc: TOTAL_STATUS_MAP[TOTAL_STATUS_TO_COMMENT_FLAG]},
		OrderStatus:   Status{StatusID: ORDER_STATUS_TO_COMMENT_FLAG, StatusDesc: ORDER_STATUS_MAP[ORDER_STATUS_TO_COMMENT_FLAG]},
		DeliverStatus: Status{StatusID: DELIVER_STATUS_RECEIVED_FLAG, StatusDesc: DELIVER_STATUS_DELIVER_MAP[DELIVER_STATUS_RECEIVED_FLAG]},
	}
)


// 因为会用到FSM接口的方法，所以要实现FSMStatus接口才行
func (e ModelStatus) FSMStatusID() int {
	return e.TotalStatus.StatusID
}

func (e ModelStatus) FSMStatusDesc() string {
	return e.TotalStatus.StatusDesc
}