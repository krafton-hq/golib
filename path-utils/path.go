package path_utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ResolvePathToAbs(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = strings.Replace(path, "~", home, 1)
	}
	return filepath.Abs(path)
}

func IsBinaryExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func ExistsFile(absPath string) bool {
	_, err := os.Stat(absPath)
	return !errors.Is(err, os.ErrNotExist)
}
