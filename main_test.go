package main

import (
	"os"
	"testing"
)

func TestGetBooklogInfo(t *testing.T) {
	os.Setenv("ACCOUNT_ID", "vtryo")
	account := getBooklogInfo().Tana.Account
	name := getBooklogInfo().Tana.Name

	accountExpected := "vtryo"
	nameExpected := "Ryoの本棚"

	if account != accountExpected {
		t.Errorf("got: %v\nwant: %v", account, accountExpected)
	}
	if name != nameExpected {
		t.Errorf("got: %v\nwant: %v", name, nameExpected)
	}
}
