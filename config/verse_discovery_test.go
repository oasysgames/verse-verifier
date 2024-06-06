package config

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	httphelper "github.com/oasysgames/oasys-optimism-verifier/testhelper/http"
	"github.com/stretchr/testify/suite"
)

type VerseDiscoveryTestSuite struct {
	suite.Suite
}

func TestVerseDiscovery(t *testing.T) {
	suite.Run(t, new(VerseDiscoveryTestSuite))
}

func (s *VerseDiscoveryTestSuite) TestDiscover() {
	// setup test client
	client := httphelper.NewTestHTTPClient(func(req *http.Request) *http.Response {
		s.Equal("https://example.com/", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body: ioutil.NopCloser(bytes.NewBufferString(`
				[
					{
						"chain_id": 1,
						"rpc": "https://verse1.example.com/",
						"l1_contracts": {
							"StateCommitmentChain": "0x6B7e39db6638be17eBF8a9e64120c62a707982c2"
						}
					},
					{
						"chain_id": 2,
						"rpc": "https://verse2.example.com/",
						"l1_contracts": {
							"StateCommitmentChain": "0xCe97A2618d6990e19741E61fa5DCD896D759E641"
						}
					}
				]
			`)),
		}
	})

	// setup pubsub
	discovery := NewVerseDiscovery(client, "https://example.com/", time.Second)
	sub0 := discovery.Subscribe(context.Background())
	sub1 := discovery.Subscribe(context.Background())

	var (
		wg         sync.WaitGroup
		got0, got1 []*Verse
	)
	wg.Add(2)
	go func() {
		for {
			select {
			case got := <-sub0.Next():
				got0 = got
			case got := <-sub1.Next():
				got1 = got
			}
			wg.Done()
		}
	}()

	// start discovery
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		tick := time.NewTicker(discovery.refreshInterval)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				discovery.Work(ctx)
			}
		}
	}()
	wg.Wait()
	cancel()

	// assert
	want0 := Verse{
		ChainID: 1,
		RPC:     "https://verse1.example.com/",
		L1Contracts: map[string]string{
			"StateCommitmentChain": "0x6B7e39db6638be17eBF8a9e64120c62a707982c2",
		},
	}
	want1 := Verse{
		ChainID: 2,
		RPC:     "https://verse2.example.com/",
		L1Contracts: map[string]string{
			"StateCommitmentChain": "0xCe97A2618d6990e19741E61fa5DCD896D759E641",
		},
	}

	s.Len(got0, 2)
	s.Len(got1, 2)

	s.Equal(want0, *got0[0])
	s.Equal(want0, *got1[0])

	s.Equal(want1, *got0[1])
	s.Equal(want1, *got1[1])
}
