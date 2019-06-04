package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/edge-pprof/pkg/app"
	"github.com/EDDYCJY/edge-pprof/pkg/e"
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type Goroutine struct {
	Debug int `form:"debug"` // 1 or 2
	PProf *PProf
}

func NewGoroutine() *Goroutine {
	return &Goroutine{PProf: &PProf{
		Service:    DefaultServiceInfo,
		Collection: DefaultCollectionInfo,
	}}
}

func (h *Goroutine) GetURL() string {
	return h.PProf.GetURL(setting.ProfileSetting.GoroutineUrl)
}

func (h *Goroutine) GetDebugURL(debug int) string {
	return h.GetURL() + "&debug=" + strconv.Itoa(debug)
}

func (h *Goroutine) Handle(c *gin.Context) {
	var (
		httpCode = http.StatusOK
		response = app.NewResponse()
	)
	defer func() {
		c.JSON(httpCode, response)
	}()

	debug, err := strconv.Atoi(c.DefaultQuery("debug", "0"))
	if err != nil {
		log.Printf("h.PProf.strconv.Atoi err: %v", err)
		httpCode = http.StatusBadRequest
		response.Set(e.INVALID_PARAMS)
		return
	}
	h.Debug = debug

	err = h.PProf.BindBasicData(c)
	if err != nil {
		log.Printf("h.PProf.BindBasicData err: %v", err)
		httpCode = http.StatusBadRequest
		response.Set(e.INVALID_PARAMS)
		return
	}

	var (
		path       = &profile.CompletePath{}
		saver      save.Save
		statusCode int
	)
	if h.Debug == 0 {
		path = &profile.CompletePath{
			PbGz:  h.PProf.GetPbGzCompletePath(DefaultGoroutineFile, profile.PBGZ),
			Image: h.PProf.GetImageCompletePath(DefaultGoroutineFile, profile.SVG),
		}
		saver, err = save.NewSave(setting.ProfileSetting.SaveMode, path)
		if err != nil {
			log.Printf("save.NewSave err: %v", err)
			httpCode = http.StatusInternalServerError
			response.Set(e.PROFILE_SAVE_MODE_UNKNOWN_ERROR)
			return
		}

		statusCode, err = h.PProf.HanldePzPb(h.GetURL(), saver)
		if err != nil {
			log.Printf("h.PProf.HanldePzPb err: %v", err)
			httpCode = http.StatusInternalServerError
			response.Set(statusCode)
			return
		}
	} else {
		path = &profile.CompletePath{
			PbGz:  h.PProf.GetPbGzCompletePath(DefaultGoroutineFile, profile.TXT),
			Image: &profile.Path{},
		}
		saver, err = save.NewSave(setting.ProfileSetting.SaveMode, path)
		if err != nil {
			log.Printf("save.NewSave err: %v", err)
			httpCode = http.StatusInternalServerError
			response.Set(e.PROFILE_SAVE_MODE_UNKNOWN_ERROR)
			return
		}

		statusCode, err = h.PProf.HanldePzPb(h.GetDebugURL(h.Debug), saver)
		if err != nil {
			log.Printf("h.PProf.HanldePzPb err: %v", err)
			httpCode = http.StatusInternalServerError
			response.Set(statusCode)
			return
		}
	}

	if h.Debug == 0 {
		statusCode, err = h.PProf.HandleImage(saver, []string{"-" + profile.SVG, path.PbGz.CompletePath})
		if err != nil {
			log.Printf("h.PProf.HandleImage err: %v", err)
			httpCode = http.StatusInternalServerError
			response.Set(statusCode)
			return
		}
	}

	response.Data = h.PProf.Response(path)
	return
}
