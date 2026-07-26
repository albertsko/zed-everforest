# 🌲 Everforest for Zed

[Everforest](https://github.com/sainnhe/everforest) color scheme for Zed.
Theme comes in `regular`, `material`, and `blur` variants.

### Development

The files in `assets/` are the source of truth. `generate.go` combines the palettes and templates into the checked-in files under `themes/`. Do not edit generated themes directly.

The generator requires Go.

| Command         | Purpose                                                                       |
| --------------- | ----------------------------------------------------------------------------- |
| `make generate` | Regenerate every theme file.                                                  |
| `make ci`       | Run formatting, vet, and test checks.                                         |
| `make copygen`  | Regenerate themes and copy them to `~/.config/zed/themes/` for local testing. |

When contributing, update `assets/`, run `make generate`, run `make ci`, and commit both the source and generated changes.

### License

Released under the [MIT License](LICENSE).
