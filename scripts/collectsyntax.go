package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type collector func(map[string]struct{}) error

func main() {
	keys := make(map[string]struct{})
	collectors := []collector{
		collectFromThemes("assets/themes"),
		collectFromDocumentation("docs/src/extensions/languages.md"),
		collectFromZedSyntaxToken("crates/theme_importer/src/vscode/syntax.rs"),
	}

	for _, collect := range collectors {
		if err := collect(keys); err != nil {
			if errors.Is(err, fs.ErrNotExist) || errors.Is(err, os.ErrNotExist) {
				continue
			}
			fmt.Fprintf(os.Stderr, "error collecting syntax keys: %v\n", err)
			os.Exit(1)
		}
	}

	values := make([]string, 0, len(keys))
	for key := range keys {
		values = append(values, key)
	}
	sort.Strings(values)

	out, err := json.Marshal(values)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding syntax keys: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}

func collectFromThemes(root string) collector {
	return func(keys map[string]struct{}) error {

		type themeStyle struct {
			Syntax map[string]json.RawMessage `json:"syntax"`
		}

		type themeEntry struct {
			Style *themeStyle `json:"style"`
		}

		type themeFile struct {
			Themes []themeEntry `json:"themes"`
		}

		parseFile := func(path string) error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			var theme themeFile
			if err := json.NewDecoder(file).Decode(&theme); err != nil {
				return err
			}

			for _, entry := range theme.Themes {
				if entry.Style == nil || entry.Style.Syntax == nil {
					continue
				}
				for key := range entry.Style.Syntax {
					keys[key] = struct{}{}
				}
			}

			return nil
		}

		return filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || filepath.Ext(path) != ".json" {
				return nil
			}
			if err := parseFile(path); err != nil {
				return fmt.Errorf("parse %s: %w", path, err)
			}
			return nil
		})
	}
}

func collectFromDocumentation(path string) collector {
	return func(keys map[string]struct{}) error {
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		selectSection := func(doc, heading string) string {
			index := strings.Index(doc, heading)
			if index == -1 {
				return ""
			}
			section := doc[index+len(heading):]
			if next := strings.Index(section, "\n### "); next != -1 {
				return section[:next]
			}
			return section
		}

		section := selectSection(string(content), "### Syntax highlighting")
		scanner := bufio.NewScanner(strings.NewReader(section))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if !strings.HasPrefix(line, "| @") {
				continue
			}
			parts := strings.Split(line, "|")
			if len(parts) < 2 {
				continue
			}
			token := strings.TrimSpace(parts[1])
			if token == "" {
				continue
			}
			token = strings.TrimPrefix(token, "@")
			keys[token] = struct{}{}
		}
		return scanner.Err()
	}
}

func collectFromZedSyntaxToken(path string) collector {
	return func(keys map[string]struct{}) error {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		re := regexp.MustCompile(`ZedSyntaxToken::\w+\s*=>\s*"([^"]+)"`)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			if len(matches) != 2 {
				continue
			}
			keys[matches[1]] = struct{}{}
		}
		return scanner.Err()
	}
}
