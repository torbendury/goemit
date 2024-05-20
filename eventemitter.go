package goemit

import "fmt"

type EventEmitter struct {
	events map[string][]*func(...interface{}) interface{}
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		events: make(map[string][]*func(...interface{}) interface{}, 0),
	}
}

func (ee *EventEmitter) String() string {
	return fmt.Sprintf("registered events: %v", ee.events)
}

func (ee *EventEmitter) On(event string, callback *func(...interface{}) interface{}) bool {
	ee.events[event] = append(ee.events[event], callback)
	return true
}

func (ee *EventEmitter) Emit(event string, args ...interface{}) {
	for _, cb := range ee.events[event] {
		(*cb)(args...)
	}
}

func (ee *EventEmitter) Off(event string, callback *func(...interface{}) interface{}) bool {
	for idx, cb := range ee.events[event] {
		if cb == callback {
			ee.events[event] = append(ee.events[event][:idx], ee.events[event][idx+1:]...)
			return true
		}
	}
	return false
}
