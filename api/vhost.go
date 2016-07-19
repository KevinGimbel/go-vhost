package vhost

import (
  "os"
  "text/template"
  "log"
)

// Defines the vhost structure
type Vhost struct {
  ServerName *string
  TLD *string
  Port *string
  DocumentRoot *string
  Output *string
  Template *string
}

// createConfiguration takes a template and Vhost struct object to
// create the new configuration and save it to the path.
func CreateConfiguration(path string, config *Vhost) {
	t, err := template.ParseFiles(path)
	if err != nil {
    log.Print(err)
    return
	}
  // Construct the file name
  outputPath := *config.Output + *config.ServerName + ".conf"

  f, err := os.Create(outputPath)

  if err != nil {
    log.Println("create file: ", err)
    return
  }

  err = t.Execute(f, config)
  if err != nil {
    log.Print("execute: ", err)
    return
  }
  f.Close()
}

// Main function to create a new virtual host entry
// Takes a Vhost struct as parameter
func CreateHost(config *Vhost) {
  pwd, _ := os.Getwd()
  var template = pwd + "/templates/apache.template"

  if config.Template != nil {
    template = *config.Template
  }

  CreateConfiguration(template, config)
}
