package main

import (
	"bufio"
	"strconv"
	"unicode/utf8"
)

func isSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\v' || b == '\f' || b == '\r'
}

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
		lcount, wcount := a.LineWordCounts()
		if a.L {
			ret += strconv.Itoa(lcount) + " "
		}
		if a.W {
			ret += strconv.Itoa(wcount) + " "
		}
	}

	if a.M {
		cc, err := a.CharacterCount()
		if err != nil {
			return ret, err
		}
		ret += strconv.Itoa(cc) + " "
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

// LineWordCounts counts and returns number of lines and words
func (a *App) LineWordCounts() (int, int) {
	//b := make([]byte, 1)
	fr := bufio.NewReader(a.Fd)

	var b, prev byte
	var err error
	wcount := 0
	lcount := 0
	prev = ' '

	for {
		b, err = fr.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			lcount++
		} else if !isSpace(b) && isSpace(prev) {
			wcount++
		}
		prev = b
	}
	return lcount, wcount
}

func (a *App) CharacterCount() (int, error) {
	if !a.C {
		f, err := a.ByteCount()
		if err != nil {
			return -1, err
		}
		a.fileLen = int(f)
	}

	b := make([]byte, a.fileLen)
	//if Line and Word count generated, file pointer at end
	_, err := a.Fd.ReadAt(b, 0)
	if err != nil {
		return -1, err
	}
	return utf8.RuneCount(b), nil
}
