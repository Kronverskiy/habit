package app

import (
	"strings"
	"testing"

	"github.com/nalgeon/be"
	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	yamlConfig := `
server:
  host: "127.0.0.2"
  port: 8081
`

	viper.Reset()
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(strings.NewReader(yamlConfig))
	be.Err(t, err, nil)

	var cfg appConfig
	err = viper.Unmarshal(&cfg)
	be.Err(t, err, nil)

	be.Equal(t, cfg.Server.Host, "127.0.0.2")
	be.Equal(t, cfg.Server.Port, "8081")
}

func TestSetDefaults(t *testing.T) {
	t.Run("full defaults", func(t *testing.T) {
		cfg := appConfig{
			Server: struct {
				Host string
				Port string
			}{
				Host: "",
				Port: "",
			},
		}

		actualCfg := setDefaults(cfg)

		be.Equal(t, actualCfg.Server.Host, defaultServerHost)
		be.Equal(t, actualCfg.Server.Host, defaultServerHost)
	})

	t.Run("not defaults", func(t *testing.T) {
		cfg := appConfig{
			Server: struct {
				Host string
				Port string
			}{
				Host: "127.0.0.5",
				Port: "8888",
			},
		}

		actualCfg := setDefaults(cfg)

		be.Equal(t, actualCfg.Server.Host, "127.0.0.5")
		be.Equal(t, actualCfg.Server.Port, "8888")
	})
}
