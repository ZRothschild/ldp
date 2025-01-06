package config

type (
	Config struct {
		AppName  string
		Env      string
		GrpcPort int
		HttpPort int
		JwtSk    string
		Mysql
	}

	Mysql struct {
		User      string // Username
		Passwd    string // Password (requires User)
		Net       string // Network (e.g. "tcp", "tcp6", "unix". default: "tcp")
		Addr      string // Address (default: "127.0.0.1:3306" for "tcp" and "/tmp/mysql.sock" for "unix")
		DBName    string // Database name
		Collation string // Connection collation
	}
)

const (
	FileName = "config"
	FileType = "yaml"
	FilePath = "./config"
)
