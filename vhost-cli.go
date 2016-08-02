package main

import (
	"flag"
	vhost "github.com/kevingimbel/vhost/api"
)

var (
	name     string
	tld      string
	port     string
	docroot  string
	output   string
	template string
)

func initFlags() {
	flag.StringVar(&name, "name", "test", "Server Name for the new Virtual Host")
	flag.StringVar(&tld, "tld", "local", "TLD (Top Level Domain) to use")
	flag.StringVar(&port, "port", "80", "Defines the port to serve to")
	flag.StringVar(&docroot, "docroot", "/var/www/html/", "Document Root where the files live")
	flag.StringVar(&output, "output", "./", "Where to save the file")
	flag.StringVar(&template, "template", "./templates/apache.template", "Choose which template to load. Pass full path to template")

	flag.Parse()
}

// Reads the command line flags and creates a
// Vhost configuration.
func assignVhostArguments() *vhost.Vhost {
	initFlags()

	hostConfig := new(vhost.Vhost)

	hostConfig.ServerName = name
	hostConfig.Port = port
	hostConfig.TLD = tld
	hostConfig.DocumentRoot = docroot
	hostConfig.Output = output
	hostConfig.Template = template

	return hostConfig
}

// Calls assignVhostArguments and runs vhost.CreateHost
func main() {
	config := assignVhostArguments()
	vhost.CreateHost(config)
}
