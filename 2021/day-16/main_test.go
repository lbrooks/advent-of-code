package main

import (
	"fmt"
	"testing"
)

type housing struct {
	bits string
	exp  uint64
}

func TestSumVersions(t *testing.T) {
	allTests := []housing{
		{"8A004A801A8002F478", uint64(16)},
		{"620080001611562C8802118E34", uint64(12)},
		{"C0015000016115A2E0802F182340", uint64(23)},
		{"A0016C880162017C3686B18A3D4780", uint64(31)},
	}

	for _, s := range allTests {
		reader := NewBitsReader(toBin(s.bits))
		packet := read(reader)
		if packet.version != s.exp {
			t.Errorf("read('%s') : Version = %v; want %v", s.bits, packet.version, s.exp)
		}
	}
}

func TestValues(t *testing.T) {
	allTests := []housing{
		{"C200B40A82", uint64(3)},
		{"04005AC33890", uint64(54)},
		{"880086C3E88112", uint64(7)},
		{"CE00C43D881120", uint64(9)},
		{"D8005AC2A8F0", uint64(1)},
		{"F600BC2D8F", uint64(0)},
		{"9C005AC2F8F0", uint64(0)},
		{"9C0141080250320F1802104A08", uint64(1)},
	}

	for _, s := range allTests {
		reader := NewBitsReader(toBin(s.bits))
		packet := read(reader)

		fmt.Println(packet)
		if packet.value != s.exp {
			t.Errorf("read('%s') : Value = %v; want %v", s.bits, packet.value, s.exp)
		}
	}
}
