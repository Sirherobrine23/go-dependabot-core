package godependabotcore

import (
	"errors"
	"io/fs"

	"sirherobrine23.org/Sirherobrine23/go-dependabot-core/dependabot"
	"sirherobrine23.org/Sirherobrine23/go-dependabot-core/npm"
)

var (
	ErrInvalidPackageEco error = errors.New("invalid package ecossystem")
)

type Updater interface {
	// Check update avaible
	Check() bool
	// Run updates and return diff
	Update() error
}

func GetUpdater(reg dependabot.Registry, rootDir fs.FS) (Updater, error) {
	switch reg.Package {
	case "npm":
		return &npm.NPM{
			RootFolder: rootDir,
		}, nil
	}
	return nil, ErrInvalidPackageEco
}
