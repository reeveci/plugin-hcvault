package main

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/reeveci/reeve-lib/plugin"
	"github.com/reeveci/reeve-lib/schema"
)

const PLUGIN_NAME = "hcvault"

func main() {
	log := hclog.New(&hclog.LoggerOptions{})

	plugin.Serve(&plugin.PluginConfig{
		Plugin: &VaultPlugin{
			Log: log,

			http: &http.Client{},
		},

		Logger: log,
	})
}

type VaultPlugin struct {
	Url      string
	Token    string
	Path     string
	Priority uint32
	NoSecret bool

	Log hclog.Logger

	http *http.Client
}

func (p *VaultPlugin) Name() (string, error) {
	return PLUGIN_NAME, nil
}

func (p *VaultPlugin) Register(settings map[string]string, api plugin.ReeveAPI) (capabilities plugin.Capabilities, err error) {
	api.Close()

	var enabled bool
	if enabled, err = boolSetting(settings, "ENABLED"); !enabled || err != nil {
		return
	}
	if p.Url, err = requireSetting(settings, "URL"); err != nil {
		return
	}
	if p.Token, err = requireSetting(settings, "TOKEN"); err != nil {
		return
	}
	if p.Path, err = requireSetting(settings, "PATH"); err != nil {
		return
	}
	var priority int
	if priority, err = intSetting(settings, "PRIORITY", 1); err != nil {
		return
	} else {
		p.Priority = uint32(priority)
	}
	if p.NoSecret, err = boolSetting(settings, "NO_SECRET"); err != nil {
		return
	}

	capabilities.Resolve = true
	return
}

func (p *VaultPlugin) Unregister() error {
	return nil
}

func (p *VaultPlugin) Message(source string, message schema.Message) error {
	return nil
}

func (p *VaultPlugin) Discover(trigger schema.Trigger) ([]schema.Pipeline, error) {
	return nil, nil
}

func (p *VaultPlugin) Notify(status schema.PipelineStatus) error {
	return nil
}

func (p *VaultPlugin) CLIMethod(method string, args []string) (string, error) {
	return "", nil
}
