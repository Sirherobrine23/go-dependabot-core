package npm

import (
	"io/fs"
	"os"

	"github.com/vaughan0/go-ini"
)

type NPM struct {
	RootFolder fs.FS
	NpmRC      ini.File
}

type Package struct {
	Name                 string            `json:"name"` // Package name
	Dependencies         map[string]string `json:"dependencies"`
	DevDependencies      map[string]string `json:"devDependencies"`
	PeerDependencies     map[string]string `json:"peerDependencies"`
	OptionalDependencies map[string]string `json:"optionalDependencies"`
}

func (npm *NPM) Config() error {
	file, err := npm.RootFolder.Open(".npmrc")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()
	if npm.NpmRC, err = ini.Load(file); err != nil {
		return err
	}
	return nil
}

func (npm *NPM) Check() bool {
	return false
}

func (npm *NPM) Update() error {
  return nil
}