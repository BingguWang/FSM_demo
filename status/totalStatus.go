package status

// 总状态枚举
const (
    TOTAL_STATUS_TO_PAY_FLAG     = 0
    TOTAL_STATUS_TO_DELIVER_FLAG = 1
    TOTAL_STATUS_TO_RECEIVE_FLAG = 2
    TOTAL_STATUS_TO_COMMENT_FLAG = 3
)

const (
    TOTAL_STATUS_TO_PAY_DESC     = "总状态：未支付"
    TOTAL_STATUS_TO_DELIVER_DESC = "总状态：待发货"
    TOTAL_STATUS_TO_RECEIVE_DESC = "总状态：待收货"
    TOTAL_STATUS_TO_COMMENT_DESC = "总状态：待评价"
)

// 总状态
var TOTAL_STATUS_MAP = map[int]string{
    TOTAL_STATUS_TO_PAY_FLAG:     TOTAL_STATUS_TO_PAY_DESC,
    TOTAL_STATUS_TO_DELIVER_FLAG: TOTAL_STATUS_TO_DELIVER_DESC,
    TOTAL_STATUS_TO_RECEIVE_FLAG: TOTAL_STATUS_TO_RECEIVE_DESC,
    TOTAL_STATUS_TO_COMMENT_FLAG: TOTAL_STATUS_TO_COMMENT_DESC,
}
