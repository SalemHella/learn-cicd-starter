package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey2(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: error ")
	}

}

func TestGetAPIKey1(t *testing.T) {
	header := http.Header{}

	header.Set("Authorization", "ApiKey testtest")

	result, err := GetAPIKey(header)

	if err != nil {
		t.Fatal("get an error where it shouldn't be ", err)
	}

	if result != "testtest" {
		t.Fatal("didn't get the expacted result instead get ", result)
	}
}
