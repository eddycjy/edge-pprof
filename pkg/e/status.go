package e

const (
	SUCCESS        = 1
	FAIL           = 0
	INVALID_PARAMS = -1

	PROFILE_CREATE_PBGZ_ERROR       = 10001
	PROFILE_CREATE_IMAGE_ERROR      = 10002
	PROFILE_SAVE_MODE_UNKNOWN_ERROR = 10003
)

var statusText = map[int]string{
	SUCCESS:                         "ok",
	FAIL:                            "fail",
	INVALID_PARAMS:                  "Request parameter error",
	PROFILE_CREATE_PBGZ_ERROR:       "Generate profile pb.gz error",
	PROFILE_CREATE_IMAGE_ERROR:      "Generate profile image error",
	PROFILE_SAVE_MODE_UNKNOWN_ERROR: "Unknown Profile Mode error",
}

func StatusText(code int) string {
	return statusText[code]
}
