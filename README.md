# goemit

A [EventEmitter](https://gist.github.com/abravalheri/d137cf14652eb932f398cdffe06fc7c2) implementation in Golang.

## Problem Solved

This is a slightly adjusted modularized implementation of an event emitter, which helps users implement the observer pattern. It relies on a subject (the emitter) maintaining a list of its dependents (callbacks) which notifies them automatically (event emitting) of any state changes.

The pattern is commonly used for implementing distributed event-handling systems.

## Interface

Events (respective, their "type") are stored as string values.

`goemit` offers a small amount of methods for users. More exhaustive documentation is found in the godocs.

Here's the short hand information:

- `NewEventEmitter()` returns an `EventEmitter` which maintains a list of events and their callbacks.
- `On(event, callback)` registers a new callback function for an event type.
- `Emit(event, args...)` lets the subject emit events to its dependants and calls the registered callback functions with `args`.
- `Off(event, callback)` removes a given registration for a event. It is looked up by the stored callback function pointer.
- `AllOff()` removes all registered callbacks for all event types.
- `EventOff(event)` removes all registered callbacks for a given event type.
- `Once(event, callback)` registers a callback to be executed only once as soon as the emitter emits an event for it. The callback is removed after execution.

`On()` as well as `Once()` return the EventEmitter itself, enabling chaining methods.

## License

See [LICENSE](LICENSE).

## Usage Example

See [example/](example/main.go).

## Building Blocks

![A PlantUML graphic showing the building blocks of goemit](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/torbendury/goemit/main/eventemitter.puml)
