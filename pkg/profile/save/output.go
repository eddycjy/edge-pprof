package save

import (
	"net/http"
	"os/exec"

	"github.com/EDDYCJY/edge-pprof/pkg/ehttp"
)

func GetOriginPzPb(profileUrl string) (*http.Response, error) {
	resp, err := ehttp.GetClient().Get(profileUrl)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetOriginImage(avgs []string) ([]byte, error) {
	tools := []string{"tool", "pprof"}
	tools = append(tools, avgs...)
	cmd := exec.Command("go", tools...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
