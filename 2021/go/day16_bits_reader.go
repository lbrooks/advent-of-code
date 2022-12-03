package main

type BitsReader struct {
	data   Bits
	length uint64
	index  uint64
}

func NewBitsReader(bits Bits, length uint64) *BitsReader {
	return &BitsReader{
		data:   bits,
		length: length,
		index:  0,
	}
}

func (r *BitsReader) Read() uint64 {
	res := r.data.getBit(r.index)
	r.index++
	return res
}

func (r *BitsReader) ReadBits(num uint64) uint64 {
	if num > 64 {
		return 0
	}
	if (r.index + num) > r.length {
		return 0
	}
	var val uint64
	for i := uint64(0); i < num; i++ {
		val = val<<1 | r.Read()
	}
	return val
}

func (r *BitsReader) HasNext() bool {
	return r.index < r.length
}
