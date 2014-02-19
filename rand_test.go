package rand_test

import (
	. "."
	"crypto/rand"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Rand", func() {
	Context("Runes", func() {
		It("should work normally", func() {
			runes := Runes('a', 'd')
			Expect(runes).To(Equal([]rune{'a', 'b', 'c', 'd'}))
		})
	})
	Context("Alnum", func() {
		var rs *RandomStringer
		BeforeEach(func() {
			rs = Alnum()
		})
		It("should work normally", FixSeed("abcABCxyz123", func() {
			Expect(rs.Next(10)).To(Equal("xyz123UVWN"))
		}))
	})
})

func FixSeed(seed string, fn func()) func() {
	return func() {
		orig := rand.Reader
		rand.Reader = strings.NewReader(seed) // fix seed
		defer func() {
			rand.Reader = orig
		}()
		fn()
	}
}
