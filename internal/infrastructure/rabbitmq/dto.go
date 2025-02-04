package rabbitmq

type OrderMessage struct {
	OrderID uint64 `json:"order_id"`
	Status  string `json:"status"`
}
