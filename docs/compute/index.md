---
layout: page
title: Getting Started with Compute v2
---

* [Setup](#setup)
* [Flavors](#flavors)
  * [List flavors](#list-flavors)
  * [Get flavor](#get-flavor)
  * [Get flavor ID](#get-flavor-id)
* [Images](#images)
  * [List images](#list-images)
  * [Get image](#get-image)
  * [Get image ID](#get-image-id)
  * [Delete image](#delete-image)
* [Key Pairs](#keypairs)
  * [Create a Key Pair](#create-keypair)
  * [Delete a Key Pair](#delete-keypair)
* [Security Groups](#secgroups)
  * [Create a Security Group](#create-secgroup)
  * [Delete a Security Group](#delete-secgroup)
* [Server Groups](#servergroups)
  * [Create a Server Group](#create-servergroup)
  * [Delete a Server Group](#delete-servergroup)
* [Servers](#servers)
  * [Create a server](#create-server)
  * [List servers](#list-servers)
  * [Get server](#get-server)
  * [Update server](#update-server)
  * [Delete server](#delete-server)
  * [Create image](#create-image)
  * [Change admin password](#change-password)
  * [Rebuild](#rebuild)
  * [Resize](#resize)
  * [Confirm resize](#confirm)
  * [Revert resize](#revert)
  * [Start](#start-server)
  * [Stop](#stop-server)
  * [Boot from volume](#boot-from-volume)
  * [Scheduler Hints](#schedulerhints)
* [More](#more)
* [Providers](#providers)
  * [Rackspace](#rackspace)

# <a name="setup"></a>Setup

In order to interact with OpenStack APIs, you must first pass in your auth
credentials to a `Provider` struct. Once you have this, you then retrieve
whichever service struct you're interested in - so in our case, we invoke the
`NewComputeV2` method:

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack"

authOpts, err := openstack.AuthOptionsFromEnv()

provider, err := openstack.AuthenticatedClient(authOpts)

client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
	Region: "RegionOne",
})
{% endhighlight %}

If you're unsure about how to retrieve credentials, please read our [introductory
guide](/docs) which outlines the steps you need to take.

# <a name="flavors"></a>Flavors

A flavor is a hardware configuration for a server. Each one has a unique
combination of disk space, memory capacity and priority for CPU time.

### <a name="list-flavors"></a>List all available flavors

{% highlight go %}
import (
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
)

// We have the option of filtering the flavor list. If we want the full
// collection, leave it as an empty struct
opts := flavors.ListOpts{ChangesSince: "2014-01-01T01:02:03Z", MinRAM: 4}

// Retrieve a pager (i.e. a paginated collection)
pager := flavors.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
	flavorList, err := networks.ExtractFlavors(page)

	for _, f := range flavorList {
		// "f" will be a flavors.Flavor
	}
})
{% endhighlight %}

### <a name="get-flavor"></a>Get details for a specific flavor

In order to retrieve information for a specific flavor, you need its UUID in
string form. You receive back a `flavors.Flavor` struct with `ID`, `Disk`, `RAM`,
`Name`, `RxTxFactor`, `Swap` and `VCPUs` fields.

{% highlight go %}
// Get back a flavors.Flavor struct
flavor, err := flavors.Get(client, "flavor_id").Extract()
{% endhighlight %}

### <a name="get-flavor-id"></a>Get a flavor ID from a flavor Name

You can look up the ID of a flavor by using its human-readable name.

{% highlight go %}
// Get back a flavor ID string
flavorID, err := flavors.IDFromName(client, "flavor_name").Extract()
{% endhighlight %}

# <a name="images"></a>Images

An image is the operating system for a VM - a collection of files used to
create or rebuild a server. Operators provide a number of pre-built OS images
by default, but you may also create custom images from cloud servers you have
launched.

### <a name="list-images"></a>List all available images

{% highlight go %}
import (
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
)

// We have the option of filtering the image list. If we want the full
// collection, leave it as an empty struct
opts := images.ListOpts{ChangesSince: "2014-01-01T01:02:03Z", Name: "Ubuntu 12.04"}

// Retrieve a pager (i.e. a paginated collection)
pager := images.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
	imageList, err := images.ExtractImages(page)

	for _, i := range imageList {
		// "i" will be a images.Image
	}
})
{% endhighlight %}

### <a name="get-image"></a>Get details for a specific image

In order to retrieve information for a specific flavor, you need its UUID in
string form. You receive back an `images.Image` struct with `ID`, `Created`, `MinDisk`,
`MinRAM`, `Name`, `Progress`, `Status` and `Updated` fields.

{% highlight go %}
// Get back a images.Image struct
image, err := images.Get(client, "image_id").Extract()
{% endhighlight %}

### <a name="get-image-id"></a>Get an image ID from an image Name

You can look up the ID of an image by using its human-readable name.

{% highlight go %}
// Get back an image ID string
imageID, err := images.IDFromName(client, "image_name").Extract()
{% endhighlight %}

### <a name="delete-image"></a>Delete an image

{% highlight go %}
res := images.Delete(client, "image_id")
{% endhighlight %}

# <a name="keypairs"></a>Key Pairs

### <a name="create-keypair"></a>Create a Key Pair

{% highlight go %}
import (
  "crypto/rand"
  "crypto/rsa"

  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"

  "golang.org/x/crypto/ssh"
)

privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
publicKey := privateKey.PublicKey
pub, err := ssh.NewPublicKey(&publicKey)
pubBytes := ssh.MarshalAuthorizedKey(pub)
pk := string(pubBytes)

kp, err := keypairs.Create(client, keypairs.CreateOpts{
  Name:      "keypair_name",
  PublicKey: pk,
}).Extract()
{% endhighlight %}


### <a name="delete-keypair"></a>Delete a Key Pair

{% highlight go %}
err := keypairs.Delete(client, "keypair_name").ExtractErr()
{% endhighlight %}

# <a name="secgroups"></a>Security Groups

### <a name="create-secgroup"></a>Create a Security Group

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/secgroups"
)

opts := secgroups.CreateOpts{
  Name:        "MySecGroup",
  Description: "something",
}

group, err := secgroups.Create(client, opts).Extract()

opts := secgroups.CreateRuleOpts{
  ParentGroupID: group.ID,
  FromPort:      22,
  ToPort:        22,
  IPProtocol:    "TCP",
  CIDR:          "0.0.0.0/0",
}

rule, err := secgroups.CreateRule(client, opts).Extract()
{% endhighlight %}

### <a name="delete-secgroup"></a>Delete a Security Group

{% highlight go %}
err := secgroups.Delete(client, group.ID).ExtractErr()
{% endhighlight %}

# <a name="servergroups"></a>Server Groups

### <a name="create-servergroup"></a>Create a Server Group

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/servergroups"
)

sg, err := servergroups.Create(computeClient, &servergroups.CreateOpts{
  Name:     "test",
  Policies: []string{"affinity"},
}).Extract()

`Policies` can either be `affinity` or `anti-affinity`.

{% endhighlight %}

### <a name="delete-servergroup"></a>Delete a Server Group

{% highlight go %}
err := servergroups.Delete(computeClient, "servergroup_id").ExtractErr()
{% endhighlight %}

# <a name="servers"></a>Servers

A server is either a virtual machine (VM) instance or a physical device
managed by the compute system.

### <a name="create-server"></a>Create a server

To create a server with a minimal configuration, do:

{% highlight go %}
server, err := servers.Create(client, servers.CreateOpts{
  Name:      name,
  FlavorName: "flavor_name",
  ImageName: "image_name",
}).Extract()
if err != nil {
  fmt.Println("Unable to create server: %s", err)
}
fmt.Println("Server ID: %s", server.ID)
{% endhighlight %}

See the [Go Documentation](http://godoc.org/github.com/rackspace/gophercloud/openstack/compute/v2/servers)
for a full list of Create options.

### <a name="list-servers"></a>List all available servers

{% highlight go %}
import (
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

// We have the option of filtering the server list. If we want the full
// collection, leave it as an empty struct
opts := servers.ListOpts{Name: "server_1"}

// Retrieve a pager (i.e. a paginated collection)
pager := servers.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
	serverList, err := servers.ExtractServers(page)

	for _, s := range serverList {
		// "s" will be a servers.Server
	}
})
{% endhighlight %}

### <a name="get-server"></a>Get details for a server

{% highlight go %}
// We need the UUID in string form
server, id := servers.Get(client, "server_id").Extract()
{% endhighlight %}

### <a name="update-server"></a>Update an existing server

{% highlight go %}
opts := servers.UpdateOpts{Name: "new_name"}

server, err := servers.Update(client, "server_id", opts).Extract()
{% endhighlight %}

### <a name="delete-server"></a>Delete an existing server

{% highlight go %}
result := servers.Delete(client, "server_id")
{% endhighlight %}

### <a name="create-image"></a>Create image

{% highlight go %}
opts = servers.CreateImageOpts{Name: "My Image"}
image_id, err = servers.CreateImage(client, "server_id", opts)
{% endhighlight %}

### <a name="change-password"></a>Change admin password

{% highlight go %}
result := servers.ChangeAdminPassword(client, "server_id", "newPassword_&123")
{% endhighlight %}

### <a name="reboot"></a>Reboot a server

There are two different methods for rebooting a VM: soft or hard reboot. A
soft reboot instructs the operating system to initiate its own restart procedure,
whereas a hard reboot cuts power (if a physical machine) or teminates the
instance at the hypervisor level (if a virtual machine).

{% highlight go %}
// You have a choice of two reboot methods: servers.SoftReboot or servers.HardReboot
result := servers.Reboot(client, "server_id", servers.SoftReboot)
{% endhighlight %}

### <a name="rebuild"></a>Rebuild a server

The rebuild operation removes all data on the server and replaces it with the
image specified. The server's existing ID and all IP addresses will remain the
same.

{% highlight go %}
// You have the option of specifying additional options
opts := RebuildOpts{
	Name:      "new_name",
	AdminPass: "admin_password",
	ImageID:   "image_id",
	Metadata:  map[string]string{"owner": "me"},
}

result := servers.Rebuild(client, "server_id", opts)

// You can extract a servers.Server struct from the HTTP response
server, err := result.Extract()
{% endhighlight %}

### <a name="resize"></a>Resize a server

{% highlight go %}
result := servers.Resize(client, "server_id", "new_flavor_id")
{% endhighlight %}

### <a name="confirm"></a>Confirm a resize operation

{% highlight go %}
result := servers.ConfirmResize(client, "server_id")
{% endhighlight %}

### <a name="revert"></a>Revert a resize operation

{% highlight go %}
result := servers.RevertResize(client, "server_id")
{% endhighlight %}

### <a name="stop-server"></a>Stop a Server

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
)

err := startstop.Stop(client, "server_id").ExtractErr()
{% endhighlight %}

### <a name="start-server"></a>Start a Server

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
)

err := startstop.Start(client, "server_id").ExtractErr()
{% endhighlight %}


### <a name="boot-from-volume"></a>Boot from a Volume

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
)

bd := []bootfromvolume.BlockDevice{
  bootfromvolume.BlockDevice{
    UUID:       "image_id",
    SourceType: bootfromvolume.Image,
    VolumeSize: 10,
  },
}

serverCreateOpts := servers.CreateOpts{
  Name:      name,
  FlavorRef: "flavor_id",
}
server, err := bootfromvolume.Create(client, bootfromvolume.CreateOptsExt{
  serverCreateOpts,
  bd,
}).Extract()
{% endhighlight %}

### <a name="schedulerhints"></a>Scheduler Hints

{% highlight go %}
import (
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/schedulerhints"
  "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/servergroups"
  "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

sg, err := servergroups.Create(computeClient, &servergroups.CreateOpts{
  Name:     "test",
  Policies: []string{"affinity"},
}).Extract()

serverCreateOpts := servers.CreateOpts{
  Name:      name,
  FlavorName: "flavor_name",
  ImageName:  "image_name",
}
server, err := servers.Create(computeClient, schedulerhints.CreateOptsExt{
  serverCreateOpts,
  schedulerhints.SchedulerHints{
    Group: sg.ID,
  },
}).Extract()
{% endhighlight %}

# <a name="more"></a>More

It's possible for a new feature or action to be added before the documentation is updated.
Check the [extensions](http://godoc.org/github.com/rackspace/gophercloud/openstack/compute/v2/extensions)
directory for a full list of Compute extensions.

# <a name="providers"></a>Providers

### <a name="rackspace"></a>Rackspace

* [Quickstart for Cloud Servers](https://developer.rackspace.com/docs/cloud-servers/getting-started/?lang=go)
on the Rackspace Developer portal.
