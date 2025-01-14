package utils

import (
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

const (
	E_DB_LINK        = "DATABASE_DEFAULT_LINK"
	E_PROJECT_OUTPUT = "PROJECT_OUTPUT"
	E_PROJECT_UPLOAD = "PROJECT_UPLOAD"
	E_PROJECT_SM4KEY = "PROJECT_SM4KEY"
)

var (
	DbLink = getEnvWithDefault(E_DB_LINK)
	Output = getEnvWithDefault(E_PROJECT_OUTPUT)
	Upload = getEnvWithDefault(E_PROJECT_UPLOAD)
	Sm4Key = getEnvWithDefault(E_PROJECT_SM4KEY)
)

func getEnvWithDefault(envKey string) string {
	if envStr := genv.Get(envKey); envStr != nil {
		return envStr.String()
	}
	if IsDocker() {
		return ""
	}
	r, _ := g.Cfg().Get(gctx.GetInitCtx(), convertEnvToConfig(envKey))
	return r.String()
}

func convertEnvToConfig(env string) string {
	return strings.ToLower(strings.ReplaceAll(env, "_", "."))
}
