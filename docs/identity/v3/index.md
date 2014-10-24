---
layout: page
title: Getting Started with Identity v3
---

* [Setup](#setup)
* [Tokens](#tokens)
  * [Generate token](#create-token)
  * [Get token](#get-token)
  * [Validate token](#validate-token)
  * [Revoke token](#revoke-token)
* [Service catalog](#services)
  * [Add service](#create-service)
  * [List services](#list-services)
  * [Show service details](#get-service)
  * [Delete service](#delete-service)
* [Endpoints](#endpoints)
  * [Add endpoint](#create-endpoint)
  * [List endpoints](#list-endpoints)
  * [Show endpoint details](#get-endpoint)
  * [Delete endpoint](#delete-endpoint)

## <a name="setup"></a>Setup

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

## <a name="tokens"></a>Tokens

A token is an arbitrary bit of text that is returned when you authenticate,
which you subsequently use as to access and control API resources. Each
token has a scope that describes which resources are accessible with it. A
token may be revoked at anytime and is valid for a finite duration.

### <a name="create-token"></a>Generate a token

The configuration options that you need to pass into `tokens.AuthOptions` will
mostly depend on your provider: some authenticate with API keys rather than
passwords, some require tenant names or tenant IDs; with others, both are
required. It is therefore useful to double check what the expectations are
before authenticating.

Scoping was introduced in v3, which allows you to confine the token's scope to
a particular project. If no scope is provided, and the user has a default one
set, that will be used; if there is no default, the token returned will not
have any explicit scope defined.

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"

opts := tokens.AuthOptions{
  IdentityEndpoint: "{identityEndpoint}",
  Username: "{username}",
  Password: "{password}",
  TenantID: "{tenantID}",
}

// This is completely optional
scope := tokens.Scope{ProjectName: "tmp_project"}

// Make the call
token, err := tokens.Create(client, opts, scope).Extract()
{% endhighlight %}

### <a name="get-token"></a>Get token

{% highlight go %}
token, err := tokens.Get(client, "token_id").Extract()
{% endhighlight %}

### <a name="validate-token"></a>Validate token

To check whether an existing token is still valid (i.e. whether it has expired
  or not), you can validate a token ID and get back a boolean value.

{% highlight go %}
valid, err := tokens.Validate(client, "token_id")
{% endhighlight %}

### <a name="revoke-token"></a>Revoke token

Revoking a token will prevent it being used in further API calls.

{% highlight go %}
err := tokens.Revoke(client, "token_id")
{% endhighlight %}

## <a name="services"></a>Service catalog

A service is a RESTful API that controls the functionality of an OpenStack
service - such as as Compute, Object Storage, etc. It provides one or more
endpoints through which users can access resources and perform operations.

### <a name="create-service"></a>Add service

The only parameter required when creating a new service is the "type" in string
form:

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack/identity/v3/services"

service, err := services.Create(client, "service_type").Extract()
{% endhighlight %}

### <a name="list-services"></a>List services

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/openstack/identity/v2/services"
)

// We have the option of filtering the service list. If we want the full
// collection, leave it as an empty struct
opts := services.ListOpts{ServiceType: "foo_type", Limit: 5}

// Retrieve a pager (i.e. a paginated collection)
pager := services.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  serviceList, err := services.ExtractServices(page)

  for _, s := range serviceList {
    // "s" will be a services.Service
  }
})
{% endhighlight %}

### <a name="get-service"></a>Show service details

In order to retrieve the details for a particular service, all you need is its
UUID in string form:

{% highlight go %}
service, err := services.Get(client, "service_id").Extract()
{% endhighlight %}

### <a name="update-service"></a>Update service

The only modifiable attribute for a service is its type. To set a new one:

{% highlight go %}
service, err := services.Update(client, "service_id", "new_type").Extract()
{% endhighlight %}

### <a name="delete-service"></a>Delete service

To permanently delete a service from the catalog, just pass in its UUID like so:

{% highlight go %}
err := services.Delete(client, "service_id")
{% endhighlight %}


## <a name="endpoints"></a>Endpoints

An endpoint is a network-accessible address, usually described by a URL, where
a service may be accessed. If using an extension for templates, you can create
an endpoint template, which represents the templates of all the consumable
services that are available across the regions.

### <a name="create-endpoint"></a>Add endpoint

{% highlight go %}
import (
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/openstack/identity/v3/endpoints"
)

// All fields except Region are required
opts := endpoints.EndpointOpts{
  Availability: gophercloud.AvailabilityPublic,
  Name: "backup_endpoint",
  Region: "Region4",
  URL: "backup.my-openstack.org",
  ServiceID: "service_id",
}

endpoint, err := endpoints.Create(client, opts).Extract()
{% endhighlight %}

### <a name="list-endpoints"></a>List endpoints

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/openstack/identity/v2/endpoints"
)

// We have the option of filtering the endpoint list. If we want the full
// collection, leave it as an empty struct
opts := endpoints.ListOpts{ServiceID: "service_id", Limit: 5}

// Retrieve a pager (i.e. a paginated collection)
pager := endpoints.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  endpointList, err := endpoints.ExtractEndpoints(page)

  for _, e := range endpointList {
    // "e" will be a endpoints.Endpoint
  }
})
{% endhighlight %}

### <a name="update-endpoint"></a>Update endpoint

All fields are modifiable and are optional.

{% highlight go %}
import (
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/openstack/identity/v3/endpoints"
)

opts := endpoints.EndpointOpts{Name: "new_name"}

endpoint, err := endpoints.Update(client, "endpoint_id", opts).Extract()
{% endhighlight %}

### <a name="delete-endpoint"></a>Delete endpoint

{% highlight go %}
err := endpoints.Delete(client, "endpoint_id")
{% endhighlight %}
