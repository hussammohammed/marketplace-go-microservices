package messagebroker

type Events struct {
	OrderReceived string
	OrderCreated  string
	OrderUpdated  string
}

type Topics struct {
	OrderEvents    string
	UserEvents     string
	ShipmentEvents string
}

func NewEventsEnum() Events {
	return Events{
		OrderReceived: "OrderReceived",
		OrderCreated:  "OrderCreated",
		OrderUpdated:  "OrderUpdated",
	}
}

func NewTopicsEnum() Topics {
	return Topics{
		OrderEvents:    "OrderEvents",
		UserEvents:     "UserEvents",
		ShipmentEvents: "ShipmentEvents",
	}
}
