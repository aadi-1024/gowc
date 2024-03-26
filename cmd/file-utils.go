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
	//var count int64
	//var err error

	if a.L || a.W || a.C {
		lcount, wcount, ccount := a.Counts()
		if a.L {
			ret += strconv.Itoa(lcount) + " "
		}
		if a.W {
			ret += strconv.Itoa(wcount) + " "
		}
		if a.C {
			ret += strconv.Itoa(ccount) + " "
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

// Counts counts and returns number of lines, words and characters
func (a *App) Counts() (int, int, int) {
	//b := make([]byte, 1)
	fr := bufio.NewReader(a.Fd)

	var b, prev byte
	var err error
	wcount := 0
	lcount := 0
	ccount := 0
	prev = ' '

	for {
		b, err = fr.ReadByte()
		if err != nil {
			break
		}
		ccount++
		if b == '\n' {
			lcount++
		} else if !isSpace(b) && isSpace(prev) {
			wcount++
		}
		prev = b
	}
	a.fileLen = ccount
	return lcount, wcount, ccount
}

func (a *App) CharacterCount() (int, error) {
	if !a.C || !a.W || !a.L {
		_, _, f := a.Counts()
		a.fileLen = f
	}

	b := make([]byte, a.fileLen)
	//if Line and Word count generated, file pointer at end
	_, err := a.Fd.ReadAt(b, 0)
	if err != nil {
		return -1, err
	}
	return utf8.RuneCount(b), nil
}
