package vhost

import (
  "os"
  "text/template"
  "log"
)

// Define the vhost structure
type Vhost struct {
  ServerName *string
  TLD *string
  Port *string
  DocumentRoot *string
  Output *string
}

// function to create the virtual host configuration file.
// takes the output path and the vhost config (Vhost struct) as arguments
func createConfiguration(path string, config *Vhost) {
	t, err := template.ParseFiles(path)
	if err != nil {
    log.Print(err)
    return
	}

  f, err := os.Create("./test.conf")

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
func CreateHost(config *Vhost) {
  pwd, _ := os.Getwd()
  createConfiguration(pwd + "/templates/apache.template", config)
}
