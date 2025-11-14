package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	Author = "Albert Skonieczny"
	Name   = "Everforest"

	EverforestTemplatePath         = "./templates/everforest.json.tmpl"
	EverforestMaterialTemplatePath = "./templates/everforest-material.json.tmpl"
	EverforestBlurTemplatePath     = "./templates/everforest-blur.json.tmpl"

	EverforestOutputPath         = "./themes/everforest.json"
	EverforestMaterialOutputPath = "./themes/everforest-material.json"
	EverforestBlurOutputPath     = "./themes/everforest-blur.json"

	EverforestDarkHardPalette    = "./palettes/everforest-dark-hard.json"
	EverforestDarkMediumPalette  = "./palettes/everforest-dark-medium.json"
	EverforestDarkSoftPalette    = "./palettes/everforest-dark-soft.json"
	EverforestLightHardPalette   = "./palettes/everforest-light-hard.json"
	EverforestLightMediumPalette = "./palettes/everforest-light-medium.json"
	EverforestLightSoftPalette   = "./palettes/everforest-light-soft.json"

	DarkAppearance  = "dark"
	LightAppearance = "light"

	BaseTemplatePath = "./templates/base.json.tmpl"
	ThemeIndent      = "    "
)

var zedThemes = []ZedTheme{
	{
		Author:       Author,
		Name:         Name,
		NameSuffix:   "",
		TemplatePath: EverforestTemplatePath,
		OutputPath:   EverforestOutputPath,
		Themes:       themeVariants,
	},
	{
		Author:       Author,
		Name:         Name,
		NameSuffix:   " (material)",
		TemplatePath: EverforestMaterialTemplatePath,
		OutputPath:   EverforestMaterialOutputPath,
		Themes:       themeVariants,
	},
	{
		Author:       Author,
		Name:         Name,
		NameSuffix:   " (blur)",
		TemplatePath: EverforestBlurTemplatePath,
		OutputPath:   EverforestBlurOutputPath,
		Themes:       themeVariants,
	},
}

var themeVariants = []ThemeVariant{
	{Name: "Everforest Dark Hard", Appearance: DarkAppearance, PalettePath: EverforestDarkHardPalette},
	{Name: "Everforest Dark Medium", Appearance: DarkAppearance, PalettePath: EverforestDarkMediumPalette},
	{Name: "Everforest Dark Soft", Appearance: DarkAppearance, PalettePath: EverforestDarkSoftPalette},
	{Name: "Everforest Light Hard", Appearance: LightAppearance, PalettePath: EverforestLightHardPalette},
	{Name: "Everforest Light Medium", Appearance: LightAppearance, PalettePath: EverforestLightMediumPalette},
	{Name: "Everforest Light Soft", Appearance: LightAppearance, PalettePath: EverforestLightSoftPalette},
}

func Generate() {
	baseTemplateContent, err := os.ReadFile(BaseTemplatePath)
	if err != nil {
		log.Fatalf("fatal: could not read base template: %v\n", err)
	}
	baseTemplate, err := template.New("base").Funcs(template.FuncMap{
		"renderThemeTemplates": func(themes []string) string {
			return renderThemeTemplates(themes, ThemeIndent)
		},
	}).Parse(string(baseTemplateContent))
	if err != nil {
		log.Fatalf("fatal: could not parse base template: %v\n", err)
	}

	for i := range zedThemes {
		zt := &zedThemes[i]

		if err := zt.loadPalettes(); err != nil {
			log.Fatalf("fatal: %v\n", err)
		}

		displayName := fmt.Sprintf("%s%s", zt.Name, zt.NameSuffix)

		renderedThemes, err := zt.ThemeTemplates()
		if err != nil {
			log.Fatalf("fatal: %v\n", err)
		}

		zedThemeTemplateData := struct {
			Author         string
			Name           string
			ThemeTemplates []string
		}{
			Author:         zt.Author,
			Name:           displayName,
			ThemeTemplates: renderedThemes,
		}

		var output bytes.Buffer
		if err := baseTemplate.Execute(&output, zedThemeTemplateData); err != nil {
			log.Fatalf("fatal: could not execute base template for %s: %v\n", displayName, err)
		}

		if err := os.WriteFile(zt.OutputPath, output.Bytes(), 0o644); err != nil {
			log.Fatalf("fatal: could not write theme file %s: %v\n", zt.OutputPath, err)
		}
	}
}

type ZedTheme struct {
	Author       string
	Name         string
	NameSuffix   string
	TemplatePath string
	OutputPath   string
	Themes       []ThemeVariant
}

type ThemeVariant struct {
	Name        string
	Appearance  string
	PalettePath string
	Palette     Palette
}

func (zt *ZedTheme) ThemeTemplates() ([]string, error) {
	content, err := os.ReadFile(zt.TemplatePath)
	if err != nil {
		return nil, fmt.Errorf("could not read theme template %s: %w", zt.TemplatePath, err)
	}

	themeTemplate, err := template.New("theme").Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("could not parse theme template %s: %w", zt.TemplatePath, err)
	}

	rendered := make([]string, 0, len(zt.Themes))
	for i := range zt.Themes {
		theme := zt.Themes[i]
		theme.Name = fmt.Sprintf("%s%s", theme.Name, zt.NameSuffix)

		var buf bytes.Buffer
		if err := themeTemplate.Execute(&buf, theme); err != nil {
			return nil, fmt.Errorf("could not execute theme template for %s: %w", theme.Name, err)
		}

		rendered = append(rendered, strings.TrimSpace(buf.String()))
	}

	return rendered, nil
}

func renderThemeTemplates(themes []string, indent string) string {
	if len(themes) == 0 {
		return ""
	}

	var blocks []string
	for _, theme := range themes {
		theme = strings.TrimSpace(theme)
		if theme == "" {
			continue
		}

		lines := strings.Split(theme, "\n")
		for i := range lines {
			if lines[i] == "" {
				lines[i] = indent
				continue
			}
			lines[i] = indent + lines[i]
		}

		blocks = append(blocks, strings.Join(lines, "\n"))
	}

	if len(blocks) == 0 {
		return ""
	}

	return strings.Join(blocks, ",\n")
}

type Palette struct {
	BgDim       string `json:"bg_dim"`
	Bg0         string `json:"bg0"`
	Bg1         string `json:"bg1"`
	Bg2         string `json:"bg2"`
	Bg3         string `json:"bg3"`
	Bg4         string `json:"bg4"`
	Bg5         string `json:"bg5"`
	BgVisual    string `json:"bg_visual"`
	BgRed       string `json:"bg_red"`
	BgYellow    string `json:"bg_yellow"`
	BgGreen     string `json:"bg_green"`
	BgBlue      string `json:"bg_blue"`
	BgPurple    string `json:"bg_purple"`
	Fg          string `json:"fg"`
	Red         string `json:"red"`
	Orange      string `json:"orange"`
	Yellow      string `json:"yellow"`
	Green       string `json:"green"`
	Aqua        string `json:"aqua"`
	Blue        string `json:"blue"`
	Purple      string `json:"purple"`
	Grey0       string `json:"grey0"`
	Grey1       string `json:"grey1"`
	Grey2       string `json:"grey2"`
	Statusline1 string `json:"statusline1"`
	Statusline2 string `json:"statusline2"`
	Statusline3 string `json:"statusline3"`
}

func (zt *ZedTheme) loadPalettes() error {
	for i := range zt.Themes {
		paletteFile, err := os.Open(zt.Themes[i].PalettePath)
		if err != nil {
			return fmt.Errorf("could not open palette %s: %w", zt.Themes[i].PalettePath, err)
		}

		decodeErr := json.NewDecoder(paletteFile).Decode(&zt.Themes[i].Palette)

		closeErr := paletteFile.Close()

		if decodeErr != nil {
			return fmt.Errorf("failed to decode palette for %s: %w", zt.Themes[i].Name, decodeErr)
		}

		if closeErr != nil {
			return fmt.Errorf("could not close palette %s: %w", zt.Themes[i].PalettePath, closeErr)
		}
	}

	return nil
}
