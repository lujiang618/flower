package kernel

type Kernel struct {
	apps []Application
}

func NewKernel() *Kernel {
	return &Kernel{}
}

func (k *Kernel) AddApplication(app Application) {
	k.apps = append(k.apps, app)
}

func (k *Kernel) RemoveApplication() {

}

func (k *Kernel) Run() {
	for _, app := range k.apps {
		app.Start()
	}
}

func (k *Kernel) Stop() {
	for _, app := range k.apps {
		app.Stop()
	}
}
