package login

type ConfigDB struct {
	Port          string   `env:"PORT" env-default:"5432"`
	Host          string   `env:"HOST" env-default:"localhost"`
	Name          string   `env:"NAME" env-default:"db2"`
	User          string   `env:"USER" env-default:""`
	Password      string   `env:"PASSWORD" env-default:"Ankita@2307"`
	Topicname     string   `env:"TOPICNAME" env-default:"quickstart-events"`
	BrokerAddress []string `env:"BROKERADDRESS" env-default:"localhost:9092"`
}
