package observer

import "testing"

const action1 = "A1"
const action2 = "A2"

type event struct {
	action string
}

func (e event) Action() string {
	return e.action
}

func TestNotifyBasic(t *testing.T) {
	shouldChange := false
	shouldNotChange := false

	observable := CreateObservable()
	observable.Subscribe(&EventHandler{action1, func(event Event) { shouldChange = true }})
	observable.Subscribe(&EventHandler{action2, func(event Event) { shouldNotChange = true }})

	observable.Notify(event{action1})

	if shouldChange != true || shouldNotChange == true {
		t.Error("notify")
	}
}

func TestUnsubscribeBasic(t *testing.T) {
	shouldChange := false
	shouldNotChange := false

	observable := CreateObservable()
	eventHandler1 := &EventHandler{action1, func(event Event) { shouldChange = true }}
	observable.Subscribe(eventHandler1)
	observable.Subscribe(&EventHandler{action2, func(event Event) { shouldNotChange = true }})

	observable.Unsubscribe(eventHandler1)
	observable.Notify(event{action1})

	if shouldChange == true || shouldNotChange == true {
		t.Error("unsubscribe")
	}
}

func TestUnsubscribeMultiple(t *testing.T) {
	shouldChange := false
	shouldNotChange := false

	observable := CreateObservable()
	eventHandler1 := &EventHandler{action1, func(event Event) { shouldNotChange = true }}
	eventHandler2 := &EventHandler{action1, func(event Event) { shouldChange = true }}
	observable.Subscribe(eventHandler1)
	observable.Subscribe(eventHandler2)

	observable.Unsubscribe(eventHandler1)
	observable.Notify(event{action1})

	if shouldChange != true || shouldNotChange == true {
		t.Error("unsubscribe multiple")
	}
}
