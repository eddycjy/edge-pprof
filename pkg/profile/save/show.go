package save

import (
	"fmt"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save/savefile"
)

type Show interface {
	PzPb() string
	Image() string
}

func NewShow(mode int32, path *profile.CompletePath) (Show, error) {
	var (
		shower Show
		err    error
	)

	switch mode {
	case 1:
		shower = savefile.NewShow(path)
	default:
		err = fmt.Errorf("save.NewShow mode: %d, path: %v", mode, path)
	}

	return shower, err
}
