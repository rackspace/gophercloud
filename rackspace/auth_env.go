package rackspace

import (
	"os"

	"github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"
)

var nilOptions = gophercloud.AuthOptions{}

func prefixedEnv(base string) string {
	value := os.Getenv("RS_" + base)
	if value == "" {
		value = os.Getenv("OS_" + base)
	}
	return value
}

// AuthOptionsFromEnv fills out an identity.AuthOptions structure with the
// settings found on the various Rackspace RS_* environment variables.
func AuthOptionsFromEnv() (gophercloud.AuthOptions, error) {
	authURL := prefixedEnv("AUTH_URL")
	username := prefixedEnv("USERNAME")
	password := prefixedEnv("PASSWORD")
	apiKey := prefixedEnv("API_KEY")

	if authURL == "" {
		return nilOptions, &ErrNoAuthURL{
			&gophercloud.InvalidInputError{
				BaseError: gophercloud.BaseError{
					Function: "rackspace.AuthOptionsFromEnv",
				},
				Argument: "authURL",
			},
		}
	}

	if username == "" {
		return nilOptions, &ErrNoUsername{
			&gophercloud.InvalidInputError{
				BaseError: gophercloud.BaseError{
					Function: "rackspace.AuthOptionsFromEnv",
				},
				Argument: "username",
			},
		}
	}

	if password == "" && apiKey == "" {
		return nilOptions, &ErrNoPassword{
			&gophercloud.InvalidInputError{
				BaseError: gophercloud.BaseError{
					Function: "rackspace.AuthOptionsFromEnv",
				},
				Argument: "password",
			},
		}
	}

	ao := gophercloud.AuthOptions{
		IdentityEndpoint: authURL,
		Username:         username,
		Password:         password,
		APIKey:           apiKey,
	}

	return ao, nil
}
