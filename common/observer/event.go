package observer

type Event interface {
	Action() string
}

type Action string

func (a Action) Action() string {
	return string(a)
}
