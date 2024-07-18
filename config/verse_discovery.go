package config

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

type VerseDiscovery struct {
	client          *http.Client
	url             string
	refreshInterval time.Duration
	topic           *util.Topic
	log             log.Logger
}

func NewVerseDiscovery(
	ctx context.Context,
	client *http.Client,
	url string,
	refreshInterval time.Duration,
) (disc *VerseDiscovery, err error) {
	disc = &VerseDiscovery{
		client:          client,
		url:             url,
		refreshInterval: refreshInterval,
		topic:           util.NewTopic(),
		log:             log.New("worker", "verse-discovery"),
	}
	// Commented out the initial fetch, as it will be done in the worker
	// if _, err = disc.fetch(ctx); err != nil {
	// 	return nil, fmt.Errorf("the inital verse discovery failed, make sure the url(%s) is reachable: %w", url, err)
	// }
	return
}

func (w *VerseDiscovery) Subscribe(ctx context.Context) *VerseSubscription {
	ch := make(chan []*Verse)
	cancel := w.topic.Subscribe(ctx, func(ctx context.Context, data interface{}) {
		if t, ok := data.([]*Verse); ok {
			ch <- t
		}
	})
	return &VerseSubscription{Cancel: cancel, ch: ch}
}

func (w *VerseDiscovery) Work(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	data, err := w.fetch(ctx)
	if err != nil {
		w.log.Error("Discovery request failed", "err", err)
		return err
	}

	verses, err := w.unmarshal(data)
	if err != nil {
		w.log.Error("Failed to unmarshal response body", "err", err)
		return err
	}

	w.topic.Publish(verses)
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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (w *VerseDiscovery) unmarshal(data []byte) (verses []*Verse, err error) {
	err = json.Unmarshal(data, &verses)
	if err != nil {
		return nil, err
	}
	return verses, nil
}

type VerseSubscription struct {
	Cancel context.CancelFunc
	ch     chan []*Verse
}

func (s *VerseSubscription) Next() <-chan []*Verse {
	return s.ch
}
