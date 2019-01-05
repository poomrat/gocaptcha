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
