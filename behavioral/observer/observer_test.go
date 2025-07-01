package observer

import "testing"

type MockEvent struct {
	name string
	data any
}

type MockObserver struct {
	ReceivedEvents []MockEvent
}

func (m *MockObserver) OnEvent(event string, payload any) {
	m.ReceivedEvents = append(m.ReceivedEvents, MockEvent{
		name: event,
		data: payload,
	})
}

func Test_EventManager_Publish(t *testing.T) {
	tests := []struct {
		name    string
		event   string
		payload any
	}{
		{
			name:    "Send an event",
			event:   "product:in_stock",
			payload: "Keyboard",
		},
		{
			name:    "Send another event",
			event:   "product:in_stock",
			payload: "Gaming Console",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mo := &MockObserver{}

			em := NewEventManager()
			em.Subscribe(tt.event, mo)
			em.Publish(tt.event, tt.payload)

			if len(mo.ReceivedEvents) != 1 {
				t.Fatalf("Expected 1 event, got %d", len(mo.ReceivedEvents))
			}

			got := mo.ReceivedEvents[0]
			if got.name != tt.event {
				t.Errorf("Expected event %q, got %q", tt.event, got.name)
			}

			if got.data != tt.payload {
				t.Errorf("Expected payload %q, got %q", tt.payload, got.data)
			}
		})
	}
}
