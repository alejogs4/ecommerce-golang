package aggregate

import (
	"log"
	"strings"

	"github.com/alejogs4/hn-website/src/shared/domain/domainevent"
)

// Aggregate interface to accomplish aggregate features such as dispatch events
type Aggregate interface {
	DispatchRegisteredEvents(handlers map[string][]domainevent.DomainEventHandler)
}

// CommonAggregate is the main functionality an aggregate could have
// as my desire is to practice domain events dispatching an aggregate is supposed to be capable of register and dispatch those events
// to any part of the system which wants to react with a handler of function to a event in the entity
type CommonAggregate struct {
	events []domainevent.DomainEvent
}

// RegisterEvent register when a event happened, this registered event will be therefore prepared to be dispacthed to any handler interested
// to its
func (ca *CommonAggregate) RegisterEvent(event domainevent.DomainEvent) {
	ca.events = append(ca.events, event)
}

// DispatchRegisteredEvents implementation to execute user events handlers
func (ca *CommonAggregate) DispatchRegisteredEvents(eventHandlers map[string][]domainevent.DomainEventHandler, targetEvents []string) {
	eventsString := strings.Join(targetEvents, " ")

	for _, e := range ca.events {
		if !strings.Contains(eventsString, e.Name()) {
			continue
		}

		if eventHandlers, ok := eventHandlers[e.Name()]; ok {
			for _, handler := range eventHandlers {
				go func(hn domainevent.DomainEventHandler, event domainevent.DomainEvent) {
					err := hn.Run(event)
					if err != nil {
						// It's a decision that event handler errors will not stop application
						log.Printf("Error: %s in event: %v", err, event)
					}
				}(handler, e)
			}
		}
	}
}
