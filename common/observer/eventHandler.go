package observer

type EventHandler struct {
	Action   string
	callback func(event Event)
}
