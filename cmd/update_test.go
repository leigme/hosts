package cmd

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

func Test_readSource(t *testing.T) {
	type args struct {
		header []string
		footer []string
	}
	viper.SetConfigFile("/home/leig/.___config_i/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testReadSource_1",
			args: args{
				header: make([]string, 0),
				footer: make([]string, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readSource(&tt.args.header, &tt.args.footer)
		})
	}
}

func Test_updateSource(t *testing.T) {
	viper.SetConfigFile("/home/leig/.___config_i/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	tests := []struct {
		name string
	}{
		{
			name: "testUpdateSource_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateSource()
		})
	}
}
