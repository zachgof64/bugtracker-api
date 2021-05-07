package gotils

import (
	"io/ioutil"
	"log"
	"os"
)

//GetByte will Return []byte version of the filepath given
func GetByte(filePath string) []byte {
	file, fileErr := os.Open(filePath)
	if (fileErr != nil) {
		log.Fatal(fileErr)
	}
	b, bErr := ioutil.ReadAll(file)
	if (bErr != nil){
		log.Fatal(bErr)
	}
	file.Close()
	return b
}

func FileDirExists(path string) bool {
	stat, _ := os.Stat(path)
	return !(stat == nil)
}