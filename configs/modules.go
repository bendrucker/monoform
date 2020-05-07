package configs

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform/configs"
)

// RootModulesDirs walks the filesystem beginning at path and returns all root Terraform module directories it encounters.
// A module is assumed to be a root module if defines a backend.
func RootModulesDirs(path string) ([]string, error) {
	dirs := make([]string, 0)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		parser := configs.NewParser(nil)

		if !parser.IsConfigDir(path) {
			return nil
		}

		module, diags := parser.LoadConfigDir(path)
		if diags.HasErrors() {
			return diags
		}

		if module.Backend != nil {
			dirs = append(dirs, module.SourceDir)
		}

		return nil
	})

	return dirs, err
}
