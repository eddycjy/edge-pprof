package setting

import (
	"log"

	"github.com/go-ini/ini"

	"github.com/EDDYCJY/edge-pprof/pkg/bindata"
)

type App struct {
	AppName    string
	AppVersion string
	RunMode    string
}

type Server struct {
	HttpAddr string
	HttpPort string
}

type Profile struct {
	SaveMode       int32
	Protocol       string
	ProfileUrl     string
	HeapUrl        string
	BlockUrl       string
	MutexUrl       string
	SuffixUrl      string
	DefaultSeconds int32
	DefaultTimeout int32
	MaxSeconds     int32
}

type ProfileFile struct {
	ProfileFilePrefix string
	HeapFilePrefix    string
	BlockFilePrefix   string
	MutexFilePrefix   string
}

type ProfileFileStatic struct {
	FSProtocol     string
	FSDomain       string
	FSRelativePath string
}

type ProfileRoute struct {
	RootSavePath    string
	ProfileSavePath string
	HeapSavePath    string
	BlockSavePath   string
	MutexSavePath   string
}

var (
	Cfg                      *ini.File
	AppSetting               = &App{}
	ServerSetting            = &Server{}
	ProfileSetting           = &Profile{}
	ProfileFileSetting       = &ProfileFile{}
	ProfileFileStaticSetting = &ProfileFileStatic{}
	ProfileRouteSetting      = &ProfileRoute{}
)

func NewSetting() {
	var err error
	Cfg, err = ini.LoadSources(ini.LoadOptions{}, bindata.MustAsset("conf/app.ini"))
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatal(2, "Fail to map section 'app': %v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatal(2, "Fail to map section 'server': %v", err)
	}

	err = Cfg.Section("profile").MapTo(ProfileSetting)
	if err != nil {
		log.Fatal(2, "Fail to map section 'profile': %v", err)
	}

	err = Cfg.Section("profile.file").MapTo(ProfileFileSetting)
	if err != nil {
		log.Fatal(2, "Fail to map section 'profile.file': %v", err)
	}

	err = Cfg.Section("profile.file.static").MapTo(ProfileFileStaticSetting)
	if err != nil {
		log.Fatal(2, "Fail to map section 'profile.file.static': %v", err)
	}

	err = Cfg.Section("profile.route").MapTo(ProfileRouteSetting)
	if err != nil {
		log.Fatal(2, "Fail to map section 'profile.route': %v", err)
	}
}
