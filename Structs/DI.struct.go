// Package structs provides struct definitions for dependency injection and module management.
package structs

import "reflect"

type ServiceEntry struct {
	Service any
	Type    reflect.Type
}
type DI struct {
	Modules  []string
	Services map[string]ServiceEntry
}
