package goembed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed testing.text
var version string
/* 
comand diatas, untuk membaca file dari testing string
 */
func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}
