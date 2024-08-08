package messagebroker

type Event struct {
	Topic string
	Key   string
	Text  string
}
