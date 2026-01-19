package imagetoterminal

import (
	"encoding/base64"
	"os"
)

func RunParse() {
	println("in parser")

	file, err := os.OpenFile("wifey.png", os.O_RDONLY, 0664)
	if err != nil {
		println("error loading file", err.Error())
	}
	file_data := make([]byte, 2048)

	file.Read(file_data)
	base := base64.StdEncoding.EncodeToString(file_data)
	println(base)
}
