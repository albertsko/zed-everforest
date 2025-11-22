# 🌲 Everforest for Zed

A faithful, exact port of sainnhe's [Everforest](https://github.com/sainnhe/everforest) theme for Zed.

The theme comes in regular, material, and blur variants.

## Palettes

The swatches below are pulled from `./palettes/*.json`. Each cell shows a `█` tinted with the palette color plus its hex code, covering the core color keys used by the generator.

| Theme                   | bg0                              | bg1                              | bg2                              | bg_visual                        | fg                               | red                              | orange                           | yellow                           | green                            | aqua                             | blue                             | purple                           |
| ----------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- |
| Everforest Dark Hard    | <a style="color:#272e33">███</a> | <a style="color:#2e383c">███</a> | <a style="color:#374145">███</a> | <a style="color:#4c3743">███</a> | <a style="color:#d3c6aa">███</a> | <a style="color:#e67e80">███</a> | <a style="color:#e69875">███</a> | <a style="color:#dbbc7f">███</a> | <a style="color:#a7c080">███</a> | <a style="color:#83c092">███</a> | <a style="color:#7fbbb3">███</a> | <a style="color:#d699b6">███</a> |
| Everforest Dark Medium  | <a style="color:#2d353b">███</a> | <a style="color:#343f44">███</a> | <a style="color:#3d484d">███</a> | <a style="color:#543a48">███</a> | <a style="color:#d3c6aa">███</a> | <a style="color:#e67e80">███</a> | <a style="color:#e69875">███</a> | <a style="color:#dbbc7f">███</a> | <a style="color:#a7c080">███</a> | <a style="color:#83c092">███</a> | <a style="color:#7fbbb3">███</a> | <a style="color:#d699b6">███</a> |
| Everforest Dark Soft    | <a style="color:#333c43">███</a> | <a style="color:#3a464c">███</a> | <a style="color:#434f55">███</a> | <a style="color:#5c3f4f">███</a> | <a style="color:#d3c6aa">███</a> | <a style="color:#e67e80">███</a> | <a style="color:#e69875">███</a> | <a style="color:#dbbc7f">███</a> | <a style="color:#a7c080">███</a> | <a style="color:#83c092">███</a> | <a style="color:#7fbbb3">███</a> | <a style="color:#d699b6">███</a> |
| Everforest Light Hard   | <a style="color:#fffbef">███</a> | <a style="color:#f8f5e4">███</a> | <a style="color:#f2efdf">███</a> | <a style="color:#f0f2d4">███</a> | <a style="color:#5c6a72">███</a> | <a style="color:#f85552">███</a> | <a style="color:#f57d26">███</a> | <a style="color:#dfa000">███</a> | <a style="color:#8da101">███</a> | <a style="color:#35a77c">███</a> | <a style="color:#3a94c5">███</a> | <a style="color:#df69ba">███</a> |
| Everforest Light Medium | <a style="color:#fdf6e3">███</a> | <a style="color:#f4f0d9">███</a> | <a style="color:#efebd4">███</a> | <a style="color:#eaedc8">███</a> | <a style="color:#5c6a72">███</a> | <a style="color:#f85552">███</a> | <a style="color:#f57d26">███</a> | <a style="color:#dfa000">███</a> | <a style="color:#8da101">███</a> | <a style="color:#35a77c">███</a> | <a style="color:#3a94c5">███</a> | <a style="color:#df69ba">███</a> |
| Everforest Light Soft   | <a style="color:#f3ead3">███</a> | <a style="color:#eae4ca">███</a> | <a style="color:#e5dfc5">███</a> | <a style="color:#e1e4bd">███</a> | <a style="color:#5c6a72">███</a> | <a style="color:#f85552">███</a> | <a style="color:#f57d26">███</a> | <a style="color:#dfa000">███</a> | <a style="color:#8da101">███</a> | <a style="color:#35a77c">███</a> | <a style="color:#3a94c5">███</a> | <a style="color:#df69ba">███</a> |

## Development

The repository includes a tiny Go entry point `main.go`.
Use the provided `Makefile` to keep generated assets in sync with any palette or template changes.

### Common tasks

- `make generate` - generate `./themes/*.json` with `./scripts/generate.go`
- `make lint` - run `golangci-lint`

## Inspiration

- [golang-templates/seed](https://github.com/golang-templates/seed) - `Makefile`, CI, and `.gitignore`
- [TheStandup - DHH Talks Omarchy](https://www.youtube.com/watch?v=ljGPvgMPOn8) - motivation behind doing my best in porting fav color scheme to fav code editor <3
