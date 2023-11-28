package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

func main() {
	const pngHeader = "\x89PNG\r\n\x1a\n"
	data := make([]byte, 13)
	binary.BigEndian.PutUint32(data[0:4], 800)
	binary.BigEndian.PutUint32(data[4:8], 600)
	// depth and color type
	data[8] = 8
	data[9] = 6
	// compression type
	data[10] = 0
	data[11] = 0
	data[12] = 0

	name := "IHDR"

	head := make([]byte, 8)
	foot := make([]byte, 4)
	n := uint32(len(data))
	binary.BigEndian.PutUint32(head[:4], n)
	head[4] = name[0]
	head[5] = name[1]
	head[6] = name[2]
	head[7] = name[3]
	crc := crc32.NewIEEE()
	crc.Write(head[4:8])
	crc.Write(data)
	binary.BigEndian.PutUint32(foot[:4], crc.Sum32())

	header := make([]byte, 12)
	header = append(header, head...)
	header = append(header, foot...)
	fmt.Printf("header: %v\n", header)
}
