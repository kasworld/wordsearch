// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wordsearch

import "testing"

func TestNew(t *testing.T) {
	bn := New(0)
	bn.Add([]byte("hello world"))
	bn.Add([]byte("hello"))
	bn.Add([]byte("hel"))
	bn.Add([]byte("0123456789"))
	bn.Add([]byte("world hello"))
	// t.Logf("%s", bn)

	// t.Logf("%v %v %v %v", arg, len(arg), n, m)

	bn.Del([]byte("hel"))
	t.Logf("%s", bn)

	arg := []byte("helloo asfas")

	bn.Find(arg, nil, func(in ByteNodeList) bool {
		t.Logf("%v", in)
		return false
	})

	if bn.IsExist("hello world hhdd") {
		t.Logf("found")
	}
}
