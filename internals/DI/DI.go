// Package di provides various things like DI system etc
package di

import (
	"fmt"
	"reflect"
)

func NewDI() *DI {
	return &DI{
		Modules:  []string{},
		Services: make(map[string]ServiceEntry),
	}
}

func (d *DI) AddService(key string, service any) {
	d.Services[key] = ServiceEntry{
		Service: service,
		Type:    reflect.TypeOf(service),
	}
}

//GetService function returns bool,service ,service type , pass it key of the service and also moduleName to store modules which use DI to ensure we can check which modules need Di
func (d *DI) GetService(key string, moduleName string) (bool, any, reflect.Type) {
	d.Modules = append(d.Modules, moduleName)
	serviceEntry, ok := d.Services[key]
	return ok, serviceEntry.Service, serviceEntry.Type
}

func (d *DI) CheckService(key string) bool {
	_, ok := d.Services[key]
	return ok
}

func (d *DI) PrintDIModules() {
	fmt.Println("Printing which modules are using dependency injection")
	for _, module := range d.Modules {
		fmt.Println(module)
	}
}

func (d *DI) DeleteDI() {
	d.Modules = nil
	d.Services = nil
}
