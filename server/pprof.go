package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/edge-pprof/pkg/e"
	"github.com/EDDYCJY/edge-pprof/pkg/profile"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save"
	"github.com/EDDYCJY/edge-pprof/pkg/profile/save/savefile"
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
)

type PProf struct {
	Service    *ServiceInfo
	Collection *CollectionInfo
}

type ServiceInfo struct {
	Name      string `form:"service_name" binding:"required"`
	Port      int64  `form:"service_port" binding:"required"`
	Namespace string `form:"namespace" `
}

type CollectionInfo struct {
	Seconds int32 `form:"collection_seconds"`
	Timeout int32 `form:"collection_timeout"`
}

type DebugPProf interface {
	GetURL() string
	Handle(c *gin.Context)
}

var (
	DefaultServiceInfo    = &ServiceInfo{}
	DefaultCollectionInfo = &CollectionInfo{}
	DefaultProfileFile    savefile.PPfile
	DefaultHeapFile       savefile.PPfile
	DefaultBlockFile      savefile.PPfile
	DefaultMutexFile      savefile.PPfile
)

func NewPProf() {
	DefaultServiceInfo = &ServiceInfo{
		Namespace: "default",
	}
	DefaultCollectionInfo = &CollectionInfo{
		Seconds: setting.ProfileSetting.DefaultSeconds,
		Timeout: setting.ProfileSetting.DefaultTimeout,
	}
	DefaultProfileFile = savefile.NewProfileFile()
	DefaultHeapFile = savefile.NewHeapFile()
	DefaultBlockFile = savefile.NewBlockFile()
	DefaultMutexFile = savefile.NewMutexFile()
}

func (p *PProf) GetPbGzCompletePath(ppf savefile.PPfile, ext string) *profile.Path {
	return savefile.GetCompletePath(ppf, p.Service.Name, ext)
}

func (p *PProf) GetImageCompletePath(ppf savefile.PPfile, ext string) *profile.Path {
	return savefile.GetCompletePath(ppf, p.Service.Name, ext)
}

func (p *PProf) GetURL(profilingUrl string) string {
	if p.Collection.Seconds > setting.ProfileSetting.MaxSeconds {
		p.Collection.Seconds = setting.ProfileSetting.MaxSeconds
	}

	return fmt.Sprintf(
		profilingUrl+setting.ProfileSetting.SuffixUrl,
		setting.ProfileSetting.Protocol,
		p.Service.Name,
		p.Service.Namespace,
		p.Service.Port,
		p.Collection.Seconds,
		p.Collection.Timeout,
	)
}

func (p *PProf) BindBasicData(c *gin.Context) error {
	err := c.ShouldBind(p.Service)
	if err != nil {
		return err
	}

	err = c.ShouldBind(p.Collection)
	if err != nil {
		return err
	}

	return nil
}

func (p *PProf) HanldePzPb(d DebugPProf, saver save.Save) (int, error) {
	pzpbResp, err := save.GetOriginPzPb(d.GetURL())
	if err != nil {
		return e.PROFILE_CREATE_PBGZ_ERROR, err
	}
	defer pzpbResp.Body.Close()

	err = saver.GetPzPb(pzpbResp.Body)
	if err != nil {
		return e.PROFILE_CREATE_PBGZ_ERROR, err
	}

	return 0, nil
}

func (p *PProf) HandleImage(saver save.Save, avgs []string) (int, error) {
	output, err := save.GetOriginImage(avgs)
	if err != nil {
		return e.PROFILE_CREATE_IMAGE_ERROR, err
	}

	err = saver.GetImage(output)
	if err != nil {
		return e.PROFILE_CREATE_IMAGE_ERROR, err
	}

	return 0, nil
}

func (p *PProf) Response(path *profile.CompletePath) map[string]string {
	shower := savefile.NewShow(path)
	return map[string]string{
		"pzpb_url":  shower.PzPb(),
		"image_url": shower.Image(),
	}
}
