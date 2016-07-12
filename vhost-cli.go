package main

// Import the vhost "api"
// and the flags module(?)
import (
  vhost "github.com/kevingimbel/vhost/api"
  "flag"
)

var name = flag.String("name", "test", "Server Name for the new Virtual Host")
var tld = flag.String("tld", "local", "TLD (Top Level Domain) to use.")
var port  = flag.String("port", "80", "Defines the port to serve to.")
var docroot = flag.String("docroot", "/var/www/html/", "Document Root where the files live")
var output = flag.String("output", "./", "Where to save the file")

func assignVhostArguments() *vhost.Vhost {
  flag.Parse()

  hostConfig := new(vhost.Vhost)

  hostConfig.ServerName = name
  hostConfig.Port = port
  hostConfig.TLD = tld
  hostConfig.DocumentRoot = docroot
  hostConfig.Output = output;

  return hostConfig
}

func main() {
  config := assignVhostArguments()
  vhost.CreateHost(config)
}
