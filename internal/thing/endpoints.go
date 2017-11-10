package thing

import (
	"context"
)

// Endpoints defines an interface for all Entitlements endpoints
type Endpoints interface {
	Thing(context.Context, interface{}) (interface{}, error)
}

// endpoints implements the Entitlements interface and holds any required data
type endpoints struct {
	service Service
}

// NewEndpoints returns an endpoints struct
func NewEndpoints(s Service) Endpoints {
	return endpoints{
		service: s,
	}
}
