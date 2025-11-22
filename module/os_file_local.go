package module

import (
	"errors"
	"os/user"
	"path/filepath"
)

func Get_file_path(path string) (string, error) {
	if path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return path, errors.New("User path load error!!!")
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}
