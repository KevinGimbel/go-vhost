package main

import (
  "os"
  "text/template"
  "flag"
  "log"
  "bytes"
)

type Vhost struct {
  ServerName *string
  TLD *string
  Port *string
  DocumentRoot *string
  Output *string
}

var name = flag.String("name", "test", "Server Name for the new Virtual Host")
var tld = flag.String("tld", "local", "TLD (Top Level Domain) to use.")
var port  = flag.String("port", "80", "Defines the port to serve to.")
var docroot = flag.String("docroot", "/var/www/html/", "Document Root where the files live")
var output = flag.String("output", "./", "Where to save the file")

func assignVhostArguments() *Vhost {
  flag.Parse()

  hostConfig := new(Vhost)

  hostConfig.ServerName = name
  hostConfig.Port = port
  hostConfig.TLD = tld
  hostConfig.DocumentRoot = docroot
  hostConfig.Output = output;

  return hostConfig
}

func createConfiguration(path string, config *Vhost) {
  var buffer bytes.Buffer

	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return
	}
  buffer.WriteString(*config.Output)
  buffer.WriteString(*config.ServerName)
  buffer.WriteString(".conf")
  filename := buffer.String()

  f, err := os.Create(filename)
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

func main() {
  config := assignVhostArguments()
  pwd, _ := os.Getwd()
  createConfiguration(pwd + "/templates/apache.template", config)
}
