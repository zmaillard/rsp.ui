package generator

import (
	"path"

	"github.com/spf13/afero"
)

type Lookup interface {
	GetLookup() ([]byte, error)
	OutLookupFiles() []string
}

func SaveLookup(appFs afero.Fs, basePath string, v Lookup) error {
	data, err := v.GetLookup()
	for _, f := range v.OutLookupFiles() {
		newFile := path.Join(basePath, f)
		if err != nil {
			return err
		}
		err = afero.WriteFile(appFs, newFile, data, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
