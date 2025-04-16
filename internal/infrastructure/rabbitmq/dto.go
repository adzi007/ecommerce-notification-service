package rabbitmq

type OrderMessage struct {
	OrderID uint64 `json:"order_id"`
	UserId  string `json:"user_id"`
	Status  string `json:"status"`
}
