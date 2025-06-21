package auth

import (
	"encoding/json"
	"filmPrice/config"
	"testing"
)

func TestParseToken(t *testing.T) {
	config.LoadConfigYaml("")
	token, err := GenToken("test")
	if err != nil {
		t.Error(err)
	}

	t.Log(token)

	claims, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}

	s, _ := json.Marshal(claims)
	t.Log(string(s))
}
