package utils

import (
	"hash/crc32"
	"strconv"
)

func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

//计算key的插槽.
func Slot(key string) uint32 {
	return CRC32(key) % 3
}

//ParseInt.
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return n
}
