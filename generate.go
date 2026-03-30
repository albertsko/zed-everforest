package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
	"text/template"
)

type (
	Palette struct {
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
	themeVariant struct {
		Name       string
		Appearance string
		Palette    Palette
	}
)

var (
	variants = []struct {
		name       string
		appearance string
		palette    string
	}{
		{"Everforest Dark Hard", "dark", "./assets/everforest-dark-hard.json"},
		{"Everforest Dark Medium", "dark", "./assets/everforest-dark-medium.json"},
		{"Everforest Dark Soft", "dark", "./assets/everforest-dark-soft.json"},
		{"Everforest Light Hard", "light", "./assets/everforest-light-hard.json"},
		{"Everforest Light Medium", "light", "./assets/everforest-light-medium.json"},
		{"Everforest Light Soft", "light", "./assets/everforest-light-soft.json"},
	}
	zedThemes = []struct {
		suffix       string
		templatePath string
		outputPath   string
	}{
		{" (regular)", "./assets/_everforest-regular.json.tmpl", "./themes/everforest-regular.json"},
		{" (material)", "./assets/_everforest-material.json.tmpl", "./themes/everforest-material.json"},
		{" (blur)", "./assets/_everforest-blur.json.tmpl", "./themes/everforest-blur.json"},
	}
)

func main() {
	// load all palettes once
	loaded := make([]themeVariant, len(variants))
	for i, v := range variants {
		loaded[i] = themeVariant{Name: v.name, Appearance: v.appearance}
		loaded[i].Palette = loadPalette(v.palette)
	}

	// parse base template
	baseContent, err := os.ReadFile("./assets/_base.json.tmpl")
	if err != nil {
		log.Fatalf("could not read base template: %v", err)
	}
	baseTmpl, err := template.New("base").Funcs(template.FuncMap{
		"renderThemeTemplates": func(themes []string) string {
			return indentAndJoin(themes, "    ")
		},
	}).Parse(string(baseContent))
	if err != nil {
		log.Fatalf("could not parse base template: %v", err)
	}

	// generate each theme file
	for _, zt := range zedThemes {
		displayName := "Everforest Theme" + zt.suffix

		renderedThemes := renderVariants(zt.templatePath, zt.suffix, loaded)

		var output bytes.Buffer
		err := baseTmpl.Execute(&output, struct {
			Name           string
			ThemeTemplates []string
		}{displayName, renderedThemes})
		if err != nil {
			log.Fatalf("could not execute base template for %s: %v", displayName, err)
		}

		if err := os.WriteFile(zt.outputPath, output.Bytes(), 0o644); err != nil {
			log.Fatalf("could not write %s: %v", zt.outputPath, err)
		}
	}
}

func loadPalette(path string) Palette {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read palette %s: %v", path, err)
	}
	var p Palette
	if err := json.Unmarshal(data, &p); err != nil {
		log.Fatalf("could not decode palette %s: %v", path, err)
	}
	return p
}

func renderVariants(templatePath, suffix string, themes []themeVariant) []string {
	content, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("could not read template %s: %v", templatePath, err)
	}
	tmpl, err := template.New("theme").Parse(string(content))
	if err != nil {
		log.Fatalf("could not parse template %s: %v", templatePath, err)
	}

	rendered := make([]string, 0, len(themes))
	for _, theme := range themes {
		theme.Name = theme.Name + suffix

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, theme); err != nil {
			log.Fatalf("could not execute template for %s: %v", theme.Name, err)
		}
		rendered = append(rendered, strings.TrimSpace(buf.String()))
	}
	return rendered
}

func indentAndJoin(themes []string, indent string) string {
	var blocks []string
	for _, theme := range themes {
		theme = strings.TrimSpace(theme)
		if theme == "" {
			continue
		}

		lines := strings.Split(theme, "\n")
		for i, line := range lines {
			if line == "" {
				lines[i] = indent
				continue
			}
			lines[i] = indent + line
		}
		blocks = append(blocks, strings.Join(lines, "\n"))
	}
	return strings.Join(blocks, ",\n")
}
