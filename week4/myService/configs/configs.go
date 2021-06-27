package configs

var httpAddress string

func InitConfig() error {
	//todo: load config file
	return nil
}

func GetHttpAddress() string {
	return httpAddress
}
