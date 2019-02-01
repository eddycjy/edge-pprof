package savefile

import (
	"io"

	"github.com/EDDYCJY/edge-pprof/pkg/file"
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
)

type FileOutput struct {
	Path *profile.CompletePath
}

func NewFileOutput(path *profile.CompletePath) *FileOutput {
	return &FileOutput{
		Path: path,
	}
}

func (f *FileOutput) GetPzPb(body io.ReadCloser) error {
	fd, err := file.Open(f.Path.PbGz.SaveCompletePath, f.Path.PbGz.CompletePath, 0755)
	if err != nil {
		return err
	}
	defer fd.Close()
	io.Copy(fd, body)

	return nil
}

func (f *FileOutput) GetImage(output []byte) error {
	fd, err := file.Open(f.Path.Image.SaveCompletePath, f.Path.Image.CompletePath, 0755)
	defer fd.Close()
	_, err = fd.Write(output)
	if err != nil {
		return err
	}

	return nil
}
