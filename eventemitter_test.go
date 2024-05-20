package goemit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewEventEmitter(t *testing.T) {
	ee := NewEventEmitter()

	got := ee.String()
	want := "registered events: map[]"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestOn(t *testing.T) {
	ee := NewEventEmitter()

	cb := func(input ...interface{}) interface{} {

		return fmt.Sprintf("Event emitted, input registered: %s", input)
	}

	ok := ee.On("test", &cb)
	if !ok {
		t.Errorf("could not register callback %T", cb)
	}
}

func TestEmit(t *testing.T) {
	ee := NewEventEmitter()

	want := "Event emitted"

	ch := make(chan string)

	cb := func(input ...interface{}) interface{} {
		ch, ok := input[0].(chan string)
		if !ok {
			t.Error("first argument is not a string channel")
		}
		ch <- "Event emitted"
		return nil
	}

	ee.On("test", &cb)

	var got string

	done := make(chan bool)

	go func() {
		select {
		case got = <-ch:
			done <- true
		case <-time.After(time.Second * 5):
			done <- false
		}
	}()

	ee.Emit("test", ch)

	close(ch)

	if ok := <-done; !ok {
		t.Fatal("timed out")
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestOffOne(t *testing.T) {
	ee := NewEventEmitter()

	cb := func(i ...interface{}) interface{} {
		fmt.Println(i)
		return nil
	}

	cb2 := func(i ...interface{}) interface{} {
		fmt.Printf("Event %v\n", i)
		return nil
	}

	ee.On("test", &cb)

	ok := ee.Off("test", &cb)

	if !ok {
		t.Error("could not remove callback")
	}

	ok = ee.Off("test", &cb2)

	if ok {
		t.Error("non-existing callback should not be removable")
	}
}

func TestOffAll(t *testing.T) {
	ee := NewEventEmitter()

	cb := func(...interface{}) interface{} {
		return nil
	}

	ee.On("test1", &cb)
	ee.On("test2", &cb)
	ee.On("test3", &cb)

	got := len(ee.events)

	want := 3

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	ee.AllOff()

	got = len(ee.events)

	want = 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
