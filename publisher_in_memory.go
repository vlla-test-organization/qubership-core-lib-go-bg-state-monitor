package blue_green_state_monitor_go

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
)

type InMemoryBlueGreenStatePublisher struct {
	statePointer *atomic.Pointer[BlueGreenState]
	subscribers  map[int64]func(state BlueGreenState)
	mutex        *sync.RWMutex
}

func NewInMemoryPublisher(bgState BlueGreenState) (*InMemoryBlueGreenStatePublisher, error) {
	statePointer := &atomic.Pointer[BlueGreenState]{}
	statePointer.Store(&bgState)
	return &InMemoryBlueGreenStatePublisher{statePointer: statePointer, subscribers: map[int64]func(state BlueGreenState){}, mutex: &sync.RWMutex{}}, nil
}

func (p *InMemoryBlueGreenStatePublisher) Subscribe(ctx context.Context, callback func(state BlueGreenState)) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	id := rand.Int63n(math.MaxInt64)
	p.subscribers[id] = callback
	go func() {
		for {
			select {
			case <-ctx.Done():
				p.unsubscribe(id)
				return
			}
		}
	}()
	callback(*p.statePointer.Load())
}

func (p *InMemoryBlueGreenStatePublisher) GetState() BlueGreenState {
	state := p.statePointer.Load()
	return *state
}

func (p *InMemoryBlueGreenStatePublisher) SetState(bgState BlueGreenState) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	p.statePointer.Store(&bgState)
	for _, subscriber := range p.subscribers {
		subscriber(bgState)
	}
}

func (p *InMemoryBlueGreenStatePublisher) unsubscribe(id int64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.subscribers, id)
}
