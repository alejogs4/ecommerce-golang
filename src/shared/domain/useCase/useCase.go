package usecase

import "github.com/alejogs4/hn-website/src/shared/domain/domainevent"

type UseCase interface {
	RegisterEventHandler(eventName string, handler domainevent.DomainEventHandler)
}

type EventScheduler struct {
	handlers map[string][]domainevent.DomainEventHandler
}

func NewEventScheduler() EventScheduler {
	return EventScheduler{handlers: make(map[string][]domainevent.DomainEventHandler)}
}

func (uc *EventScheduler) Handlers() map[string][]domainevent.DomainEventHandler {
	return uc.handlers
}

// RegisterEventHandler UseCase interface implementation for register event handlers
func (uc *EventScheduler) RegisterEventHandler(eventName string, handler domainevent.DomainEventHandler) {
	uc.handlers[eventName] = append(uc.handlers[eventName], handler)
}
