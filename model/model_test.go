package model

import "testing"

func TestRecordValidation(t *testing.T) {
	if NewRecord("1", "vase", "A", "B").Validate() != nil {
		t.Fatal("valid record rejected")
	}
	if NewRecord("", "", "", "").Validate() == nil {
		t.Fatal("invalid record accepted")
	}
}
