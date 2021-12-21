package config

import (
	"reflect"
	"testing"
)

func Test_getFileName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Normal Test",
			want: "/root/.ssh/ssh_config.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileName(); got != tt.want {
				t.Errorf("getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configLoad(t *testing.T) {
	var (
		cfg     = configExampleLoad()
		mcfg    = []configParams{cfg[0], cfg[0]}
		jumpcfg = configExampleLoadJump()
	)

	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want []configParams
	}{
		{
			name: "Single node Test",
			args: args{
				fileName: "/root/.ssh/ssh_config.yaml",
			},
			want: cfg,
		},
		{
			name: "Multipul Node Test",
			args: args{
				fileName: "/root/.ssh/ssh_config_multi.yaml",
			},
			want: mcfg,
		},
		{
			name: "Jump Node Test",
			args: args{
				fileName: "/root/.ssh/ssh_config_jump.yaml",
			},
			want: jumpcfg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configLoad(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
