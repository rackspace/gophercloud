package openstack

import (
	"os"

	"github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"
)

var nilOptions = gophercloud.AuthOptions{}

// AuthOptionsFromEnv fills out an identity.AuthOptions structure with the settings found on the various OpenStack
// OS_* environment variables.  The following variables provide sources of truth: OS_AUTH_URL, OS_USERNAME,
// OS_PASSWORD, OS_TENANT_ID, and OS_TENANT_NAME.  Of these, OS_USERNAME, OS_PASSWORD, and OS_AUTH_URL must
// have settings, or an error will result.  OS_TENANT_ID and OS_TENANT_NAME are optional.
func AuthOptionsFromEnv() (gophercloud.AuthOptions, error) {
	authURL := os.Getenv("OS_AUTH_URL")
	username := os.Getenv("OS_USERNAME")
	userID := os.Getenv("OS_USERID")
	password := os.Getenv("OS_PASSWORD")
	tenantID := os.Getenv("OS_TENANT_ID")
	tenantName := os.Getenv("OS_TENANT_NAME")
	domainID := os.Getenv("OS_DOMAIN_ID")
	domainName := os.Getenv("OS_DOMAIN_NAME")

	if authURL == "" {
		return nilOptions, &ErrNoAuthURL{
			&gophercloud.InvalidInputError{
				BaseError: gophercloud.BaseError{
					Function: "openstack.AuthOptionsFromEnv",
				},
				Argument: "authURL",
			},
		}
	}

	if username == "" && userID == "" {
		return nilOptions, &ErrNoUsername{
			&gophercloud.InvalidInputError{
				BaseError: gophercloud.BaseError{
					Function: "openstack.AuthOptionsFromEnv",
				},
				Argument: "username",
			},
		}
	}

	if password == "" {
		return nilOptions, &ErrNoPassword{
			&gophercloud.InvalidInputError{
				BaseError: gophercloud.BaseError{
					Function: "openstack.AuthOptionsFromEnv",
				},
				Argument: "password",
			},
		}
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
	}

	return ao, nil
}
