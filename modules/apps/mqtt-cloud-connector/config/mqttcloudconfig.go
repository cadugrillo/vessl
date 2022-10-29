package mqttcloudconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ClientSub struct {
		ClientId           string `yaml:"clientId"`
		ServerAddress      string `yaml:"serverAddress"`
		Qos                int    `yaml:"qos"`
		ConnectionTimeout  int    `yaml:"connectionTimeout"`
		WriteTimeout       int    `yaml:"writeTimeout"`
		KeepAlive          int    `yaml:"keepAlive"`
		PingTimeout        int    `yaml:"pingTimeout"`
		ConnectRetry       bool   `yaml:"connectRetry"`
		AutoConnect        bool   `yaml:"autoConnect"`
		OrderMaters        bool   `yaml:"orderMaters"`
		UserName           string `yaml:"userName"`
		Password           string `yaml:"password"`
		TlsConn            bool   `yaml:"tlsConn"`
		RootCAPath         string `yaml:"rootCAPath"`
		ClientKeyPath      string `yaml:"clientKeyPath"`
		PrivateKeyPath     string `yaml:"privateKeyPath"`
		InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
	} `yaml:"clientSub"`
	ClientPub struct {
		ClientId           string `yaml:"clientId"`
		ServerAddress      string `yaml:"serverAddress"`
		Qos                int    `yaml:"qos"`
		ConnectionTimeout  int    `yaml:"connectionTimeout"`
		WriteTimeout       int    `yaml:"writeTimeout"`
		KeepAlive          int    `yaml:"keepAlive"`
		PingTimeout        int    `yaml:"pingTimeout"`
		ConnectRetry       bool   `yaml:"connectRetry"`
		AutoConnect        bool   `yaml:"autoConnect"`
		OrderMaters        bool   `yaml:"orderMaters"`
		UserName           string `yaml:"userName"`
		Password           string `yaml:"password"`
		TlsConn            bool   `yaml:"tlsConn"`
		RootCAPath         string `yaml:"rootCAPath"`
		ClientKeyPath      string `yaml:"clientKeyPath"`
		PrivateKeyPath     string `yaml:"privateKeyPath"`
		InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
		TranslateTopic     bool   `yaml:"translateTopic"`
		PublishInterval    int    `yaml:"publishInterval"`
	} `yaml:"clientPub"`
	Logs struct {
		SubPayload bool `yaml:"subPayload"`
		Debug      bool `yaml:"debug"`
		Warning    bool `yaml:"warning"`
		Error      bool `yaml:"error"`
		Critical   bool `yaml:"critical"`
	} `yaml:"logs"`
	TopicsSub struct {
		Topic []string
	} `yaml:"topicsSub"`
	TopicsPub struct {
		Topic []string
	} `yaml:"topicsPub"`
}

func ReadConfig() Config {
	f, err := os.Open("./apps/mqtt-cloud-connector/config/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func WriteConfig(ConfigFile Config) error {
	f, err := os.Create("./apps/mqtt-cloud-connector/config/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(&ConfigFile)
	if err != nil {
		panic(err)
	}
	return err

}
