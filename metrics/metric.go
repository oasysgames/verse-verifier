package metrics

type Counter interface {
	Incr()
	Add(value float64)
}

type Gauge interface {
	Incr()
	Decr()
	Set(value float64)
	Add(value float64)
	Sub(value float64)
}
