package vhost

import (
	"os"
	"testing"
)

// Dummy Vhost configuration
var dummyConfig = &Vhost{
	ServerName:   "test",
	TLD:          "local",
	Port:         "8080",
	DocumentRoot: "/var/www/",
	Output:       "/tmp/vhost_test",
	Template:     "../../template/apache.template"}

func TestCreateDirectory(t *testing.T) {
	msg, err := CreateOutputDirectory("/tmp/vhost_test/")
	if err != nil {
		t.Error(msg, err)
	}
}

func TestCreateConfigurationFile(t *testing.T) {
	pwd, _ := os.Getwd()
	var templateFile = pwd + "/../templates/apache.template"

	msg, err := CreateConfigurationFile(templateFile, dummyConfig)
	if err != nil {
		t.Error(msg, err)
	}
}
