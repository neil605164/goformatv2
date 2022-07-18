package global

import (
	"embed"
	"goformatv2/app/global/structs"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config 讀取dev.yaml檔案
var Config *structs.EnvConfig

// Lang 各語系
var Lang = []string{"en", "tw", "cn"}

func getEnv() string {
	if len(os.Getenv("ENV")) <= 0 {
		log.Fatalf("🔔🔔🔔  Can not get ENV value 🔔🔔🔔")
	}
	return os.Getenv("ENV")
}

// Start 執行main.go的第一步驟，載入各環境設定檔
func Start(f embed.FS) {
	env := getEnv()

	envPathList := []string{
		"env/" + env + "/db.yaml",
		"env/" + env + "/api.yaml",
		"env/" + env + "/other.yaml",
	}

	for k := range envPathList {
		configFile, err := f.ReadFile(envPathList[k])
		if err != nil {
			log.Fatalf("🔔🔔🔔  Can not find Yaml file %v 🔔🔔🔔", err)
		}
		// 塞值進入struct
		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			panic(err)
		}

	}

}
