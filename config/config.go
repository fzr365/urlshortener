package config

type Config struct {
	DB struct {
		DSN string `yaml:"dsn"`
	} `yaml:"db"`

}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	//数据库连接池属性
	MaxIdleConns int `mapstructure:"max_idle_conns"`
	MaxOpenConns int `mapstructure:"max_open_conns"`

}

//数据库配置
type RedisConfig struct {
	Address string `mapstructure:"address"`
	Password string `mapstructure:"password"`	
	DB int `mapstructure:"db"`
}