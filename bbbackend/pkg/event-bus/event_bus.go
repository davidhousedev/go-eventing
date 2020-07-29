package eventbus

import (
	"encoding/json"
	"io"
)

// Dispatcher sends an event to a specific topic namespace
type Dispatcher interface {
	Send(string, Event) error
}

// Receiver reads data from an IO source
type Receiver interface {
	Receive(io.Reader) error
}

// Bus receives events at namespaces and makes them available in topics
type Bus struct {
	topics []*Topic
}

// Client sends events to an EventBus
type Client struct {
	destination io.Writer
}

// Topic collects multiple similar events and publishes them to subscribers
type Topic struct {
	name string
}

// Event contrains the data for a particular event
type Event struct {
	Name    string
	Payload map[string]string
}

func (client *Client) Send(namespace string, event *Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	client.destination.Write(data)
	return nil
}

func (bus *Bus) Subscribe(name string, source *Topic) error {
	bus.topics = append(bus.topics, source)

	return nil
}

func (bus *Bus) get_subscriptions() ([]*Topic, error) {
	return bus.topics, nil
}
