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

	lcount, wcount, ccount, rcount := a.Counts()
	if a.L {
		ret += strconv.Itoa(lcount) + " "
	}
	if a.W {
		ret += strconv.Itoa(wcount) + " "
	}
	if a.C {
		ret += strconv.Itoa(ccount) + " "
	}
	if a.M {
		ret += strconv.Itoa(rcount) + " "
	}

	ret += a.Fd.Name()
	return ret, nil
}

// Counts counts and returns number of lines, words, characters and bytes
func (a *App) Counts() (int, int, int, int) {
	fr := bufio.NewReader(a.Fd)
	buf := make([]byte, a.Max_buf)

	var prev byte
	wcount := 0
	lcount := 0
	ccount := 0
	rcount := 0
	prev = ' '

	for {
		n, err := fr.Read(buf)
		ccount += n
		rcount += utf8.RuneCount(buf[0:n])
		for i := 0; i < n; i++ {
			if buf[i] == '\n' {
				lcount++
			}
			if !isSpace(buf[i]) && isSpace(prev) {
				wcount++
			}
			prev = buf[i]
		}
		if err != nil {
			break
		}

	}
	return lcount, wcount, ccount, rcount
}
