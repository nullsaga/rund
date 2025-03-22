package conf

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCanLoadConfFile(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFile := filepath.Join(filepath.Join(filepath.Dir(filename), "..", ".."), "testdata", "conf.yml")
	projectsConf, err := NewLoader().LoadConf(testFile)

	if _, err = os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Test file doesn't exist: %s", testFile)
	}

	if err != nil {
		t.Error(err)
	}

	if projectsConf == nil {
		t.Fatal("projectsConf is nil")
	}

	defaultProject := projectsConf.Projects["default"]

	if defaultProject == nil {
		t.Fatal("default project is nil")
	}

	if defaultProject.Secret != "api-secret" {
		t.Errorf("expected project config to contain 'api-secret', got %s", defaultProject.Secret)
	}

	if defaultProject.Deploy == nil {
		t.Fatal("config project does not have a deploy key")
	}

	if defaultProject.Deploy.Exec != "sh::deploy_script.sh" {
		t.Errorf("expected project deploy config to contain 'sh::deploy_script.sh', got %s", defaultProject.Deploy.Exec)
	}

	if defaultProject.Webhook == nil {
		t.Fatal("project config does not have a webhook key")
	}

	if defaultProject.Webhook.Secret != "webhook-secret" {
		t.Errorf("expected project webhook config to contain 'webhook-secret', got %s", defaultProject.Webhook.Secret)
	}

	if defaultProject.Webhook.Provider != "auto" {
		t.Errorf("expected project webhook provider to be auto, got %s", defaultProject.Webhook.Provider)
	}
}
