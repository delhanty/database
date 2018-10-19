package secrets

import (
	"fmt"
	"strconv"
)

// DBSecrets for accessing database connection strings
type DBSecrets struct {
	Region      string
	Environment string
	secrets     Secrets
}

// NewDBSecrets returns an instance for acquiring the database connection string from
// either local env vars or AWS
func NewDBSecrets(env, region string) *DBSecrets {
	s := &DBSecrets{Environment: env, Region: region}
	if s.Environment != "local" {
		s.secrets = NewAWSSecrets(region)
	} else {
		s.secrets = NewEnvSecrets()
	}
	return s
}

// DBString returns the database connection string for the environment and service
func (s *DBSecrets) DBString(serviceKey string) (string, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/db/%s/dbstring", s.Environment, serviceKey))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ServicePassword retrieves the password for the initialized servicekey
func (s *DBSecrets) ServicePassword(serviceKey string) (string, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/db/%s/pwd", s.Environment, serviceKey))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *DBSecrets) CacheConfig() (string, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/cache/config", s.Environment))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *DBSecrets) DiscoveryAddr() (string, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/discovery/config", s.Environment))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *DBSecrets) LoadBalancerAddr() (string, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/loadbalancer/config", s.Environment))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *DBSecrets) SystemOrgID() (int, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/system/orgid", s.Environment))
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(data))
}

func (s *DBSecrets) SystemUserID() (int, error) {
	data, err := s.secrets.GetSecureParameter(fmt.Sprintf("/am/%s/system/userid", s.Environment))
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(data))
}
