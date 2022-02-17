package observer

const GLOBAL = "GLOBAL"

type Observable struct {
	eventHandlerGroups map[string][]*EventHandler
}

func CreateObservable() *Observable {
	return &Observable{
		eventHandlerGroups: make(map[string][]*EventHandler),
	}
}

func (o *Observable) Notify(event Event) {
	for _, eventHandler := range o.eventHandlerGroups[event.Action()] {
		eventHandler.callback(event)
	}
}

func (o *Observable) Subscribe(eventHandler *EventHandler) {
	eventHandlers, ok := o.eventHandlerGroups[eventHandler.Action]
	if !ok {
		o.eventHandlerGroups[eventHandler.Action] = []*EventHandler{eventHandler}
	} else {
		o.eventHandlerGroups[eventHandler.Action] = append(eventHandlers, eventHandler)
	}
}

func (o *Observable) Unsubscribe(eventHandler *EventHandler) {
	eventHandlers, ok := o.eventHandlerGroups[eventHandler.Action]

	if !ok {
		return
	}

	for i, v := range eventHandlers {
		if v == eventHandler {
			eventHandlers[i] = nil
			eventHandlers[i] = eventHandlers[len(eventHandlers)-1]
			eventHandlers = eventHandlers[:len(eventHandlers)-1]
		}
	}

	if len(eventHandlers) == 0 {
		delete(o.eventHandlerGroups, eventHandler.Action)
	}
}
