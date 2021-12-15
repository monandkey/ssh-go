BINNAME:=ssh
BINDIR:=../bin
GOARCH:=amd64
GO_FILES:=../main.go

.PHONY: build
build:
	GOOS=windows GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINNAME).exe $(GO_FILES)
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINNAME) $(GO_FILES)
