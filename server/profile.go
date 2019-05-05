package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/edge-pprof/pkg/app"
	"github.com/EDDYCJY/edge-pprof/pkg/e"
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
	"log"
)

type Profile struct {
	PProf *PProf
}

func NewProfile() *Profile {
	return &Profile{PProf: &PProf{
		Service:    DefaultServiceInfo,
		Collection: DefaultCollectionInfo,
	}}
}

func (p *Profile) GetURL() string {
	log.Printf("url: %s", p.PProf.GetURL(setting.ProfileSetting.ProfileUrl))
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
		log.Printf("p.PProf.BindBasicData err: %v", err)
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
		log.Printf("save.NewSave err: %v", err)
		httpCode = http.StatusInternalServerError
		response.Set(e.PROFILE_SAVE_MODE_UNKNOWN_ERROR)
		return
	}

	statusCode, err := p.PProf.HanldePzPb(p, saver)
	if err != nil {
		log.Printf("p.PProf.HanldePzPb err: %v", err)
		httpCode = http.StatusInternalServerError
		response.Set(statusCode)
		return
	}

	statusCode, err = p.PProf.HandleImage(saver, []string{"-" + profile.SVG, path.PbGz.CompletePath})
	if err != nil {
		log.Printf("p.PProf.HandleImage err: %v", err)
		httpCode = http.StatusInternalServerError
		response.Set(statusCode)
		return
	}

	response.Data = p.PProf.Response(path)
	return
}
