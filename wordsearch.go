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

// package wordsearch fast word search for prepared wordlist
// use 256 radix(byte) search
package wordsearch

import (
	"bytes"
	"fmt"
)

type ByteNodeList []*ByteNode

func (bnl ByteNodeList) String() string {
	var buff bytes.Buffer

	for _, v := range bnl {
		fmt.Fprintf(&buff, "%c", v.Value)
	}
	fmt.Fprintf(&buff, "\n")
	return buff.String()
}

type ByteNode struct {
	Exist bool
	Value byte
	Child [256]*ByteNode
}

func New(v byte) *ByteNode {
	rt := &ByteNode{
		Value: v,
	}
	return rt
}

func (bn ByteNode) String() string {
	var buff bytes.Buffer

	bn.Traverse(nil, func(in ByteNodeList) bool {
		for _, v := range in {
			fmt.Fprintf(&buff, "%c", v.Value)
		}
		fmt.Fprintf(&buff, "\n")
		return false
	})
	return buff.String()
}

func (bn *ByteNode) Traverse(in ByteNodeList, fn func(b ByteNodeList) bool) bool {
	self := append(in, bn)
	if bn.Exist {
		if fn(self) { // stop
			return true
		}
	}
	for _, v := range bn.Child {
		if v != nil {
			if v.Traverse(self, fn) {
				return true
			}
		}
	}
	return false
}

func (bn *ByteNode) Add(s []byte) bool {
	cn := bn
	for i := 0; i < len(s); i++ {
		if cn.Child[s[i]] == nil {
			cn.Child[s[i]] = New(s[i])
			cn = cn.Child[s[i]]
		} else {
			cn = cn.Child[s[i]]
		}
	}
	if cn.Exist {
		return false
	}
	cn.Exist = true
	return true
}

func (bn *ByteNode) Del(s []byte) bool {
	cn := bn
	for i := 0; i < len(s); i++ {
		if cn.Child[s[i]] == nil {
			return false
		} else {
			cn = cn.Child[s[i]]
		}
	}
	if cn.Exist == false {
		return false
	} else {
		cn.Exist = false
		return true
	}
}

func (bn *ByteNode) Find(s []byte, in ByteNodeList, fn func(b ByteNodeList) bool) {
	self := append(in, bn)
	if bn.Exist {
		if fn(self) { // stop
			return
		}
	}
	if len(s) == 0 || bn.Child[s[0]] == nil {
		return
	} else {
		bn.Child[s[0]].Find(s[1:], self, fn)
	}
}

func (bn *ByteNode) IsExist(s string) bool {
	arg := []byte(s)
	for i := 0; i < len(arg)-1; i++ {
		found := false
		bn.Find(arg[i:], nil, func(in ByteNodeList) bool {
			found = true
			return true // stop traverse
		})
		if found {
			return true
		}
	}
	return false
}
