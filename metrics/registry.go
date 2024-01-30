package metrics

type Registry interface {
	GetOrRegisterCounter(name []string, description string) Counter
	GetOrRegisterGauge(name []string, description string) Gauge
}

func GetOrRegisterCounter(name []string, description string) Counter {
	return registry.GetOrRegisterCounter(append([]string{cfg.Prefix}, name...), description)
}

func GetOrRegisterGauge(name []string, description string) Gauge {
	return registry.GetOrRegisterGauge(append([]string{cfg.Prefix}, name...), description)
}
