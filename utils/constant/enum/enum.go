package enum

const (
	Admin = "admin"
	User  = "user"
)

type OrderStatus int

func (os OrderStatus) Int() int {
	return int(os)
}

func (os OrderStatus) String() string {
	switch os {
	case DeliveringOrder:
		return "Delivering Order"
	case OrderConfirmed:
		return "Order Confirmed"
	case OrderCompleted:
		return "Order Completed"
	}

	return ""
}

var (
	OrderConfirmed  OrderStatus = 1
	DeliveringOrder OrderStatus = 2
	OrderCompleted  OrderStatus = 3
)

type BalanceHistoryType int

func (b BalanceHistoryType) Int() int {
	return int(b)
}

func (b BalanceHistoryType) String() string {
	switch b {
	case TopUp:
		return "Top Up Balance"
	case Payment:
		return "Payment"
	}

	return ""
}

var (
	Payment BalanceHistoryType = 1
	TopUp   BalanceHistoryType = 2
)

type MidtransStatus string

func (m MidtransStatus) ToString() string {
	return string(m)
}

var (
	Settlement MidtransStatus = "settlement"
)

type BalanceHistoryStatus int

var (
	Pending BalanceHistoryStatus = 1
	Success BalanceHistoryStatus = 2
	Failed  BalanceHistoryStatus = 3
)

func (b BalanceHistoryStatus) String() string {
	switch b {
	case Pending:
		return "Pending"
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	}

	return ""
}

func (b BalanceHistoryStatus) Int() int {
	return int(b)
}
