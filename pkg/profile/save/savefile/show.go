package savefile

import (
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type Show struct {
	Path *profile.CompletePath
}

func NewShow(path *profile.CompletePath) *Show {
	return &Show{
		Path: path,
	}
}

func (c *Show) PzPb() string {
	if c.Path.PbGz.SavePath == "" {
		return ""
	}

	host := setting.ProfileFileStaticSetting.FSProtocol + "://" + setting.ProfileFileStaticSetting.FSDomain
	path := "/" + setting.ProfileFileStaticSetting.FSRelativePath + "/" + c.Path.PbGz.SavePath + "/" + c.Path.PbGz.FileName
	if setting.ProfileFileStaticSetting.FSPort == "" {
		return host + path
	}

	return host + ":" + setting.ProfileFileStaticSetting.FSPort + path
}

func (c *Show) Image() string {
	if c.Path.Image.SavePath == "" {
		return ""
	}

	host := setting.ProfileFileStaticSetting.FSProtocol + "://" + setting.ProfileFileStaticSetting.FSDomain
	path := "/" + setting.ProfileFileStaticSetting.FSRelativePath + "/" + c.Path.Image.SavePath + "/" + c.Path.Image.FileName
	if setting.ProfileFileStaticSetting.FSPort == "" {
		return host + path
	}

	return host + ":" + setting.ProfileFileStaticSetting.FSPort + path
}
