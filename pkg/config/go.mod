module github.com/monandkey/ssh/pkg/config

go 1.17

replace github.com/monandkey/ssh/pkg/fileutil => ../fileutil

require (
	github.com/monandkey/ssh/pkg/fileutil v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)
