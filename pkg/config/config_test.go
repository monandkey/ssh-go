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
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want []configParams
	}{
		{
			name: "Normal Test",
			args: args{
				fileName: "/root/.ssh/ssh_config.yaml",
			},
			want: configExampleLoad(),
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
