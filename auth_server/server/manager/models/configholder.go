package models

import (
	"github.com/cesanta/docker_auth/auth_server/authn"
	"github.com/cesanta/docker_auth/auth_server/authz"
	"github.com/cesanta/docker_auth/auth_server/server/config"

	"sort"
)

type ByString []string

func (s ByString) Len() int {
	return len(s)
}

func (s ByString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByString) Less(i, j int) bool {
	return s[i] < s[j]
}

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
	userMap := acm.authConfig.Config.Users
	users := make([]string, len(userMap))
	i := 0
	for key, _ := range userMap {
		users[i] = key
		i = i + 1
	}

	sort.Sort(ByString(users))

	return users
}
