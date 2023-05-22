package main

// Tool list from https://github.com/golang/vscode-go/wiki/tools

import (
	_ "github.com/cweill/gotests/gotests"
	_ "github.com/fatih/gomodifytags"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/haya14busa/goplay"
	_ "github.com/josharian/impl"
	_ "golang.org/x/tools/gopls"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
