package observer

type EventHandler struct {
	Action   string
	Callback func(event Event)
}
