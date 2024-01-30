package ipccmd

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/wallet"
)

var WalletUnlockCmd = &walletUnlock{handlerID: WALLET_UNLOCK}

type walletUnlock struct {
	handlerID int
}

type walletUnlockMsg struct {
	Address  string
	Password string
}

func (c *walletUnlock) Run(listen, address, password string) {
	// attach to ipc
	s, err := ipc.NewClient(listen, c.handlerID)
	if err != nil {
		util.Exit(1, "connection failure: %s\n", err)
	}

	// start read loop
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.read(s)
	}()

	// send message
	data, err := json.Marshal(&walletUnlockMsg{address, password})
	if err != nil {
		util.Exit(1, "failed to create message: %s\n", err)
	}
	err = s.Write(data)
	if err != nil {
		util.Exit(1, "failed to write message: %s\n", err)
	}

	wg.Wait()
}

func (c *walletUnlock) read(s *ipc.Client) {
	data, err := s.Read()
	if err != nil {
		util.Exit(1, "failed to read message: %s\n", err)
	}

	fmt.Println(string(data))
}

func (c *walletUnlock) NewHandler(ks *wallet.KeyStore) (int, ipc.Handler) {
	return c.handlerID, func(s *ipc.IPCServer, data []byte) {
		var m walletUnlockMsg
		err := json.Unmarshal(data, &m)
		if err != nil {
			s.Write(c.handlerID, []byte(err.Error()))
			return
		}

		_, account, err := ks.FindWallet(common.HexToAddress(m.Address))
		if err != nil {
			s.Write(c.handlerID, []byte(err.Error()))
			return
		}

		err = ks.Unlock(*account, m.Password)
		if err != nil {
			s.Write(c.handlerID, []byte(err.Error()))
			return
		}

		s.Write(c.handlerID, []byte("success!"))
	}
}
