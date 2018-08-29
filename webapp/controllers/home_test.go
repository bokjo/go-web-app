package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestLoginExecutesCorrectTemplate(t *testing.T) {
	h := new(home)

	expected := `login template`

	h.loginTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	h.handleLogin(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	fmt.Printf("[TEST]: %#v", w)

	if string(actual) != expected {
		t.Errorf("Failed to execute correct login template\n Expected: `%s` - got: `%s`\n", expected, actual)
	}
}
