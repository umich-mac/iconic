package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"howett.net/plist"
)

func HashFile(filename string) string {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filename)
	hasher.Write(s)
	if err != nil {
		fmt.Println(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	encoder := plist.NewEncoderForFormat(os.Stdout, plist.XMLFormat)
	encoder.Indent("\t")

	files, err := filepath.Glob("*.png")
	if err != nil {
		fmt.Println(err)
	}

	hashes := make(map[string]string)

	for _, file := range files {
		hashes[file] = HashFile(file)
	}

	encoder.Encode(hashes)
	fmt.Println("")

}
