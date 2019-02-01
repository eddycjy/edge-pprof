package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save/savefile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
	"github.com/EDDYCJY/edge-pprof/server"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if setting.ProfileSetting.SaveMode == profile.SAVE_FILE_MODE {
		r.StaticFS(setting.ProfileFileStaticSetting.FSRelativePath, http.Dir(savefile.GetRootPath()))
	}

	apiv1 := r.Group("/api/v1")
	apiv1.GET("debug/pprof/profile", server.NewProfile().Handle)
	apiv1.GET("debug/pprof/heap", server.NewHeap().Handle)
	apiv1.GET("debug/pprof/block", server.NewBlock().Handle)
	apiv1.GET("debug/pprof/mutex", server.NewMutex().Handle)
	return r
}
