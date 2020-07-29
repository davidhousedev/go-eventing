package eventbus

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBus(t *testing.T) {
	t.Run("Bus can subscribe to event publishers", func(t *testing.T) {
		bus := Bus{topics: []*Topic{}}
		topic := Topic{name: "foo"}
		err := bus.Subscribe("foo", &topic)

		if err != nil {
			t.Fatal("should have allowed subscription:", err)
		}

		want := []*Topic{&topic}
		got, err := bus.get_subscriptions()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted subscription list %v but got %v", want, got)
		}
	})

	t.Run("Bus can receive events at specific topics", func(t *testing.T) {
		bus := Bus{topics: []*Topic{}}
		topic := Topic{name: "foo"}

		err := bus.Subscribe("foo", &topic)

		if err != nil {
			t.Fatal("should have allowed subscription:", err)
		}

	})
}

func TestClient(t *testing.T) {
	t.Run("Client can send events", func(t *testing.T) {
		dest := bytes.Buffer{}
		client := Client{destination: &dest}
		event := Event{Name: "event-sent", Payload: make(map[string]string)}
		err := client.Send("test-event", &event)

		if err != nil {
			t.Fatal("should have sent event:", event)
		}

		want := "{\"Name\":\"event-sent\",\"Payload\":{}}"
		got := dest.String()
		if want != got {
			t.Errorf("wanted event to be sent as %v but got %v", want, got)
		}
	})
}
