package domainevent

import "time"

type DomainEvent interface {
	Name() string
	OccurredTimed() time.Time
	EventInformation() interface{}
}

type DomainEventHandler interface {
	Run(DomainEvent) error
}
