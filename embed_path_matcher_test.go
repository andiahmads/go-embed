package goembed

import (
	"embed"
	"fmt"
	"testing"
)

/* patch matcher berfungsi untuk membaca multi file yang kita inginkan.
ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca.
caranya, kita perlu menggunakan path matcher seperti package function path.Math
*/

//go:embed files/*.text

var path embed.FS

func TestEmbedMulitFile(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("content:", string(content))
		}
	}
}
