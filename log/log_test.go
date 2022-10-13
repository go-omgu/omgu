package log

import (
	"os"
	"testing"
)

type Dog struct {
	Name string
	Age  int
}

type Man struct {
	Name string
	Pets []*Dog
}

func TestLog(t *testing.T) {
	log := New(os.Stderr, Normal)
	log.Log("test", "log")
	log.Verbose("test", "verbose")
	log.Debug("test", "debug")

	log = New(os.Stderr, Verbose)
	log.Log("test", "log")
	log.Verbose("test", "verbose")
	log.Debug("test", "debug")

	log = New(os.Stderr, Debug)
	log.Log("test", "log")
	log.Verbose("test", "verbose")
	log.Debug("test", "debug")
}

func TestDump(t *testing.T) {
	log := New(os.Stderr, Debug)
	log.Dump(struct {
		Name string
		Age  int
	}{"hello", 13})
	d := Dog{"Lucy", 14}
	log.Dump(d)

	mike := new(Man)
	mike.Name = "Mike"
	mike.Pets = []*Dog{&d}
	log.Dump(mike)
}
