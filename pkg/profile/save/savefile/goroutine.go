package savefile

import (
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
	"strconv"
	"time"
)

type GoroutineFile struct{}

func NewGoroutineFile() *GoroutineFile {
	return &GoroutineFile{}
}

func (p *GoroutineFile) GetSuffix() string {
	return strconv.Itoa(int(time.Now().Local().Unix()))
}

func (p *GoroutineFile) GetSaveCompletePath() string {
	return GetRootPath() + profile.SEP + setting.ProfileRouteSetting.GoroutineSavePath
}

func (p *GoroutineFile) GetSavePath() string {
	return setting.ProfileRouteSetting.GoroutineSavePath
}

func (p *GoroutineFile) GetPrefix() string {
	return setting.ProfileFileSetting.BlockFilePrefix
}
