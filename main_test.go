package main

import (
	"strings"
	"testing"
)

func TestPipe1(t *testing.T) {
	result := pipe1()

	if result == "" {
		t.Error("Result is missing")
	}
	if !strings.Contains(result, "ABC") {
		t.Error("Result is not valid, no ABC")
	}
	t.Error("Sucks 1")
	t.Error("Sucks 2")
	//t.Log(result)
}

func TestVersion(t *testing.T) {
	result := buildVersion

	if result == "" {
		t.Error("BuildVersion is missing")
	}
	t.Log("BuildVersion = " + result)
}

func TestZapp(t *testing.T) {
	main()
	result := buildVersion
	t.Log("BuildVersion = " + result)
}
