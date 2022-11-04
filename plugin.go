package types

import "plugin"

type Manifest struct {
	ID           string
	Name         string
	Description  string
	Author       string
	Version      string
	Dependencies []string
}

type PluginBase struct {
	Runtime  *plugin.Plugin
	Manifest *Manifest
}

type Symbols string

const (
	MANIFEST_SYMBOL Symbols = "GetManifest"
)

type GetManifest func() *Manifest
