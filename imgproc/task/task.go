package task

import (
	"fmt"
	"path/filepath"
)

type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFileList(srcDir string) []string {
	files, err := filepath.Glob(srcDir + "/*.jpg")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}
