module github.com/monandkey/ssh

go 1.17

replace github.com/monandkey/ssh/pkg/ssh => ./pkg/ssh

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/monandkey/ssh/pkg/ssh v0.0.0-00010101000000-000000000000
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)

require github.com/spf13/cobra v1.1.3
