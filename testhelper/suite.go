package testhelper

import (
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Suite struct {
	suite.Suite
}

func (s *Suite) Range(start, stop int) []int {
	size := stop - start
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = start + i
	}
	return numbers
}

func (s *Suite) Shuffle(src []int) []int {
	cpy := make([]int, len(src))
	copy(cpy, src)
	rand.Shuffle(len(cpy), func(i, j int) {
		cpy[i], cpy[j] = cpy[j], cpy[i]
	})
	return cpy
}

func (s *Suite) Pick(src []int) int {
	return rand.Intn(len(src))
}

func (s *Suite) Intn(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (s *Suite) ItoHash(i int) common.Hash {
	return common.BigToHash(big.NewInt(int64(i)))
}

func (s *Suite) AtoHash(a string) common.Hash {
	return common.BytesToHash([]byte(a))
}

func (s *Suite) ItoAddress(i int) common.Address {
	return common.BigToAddress(big.NewInt(int64(i)))
}

func (s *Suite) RandBytes() []byte {
	var b []byte
	rand.Read(b)
	return b
}

func (s *Suite) RandHash() common.Hash {
	var hash common.Hash
	b := make([]byte, common.HashLength)
	rand.Read(b)
	copy(hash[:], b)
	return hash
}

func (s *Suite) RandAddress() common.Address {
	return common.BigToAddress(big.NewInt(int64(rand.Intn(999999999))))
}

func (s *Suite) NoDBError(tx *gorm.DB, msgAndArgs ...interface{}) bool {
	return s.NoError(tx.Error, msgAndArgs...)
}
