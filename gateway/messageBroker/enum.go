package messagebroker

type IEvents interface{}
type Events struct {
	OrderReceived string
	OrderCreated  string
	OrderUpdated  string
}

func NewEventsEnum() Events {
	return Events{
		OrderReceived: "OrderReceived",
		OrderCreated:  "OrderCreated",
		OrderUpdated:  "OrderUpdated",
	}
}
