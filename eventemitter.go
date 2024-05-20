// Package goemit implements a simple Event Emitter pattern to be used for registering callbacks and
// emitting events.
// It is derived by a lightweight design interface for original jQuery implementations in JavaScript world.
package goemit

import "fmt"

// A EventEmitter holds information about registered callbacks and the event type they should be executed on.
type EventEmitter struct {
	events     map[string][]*func(...interface{}) interface{}
	onceEvents map[string][]*func(...interface{}) interface{}
}

// NewEventEmitter returns a pointer to a new initialized EventEmitter with an empty map of events.
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		events:     make(map[string][]*func(...interface{}) interface{}, 0),
		onceEvents: make(map[string][]*func(...interface{}) interface{}, 0),
	}
}

func (ee *EventEmitter) String() string {
	return fmt.Sprintf("registered events: %v", ee.events)
}

// On registers a new callback function to a given event type. When an event is emitted,
// the function receives all arguments that were parsed when emitting.
func (ee *EventEmitter) On(event string, callback *func(...interface{}) interface{}) *EventEmitter {
	ee.events[event] = append(ee.events[event], callback)
	return ee
}

// Emit emits a new event on a event type. All existing callbacks, including the ones registered using Once() are called,
// providing all arguments that were given when the event was emitted.
// After emitting the events, the callbacks registered for only-once execution are removed.
func (ee *EventEmitter) Emit(event string, args ...interface{}) {
	for _, cb := range ee.events[event] {
		(*cb)(args...)
	}
	for _, cb := range ee.onceEvents[event] {
		(*cb)(args...)
	}
	ee.onceEvents = make(map[string][]*func(...interface{}) interface{}, 0)
}

// Off removes a registered callback for a certain event type. A callback registered for only-once execution is also removed.
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

func (ee *EventEmitter) EventOff(event string) {
	ee.events[event] = make([]*func(...interface{}) interface{}, 0)
	ee.onceEvents[event] = make([]*func(...interface{}) interface{}, 0)
}

// AllOff removes all registered callbacks for all event types. This is equivalent to re-initializing the EventEmitter.
func (ee *EventEmitter) AllOff() {
	ee.events = make(map[string][]*func(...interface{}) interface{}, 0)
	ee.onceEvents = make(map[string][]*func(...interface{}) interface{}, 0)
}

// Once registers a callback to be executed only once when a event is emitted.
// After execution, the callback is automatically removed from the list of registered callbacks.
func (ee *EventEmitter) Once(event string, callback *func(...interface{}) interface{}) *EventEmitter {
	ee.onceEvents[event] = append(ee.onceEvents[event], callback)
	return ee
}
