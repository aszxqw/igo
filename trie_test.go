package igo

import (
	//"fmt"
	"testing"
	//"reflect"
)

func TestTrie0(t *testing.T) {
	trie := NewTrie()
	trie.Insert("123")
	trie.Insert("abc")
	trie.Insert("hello")
	results := trie.Find("123 abc hello")
	if len(results) != 3 {
		t.Fatalf("len(%d) != 3", len(results))
	}
	if results[0].Pattern != "123" || results[1].Pattern != "abc" || results[2].Pattern != "hello" {
		t.Fatal(results)
	}
	if results[0].Data != true || results[1].Data != true || results[2].Data != true {
		t.Fatal(results)
	}
}

type TestTrie1Data struct {
	Str string
	Cnt int
}

func TestTrie1(t *testing.T) {
	var data TestTrie1Data
	trie := NewTrie()
	data.Str = "hehe"
	data.Cnt = 0
	trie.Insert("123", data)
	data.Cnt = 1
	trie.Insert("abc", data)
	data.Cnt = 2
	trie.Insert("hello", data)
	results := trie.Find("123 abc hello")
	if len(results) != 3 {
		t.Fatalf("len(%d) != 3", len(results))
	}
	for i := 0; i < len(results); i++ {
		if results[i].Data.(TestTrie1Data).Cnt != i {
			t.Fatal("error")
		}
	}
}

func TestTrie2(t *testing.T) {
	trie := NewTrie()
	trie.Insert("123")
	trie.Insert("abc")
	trie.Insert("你好")
	results := trie.Find("123 abc hello你好")
	if len(results) != 3 {
		t.Fatalf("len(%d) != 3", len(results))
	}
	if results[0].Pattern != "123" || results[1].Pattern != "abc" || results[2].Pattern != "你好" {
		t.Fatal(results)
	}
}
