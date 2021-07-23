package kernel

type Application interface {
	AddModule(module Module)
	RemoveModule(module Module)
	Start()
	Stop()
}
