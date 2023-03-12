package beacon

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type BeaconWorkerTestSuite struct {
	testhelper.Suite
}

func TestBeaconWorker(t *testing.T) {
	suite.Run(t, new(BeaconWorkerTestSuite))
}

func (s *BeaconWorkerTestSuite) TestBeaconWorker() {
	var (
		url  string
		data []byte
	)

	// setup test client
	client := testhelper.NewTestHTTPClient(func(req *http.Request) *http.Response {
		url = req.URL.String()
		data, _ = io.ReadAll(req.Body)

		return &http.Response{StatusCode: 200, Header: make(http.Header)}
	})

	conf := &config.Beacon{Endpoint: "https://example.com/", Interval: time.Hour}
	signer := common.HexToAddress("0xa379E5DeD16Bfc8D8B36449c45966648A2cE2AAE")
	worker := NewBeaconWorker(conf, client, Beacon{
		Signer:  signer.String(),
		Version: "0.0.1",
		PeerID:  "12D3KooWGqAVCkVwK2V8rrYS4GerjeUrGgzYUWZHp7SUng5LUMPY",
	})

	go worker.Start(context.Background())
	time.Sleep(time.Second / 10)

	// assert
	s.Equal("https://example.com/", url)
	s.Equal(
		`{"signer":"0xa379E5DeD16Bfc8D8B36449c45966648A2cE2AAE","version":"0.0.1","peer_id":"12D3KooWGqAVCkVwK2V8rrYS4GerjeUrGgzYUWZHp7SUng5LUMPY"}`,
		string(data),
	)
}
