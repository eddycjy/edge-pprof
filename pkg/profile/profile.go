package profile

import "os"

const (
	SEP  = string(os.PathSeparator)
	DOT  = "."
	PBGZ = "pb.gz"
	SVG  = "svg"
	TXT  = "txt"

	SAVE_FILE_MODE = 1
)

type Path struct {
	SaveCompletePath string
	SavePath         string
	FileName         string
	CompletePath     string
}

type CompletePath struct {
	PbGz  *Path
	Image *Path
}
