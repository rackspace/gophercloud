package trust

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/identity/v3/extensions"
	trusttokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/extensions/tokens"
)

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3Trust(client *gophercloud.ProviderClient, options extensions.AuthOptions) error {
	if options.TrustID != "" {
		return trustv3auth(client, "", options)
	} else {
		return openstack.AuthenticateV3(client, *options.AuthOptions)
	}
}

func trustv3auth(client *gophercloud.ProviderClient, endpoint string, options extensions.AuthOptions) error {
	//In case of Trust TokenId would be Provided so we have to populate the value in service client
	//to not throw password error,also if it is not provided it will be empty which maintains
	//the current implementation.
	client.TokenID = options.TokenID
	// Override the generated service endpoint with the one returned by the version endpoint.
	v3Client := openstack.NewIdentityV3(client)
	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	// copy the auth options to a local variable that we can change. `options`
	// needs to stay as-is for reauth purposes
	v3Options := options

	var scope *trusttokens3.Scope

	scope = &trusttokens3.Scope{
		TrustID: options.TrustID,
	}

	result := trusttokens3.Create(v3Client, v3Options, scope)

	token, err := result.ExtractToken()
	if err != nil {
		return err
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return err
	}

	client.TokenID = token.ID

	if options.AllowReauth {
		client.ReauthFunc = func() error {
			client.TokenID = ""
			return trustv3auth(client, endpoint, options)
		}
	}
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return trusttokens3.TrustV3EndpointURL(catalog, opts)
	}

	return nil
}
