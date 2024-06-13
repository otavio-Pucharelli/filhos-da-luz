package forms

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestForm_valid tests the Valid method of the Form type
func TestForm_valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

// TestForm_Required tests the Required method of the Form type
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("got valid when required fields were missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", strings.NewReader(postData.Encode()))

	r.PostForm = postData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("got invalid when required fields were present")
	}
}

// TestForm_Has tests the Has method of the Form type
func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	if form.Has("whatever", r) {
		t.Error("form shows has field when it does not")
	}

	postData := url.Values{}
	postData.Add("a", "a")

	r = httptest.NewRequest("POST", "/whatever", strings.NewReader(postData.Encode()))
	r.PostForm = postData

	form = New(r.PostForm)
	if !form.Has("a", r) {
		t.Error("form shows no field when it does")
	}
}

// TestForm_IsEmail tests the IsEmail method of the Form type
func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for a non-email field")
	}

	postData := url.Values{}
	postData.Add("email", "a")
	r = httptest.NewRequest("POST", "/whatever", strings.NewReader(postData.Encode()))
	r.PostForm = postData
	form = New(r.PostForm)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid email for invalid email")
	}

	postData = url.Values{}
	postData.Add("email", "a@a.com")
	r = httptest.NewRequest("POST", "/whatever", strings.NewReader(postData.Encode()))
	r.PostForm = postData
	form = New(r.PostForm)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows invalid email for valid email")
	}
}
