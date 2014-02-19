/* Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
