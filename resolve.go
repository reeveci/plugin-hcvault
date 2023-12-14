package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/reeveci/reeve-lib/schema"
)

func (p *VaultPlugin) Resolve(env []string) (map[string]schema.Env, error) {
	result := make(map[string]schema.Env, len(env))
	secrets := make(map[string]map[string]interface{}, len(env))

	for _, key := range env {
		parts := strings.Split(strings.Trim(key, "/"), "/")
		var path, field string
		if len(parts) == 1 {
			path = key
			field = "value"
		} else {
			path = strings.Join(parts[:len(parts)-1], "/")
			field = parts[len(parts)-1]
		}

		secret, ok := secrets[path]
		if !ok {
			secret = fetchSecret(p, path)
			secrets[path] = secret
		}

		if secret == nil {
			continue
		}

		if value, ok := secret[field].(string); ok {
			result[key] = schema.Env{
				Value:    value,
				Priority: p.Priority,
				Secret:   !p.NoSecret,
			}
		}
	}

	return result, nil
}

func fetchSecret(p *VaultPlugin, path string) map[string]interface{} {
	url := fmt.Sprintf("%s/v1/%s/data/%s", p.Url, strings.Trim(p.Path, "/"), url.PathEscape(path))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		p.Log.Error(fmt.Sprintf(`fetching secret "%s" failed - %s`, path, err))
		return nil
	}

	req.Header.Set("X-Vault-Token", p.Token)

	resp, err := p.http.Do(req)
	if err != nil {
		p.Log.Error(fmt.Sprintf(`fetching secret "%s" failed - %s`, path, err))
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var secretResponse SecretResponse
	err = json.NewDecoder(resp.Body).Decode(&secretResponse)
	if err != nil {
		p.Log.Error(fmt.Sprintf(`error parsing vault response for secret "%s" - %s`, path, err))
		return nil
	}

	return secretResponse.Data.Data
}
