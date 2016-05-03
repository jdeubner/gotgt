/*
Copyright 2015 The GoStor Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import "encoding/binary"

func GetUnalignedUint16(u8 []uint8) uint16 {
	return binary.BigEndian.Uint16(u8)
}

func GetUnalignedUint32(u8 []uint8) uint32 {
	return binary.BigEndian.Uint32(u8)
}

func GetUnalignedUint64(u8 []uint8) uint64 {
	return binary.BigEndian.Uint64(u8)
}

// ParseKVText parses iSCSI key value data.
func ParseKVText(txt []byte) map[string]string {
	m := make(map[string]string)
	var kv, sep int
	var key string
	for i := 0; i < len(txt); i++ {
		if txt[i] == '=' {
			if key == "" {
				sep = i
				key = string(txt[kv:sep])
			}
			continue
		}
		if txt[i] == 0 && key != "" {
			m[key] = string(txt[sep+1 : i])
			key = ""
			kv = i + 1
		}
	}
	return m
}

func MarshalKVText(kv map[string]string) []byte {
	var data []byte
	for k, v := range kv {
		data = append(data, []byte(k)...)
		data = append(data, '=')
		data = append(data, []byte(v)...)
		data = append(data, 0)
	}
	return data
}