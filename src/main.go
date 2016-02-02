package main

import (
	"runtime"

	"config"
	"query"
	"server"
	"utils"

	"github.com/pkg/profile"
)

// param: lower case + Upper Case ,No _ spliter
// Struct unit: Upper Case
// Func: golang style

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 3)
	defer profile.Start(profile.CPUProfile).Stop()

	config.InitConfig()
	utils.InitLogger()
	if config.RC.MySQLEnabled {
		query.RC_MySQLConf = config.RC.MySQLConf
		query.InitMySQL(query.RC_MySQLConf)
	}
	server.NewServer()

}