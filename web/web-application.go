package web

import "flower/kernel"

type WebApplication struct {
	modules []kernel.Module
}

func NewWebApplication() *WebApplication {
	return &WebApplication{}
}

func (a *WebApplication) AddModule(module kernel.Module) {
	a.modules = append(a.modules, module)
}

func (a *WebApplication) RemoveModule(module kernel.Module) {

}

func (a *WebApplication) Start() {
	for _, module := range a.modules {
		module.SetupRouter()
		module.RegisterCommand()
	}
}

func (a *WebApplication) Stop() {

}
