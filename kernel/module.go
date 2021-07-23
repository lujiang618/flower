package kernel

type Module interface {
	SetupRouter()
	RegisterCommand()
}
