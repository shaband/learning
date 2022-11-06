package templates

import "embed"

//go:embed *.gohtml layout/*
var FS embed.FS
