package generator

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/afero"
)

func SaveItem(fs afero.Fs, basePath string, v Generator) error {
	newFile := path.Join(basePath, v.OutFile())
	data, err := v.ToMarkdown()
	if err != nil {
		return err
	}

	outDir := path.Dir(newFile)
	if _, err := fs.Stat(outDir); os.IsNotExist(err) {
		fs.MkdirAll(outDir, 0755)
	}
	err = afero.WriteFile(fs, newFile, data, 0755)

	fmt.Println(newFile)
	return err
}
