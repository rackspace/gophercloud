// +build acceptance

package v3

import (
	"testing"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	tokens3 "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

func TestGetToken(t *testing.T) {
	// Obtain credentials from the environment.
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to acquire credentials: %v", err)
	}

	// Trim out unused fields. Skip if we don't have a UserID.
	ao.Username, ao.TenantID, ao.TenantName = "", "", ""
	if ao.UserID == "" {
		t.Logf("Skipping identity v3 tests because no OS_USERID is present.")
		return
	}

	// Create an unauthenticated client.
	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	if err != nil {
		t.Fatalf("Unable to instantiate client: %v", err)
	}

	// Create a service client.
	service := openstack.NewIdentityV3(provider)

	// Use the service to create a token.
	token, err := tokens3.Create(service, ao, nil).Extract()
	if err != nil {
		t.Fatalf("Unable to get token: %v", err)
	}

	t.Logf("Acquired token: %s", token.ID)
}

func TestTokenAuth(t *testing.T) {
	// Create a service client.
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to acquire credentials: %v", err)
	}

	// Save the tenant name
	tenantName := ao.TenantName

	// Trim out unused fields.
	ao.TenantID, ao.TenantName = "", ""

	if ao.DomainID == "" {
		ao.DomainID = "default"
	}

	providerClient, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		t.Fatalf("%s", err)
	}

	// Create a service client.
	serviceClient := openstack.NewIdentityV3(providerClient)

	// Get a token
	scope := &tokens3.Scope{
		ProjectName: tenantName,
		DomainID:    ao.DomainID,
	}

	result := tokens3.Create(serviceClient, ao, scope)
	token, err := result.ExtractToken()
	if err != nil {
		t.Fatalf("Could not generate a token: %v", err)
	}

	// Now create a new session with the token
	newAO := gophercloud.AuthOptions{
		TokenID:          token.ID,
		IdentityEndpoint: ao.IdentityEndpoint,
	}

	newProviderClient, err := openstack.AuthenticatedClient(newAO)
	if err != nil {
		t.Fatalf("%s", err)
	}

	newClient := openstack.NewIdentityV3(newProviderClient)
	if newClient == nil {
		t.Fatalf("Could not create a client based on a token")
	}
}
