module github.com/monandkey/ssh/pkg/ssh

go 1.17

replace github.com/monandkey/ssh/pkg/log => ../log

require (
	github.com/monandkey/ssh/pkg/log v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
)

require (
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
)
