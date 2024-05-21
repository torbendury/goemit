package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/torbendury/goemit"
)

func main() {

	// Start vendor application
	stockNotifier := NewStockNotifier()

	// Register a customer
	c := Customer{
		itemNotifications: make([]string, 0),
	}

	// Customer hits notification button
	c.NotifyAboutRestock(stockNotifier, stockNotifier.items[0])

	// Listener starts waiting for new notifications
	var notification string
	notificationChannel := make(chan string)

	go func() {
		select {
		case notification = <-notificationChannel:
			fmt.Println("Update received!")
			fmt.Println(notification)
		// Dummy for example
		case <-time.After(time.Second * 60):
			fmt.Println("Did not receive stock notification about restock within 1 minute")
		}
	}()

	// Vendor notifies all customers which are interested in nails about new stock
	stockNotifier.notifier.Emit("Nail", notificationChannel, "Stock Update: 10")

	// Shutdown
	close(notificationChannel)
}

// --------------------------------------------------
// Mock: stock notification service
type StockNotifier struct {
	notifier *goemit.EventEmitter
	items    []string
}

func NewStockNotifier() *StockNotifier {
	return &StockNotifier{
		notifier: goemit.NewEventEmitter(),
		// Mocking purposes
		items: []string{"Nail", "Screw", "Hammer", "Drill"},
	}
}

// --------------------------------------------------
// Mock: customer that can register his/her interest in a product
type Customer struct {
	itemNotifications []string
}

func (c *Customer) NotifyAboutRestock(sn *StockNotifier, item string) {
	c.itemNotifications = append(c.itemNotifications, item)
	cb := func(input ...interface{}) interface{} {
		ch, ok := input[0].(chan string)
		if !ok {
			return errors.New("input does not contain event channel")
		}
		ch <- input[1].(string)
		return nil
	}
	sn.notifier.On(item, &cb)
}
