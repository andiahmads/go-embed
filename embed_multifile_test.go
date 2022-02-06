package goembed

import (
	"embed"
	"fmt"
	"testing"
)

/*
terkadang terdapat sebuah case kita ingin mengEmbed beberapa file sekaligu.
hal ini bisa dilakukan menggunakan package embed.
kita bisa menambahkan komentar //go:embed lebih dari 1 baris
selain itu variabelnya bisa kita gunakan tipe data embed.f5
*/

//go:embed files/embed_multifile1.text
//go:embed files/embed_multifile2.text
var files embed.FS

func TestEmbedMultiFIle(t *testing.T) {
	a, _ := files.ReadFile("files/embed_multifile1.text")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/embed_multifile2.text")
	fmt.Println(string(b))

}
