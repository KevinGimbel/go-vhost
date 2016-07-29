package vhost

import (
	"log"
	"os"
	"text/template"
)

// Defines the vhost structure
type Vhost struct {
	ServerName   string
	TLD          string
	Port         string
	DocumentRoot string
	Output       string
	Template     string
}

// Creates the output directory if it doesn't exist already
func CreateOutputDirectory(path string) (string, error) {
	err := os.MkdirAll(path, 0755)

	if err != nil {
		return "Unable to create directory", err
	}

	return "success", nil
}

// Creates the output file and populates it with the variables from the
// Vhost type.
func CreateConfigurationFile(path string, config *Vhost) (string, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return "Unable to parse template file", err
	}

	outputPath := config.Output + config.ServerName + ".conf"

	f, err := os.Create(outputPath)
	if err != nil {
		return "Unable to create tempalte file", err
	}

	err = t.Execute(f, config)
	if err != nil {
		return "Unable to execute template", err
	}

	f.Close()
	return "Successfully created file", nil
}

// createConfiguration takes a template and Vhost struct object to
// create the new configuration and save it to the path.
func CreateConfiguration(path string, config *Vhost) {
	msg, err := CreateOutputDirectory(config.Output)
	if err != nil {
		log.Fatal(msg, err)
	}

	msg, err = CreateConfigurationFile(path, config)
	if err != nil {
		log.Fatal(msg, err)
	}

	log.Println(msg)
}

// Main function to create a new virtual host entry
// Takes a Vhost struct as parameter
func CreateHost(config *Vhost) {
	pwd, _ := os.Getwd()
	var template = pwd + "/templates/apache.template"

	if config.Template != "" {
		template = config.Template
	}

	CreateConfiguration(template, config)
}
