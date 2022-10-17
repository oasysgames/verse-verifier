package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

type response []*Verse

type VerseDiscovery struct {
	client          *http.Client
	url             string
	refreshInterval time.Duration

	topic *util.Topic
	log   log.Logger
}

func NewVerseDiscovery(
	client *http.Client,
	url string,
	refreshInterval time.Duration,
) *VerseDiscovery {
	return &VerseDiscovery{
		client:          client,
		url:             url,
		refreshInterval: refreshInterval,
		topic:           util.NewTopic(),
		log:             log.New("worker", "verse-discovery"),
	}
}

func (w *VerseDiscovery) Start(ctx context.Context) {
	w.log.Info("Worker started", "endpoint", w.url, "interval", w.refreshInterval)

	for {
		if w.work(ctx) == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	tick := time.NewTicker(w.refreshInterval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-tick.C:
			w.work(ctx)
		}
	}
}

func (w *VerseDiscovery) Subscribe(ctx context.Context) *VerseSubscription {
	ch := make(chan *Verse)
	cancel := w.topic.Subscribe(ctx, func(ctx context.Context, data interface{}) {
		if t, ok := data.(*Verse); ok {
			ch <- t
		}
	})
	return &VerseSubscription{Cancel: cancel, ch: ch}
}

func (w *VerseDiscovery) work(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	data, err := w.fetch(ctx)
	if err != nil {
		w.log.Error("Request failed", "err", err)
		return err
	}

	verses, err := w.unmarshal(data)
	if err != nil {
		w.log.Error("Failed to unmarshal response body", "err", err)
		return err
	}

	for _, v := range verses {
		w.topic.Publish(v)
	}

	return nil
}

func (w *VerseDiscovery) fetch(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequest("GET", w.url, nil)
	if err != nil {
		return nil, err
	}

	res, err := w.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (w *VerseDiscovery) unmarshal(data []byte) (response, error) {
	var resp response
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type VerseSubscription struct {
	Cancel context.CancelFunc
	ch     chan *Verse
}

func (s *VerseSubscription) Next() <-chan *Verse {
	return s.ch
}
