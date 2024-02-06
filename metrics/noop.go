package metrics

type noopRegistry struct{}

func (r *noopRegistry) GetOrRegisterCounter(name []string, description string) Counter {
	return &noopCounter{}
}

func (r *noopRegistry) GetOrRegisterGauge(name []string, description string) Gauge {
	return &noopGauge{}
}

type noopCounter struct{}

func (p *noopCounter) Incr()             {}
func (p *noopCounter) Add(value float64) {}

type noopGauge struct{}

func (p *noopGauge) Incr()             {}
func (p *noopGauge) Decr()             {}
func (p *noopGauge) Set(value float64) {}
func (p *noopGauge) Add(value float64) {}
func (p *noopGauge) Sub(value float64) {}
