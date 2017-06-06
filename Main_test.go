package main

import (
	"testing"
)

func TestStringchecker(t *testing.T) {
	v := Stringchecker("")
	if v != "0" {
		t.Error("Expected 0, got ", v)
	}
	v = Stringchecker("Hello!")
	if v != "Hello!" {
		t.Error("Expected Hello!, got ", v)
	}
}

func TestConnectdb(t *testing.T) {
	tdb := Connectdb()
	if tdb == nil {
		t.Error("Expected tdb to be not nil, got ", tdb)
		return
	}
	err := tdb.Close()
	if err != nil {
		t.Error("Problem closing connection while tdb != nil. Error: ", err)
	}
}
