package models

import (
	"github.com/cesanta/docker_auth/auth_server/authn"
	"github.com/cesanta/docker_auth/auth_server/authz"
	"github.com/cesanta/docker_auth/auth_server/server/config"
)

type AuthConfig struct {
	Config         *config.Config
	Authenticators []authn.Authenticator
	Authorizers    []authz.Authorizer
}

type AuthConfigManager struct {
	authConfig *AuthConfig
}

var ACManager *AuthConfigManager

func InitAuthConfigManager(config *config.Config, authenticators []authn.Authenticator, authorizers []authz.Authorizer) {
	authConfig := &AuthConfig{Config: config, Authenticators: authenticators, Authorizers: authorizers}
	ACManager = &AuthConfigManager{authConfig: authConfig}
}

func (acm *AuthConfigManager) QueryAllUser() []string {
	authenticators := acm.authConfig.Authenticators
	users := make([]string, len(authenticators))
	for _, authenticator := range authenticators {
		users = append(users, authenticator.Name())
	}

	return users
}
