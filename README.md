# 🌲 Everforest for Zed

An exact port of sainnhe's [Everforest](https://github.com/sainnhe/everforest) vim color scheme for Zed.
The theme comes in regular, material, and blur variants.

## Development

The repository includes a tiny Go entry point (`main.go`) that wires the generator into `go run .`.
Use the provided `Makefile` to keep generated assets in sync with any palette or template changes.

### Common tasks

- `make generate` - generate `./themes/*.json` with `./scripts/generate.go`
- `make lint` - run `golangci-lint`
