package aggregate

import "github.com/alejogs4/hn-website/src/shared/domain/domainevent"

// Aggregate interface to accomplish aggregate features such as dispatch events
type Aggregate interface {
	DispatchRegisteredEvents(handlers map[string][]domainevent.DomainEventHandler)
}
