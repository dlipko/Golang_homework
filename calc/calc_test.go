package main

import (
	"bytes"
	"testing"
)

func TestAdd(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "5"
	err := calc(out, "32+=")
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestSub(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "1"
	err := calc(out, "32-=")
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDiv(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "4"
	err := calc(out, "82/=")
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestMult(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "16"
	err := calc(out, "82*=")
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDiverseExpr(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "6"
	err := calc(out, "825*+132*+4-/=")
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestWrongSyntax(t *testing.T) {
	out := new(bytes.Buffer)
	err := calc(out, "82(5*+132*+4-/=")
	if err == nil {
		t.Errorf("results not match expected error")
	}
}
