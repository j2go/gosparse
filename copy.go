package gosparse

import (
	"io"
	"os"
)

const ioBufferSize int = 4 * 1024
const ioReadBufferSize int = 100 * ioBufferSize

func Copy(src, dest string) error {
	isSparse, err := IsSparseFile(src)
	if err != nil {
		return err
	}
	if isSparse {
		return sparseCopy(src, dest)
	}
	return normalCopy(src, dest)
}

func normalCopy(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = io.Copy(destFile, srcFile)
	return err
}

func sparseCopy(src, dest string) error {
	return nil
}
