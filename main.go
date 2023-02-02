package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"

	"howett.net/plist"
)

func HashFile(filename string) (string, error) {
	hasher := sha256.New()
	s, err := os.ReadFile(filename)
	hasher.Write(s)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func main() {
	icons, err := filepath.Glob("*.png")
	if err != nil {
		log.Fatalf("could not find *.png files: %v", err)
	}

	hashes := make(map[string]string)

	for _, icon := range icons {
		hash, err := HashFile(icon)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}
		hashes[icon] = hash
	}

	tempout, err := os.CreateTemp(".", "iconic-*.plist")
	if err != nil {
		log.Fatalf("could not open temp file: %v", err)
	}
	defer os.Remove(tempout.Name())

	encoder := plist.NewEncoderForFormat(tempout, plist.XMLFormat)
	encoder.Indent("\t")
	encoder.Encode(hashes)

	if err := tempout.Close(); err != nil {
		log.Fatalf("could not save temp results: %v", err)
	}

	// replace
	err = os.Rename(tempout.Name(), "_icons_hash.plist")
	if err != nil {
		log.Fatalf("could not replace _icons_hash.plist: %v", err)
	}

}
