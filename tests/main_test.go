package tests

import (
	"os/exec"
	"strings"
	"testing"
)

func TestByteCount(t *testing.T) {
	cmd := exec.Command("../build", "-c", "../test.txt")

	out, err := cmd.Output()
	if err != nil {
		t.Error(err.Error())
	}
	s := string(out)
	bef, _, _ := strings.Cut(s, " ")
	if bef != "342190" {
		t.Fail()
	}
}

func TestLineCount(t *testing.T) {
	cmd := exec.Command("../build", "-l", "../test.txt")

	out, err := cmd.Output()
	if err != nil {
		t.Error(err.Error())
	}
	s := string(out)
	bef, _, _ := strings.Cut(s, " ")
	if bef != "7145" {
		t.Fail()
	}
}

func TestWordCount(t *testing.T) {
	cmd := exec.Command("../build", "-w", "../test.txt")

	out, err := cmd.Output()
	if err != nil {
		t.Error(err.Error())
	}
	s := string(out)
	bef, _, _ := strings.Cut(s, " ")
	if bef != "58164" {
		t.Fail()
	}
}
