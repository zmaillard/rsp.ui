package generator

import (
	"fmt"
	"os"
	"path"
)

func SaveItem(basePath string, v Generator) error {
	newFile := path.Join(basePath, v.OutFile())
	data, err := v.ToMarkdown()
	if err != nil {
		return err
	}

	outDir := path.Dir(newFile)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.MkdirAll(outDir, 0755)
	}
	err = os.WriteFile(newFile, data, 0755)

	fmt.Println(newFile)
	return err
}
