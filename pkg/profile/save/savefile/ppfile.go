package savefile

import (
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type PPfile interface {
	GetSuffix() string
	GetSaveCompletePath() string
	GetSavePath() string
	GetPrefix() string
}

func GetCompletePath(p PPfile, serviceName, ext string) *profile.Path {
	path := &profile.Path{}
	path.FileName = GetFileName(p, serviceName, ext)
	path.SavePath = p.GetSavePath()
	path.SaveCompletePath = p.GetSaveCompletePath()
	path.CompletePath = path.SaveCompletePath + profile.SEP + path.FileName
	return path
}

func GetFileName(p PPfile, serviceName, ext string) string {
	fileName := p.GetPrefix() + "." + serviceName + "." + p.GetSuffix() + profile.DOT + ext
	if v := p.GetPrefix(); v == "" {
		fileName = serviceName + "." + p.GetSuffix() + profile.DOT + ext
	}

	return fileName
}

func GetRootPath() string {
	return setting.ProfileRouteSetting.RootSavePath
}
