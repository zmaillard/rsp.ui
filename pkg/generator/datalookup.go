package generator

import (
	"os"
	"path"
)

type Lookup interface {
	GetLookup() ([]byte, error)
	OutLookupFiles() []string
}

func SaveLookup(basePath string, v Lookup) error {
	data, err := v.GetLookup()
	for _, f := range v.OutLookupFiles() {
		newFile := path.Join(basePath, f)
		if err != nil {
			return err
		}
		err = os.WriteFile(newFile, data, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
