# 🌲 Everforest for Zed

A faithful port of sainnhe's [Everforest](https://github.com/sainnhe/everforest) theme for Zed.
Comes in `regular`, `material`, and `blur` variants.

## Preview

<table>
  <thead>
    <tr>
      <th>Theme</th>
      <th width="33%"><code>regular</code></th>
      <th width="33%"><code>material</code></th>
      <th width="33%"><code>blur</code></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Dark Medium</td>
      <td><img src="docs/dark.png" width="320" alt="Dark Regular screenshot" /></td>
      <td><img src="docs/dark-material.png" width="320" alt="Dark Material screenshot" /></td>
      <td><img src="docs/dark-blur.png" width="320" alt="Dark Blur screenshot" /></td>
    </tr>
    <tr>
      <td>Light Medium</td>
      <td><img src="docs/light.png" width="320" alt="Light Regular screenshot" /></td>
      <td><img src="docs/light-material.png" width="320" alt="Light Material screenshot" /></td>
      <td><img src="docs/light-blur.png" width="320" alt="Light Blur screenshot" /></td>
    </tr>
  </tbody>
</table>

_Wallpaper shown in the background: [aerial photography of roadway](https://unsplash.com/photos/aerial-photography-of-roadway-Wn4ulyzVoD4)._

## Palettes

The swatches below are pulled from `./palettes/*.json`. Each cell shows an image tinted with the palette color, covering the colors used by the `./scripts/generate.go`.

| Theme        |                          `bg0`                           |                          `bg1`                           |                          `bg2`                           |                       `bg_visual`                        |                           `fg`                           |                          `red`                           |                         `orange`                         |                         `yellow`                         |                         `green`                          |                          `aqua`                          |                          `blue`                          |                         `purple`                         |
| :----------- | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: | :------------------------------------------------------: |
| Dark Hard    | ![#272e33](https://placehold.co/27x27/272e33/272e33.png) | ![#2e383c](https://placehold.co/27x27/2e383c/2e383c.png) | ![#374145](https://placehold.co/27x27/374145/374145.png) | ![#4c3743](https://placehold.co/27x27/4c3743/4c3743.png) | ![#d3c6aa](https://placehold.co/27x27/d3c6aa/d3c6aa.png) | ![#e67e80](https://placehold.co/27x27/e67e80/e67e80.png) | ![#e69875](https://placehold.co/27x27/e69875/e69875.png) | ![#dbbc7f](https://placehold.co/27x27/dbbc7f/dbbc7f.png) | ![#a7c080](https://placehold.co/27x27/a7c080/a7c080.png) | ![#83c092](https://placehold.co/27x27/83c092/83c092.png) | ![#7fbbb3](https://placehold.co/27x27/7fbbb3/7fbbb3.png) | ![#d699b6](https://placehold.co/27x27/d699b6/d699b6.png) |
| Dark Medium  | ![#2d353b](https://placehold.co/27x27/2d353b/2d353b.png) | ![#343f44](https://placehold.co/27x27/343f44/343f44.png) | ![#3d484d](https://placehold.co/27x27/3d484d/3d484d.png) | ![#543a48](https://placehold.co/27x27/543a48/543a48.png) | ![#d3c6aa](https://placehold.co/27x27/d3c6aa/d3c6aa.png) | ![#e67e80](https://placehold.co/27x27/e67e80/e67e80.png) | ![#e69875](https://placehold.co/27x27/e69875/e69875.png) | ![#dbbc7f](https://placehold.co/27x27/dbbc7f/dbbc7f.png) | ![#a7c080](https://placehold.co/27x27/a7c080/a7c080.png) | ![#83c092](https://placehold.co/27x27/83c092/83c092.png) | ![#7fbbb3](https://placehold.co/27x27/7fbbb3/7fbbb3.png) | ![#d699b6](https://placehold.co/27x27/d699b6/d699b6.png) |
| Dark Soft    | ![#333c43](https://placehold.co/27x27/333c43/333c43.png) | ![#3a464c](https://placehold.co/27x27/3a464c/3a464c.png) | ![#434f55](https://placehold.co/27x27/434f55/434f55.png) | ![#5c3f4f](https://placehold.co/27x27/5c3f4f/5c3f4f.png) | ![#d3c6aa](https://placehold.co/27x27/d3c6aa/d3c6aa.png) | ![#e67e80](https://placehold.co/27x27/e67e80/e67e80.png) | ![#e69875](https://placehold.co/27x27/e69875/e69875.png) | ![#dbbc7f](https://placehold.co/27x27/dbbc7f/dbbc7f.png) | ![#a7c080](https://placehold.co/27x27/a7c080/a7c080.png) | ![#83c092](https://placehold.co/27x27/83c092/83c092.png) | ![#7fbbb3](https://placehold.co/27x27/7fbbb3/7fbbb3.png) | ![#d699b6](https://placehold.co/27x27/d699b6/d699b6.png) |
| Light Hard   | ![#fffbef](https://placehold.co/27x27/fffbef/fffbef.png) | ![#f8f5e4](https://placehold.co/27x27/f8f5e4/f8f5e4.png) | ![#f2efdf](https://placehold.co/27x27/f2efdf/f2efdf.png) | ![#f0f2d4](https://placehold.co/27x27/f0f2d4/f0f2d4.png) | ![#5c6a72](https://placehold.co/27x27/5c6a72/5c6a72.png) | ![#f85552](https://placehold.co/27x27/f85552/f85552.png) | ![#f57d26](https://placehold.co/27x27/f57d26/f57d26.png) | ![#dfa000](https://placehold.co/27x27/dfa000/dfa000.png) | ![#8da101](https://placehold.co/27x27/8da101/8da101.png) | ![#35a77c](https://placehold.co/27x27/35a77c/35a77c.png) | ![#3a94c5](https://placehold.co/27x27/3a94c5/3a94c5.png) | ![#df69ba](https://placehold.co/27x27/df69ba/df69ba.png) |
| Light Medium | ![#fdf6e3](https://placehold.co/27x27/fdf6e3/fdf6e3.png) | ![#f4f0d9](https://placehold.co/27x27/f4f0d9/f4f0d9.png) | ![#efebd4](https://placehold.co/27x27/efebd4/efebd4.png) | ![#eaedc8](https://placehold.co/27x27/eaedc8/eaedc8.png) | ![#5c6a72](https://placehold.co/27x27/5c6a72/5c6a72.png) | ![#f85552](https://placehold.co/27x27/f85552/f85552.png) | ![#f57d26](https://placehold.co/27x27/f57d26/f57d26.png) | ![#dfa000](https://placehold.co/27x27/dfa000/dfa000.png) | ![#8da101](https://placehold.co/27x27/8da101/8da101.png) | ![#35a77c](https://placehold.co/27x27/35a77c/35a77c.png) | ![#3a94c5](https://placehold.co/27x27/3a94c5/3a94c5.png) | ![#df69ba](https://placehold.co/27x27/df69ba/df69ba.png) |
| Light Soft   | ![#f3ead3](https://placehold.co/27x27/f3ead3/f3ead3.png) | ![#eae4ca](https://placehold.co/27x27/eae4ca/eae4ca.png) | ![#e5dfc5](https://placehold.co/27x27/e5dfc5/e5dfc5.png) | ![#e1e4bd](https://placehold.co/27x27/e1e4bd/e1e4bd.png) | ![#5c6a72](https://placehold.co/27x27/5c6a72/5c6a72.png) | ![#f85552](https://placehold.co/27x27/f85552/f85552.png) | ![#f57d26](https://placehold.co/27x27/f57d26/f57d26.png) | ![#dfa000](https://placehold.co/27x27/dfa000/dfa000.png) | ![#8da101](https://placehold.co/27x27/8da101/8da101.png) | ![#35a77c](https://placehold.co/27x27/35a77c/35a77c.png) | ![#3a94c5](https://placehold.co/27x27/3a94c5/3a94c5.png) | ![#df69ba](https://placehold.co/27x27/df69ba/df69ba.png) |

## Development

The repository includes a tiny Go entry point `main.go`.
Use the provided `Makefile` to keep generated assets in sync with any palette or template changes.

### Common tasks

- `make generate` - generate `./themes/*.json` with `./scripts/generate.go`

## Inspiration

- [golang-templates/seed](https://github.com/golang-templates/seed) - `Makefile`, CI, and `.gitignore`
- [TheStandup - DHH Talks Omarchy](https://www.youtube.com/watch?v=ljGPvgMPOn8) - motivation behind porting fav color scheme to fav code editor <3
