package captcha_test

import (
	"testing"

	captcha "github.com/poomrat/gocaptcha"
)

func Test1PlusOne(t *testing.T) {
	result := captcha.Generate(1, 1, 1, 1)
	expected := "1 + One"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func Test2PlusOne(t *testing.T) {
	result := captcha.Generate(1, 2, 1, 1)
	expected := "2 + One"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestOnePlus1(t *testing.T) {
	result := captcha.Generate(2, 1, 1, 1)
	expected := "One + 1"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestTwoPlus1(t *testing.T) {
	result := captcha.Generate(2, 2, 1, 1)
	expected := "Two + 1"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestTwoMinus1(t *testing.T) {
	result := captcha.Generate(2, 2, 2, 1)
	expected := "Two - 1"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func Test2MinusOne(t *testing.T) {
	result := captcha.Generate(1, 2, 2, 1)
	expected := "2 - One"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func Test7MultiplyOne(t *testing.T) {
	result := captcha.Generate(1, 7, 3, 1)
	expected := "7 * One"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestTwoMultiply9(t *testing.T) {
	result := captcha.Generate(2, 2, 3, 9)
	expected := "Two * 9"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
