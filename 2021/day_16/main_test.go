package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadPacket(t *testing.T) {
	tts := []struct {
		name string
		set  string
		exp  Packet
	}{
		{"literal", "110100101111111000101000", Packet{version: 6, typeID: 4, content: "101111111000101"}},
		{"operatorl0", "00111000000000000110111101000101001010010001001000000000", Packet{version: 1, typeID: 6, content: "110100010100101001000100100", lengthType: 0, length: 27, subPackets: []Packet{{version: 6, typeID: 4, content: "01010"}, {version: 2, typeID: 4, content: "1000100100"}}}},
		{"operatorl1", "11101110000000001101010000001100100000100011000001100000", Packet{version: 7, typeID: 3, content: "010100000011001000001000110000011", lengthType: 1, length: 3, subPackets: []Packet{{version: 3, typeID: 4, content: "00001"}, {version: 4, typeID: 4, content: "00010"}, {version: 1, typeID: 4, content: "00011"}}}},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			res, br := readPacket(tt.set, 0, true)
			assert.Equal(t, tt.exp.version, res.version, "version should be equal")
			assert.Equal(t, tt.exp.typeID, res.typeID, "typeID should be equal")
			assert.Equal(t, tt.exp.content, res.content, "content should be equal")

			if tt.exp.typeID != 4 {
				assert.Equal(t, len(tt.exp.subPackets), len(res.subPackets))
			}
			if tt.exp.typeID == 4 {
				assert.Equal(t, len(tt.set), br, "should have consumed all bits")
			}
		})
	}
}

func TestOperate(t *testing.T) {
	tts := []struct {
		name string
		set  string
		exp  int64
	}{
		{"example-1", "C200B40A82", 3},
		{"ex-2", "04005AC33890", 54},
		{"ex-3", "880086C3E88112", 7},
		{"ex-4", "CE00C43D881120", 9},
		{"ex-5", "D8005AC2A8F0", 1},
		{"ex-6", "F600BC2D8F", 0},
		{"ex-7", "9C005AC2F8F0", 0},
		{"ex-8", "9C0141080250320F1802104A08", 1},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			tm := parseInputs([]string{tt.set})
			pkt, _ := readPacket(tm, 0, true)
			assert.Equal(t, tt.exp, operate(pkt))
		})
	}
}
