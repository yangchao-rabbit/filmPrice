package config

import "testing"

func TestLoadConfigYaml(t *testing.T) {
	LoadConfigYaml("E:\\gitClone\\k8s-scale-go\\config.yaml")

	t.Log(Get().App.Name)
}
