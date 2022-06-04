package bitmap

import (
	"unsafe"
)

//go:noescape
func _x64count_avx2(data unsafe.Pointer, len uint64, result unsafe.Pointer)
func x64count_avx2(data []uint64) int {
	if len(data) == 0 {
		return 0
	}

	var res uint64
	_x64count_avx2(unsafe.Pointer(&data[0]), uint64(len(data)), unsafe.Pointer(&res))
	return int(res)
}
