package calab

import (
	"time"
)

// VirtualComputationalModel ...
type VirtualComputationalModel struct {
	System *DynamicalSystem
	// rendersPerSecond int
	renderers []Renderer
}

// NewVM ...
func NewVM(model *DynamicalSystem, renderers ...Renderer) *VirtualComputationalModel {
	return &VirtualComputationalModel{
		// rendersPerSecond: 60,
		System:    model,
		renderers: renderers,
	}
}

// AddRenderer ...
func (vm *VirtualComputationalModel) AddRenderer(r Renderer) {
	vm.renderers = append(vm.renderers, r)
}

// SetRPS sets renders per second rate.
// func (vm *VirtualComputationalModel) SetRPS(rendersPerSecond int) {
// 	vm.rendersPerSecond = rendersPerSecond
// }

// Run ...
func (vm *VirtualComputationalModel) Run(dt time.Duration) {
	ticks := make(chan uint64)
	done := make(chan struct{})
	// lastTime := time.Now()

	go func(done chan struct{}) {
		time.Sleep(dt)
		done <- struct{}{}
	}(done)

	vm.System.RunInfiniteSimulation(ticks, done)

	vm.System.Observe(ticks, func(n uint64, s Space) {
		// Limiting the renders per second.
		// elapsedTime := time.Since(lastTime)
		// expectedDuration := 1000 / time.Duration(vm.rendersPerSecond) * time.Millisecond

		// if elapsedTime < expectedDuration {
		// 	time.Sleep(expectedDuration - elapsedTime)
		// }

		// TODO: Update this rps limiter with an array of its, that's necessary for many renderers.
		for _, renderer := range vm.renderers {
			renderer(n, s)
		}

		// lastTime = time.Now()
	})
}

// RunTicks runs your simulation for n ticks.
func (vm *VirtualComputationalModel) RunTicks(ticks uint64) {
	ticksChannel := make(chan uint64)

	vm.System.RunSimulation(ticks, ticksChannel)

	vm.System.Observe(ticksChannel, func(n uint64, s Space) {
		for _, renderer := range vm.renderers {
			renderer(n, s)
		}
	})
}
