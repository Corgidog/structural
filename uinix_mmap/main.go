package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

const n = 1e3

func main() {
	f, err := os.OpenFile("./data.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	length := int(fi.Size())
	fmt.Printf("length:%d\n", length)

	b, err := unix.Mmap(int(f.Fd()), 0, 1<<20, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(b), cap(b))
	b[0] = '1'
	b[length-1] = 'q'
	b[length] = 'f'

	for i, v := range "i love you forever" {
		b[length+i] = byte(v)
	}

	unix.Msync(b, unix.MS_SYNC)

	f.Seek(int64(length), 0)
	f.WriteString("\nappend content")

	fmt.Println(string(b[0 : 30+length]))

	err = unix.Munmap(b)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(b))
}
