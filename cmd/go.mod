module github.com/monandkey/ssh/cmd

go 1.17

replace (
	github.com/monandkey/ssh/pkg/fileutil => ../pkg/fileutil
	github.com/monandkey/ssh/pkg/log => ../pkg/log
	github.com/monandkey/ssh/pkg/ssh => ../pkg/ssh
)

require (
	github.com/monandkey/ssh/pkg/ssh v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.3.0
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/monandkey/ssh/pkg/log v0.0.0-00010101000000-000000000000 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)
