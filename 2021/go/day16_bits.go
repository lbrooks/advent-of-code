package main

type Bits []uint64

func (b Bits) getBit(index uint64) uint64 {
	i := index / 64
	p := 63 - index%64
	v := (b[i] >> p) & 1
	return v
}

func (b Bits) getBits(index, length uint64) []uint64 {
	res := make([]uint64, 0)

	endIdx := index + length

	numbersNeeded := (length / 64)

	idx := index
	for n := uint64(0); n <= numbersNeeded; n++ {
		var num uint64
		for i := 0; i < 64; i++ {
			num = (num << 1) | b.getBit(idx)
			idx++
			if idx == endIdx {
				break
			}
		}
		res = append(res, num)
	}

	return res
}

func (b Bits) subset(index, length uint64) uint64 {
	var result uint64

	endIdx := index + length
	for iFor := index; iFor < endIdx; iFor++ {
		i := iFor / 64
		p := 63 - iFor%64
		v := (b[i] >> p) & 1

		result = (result << 1) | v
	}

	return result
}
