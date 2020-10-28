package binary

import (
	bin "encoding/binary"
	"io"
	"math"
)

func ReadByte(r io.Reader) (b byte, e error) {
	t := make([]byte, 1)
	n, e := r.Read(t)
	if n != 1 {
		return 0, ErrInvalidBytes
	}
	return t[0], e
}

func ReadBytes(r io.Reader, n int) (b []byte, e error) {
	t := make([]byte, n)
	d, e := r.Read(t)
	if d != n {
		return nil, ErrInvalidBytes
	}
	return t, e
}

func ReadString(r io.Reader, n int) (string, error) {
	b, e := ReadBytes(r, n)
	return string(b), e
}

func ReadRune(r io.Reader) (rune, error) {
	return ReadInt32(r)
}

func ReadUint8(r io.Reader) (uint8, error) {
	return ReadByte(r)
}

func ReadUint16(r io.Reader) (uint16, error) {
	b, e := ReadBytes(r, 2)
	return bin.LittleEndian.Uint16(b), e
}

func ReadUint32(r io.Reader) (uint32, error) {
	b, e := ReadBytes(r, 4)
	return bin.LittleEndian.Uint32(b), e
}

func ReadUint64(r io.Reader) (uint64, error) {
	b, e := ReadBytes(r, 8)
	return bin.LittleEndian.Uint64(b), e
}

func ReadInt8(r io.Reader) (int8, error) {
	b, e := ReadByte(r)
	return int8(b), e
}

func ReadInt16(r io.Reader) (int16, error) {
	b, e := ReadBytes(r, 2)
	return int16(bin.LittleEndian.Uint16(b)), e
}

func ReadInt32(r io.Reader) (int32, error) {
	b, e := ReadBytes(r, 4)
	return int32(bin.LittleEndian.Uint32(b)), e
}

func ReadInt64(r io.Reader) (int64, error) {
	b, e := ReadBytes(r, 8)
	return int64(bin.LittleEndian.Uint64(b)), e
}

func ReadFloat32(r io.Reader) (float32, error) {
	b, e := ReadBytes(r, 4)
	return math.Float32frombits(bin.LittleEndian.Uint32(b)), e
}

func ReadFloat64(r io.Reader) (float64, error) {
	b, e := ReadBytes(r, 8)
	return math.Float64frombits(bin.LittleEndian.Uint64(b)), e
}

func ReadUint8s(r io.Reader, n int) ([]uint8, error) {
	vs := make([]uint8, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadUint8(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadUint16s(r io.Reader, n int) ([]uint16, error) {
	vs := make([]uint16, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadUint16(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadUint32s(r io.Reader, n int) ([]uint32, error) {
	vs := make([]uint32, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadUint32(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadUint64s(r io.Reader, n int) ([]uint64, error) {
	vs := make([]uint64, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadUint64(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadInt8s(r io.Reader, n int) ([]int8, error) {
	vs := make([]int8, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadInt8(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadInt16s(r io.Reader, n int) ([]int16, error) {
	vs := make([]int16, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadInt16(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadInt32s(r io.Reader, n int) ([]int32, error) {
	vs := make([]int32, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadInt32(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadInt64s(r io.Reader, n int) ([]int64, error) {
	vs := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadInt64(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadFloat32s(r io.Reader, n int) ([]float32, error) {
	vs := make([]float32, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadFloat32(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadFloat64s(r io.Reader, n int) ([]float64, error) {
	vs := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadFloat64(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func ReadRunes(r io.Reader, n int) ([]rune, error) {
	vs := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		v, e := ReadInt32(r)
		if e != nil {
			return nil, e
		}
		vs = append(vs, v)
	}
	return vs, nil
}
