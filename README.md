# Reeve CI / CD - HashiCorp Vault Plugin

This is a [Reeve](https://github.com/reeveci/reeve) plugin for providing pipeline environment variables from a HashiCorp Vault KV store.

## Configuration

Currently, only the kv v2 secrets engine is supported.

If an env key is a path (meaning that it includes at least one `/`), all but the last segments are used as the secret path and the last segment is used as the secret data key.
Otherwise, `value` is used as the secret data key.

### Vault

An API token is required for this plugin.
It is recommended to use a token configured with minimal required access.

### Settings

Settings can be provided to the plugin through environment variables set to the reeve server.

Settings for this plugin should be prefixed by `REEVE_PLUGIN_HCVAULT_`.

Settings may also be shared between plugins by prefixing them with `REEVE_SHARED_` instead.

- `ENABLED` - `true` enables this plugin
- `URL` (required) - Vault URL
- `TOKEN` (required) - Vault API Token
- `PATH` (required) - The path of the secret engine
- `PRIORITY` (default `1`) - Priority of all variables returned by this plugin
- `NO_SECRET` - `true` prevents variables returned by this plugin from being marked as secret
