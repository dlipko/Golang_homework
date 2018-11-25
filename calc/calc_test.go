package main

import (
	"bytes"
	"testing"
)

func TestAdd(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "5"
	calc(out, "32+=", false)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestSub(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "1"
	calc(out, "32-=", false)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDiv(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "4"
	calc(out, "82/=", false)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestMult(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "16"
	calc(out, "82*=", false)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDiverseExpr(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "6"
	calc(out, "825*+132*+4-/=", false)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestWrongSyntax(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "6"
	calc(out, "82(5*+132*+4-/=", false)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestDiverseNumberExpr(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "-26"
	calc(out, "82 25 34*+11 36 2*+48-/=", true)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}

func TestNumbersWrongSyntax(t *testing.T) {
	out := new(bytes.Buffer)
	expected := "-26"
	calc(out, "82 ^^25 (((34*)))+11$$$### 36 2*+48-/=", true)
	result := out.String()
	if result != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
	}
}
