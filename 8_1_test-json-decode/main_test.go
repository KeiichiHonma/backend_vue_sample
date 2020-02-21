package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	post,err := Decode("test.json")
	if err != nil{
		t.Error(err)
	}
	if post.Id != 1{
		t.Error("でコード間違い")
	}
	if post.Content != "こんにちは"{
		t.Error("でコード間違い")
	}
	assert.Equal(t,1,post.Id,"id間違い")
	assert.Equal(t,"こんにちは",post.Content,"content間違い")
}
func BenchmarkDecode(b *testing.B) {
	for i:=0;i < b.N ;i++  {
		Decode("test.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i:=0;i < b.N ;i++  {
		Unmarshal("test.json")
	}
}