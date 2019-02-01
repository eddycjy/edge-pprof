package file

import (
	"fmt"
	"os"
)

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(savePath, completePath string, perm os.FileMode) (*os.File, error) {
	isPerm := CheckPermission(savePath)
	if isPerm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", savePath)
	}

	err := IsNotExistMkDir(savePath)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", savePath, err)
	}

	f, err := os.OpenFile(completePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
