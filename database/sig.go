package database

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	SignatureLength = 65
)

var (
	SignatureT = reflect.TypeOf(Signature{})
)

// Hash represents the 65 byte ethereum signature.
type Signature [SignatureLength]byte

// BytesSignature sets b to signature.
// If b is larger than len(h), b will be cropped from the left.
func BytesSignature(b []byte) Signature {
	var s Signature
	s.SetBytes(b)
	return s
}

// IntSignature sets int to signature.
func IntSignature(i int) Signature {
	var (
		sig  Signature
		hash = common.BigToHash(big.NewInt(int64(i)))
		pos  = SignatureLength - common.HashLength
	)
	copy(sig[pos:], hash[:])
	return sig
}

// RandSignature returns randomly generated signature.
func RandSignature() Signature {
	var sig Signature
	b := make([]byte, SignatureLength)
	rand.Read(b)
	copy(sig[:], b)
	return sig
}

// Bytes gets the byte representation of the underlying signature.
func (s Signature) Bytes() []byte { return s[:] }

// Hex converts a signature to a hex string.
func (s Signature) Hex() string { return hexutil.Encode(s[:]) }

// String implements the stringer interface and is used also by the logger when
// doing full logging into a file.
func (s Signature) String() string {
	return s.Hex()
}

// Format implements fmt.Formatter.
// Signature supports the %v, %s, %q, %x, %X and %d format verbs.
func (s Signature) Format(f fmt.State, c rune) {
	hexb := make([]byte, 2+len(s)*2)
	copy(hexb, "0x")
	hex.Encode(hexb[2:], s[:])

	switch c {
	case 'x', 'X':
		if !f.Flag('#') {
			hexb = hexb[2:]
		}
		if c == 'X' {
			hexb = bytes.ToUpper(hexb)
		}
		fallthrough
	case 'v', 'f':
		f.Write(hexb)
	case 'q':
		q := []byte{'"'}
		f.Write(q)
		f.Write(hexb)
		f.Write(q)
	case 'd':
		fmt.Fprint(f, ([len(s)]byte)(s))
	default:
		fmt.Fprintf(f, "%%!%c(signature=%x)", c, s)
	}
}

// UnmarshalJSON parses a signature in hex syntax.
func (s *Signature) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(SignatureT, input, s[:])
}

// SetBytes sets the signature to the value of b.
// If b is larger than len(s), b will be cropped from the left.
func (s *Signature) SetBytes(b []byte) {
	if len(b) > len(s) {
		b = b[len(b)-SignatureLength:]
	}

	copy(s[SignatureLength-len(b):], b)
}

// Scan implements Scanner for database/sql.
func (s *Signature) Scan(src interface{}) error {
	srcB, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can't scan %T into Signature", src)
	}
	if len(srcB) != SignatureLength {
		return fmt.Errorf(
			"can't scan []byte of len %d into Signature, want %d",
			len(srcB),
			SignatureLength,
		)
	}
	copy(s[:], srcB)
	return nil
}

// Value implements valuer for database/sql.
func (s Signature) Value() (driver.Value, error) {
	return s[:], nil
}
