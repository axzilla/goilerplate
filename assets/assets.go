package assets

import "embed"

//go:embed css/* img/* js/*
var Assets embed.FS
