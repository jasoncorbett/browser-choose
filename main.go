package main

import "os"

func main() {
	config, err := loadConfiguration()
	if err != nil {
		displayError(err)
		os.Exit(1)
	}



}

type Configuration struct {
	Name string `yaml:"name"`
	Prompt string `yaml:"prompt"`
	Browsers []BrowserInfo `yaml:"browsers"`
}

type BrowserInfo struct {
	Command string `yaml:"cmd"`
	Name string `yaml:"name"`
	Icon string `yaml:"icon"`

}

type UrlPreference struct {
	Url string `yaml:"url"`
	Ask bool `yaml:"ask"`
}

func loadConfiguration() (*Configuration, error) {
	return &Configuration{
		Name: "Browser Choose",
		Prompt: "Which browser would you like to open for:\n{{.Url}}",
		Browsers: []BrowserInfo{
			{
				Command: "",
				Name: "",
			},
			{
				Command: "",
				Name: "",
			},
		},
	}, nil
}

func displayError(err error) {

}