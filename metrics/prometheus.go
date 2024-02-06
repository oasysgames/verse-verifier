package metrics

import (
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type prometheusRegistry struct {
	cache sync.Map
}

func (r *prometheusRegistry) GetOrRegisterCounter(name []string, description string) Counter {
	key := strings.Join(name, "_")
	if m, ok := r.cache.Load(key); ok {
		return m.(Counter)
	}

	m := &prometheusCounter{
		prometheus.NewCounter(prometheus.CounterOpts{
			Name: key,
			Help: description,
		}),
	}

	r.cache.Store(key, m)
	prometheus.MustRegister(m.c)
	return m
}

func (r *prometheusRegistry) GetOrRegisterGauge(name []string, description string) Gauge {
	key := strings.Join(name, "_")
	if m, ok := r.cache.Load(key); ok {
		return m.(Gauge)
	}

	m := &prometheusGauge{
		prometheus.NewGauge(prometheus.GaugeOpts{
			Name: key,
			Help: description,
		}),
	}

	r.cache.Store(key, m)
	prometheus.MustRegister(m.c)
	return m
}

type prometheusCounter struct {
	c prometheus.Counter
}

func (p *prometheusCounter) Incr() {
	p.c.Inc()
}

func (p *prometheusCounter) Add(value float64) {
	p.c.Add(value)
}

type prometheusGauge struct {
	c prometheus.Gauge
}

func (p *prometheusGauge) Incr() {
	p.c.Inc()
}

func (p *prometheusGauge) Decr() {
	p.c.Dec()
}

func (p *prometheusGauge) Set(value float64) {
	p.c.Set(value)
}

func (p *prometheusGauge) Add(value float64) {
	p.c.Add(value)
}

func (p *prometheusGauge) Sub(value float64) {
	p.c.Sub(value)
}
