package conf

//Config 配置文件
type Config struct {
	DataMapPath string    `json yaml:"dataMapPath"`
	DBConfig    *DBConfig `json yaml:"dbConfig"`
}

//DBConfig 数据缓存配置
type DBConfig struct {
	IP       string `json yaml:"ip"`
	Port     string `json yaml:"port" `
	User     string `json yaml:"user"`
	Password string `json yaml:"password"`
}
