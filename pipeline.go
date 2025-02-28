package spurrow

import (
	"fmt"
	"sync"
)

type Pipeline struct {
	chains map[string]*Step
	names  map[*Step]string
	mu     sync.Mutex
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		chains: make(map[string]*Step),
		names:  make(map[*Step]string),
	}
}

func (pf *Pipeline) RegisterChain(
	name string,
	processor func(...interface{}) (interface{}, error),
	dependencies ...string,
) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	p := NewProcessor(processor)
	var inputSteps []*Step

	for _, dep := range dependencies {
		inputSteps = append(inputSteps, pf.chains[dep])
	}

	step := NewStep(p, inputSteps...)
	pf.chains[name] = step
	pf.names[step] = name
}

func (pf *Pipeline) Execute(chainName string, inputs map[string]interface{}) (interface{}, error) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	for name, data := range inputs {
		if step, exists := pf.chains[name]; exists {
			step.externalData = data
		}
	}

	if chain, exists := pf.chains[chainName]; exists {
		return chain.Execute()
	}
	return nil, fmt.Errorf("chain %s not found", chainName)
}
