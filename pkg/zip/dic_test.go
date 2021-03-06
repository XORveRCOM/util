package zip

import (
	"testing"

	testutil "github.com/xorvercom/util/pkg/testing"
)

func TestDic(t *testing.T) {
	// OpenDictionary
	dic, e := OpenDictionary(testutil.GetFilepath("../../test/notexists.zip"), true)
	if e == nil {
		t.Fail()
	}
	dic, e = OpenDictionary(testutil.GetFilepath("../../test/test.zip"), false)
	if e != nil {
		t.Fatal(e)
	}
	defer dic.Close()

	if dic.IsIgnoreCapital() {
		t.Fatal()
	}

	// FilePaths
	for _, filepath := range dic.FilePaths() {
		t.Logf("%+v\n", filepath)
	}

	// Contains
	if dic.Contains("notexists.txt") {
		t.Fail()
	}
	if dic.Contains("test1.txt") {
		t.Fail()
	}
	if false == dic.Contains("Test1.txt") {
		t.Fail()
	}

	// Get
	_, e = dic.Get("notexists.txt")
	if e == nil {
		t.Fail()
	}
	_, e = dic.Get("test1.txt")
	if e == nil {
		t.Fatal(e)
	}
	b, e := dic.Get("Test1.txt")
	if e != nil {
		t.Fatal(e)
	}
	s := string(b)
	if s != "test1.txt\r\n" {
		t.Fatal(s, b)
	}
	b, e = dic.Get("empty.txt")
	if e != nil {
		t.Fatal(e)
	}
	s = string(b)
	if s != "" {
		t.Fatal(s, b)
	}
}

func TestDicIgnoreCapital(t *testing.T) {
	// OpenDictionary
	dic, e := OpenDictionary(testutil.GetFilepath("../../test/notexists.zip"), true)
	if e == nil {
		t.Fail()
	}
	dic, e = OpenDictionary(testutil.GetFilepath("../../test/test.zip"), true)
	if e != nil {
		t.Fatal(e)
	}
	defer dic.Close()

	if false == dic.IsIgnoreCapital() {
		t.Fatal()
	}

	// FilePaths
	for _, filepath := range dic.FilePaths() {
		t.Logf("%+v\n", filepath)
	}

	// Contains
	if dic.Contains("notexists.txt") {
		t.Fail()
	}
	if false == dic.Contains("test1.txt") {
		t.Fail()
	}
	if false == dic.Contains("Test1.txt") {
		t.Fail()
	}

	// Get
	_, e = dic.Get("notexists.txt")
	if e == nil {
		t.Fail()
	}
	_, e = dic.Get("test1.txt")
	if e != nil {
		t.Fatal(e)
	}
	b, e := dic.Get("Test1.txt")
	if e != nil {
		t.Fatal(e)
	}
	s := string(b)
	if s != "test1.txt\r\n" {
		t.Fatal(s, b)
	}
	b, e = dic.Get("empty.txt")
	if e != nil {
		t.Fatal(e)
	}
	s = string(b)
	if s != "" {
		t.Fatal(s, b)
	}
}
