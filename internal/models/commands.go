package models

import "encoding/binary"

const setPacketId = 0x01
const getPacketId = 0x02
const delPacketId = 0x03
const modPacketId = 0x04

type Get struct {
	Key string
}

func (get *Get) ToBytes() []byte {
	b := make([]byte, 0)
	b = append(b, getPacketId)
	b = append(b, stringToBytes(&get.Key)...)

	return b
}

type Delete struct {
	Key string
}

func (delete *Delete) ToBytes() []byte {
	b := make([]byte, 0)
	b = append(b, delPacketId)
	b = append(b, stringToBytes(&delete.Key)...)

	return b
}

type Modify struct {
	Key   string
	Value string
}

func (modify *Modify) ToBytes() []byte {
	b := make([]byte, 0)
	b = append(b, modPacketId)
	b = append(b, stringToBytes(&modify.Key)...)
	b = append(b, stringToBytes(&modify.Value)...)

	return b
}

type Set struct {
	Key   string
	Value string
}

func (set *Set) ToBytes() []byte {
	b := make([]byte, 0)
	b = append(b, setPacketId)
	b = append(b, stringToBytes(&set.Key)...)
	b = append(b, stringToBytes(&set.Value)...)

	return b
}

func stringToBytes(s *string) []byte {
	b := make([]byte, 2)
	keyLen := len(*s)
	binary.BigEndian.PutUint16(b, uint16(keyLen))

	return append(b, []byte(*s)...)
}
