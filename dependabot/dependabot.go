package dependabot

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Dependabot struct {
	Version   int        `yaml:"version"` // Dependabot configuration files require this key, and its value must be 2
	Registrys []Registry `yaml:"updates"`
}

type Registry struct {
	Directory  string `yaml:"directory"`                // Location of package manifests
	PullsLimit int    `yaml:"open-pull-requests-limit"` // Limit number of open pull requests for version updates
	Rebase     string `yaml:"rebase-strategy"`          // Disable automatic rebasing. 'auto' is the default and Dependabot will rebase open pull requests when changes are detected. 'disabled' will disable automatic rebasing.
	Package    string `yaml:"package-ecosystem"`        // Package manager to use: `bundler`, `cargo`, `composer`, `devcontainers`, `docker`, `elm`, `gitsubmodule`, `github-actions`, `gomod`, `gradle`, `maven`, `mix`, `npm`, `nuget`, `pip`, `pip-compile`, `pub`, `swift`, `terraform`
	Schedule   struct {
		Interval string `yaml:"interval"` // `daily`, `weekly`, `monthly`
		Timezone string `yaml:"timezone"`
		Day      string `yaml:"day"`  // Specify an alternative day to check for updates
		Time     string `yaml:"time"` // Specify an alternative day to check for updates
	}
	Ignore []struct {
		Name        string   `yaml:"dependency-name"`
		Versions    []string `yaml:"versions"`
		UpdateTypes []string `yaml:"update-types"` // "version-update:semver-major", "version-update:semver-major", "version-update:semver-patch"
	} `yaml:"ignore"`
}

// Get registry to check and Update
func Parse(configFile io.Reader) ([]Registry, error) {
	var dep Dependabot
	if err := yaml.NewDecoder(configFile).Decode(&dep); err != nil {
		return nil, err
	}
	return dep.Registrys, nil
}
