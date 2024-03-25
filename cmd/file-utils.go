package main

import (
	"bufio"
	"strconv"
)

func (a *App) Generate() (string, error) {
	ret := ""
	var count int64
	var err error

	if a.C {
		count, err = a.ByteCount()
		if err != nil {
			return ret, err
		}
		ret += strconv.Itoa(int(count)) + " "
	}

	if a.L || a.W {
		lcount, wcount := a.OtherCounts()
		if a.L {
			ret += strconv.Itoa(lcount) + " "
		}
		if a.W {
			ret += strconv.Itoa(wcount) + " "
		}
	}

	ret += a.Fd.Name()
	return ret, nil
}

func (a *App) ByteCount() (int64, error) {
	fileInfo, err := a.Fd.Stat()
	if err != nil {
		return -1, err
	}
	bytes := fileInfo.Size()
	a.fileLen = int(bytes)
	return bytes, nil
}

// OtherCounts counts and returns number of lines and words
func (a *App) OtherCounts() (int, int) {
	//b := make([]byte, 1)
	fr := bufio.NewReader(a.Fd)

	var b, prev byte
	var err error
	wcount := 0
	lcount := 0

	for {
		b, err = fr.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			lcount++
		} else if b == ' ' && prev != ' ' {
			wcount++
		}
		prev = b
	}
	return lcount, wcount
}
