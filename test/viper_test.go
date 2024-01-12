package test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config = viper.New()
	assert.NotNil(t, config)
}

func TestViperConfig(t *testing.T) {
	viper.SetConfigFile(`../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	t.Log(viper.GetString(`database.host`))
	t.Log(viper.GetString(`database.port`))
	t.Log(viper.GetString(`database.user`))
	t.Log(viper.GetString(`database.name`))
	t.Log(viper.GetString(`database.password`))
}
