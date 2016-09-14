package main

import "github.com/kitech/php-go/phpgo"

var (
	verison = "0.1"
)

func init() {
	phpgo.InitExtension("protorpc-php", "0.1")

	// phpgo.RegisterInitFunctions(module_startup, module_shutdown, request_startup, request_shutdown)

	phpgo.AddFunc("protorpc_version", func() string {
		return verison
	})

	phpgo.AddFunc("protorpc_client", Protorpc)
	phpgo.AddFunc("protorpc_call", ProtorpcCall)
}

func main() {}
