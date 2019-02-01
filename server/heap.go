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

type Heap struct {
	Type  string `form:"type"` // -inuse_space or -alloc_objects
	PProf *PProf
}

func NewHeap() *Heap {
	return &Heap{PProf: &PProf{
		Service:    &ServiceInfo{},
		Collection: DefaultCollectionInfo,
	}}
}

func (h *Heap) GetURL() string {
	return h.PProf.GetURL(setting.ProfileSetting.HeapUrl)
}

func (h *Heap) Handle(c *gin.Context) {
	var (
		httpCode = http.StatusOK
		response = app.NewResponse()
	)
	defer func() {
		c.JSON(httpCode, response)
	}()

	h.Type = c.DefaultQuery("type", "inuse_space")

	err := h.PProf.BindBasicData(c)
	if err != nil {
		httpCode = http.StatusBadRequest
		response.Set(e.INVALID_PARAMS)
		return
	}

	path := &profile.CompletePath{
		PbGz:  h.PProf.GetPbGzCompletePath(DefaultHeapFile, profile.PBGZ),
		Image: h.PProf.GetImageCompletePath(DefaultHeapFile, profile.SVG),
	}
	saver, err := save.NewSave(setting.ProfileSetting.SaveMode, path)
	if err != nil {
		httpCode = http.StatusInternalServerError
		response.Set(e.PROFILE_SAVE_MODE_UNKNOWN_ERROR)
		return
	}

	statusCode, err := h.PProf.HanldePzPb(h, saver)
	if err != nil {
		httpCode = http.StatusInternalServerError
		response.Set(statusCode)
		return
	}

	statusCode, err = h.PProf.HandleImage(saver, []string{"-" + profile.SVG, "-" + h.Type, path.PbGz.CompletePath})
	if err != nil {
		httpCode = http.StatusInternalServerError
		response.Set(statusCode)
		return
	}

	response.Data = h.PProf.Response(path)
	return
}
