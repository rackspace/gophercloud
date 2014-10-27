---
layout: page
title: Getting Started with Identity v2
---

## Setup

In order to interact with OpenStack APIs, you must first pass in your auth
credentials to a `Provider` struct. Once you have this, you then retrieve
whichever service struct you're interested in - so in our case, we invoke the
`NewIdentityV2` method:

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack"

authOpts, err := openstack.AuthOptionsFromEnv()

provider, err := openstack.AuthenticatedClient(authOpts)

client, err := openstack.NewIdentityV2(provider, gophercloud.EndpointOpts{
	Region: "RegionOne",
})
{% endhighlight %}

If you're unsure about how to retrieve credentials, please read our [introductory
guide](/docs) which outlines the steps you need to take.

## Tokens

A token is an arbitrary bit of text that is returned when you authenticate,
which you subsequently use as to access and control API resources. Each
token has a scope that describes which resources are accessible with it. A
token may be revoked at anytime and is valid for a finite duration.

### Generate a token

The configuration options that you need to pass into `tokens.AuthOptions` will
mostly depend on your provider: some authenticate with API keys rather than
passwords, some require tenant names or tenant IDs; with others, both are
required. It is therefore useful to double check what the expectations are
before authenticating.

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack/identity/v2/tokens"

opts := tokens.AuthOptions{
	IdentityEndpoint: "{identityEndpoint}",
	Username: "{username}",
	Password: "{password}",
	TenantID: "{tenantID}",
}

token, err := tokens.Create(client, opts).Extract()
{% endhighlight %}

## Tenants

A tenant is a container used to group or isolate resources and/or identity
objects. Depending on the service operator, a tenant can map to a customer,
account, organization, or project.

### List tenants

{% highlight go %}
import (
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/identity/v2/tenants"
)

// We have the option of filtering the tenant list. If we want the full
// collection, leave it as an empty struct
opts := tenants.ListOpts{Limit: 5}

// Retrieve a pager (i.e. a paginated collection)
pager := tenants.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
	tenantList, err := tenants.ExtractTenants(page)

	for _, t := range tenantList {
		// "t" will be a tenants.Tenant
	}
})
{% endhighlight %}
