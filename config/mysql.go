package config

type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	DBName   string `json:"dbName" yaml:"dbName"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (m *Mysql) Dsn() string {
	const config = "?charset=utf8mb4&parseTime=True&loc=Local"
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBName + config
}
