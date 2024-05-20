package goemit

import "fmt"

type EventEmitter struct {
	events     map[string][]*func(...interface{}) interface{}
	onceEvents map[string][]*func(...interface{}) interface{}
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		events:     make(map[string][]*func(...interface{}) interface{}, 0),
		onceEvents: make(map[string][]*func(...interface{}) interface{}, 0),
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
	for _, cb := range ee.onceEvents[event] {
		(*cb)(args...)
	}
	ee.onceEvents = make(map[string][]*func(...interface{}) interface{}, 0)
}

func (ee *EventEmitter) Off(event string, callback *func(...interface{}) interface{}) bool {
	eventsDone := false
	onceEventsDone := false
	for idx, cb := range ee.events[event] {
		if cb == callback {
			ee.events[event] = append(ee.events[event][:idx], ee.events[event][idx+1:]...)
			eventsDone = true
		}
	}
	for idx, cb := range ee.onceEvents[event] {
		if cb == callback {
			ee.onceEvents[event] = append(ee.onceEvents[event][:idx], ee.onceEvents[event][idx+1:]...)
			onceEventsDone = true
		}
	}
	return eventsDone || onceEventsDone
}

func (ee *EventEmitter) AllOff() bool {
	ee.events = make(map[string][]*func(...interface{}) interface{}, 0)
	ee.onceEvents = make(map[string][]*func(...interface{}) interface{}, 0)
	return true
}

func (ee *EventEmitter) Once(event string, callback *func(...interface{}) interface{}) bool {
	ee.onceEvents[event] = append(ee.onceEvents[event], callback)
	return true
}
