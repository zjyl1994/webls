package service

import (
	"io/ioutil"
	"sort"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type fileInfo struct {
	Name    string
	IsDir   bool
	LastMod time.Time
	Size    int64
}

type FileItem struct {
	Name    string
	IsDir   bool
	LastMod string
	Size    string
}

func ListDir(path string) ([]FileItem, error) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fileInfos := make([]fileInfo, 0, len(infos))
	for _, v := range infos {
		name := v.Name()
		if !strings.HasPrefix(name, ".") { //hide file by start .
			fileInfos = append(fileInfos, fileInfo{
				Name:    name,
				IsDir:   v.IsDir(),
				LastMod: v.ModTime(),
				Size:    v.Size(),
			})
		}
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		if fileInfos[i].IsDir != fileInfos[j].IsDir {
			return fileInfos[i].IsDir
		} else {
			return fileInfos[i].Name < fileInfos[j].Name
		}
	})

	result := make([]FileItem, len(fileInfos))
	for i, v := range fileInfos {
		result[i] = FileItem{
			Name:    v.Name,
			IsDir:   v.IsDir,
			LastMod: v.LastMod.Format(timeFormat),
			Size:    humanize.Bytes(uint64(v.Size)),
		}
	}

	return result, nil
}
