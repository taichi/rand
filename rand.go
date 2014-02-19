package rand

import (
	"crypto/rand"
	"math/big"
)

type RandomStringer struct {
	proposal []rune
	length   *big.Int
}

func New(runes []rune) *RandomStringer {
	rs := &RandomStringer{
		proposal: runes,
		length:   big.NewInt(int64(len(runes))),
	}
	return rs
}

func Alnum() *RandomStringer {
	b := make([]rune, 0)
	b = append(b, Runes('0', '9')...)
	b = append(b, Runes('a', 'z')...)
	b = append(b, Runes('A', 'Z')...)
	return New(b)
}

func Runes(from, to rune) []rune {
	s := make([]rune, to-from+1)
	for i := to - from; -1 < i; i-- {
		s[i] = from + i
	}
	return s
}

func (rs *RandomStringer) Next(length int) string {
	s := make([]rune, length)
	for i := 0; i < length; i++ {
		r, e := rand.Int(rand.Reader, rs.length)
		if e == nil {
			s[i] = rs.proposal[uint(r.Uint64())]
		} else {
			panic(e)
		}
	}
	return string(s)
}
