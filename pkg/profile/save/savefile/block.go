package savefile

import (
	"strconv"
	"time"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type BlockFile struct{}

func NewBlockFile() *BlockFile {
	return &BlockFile{}
}

func (p *BlockFile) GetSuffix() string {
	return strconv.Itoa(int(time.Now().Local().Unix()))
}

func (p *BlockFile) GetSaveCompletePath() string {
	return GetRootPath() + profile.SEP + setting.ProfileRouteSetting.BlockSavePath
}

func (p *BlockFile) GetSavePath() string {
	return setting.ProfileRouteSetting.BlockSavePath
}

func (p *BlockFile) GetPrefix() string {
	return setting.ProfileFileSetting.BlockFilePrefix
}
