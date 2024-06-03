package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type packet struct {
	tookBits uint64

	version uint64
	typeID  uint64

	lengthType    uint64
	numSubPackets uint64
	numSubBits    uint64

	value uint64

	packets []*packet
}

func (p *packet) StringIndented(indent string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprint(indent,
		"Version: ", p.version,
		"\tTypeID: ", p.typeID,
		"\tLength Type: ", p.lengthType,
		"\tBits: ", p.tookBits,
		"\tNum Sub Packets: ", p.numSubPackets,
		"\tNum Sub Bits: ", p.numSubBits,
		"\tValue: ", p.value,
		"\tNum Children: ", len(p.packets),
		"\n",
	))
	for _, c := range p.packets {
		builder.WriteString(c.StringIndented("\t" + indent))
	}
	return builder.String()
}

func (p *packet) String() string {
	return fmt.Sprintf(p.StringIndented(""))
}

type combine func([]*packet) uint64

func mergeSubPackets(wrapper *packet) uint64 {
	for _, cp := range wrapper.packets {
		// For Part 1
		wrapper.version += cp.version

		wrapper.tookBits += cp.tookBits
	}

	var combiner combine
	switch wrapper.typeID {
	case 0:
		// Sum
		combiner = func(children []*packet) (val uint64) {
			for _, c := range children {
				val += c.value
			}
			return
		}
	case 1:
		// Product
		combiner = func(children []*packet) (val uint64) {
			val = 1
			for _, c := range children {
				val *= c.value
			}
			return
		}
	case 2:
		// Min
		combiner = func(children []*packet) (val uint64) {
			val = math.MaxUint64
			for _, c := range children {
				if c.value < val {
					val = c.value
				}
			}
			return
		}
	case 3:
		// Max
		combiner = func(children []*packet) (val uint64) {
			for _, c := range children {
				if c.value > val {
					val = c.value
				}
			}
			return
		}
	case 5:
		// Greater Than
		combiner = func(children []*packet) (val uint64) {
			if len(children) != 2 {
				return
			}
			if children[0].value > children[1].value {
				return uint64(1)
			}
			return
		}
	case 6:
		// Less Than
		combiner = func(children []*packet) (val uint64) {
			if len(children) != 2 {
				return
			}
			if children[0].value < children[1].value {
				return uint64(1)
			}
			return
		}
	case 7:
		// Equals
		combiner = func(children []*packet) (val uint64) {
			if len(children) != 2 {
				return
			}
			if children[0].value == children[1].value {
				return uint64(1)
			}
			return
		}
	default:
		combiner = func(children []*packet) uint64 {
			return uint64(0)
		}
	}

	return combiner(wrapper.packets)
}

func toBin(input string) (Bits, uint64) {
	numBits := uint64(0)
	data := make(Bits, 0)
	for i := 0; i < len(input); {
		var compressed uint64
		for b := 0; b < 16; b++ {
			compressed = compressed << 4
			if i < len(input) {
				num, _ := strconv.ParseUint(string(input[i]), 16, 4)
				compressed = compressed | num
				numBits += 4
			}
			i++
		}
		data = append(data, compressed)
	}
	return data, numBits
}

func read(input *BitsReader) *packet {
	p := &packet{
		version:  input.ReadBits(3),
		typeID:   input.ReadBits(3),
		tookBits: 6,
		packets:  make([]*packet, 0),
	}

	if p.typeID == 4 {
		hasMore := input.Read() == 1
		answer := input.ReadBits(4)
		p.tookBits += 5
		for hasMore {
			p.value = (p.value << 4) | answer

			hasMore = input.Read() == 1
			answer = input.ReadBits(4)
			p.tookBits += 5
		}
		p.value = (p.value << 4) | answer
	} else {
		p.lengthType = input.Read()
		p.tookBits += 1

		if p.lengthType == 1 {
			p.numSubPackets = input.ReadBits(11)
			p.tookBits += 11
		} else {
			p.numSubBits = input.ReadBits(15)
			p.tookBits += 15
		}

		if p.numSubPackets > 0 {
			for i := uint64(0); i < p.numSubPackets; i++ {
				cp := read(input)
				p.packets = append(p.packets, cp)
			}
		} else if p.numSubBits > 0 {
			var bitCount uint64
			for bitCount < p.numSubBits {
				cp := read(input)
				p.packets = append(p.packets, cp)

				bitCount += cp.tookBits
			}
		}

		p.value = mergeSubPackets(p)
	}

	return p
}

func playOne(input []string) {
	reader := NewBitsReader(toBin(input[0]))
	packet := read(reader)
	fmt.Println("Sum of versions", packet.version)
}

func playTwo(input []string) {
	reader := NewBitsReader(toBin(input[0]))
	packet := read(reader)
	fmt.Println("Value", packet.value)
}

func main() {
	input := utils.ReadPiped()

	playOne(input)
	playTwo(input)
}
