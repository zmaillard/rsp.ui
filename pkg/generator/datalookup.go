package generator

import (
	"os"
	"path"
)

type Lookup interface {
	GetLookup() ([]byte, error)
	OutLookupFile() string
}

func SaveLookup(basePath string, v Lookup) error {
	newFile := path.Join(basePath, v.OutLookupFile())
	data, err := v.GetLookup()
	if err != nil {
		return err
	}
	err = os.WriteFile(newFile, data, 0755)

	return err
}
