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
