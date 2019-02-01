package savefile

import (
	"strconv"
	"time"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type ProfileFile struct{}

func NewProfileFile() *ProfileFile {
	return &ProfileFile{}
}

func (p *ProfileFile) GetSuffix() string {
	return strconv.Itoa(int(time.Now().Local().Unix()))
}

func (p *ProfileFile) GetSaveCompletePath() string {
	return GetRootPath() + profile.SEP + setting.ProfileRouteSetting.ProfileSavePath
}

func (p *ProfileFile) GetSavePath() string {
	return setting.ProfileRouteSetting.ProfileSavePath
}

func (p *ProfileFile) GetPrefix() string {
	return setting.ProfileFileSetting.ProfileFilePrefix
}
