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
	return setting.ProfileFileStaticSetting.FSProtocol + "://" + setting.ProfileFileStaticSetting.FSDomain + ":" + setting.ServerSetting.HttpPort + "/" + setting.ProfileFileStaticSetting.FSRelativePath + "/" + c.Path.PbGz.SavePath + "/" + c.Path.PbGz.FileName
}

func (c *Show) Image() string {
	return setting.ProfileFileStaticSetting.FSProtocol + "://" + setting.ProfileFileStaticSetting.FSDomain + ":" + setting.ServerSetting.HttpPort + "/" + setting.ProfileFileStaticSetting.FSRelativePath + "/" + c.Path.Image.SavePath + "/" + c.Path.Image.FileName
}
