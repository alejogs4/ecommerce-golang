package usecase

import "github.com/alejogs4/hn-website/src/shared/domain/domainevent"

type UseCase interface {
	RegisterEventHandler(eventName string, handler domainevent.DomainEventHandler)
}
