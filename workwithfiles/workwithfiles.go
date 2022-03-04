package workwithfiles

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func CalculateHash(target string) (string, error) {
	f, err := os.Open("tmp/" + target)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Println(err)
		return "", err
	}	

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
func ListDirectory() map[string]interface{} {
	result := map[string]interface{}{}
	files, err := ioutil.ReadDir("tmp")
	if err != nil {
		log.Println(err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filename := f.Name()
		hash, _ := CalculateHash(filename)
		result[filename] = hash
	}
	return result
}
func DeleteFile(name string) error {
	return os.Remove(name)
}
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
