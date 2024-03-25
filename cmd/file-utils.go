package main

import "strconv"

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

	ret += a.Fd.Name()
	return ret, nil
}

func (a *App) ByteCount() (int64, error) {
	fileInfo, err := a.Fd.Stat()
	if err != nil {
		return -1, err
	}
	bytes := fileInfo.Size()
	return bytes, nil
}
