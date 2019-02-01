package savefile

import (
	"strconv"
	"time"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type MutexFile struct{}

func NewMutexFile() *MutexFile {
	return &MutexFile{}
}

func (p *MutexFile) GetSuffix() string {
	return strconv.Itoa(int(time.Now().Local().Unix()))
}

func (p *MutexFile) GetSaveCompletePath() string {
	return GetRootPath() + profile.SEP + setting.ProfileRouteSetting.MutexSavePath
}

func (p *MutexFile) GetSavePath() string {
	return setting.ProfileRouteSetting.MutexSavePath
}

func (p *MutexFile) GetPrefix() string {
	return setting.ProfileFileSetting.MutexFilePrefix
}
