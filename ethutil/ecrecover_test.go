package ethutil

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestEcrecover(t *testing.T) {
	want := common.HexToAddress("0x26fE8b0aE2d2CD413B09935e927729C6Ef93EDDd")

	hash, _ := hex.DecodeString("1229298dbe3a5cdc45973993865f6979d39e9b52645c55c98690ee71d9a25845")
	sig, _ := hex.DecodeString(
		"1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0" +
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c")
	got, _ := Ecrecover(hash, sig)

	if got != want {
		t.Errorf("failed ecrecover, want: %s, got: %s", want, got)
	}
}
