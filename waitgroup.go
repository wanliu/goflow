package flow

import "sync"

type waitGroupByName struct {
	sync.WaitGroup
	Name    string
	Counter int
}

func NewWaitGroup(name string) *waitGroupByName {
	return &waitGroupByName{Name: name}
}

func (wg *waitGroupByName) Add(delta int) {
	wg.Counter += delta
	wg.WaitGroup.Add(delta)
}

func (wg *waitGroupByName) Done() {
	wg.Counter += -1
	wg.WaitGroup.Done()
}

type WaitBus struct {
	sync.RWMutex
	List []*waitGroupByName
}

func NewWaitBus() *WaitBus {
	return &WaitBus{
		List: make([]*waitGroupByName, 0),
	}
}

func (ws *WaitBus) New(name string) *waitGroupByName {
	wg := NewWaitGroup(name)
	ws.List = append(ws.List, wg)
	return wg
}

var WaitList *WaitBus

func init() {
	WaitList = NewWaitBus()
}
