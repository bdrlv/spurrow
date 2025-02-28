package spurrow

type Processor struct {
	doFunc func(...interface{}) (interface{}, error)
}

func NewProcessor(f func(...interface{}) (interface{}, error)) *Processor {
	return &Processor{doFunc: f}
}

func (p *Processor) Do(inputs ...interface{}) (interface{}, error) {
	return p.doFunc(inputs...)
}
