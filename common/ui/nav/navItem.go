package nav

import "arvis/rtw/common/ui"

const nav_event = "nav_event"

type Nav struct {
	Name    string
	Creator func() ui.Fragment
}

func NewNav(name string, creator func() ui.Fragment) *Nav {
	return &Nav{name, creator}
}

func (n *Nav) Action() string {
	return nav_event
}
