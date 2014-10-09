---
layout: page
---

# Getting Started with gophercloud

Before working with OpenStack services like Compute or Object Storage, you must
first authenticate. The first step is to populate the `gophercloud.AuthOptions`
struct with your access details, and you can do this in two ways:

```go
// Option 1: Pass in the values yourself
opts := gophercloud.AuthOptions{
  IdentityEndpoint: "https://my-openstack.com:5000/v2.0",
  Username: "{username}",
  Password: "{password}",
  TenantID: "{tenant_id}",
}

// Option 2: Use a utility function to retrieve all your environment variables
import "github.com/rackspace/gophercloud/openstack/utils"
opts, err := utils.AuthOptions()
```

Once you have an `opts` variable, you can pass it in and get back a
`ProviderClient` struct:

```go
import "github.com/rackspace/gophercloud/openstack"

provider, err := openstack.AuthenticatedClient(opts)
```

The `ProviderClient` is the top-level client that all of your OpenStack services
derive from. The provider contains all of the authentication details that allow
your Go code to access the API - such as the base URL and token ID.

## Next steps

Cool! You've handled authentication and got your `ProviderClient`. You're now
ready to use an OpenStack service.

* [Getting started with Compute](./compute)
* [Getting started with Object Storage](./object-storage)
* [Getting started with Networking](./networking)
* [Getting started with Block Storage](./block-storage)
* [Getting started with Identity](./identity)
