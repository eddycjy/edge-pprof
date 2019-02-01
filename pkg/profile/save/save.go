package save

import (
	"fmt"
	"io"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save/savefile"
)

type Save interface {
	GetPzPb(body io.ReadCloser) error
	GetImage(output []byte) error
}

func NewSave(mode int32, path *profile.CompletePath) (Save, error) {
	var (
		saver Save
		err   error
	)

	switch mode {
	case 1:
		saver = savefile.NewFileOutput(path)
	default:
		err = fmt.Errorf("save.NewSave mode: %d, path: %v", mode, path)
	}

	return saver, err
}
