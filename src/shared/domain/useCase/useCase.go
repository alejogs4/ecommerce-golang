package usecase

import "github.com/alejogs4/hn-website/src/shared/domain/domainevent"

// UseCase interface
type UseCase interface {
	RegisterEventHandler(eventName string, handler domainevent.DomainEventHandler)
}

// EventScheduler struct define methods to register domain event and its handlers, those which will be executed when a
// meaningful domain even occurs
type EventScheduler struct {
	handlers map[string][]domainevent.DomainEventHandler
}

// NewEventScheduler .
func NewEventScheduler() EventScheduler {
	return EventScheduler{handlers: make(map[string][]domainevent.DomainEventHandler)}
}

// Handlers returns non exported handlers field
func (uc *EventScheduler) Handlers() map[string][]domainevent.DomainEventHandler {
	return uc.handlers
}

// RegisterEventHandler UseCase interface implementation for register event handlers
func (uc *EventScheduler) RegisterEventHandler(eventName string, handler domainevent.DomainEventHandler) {
	uc.handlers[eventName] = append(uc.handlers[eventName], handler)
}
