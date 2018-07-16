package main

import (
	"testing"
)

func TestNotExist(t *testing.T) {
	notFound := "霓虹"
	document := getDocument(notFound)
	result := getResult(document)
	if result != nil {
		t.Error("it should be nil")
		t.Fail()
	}
}

func TestExist(t *testing.T) {
	word := "幽灵"
	document := getDocument(word)
	result := getResult(document)
	result.Show()

}

func TestMutiResult(t *testing.T) {
	word := "大根"
	document := getDocument(word)
	result := getResult(document)
	result.Show()
}

func TestSimilarity(t *testing.T) {
	word := "习"
	document := getDocument(word)
	result := getResult(document)
	if result != nil {
		t.Error("it should be nil")
		t.Fail()
	}
}

