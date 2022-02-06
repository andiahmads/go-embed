package goembed

import (
	_ "embed"

	"io/fs"
	"io/ioutil"
	"testing"
)

/*
selain tipe data string, embed file juga bisa dilakukan kevaribel tipe data []byte
ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain2. */

//go:embed manjaro.png
var logo []byte

func TestEmbedArray(t *testing.T) {
	err := ioutil.WriteFile("manjaro_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
