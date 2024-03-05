package graph

import "github.com/GodlyDad/Backend/pkg/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IBibleService
}

// func NewResolver(s service.IBibleService) *Resolver {
// 	return &Resolver{
// 		Service: s,
// 	}
// }
