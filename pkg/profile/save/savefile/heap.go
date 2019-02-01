package savefile

import (
	"strconv"
	"time"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type HeapFile struct{}

func NewHeapFile() *HeapFile {
	return &HeapFile{}
}

func (p *HeapFile) GetSuffix() string {
	return strconv.Itoa(int(time.Now().Local().Unix()))
}

func (p *HeapFile) GetSaveCompletePath() string {
	return GetRootPath() + profile.SEP + setting.ProfileRouteSetting.HeapSavePath
}

func (p *HeapFile) GetSavePath() string {
	return setting.ProfileRouteSetting.HeapSavePath
}

func (p *HeapFile) GetPrefix() string {
	return setting.ProfileFileSetting.HeapFilePrefix
}
