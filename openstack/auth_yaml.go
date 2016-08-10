package openstack

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rackspace/gophercloud"
	"gopkg.in/yaml.v2"
)

var nilOptions = gophercloud.AuthOptions{}

var (
	ErrNoCloudYaml = fmt.Errorf("Clouds.yaml file not found.")
)

type Clouds struct {
	Clouds map[string]struct {
		Profile       string   `yaml:"profile"`
		RegionName    string   `yaml:"region_name"`
		Regions       []string `yaml:"regions"`
		DNSAPIVersion int      `yaml:"dns_api_version"`
		Auth          struct {
			AuthURL     string `yaml:"auth_url"`
			Username    string `yaml:"username"`
			Password    string `yaml:"password"`
			ProjectName string `yaml:"project_name"`
			ProjectID   int    `yaml:"project_id"`
		} `yaml:"auth,omitempty"`
	}
}

func AuthOptionsFromYaml() (gophercloud.AuthOptions, error) {
	clouds := &Clouds{}
	cloudsContent, err := cloudsYaml()
	if err != nil {
		return nilOptions, err
	}
	err = yaml.Unmarshal(cloudsContent, clouds)
	if err != nil {
		return nilOptions, fmt.Errorf("Failed to unmarshal YAML: %v", err)
	}

	// Where do we get that?
	profile := "someprofile"
	auth := clouds.Clouds[profile].Auth

	authURL := auth.AuthURL
	username := auth.Username
	//userID := auth.UserID
	password := auth.Password
	//tenantID := os.Getenv("OS_TENANT_ID")
	//tenantName := os.Getenv("OS_TENANT_NAME")
	//domainID := os.Getenv("OS_DOMAIN_ID")
	//domainName := os.Getenv("OS_DOMAIN_NAME")
	//tokenID := os.Getenv("OS_TOKEN")

	if authURL == "" {
		return nilOptions, ErrNoAuthURL
	}

	if username == "" && userID == "" && tokenID == "" {
		return nilOptions, ErrNoUsername
	}

	if password == "" && tokenID == "" {
		return nilOptions, ErrNoPassword
	}

	ao := gophercloud.AuthOptions{
		IdentityEndpoint: authURL,
		UserID:           userID,
		Username:         username,
		Password:         password,
		TenantID:         tenantID,
		TenantName:       tenantName,
		DomainID:         domainID,
		DomainName:       domainName,
		TokenID:          tokenID,
	}

	return ao, nil
}

func cloudsYaml() ([]bytes, error) {
	var err error
	if f, ok := cloudsYamlFile(os.Getenv("USER_CONFIG_DIR")); ok {
		return ioutil.ReadFile(f)
	}

	if f, ok := cloudsYamlFile(fmt.Sprintf("%s/.config", os.Getenv("HOME"))); ok {
		return ioutil.ReadFile(f)
	}

	if f, ok := cloudsYamlFile(os.Getenv("SITE_CONFIG_DIR")); ok {
		return ioutil.ReadFile(f)
	}

	if f, ok := cloudsYamlFile("/etc/openstack"); ok {
		return ioutil.ReadFile(f)
	}

	return []bytes{}, ErrNoCloudYaml
}

func cloudsYamlFile(dirname string) (string, bool) {
	filename := fmt.Sprintf("%s/clouds.yaml", dirname)
	if f, err := os.Stat(filename); err == nil && f.Mode().IsRegular() {
		return filename, true
	}
	return "", false
}
