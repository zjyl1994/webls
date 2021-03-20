package service

import (
	"io/ioutil"
	"sort"
	"strings"
	"time"

	prettytime "github.com/andanhm/go-prettytime"
	humanize "github.com/dustin/go-humanize"
	"github.com/zjyl1994/webls/config"
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
	TimeAgo string
}

func ListDir(path string) (result []FileItem, hasREADME bool, err error) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, false, err
	}

	fileInfos := make([]fileInfo, 0, len(infos))
	for _, v := range infos {
		name := v.Name()
		if name == config.ReadmeFilename {
			hasREADME = true
		}
		if !strings.HasPrefix(name, ".") {
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

	result = make([]FileItem, len(fileInfos))
	for i, v := range fileInfos {
		result[i] = FileItem{
			Name:    v.Name,
			IsDir:   v.IsDir,
			LastMod: v.LastMod.Format(timeFormat),
			Size:    humanize.Bytes(uint64(v.Size)),
			TimeAgo: prettytime.Format(v.LastMod),
		}
	}

	return result, hasREADME, nil
}
