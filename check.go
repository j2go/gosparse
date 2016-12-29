package gosparse

import "syscall"

func IsSparseFile(paht string) (bool, error) {
	allocated, taken, err := FileSize(path)
	if err != nil {
		return false, err
	}
	return allocated > taken, nil
}

func FileSize(path string) (allocated int64, taken int64, err error) {
	s := syscall.Stat_t{}
	err = syscall.Stat(path, &s)
	if err != nil {
		return 0, 0, err
	}
	blockSize, err := FSBlockSize(path)
	if err != nil {
		return 0, 0, err
	}
	return s.Size, s.Blocks * int64(blockSize), nil
}

func FSBlockSize(path string) (blockSize int, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return 0, err
	}
	return int(s.Bsize), nil
}
