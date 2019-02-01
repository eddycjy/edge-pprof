package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/edge-pprof/pkg/app"
	"github.com/EDDYCJY/edge-pprof/pkg/e"
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type Profile struct {
	PProf *PProf
}

func NewProfile() *Profile {
	return &Profile{PProf: &PProf{
		Service:    &ServiceInfo{},
		Collection: DefaultCollectionInfo,
	}}
}

func (p *Profile) GetURL() string {
	return p.PProf.GetURL(setting.ProfileSetting.ProfileUrl)
}

func (p *Profile) Handle(c *gin.Context) {
	var (
		httpCode = http.StatusOK
		response = app.NewResponse()
	)
	defer func() {
		c.JSON(httpCode, response)
	}()

	err := p.PProf.BindBasicData(c)
	if err != nil {
		httpCode = http.StatusBadRequest
		response.Set(e.INVALID_PARAMS)
		return
	}

	path := &profile.CompletePath{
		PbGz:  p.PProf.GetPbGzCompletePath(DefaultProfileFile, profile.PBGZ),
		Image: p.PProf.GetImageCompletePath(DefaultProfileFile, profile.SVG),
	}
	saver, err := save.NewSave(setting.ProfileSetting.SaveMode, path)
	if err != nil {
		httpCode = http.StatusInternalServerError
		response.Set(e.PROFILE_SAVE_MODE_UNKNOWN_ERROR)
		return
	}

	statusCode, err := p.PProf.HanldePzPb(p, saver)
	if err != nil {
		httpCode = http.StatusInternalServerError
		response.Set(statusCode)
		return
	}

	statusCode, err = p.PProf.HandleImage(saver, []string{"-" + profile.SVG, path.PbGz.CompletePath})
	if err != nil {
		httpCode = http.StatusInternalServerError
		response.Set(statusCode)
		return
	}

	response.Data = p.PProf.Response(path)
	return
}
