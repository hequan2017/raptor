package response

import "raptor/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
