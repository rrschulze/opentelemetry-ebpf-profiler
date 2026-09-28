package pfunsafe

import (
	"encoding/binary"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

var nativeEndian binary.ByteOrder

func init() {
	// Detect native byte order at runtime.
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)
	if buf[0] == 0xAB {
		nativeEndian = binary.BigEndian
	} else {
		nativeEndian = binary.LittleEndian
	}
}

func TestSliceFromPointer(t *testing.T) {
	s := uint64(0xcafebabe)
	p := &s
	actual := FromPointer(p)
	expected := make([]byte, 8)
	nativeEndian.PutUint64(expected, 0xcafebabe)
	assert.Equal(t, expected, actual)
	assert.Panics(t, func() {
		p = nil
		FromPointer(p)
	})
}

func TestSliceFromSlice(t *testing.T) {
	s := []uint64{0xcafebabe, 0xdeadbeef}
	actual := FromSlice(s)
	expected := make([]byte, 16)
	nativeEndian.PutUint64(expected[0:], 0xcafebabe)
	nativeEndian.PutUint64(expected[8:], 0xdeadbeef)
	assert.Equal(t, expected, actual)
	assert.NotPanics(t, func() {
		s = nil
		actual = FromSlice(s)
		expected = nil
		assert.Equal(t, expected, actual)
		s = []uint64{}
		actual = FromSlice(s)
		assert.Equal(t, expected, actual)
	})
}

func TestByteSlice2String(t *testing.T) {
	var b [4]byte
	s := ToString(b[:1]) // create s with length 1 and a 0 byte inside
	assert.Equal(t, "\x00", s)

	b[0] = 'a'
	assert.Equal(t, "a", s)
}
