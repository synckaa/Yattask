package bootstrap

import "Yattask/configs"

func InitApp() {
	configs.LoadEnv()
	configs.GetConnDB()
	configs.SyncTables(configs.DB)
	configs.NewValidator()
}
