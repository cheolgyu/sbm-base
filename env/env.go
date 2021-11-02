package env

import (
	"github.com/cheolgyu/stock-write-common/logging"
	"github.com/joho/godotenv"
)

var MyEnv map[string]string

func init() {
	var err error
	MyEnv, err = godotenv.Read(".env.local")
	if err != nil {
		logging.Log.Panic("env 파일읽기 실패:", err)
	}

	logging.Log.Debug("env init 실행", MyEnv)
}
