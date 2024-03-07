package config

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	thttp "github.com/oasysgames/oasys-optimism-verifier/testhelper/http"
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
	client := thttp.NewTestHTTPClient(func(req *http.Request) *http.Response {
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
		wg           sync.WaitGroup
		gots0, gots1 []*Verse
	)
	wg.Add(4)
	go func() {
		for {
			select {
			case got := <-sub0.Next():
				gots0 = append(gots0, got)
			case got := <-sub1.Next():
				gots1 = append(gots1, got)
			}
			wg.Done()
		}
	}()

	// start discovery
	go discovery.Start(context.Background())
	wg.Wait()

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

	s.Len(gots0, 2)
	s.Len(gots1, 2)

	s.Equal(want0, *gots0[0])
	s.Equal(want0, *gots1[0])

	s.Equal(want1, *gots0[1])
	s.Equal(want1, *gots1[1])
}
