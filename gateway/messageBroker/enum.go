package messagebroker

type IEvents interface{}
type Events struct {
	OrderReceived string
	OrderCreated  string
	OrderUpdated  string
}

func NewEventsEnum() Events {
	return Events{
		OrderReceived: "order received",
		OrderCreated:  "order created",
		OrderUpdated:  "order updated",
	}
}
