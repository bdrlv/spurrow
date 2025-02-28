package spurrow

import "sync"

type Step struct {
	processor    *Processor
	inputs       []*Step
	result       interface{}
	err          error
	once         sync.Once
	externalData interface{}
}

func NewStep(p *Processor, inputs ...*Step) *Step {
	return &Step{
		processor: p,
		inputs:    inputs,
	}
}

func (s *Step) Execute() (interface{}, error) {
	s.once.Do(func() {
		var wg sync.WaitGroup
		results := make([]interface{}, len(s.inputs))
		errors := make([]error, len(s.inputs))

		for i, step := range s.inputs {
			wg.Add(1)
			go func(idx int, st *Step) {
				defer wg.Done()
				res, err := st.Execute()
				results[idx] = res
				errors[idx] = err
			}(i, step)
		}
		wg.Wait()

		for _, err := range errors {
			if err != nil {
				s.err = err
				return
			}
		}

		if len(s.inputs) == 0 {
			results = []interface{}{s.externalData}
		}

		res, err := s.processor.Do(results...)
		if err != nil {
			s.err = err
			return
		}
		s.result = res
	})
	return s.result, s.err
}
