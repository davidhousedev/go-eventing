package eventbus

import "testing"

func TestEventSource(t *testing.T) {
	t.Run("EventSource can send events", func(t *testing.T) {
		source := Topic{name: "foo"}
		event := Event{Name: "event", Payload: map[string]string{"Foo": "Bar"}}

		err := source.Send(event)

		if err != nil {
			t.Fatal("should have been able to send event:", event)
		}

	})
}
