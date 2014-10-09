---
layout: docpage
title: Getting Started with Networking
---

* [Setup](#setup)
* [Networks](#networks)
* [Subnets](#subnets)
* [Ports](#ports)

## <a name="setup"></a>Setup

```go
import "github.com/rackspace/gophercloud/openstack"

authOpts, err := utils.AuthOptions()

provider, err := openstack.AuthenticatedClient(authOpts)

client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
	Name:   "neutron",
	Region: "RegionOne",
})
```

## <a name="networks"></a>Networks

A network is the central resource of the OpenStack Neutron API. If you were to
compare it to physical networking, a Neutron network would be analagous to a
VLAN - which is an isolated [broadcast domain](http://en.wikipedia.org/wiki/Broadcast_domain)
inside a larger [layer-2 network](http://en.wikipedia.org/wiki/Data_link_layer).
Because of this virtualized partitioning, a virtual network can only share
packets with other networks through one or more routers.

### Create a network

```go
import "github.com/rackspace/gophercloud/openstack/networking/v2/networks"

// We specify a name and that it should forward packets
opts := networks.CreateOpts{Name: "main_network", AdminStateUp: networks.Up}

// Execute the operation and get back a networks.Network struct
network, err := networks.Create(client, opts).Extract()
```

### List networks

```go
import "github.com/rackspace/gophercloud/pagination"

// We have the option of filtering the network list. If we want the full
// collection, leave it as an empty struct
opts := networks.ListOpts{Shared: false}

// Retrieve a pager (i.e. a paginated collection)
pager := networks.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  networkList, err := networks.ExtractNetworks(page)

  for _, n := range networkList {
    // "n" will be a networks.Network
  }
})
```

### Get details for an existing network

```go
// We need to know what the UUID of our network is and pass it in as a string
network, err := networks.Get(client, "id").Extract()
```

### Update an existing network

You can update a network's name, along with its "shared" or "admin" status:

```go
opts := networks.UpdateOpts{Name: "new_name", Shared: true}

// Like Get(), we need the UUID in string form
network, err := networks.Update(client, "id", opts)
```

### Delete a network

```go
result := networks.Delete(client, "id")
```

## <a name="subnets"></a>Subnets

A subnet is a block of IP addresses (either version 4 or 6) that are assigned
to devices in a particular network. A device in the context of Neutron
specifically means a virtual machine (Compute instance). For this reason, each
subnet must have a [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
and be associated with a network.

### Create a subnet

```go
import "github.com/rackspace/gophercloud/openstack/networking/v2/subnets"

// You must associate a new subnet with an existing network - to do this you
// need its UUID. You must also provide a well-formed CIDR value.
opts := subnets.CreateOpts{
	NetworkID:  "network_id",
	CIDR:       "192.168.199.0/24",
	IPVersion:  subnets.IPv4,
	Name:       "my_subnet",
}

// Execute the operation and get back a subnets.Subnet struct
subnet, err := subnets.Create(client, opts).Extract()
```

### List all subnets

```go
import "github.com/rackspace/gophercloud/pagination"

// We have the option of filtering subnets. For example, we may want to return
// every subnet that belongs to a specific network. Or filter again by name.
opts := subnets.ListOpts{NetworkID: "some_uuid"}

// Retrieve a pager (i.e. a paginated collection)
pager := subnets.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  subnetList, err := subnets.ExtractSubnets(page)

  for _, s := range subnetList {
    // "s" will be a subnets.Subnet
  }
})
```

### Get details for an existing subnet

```go
// We need to know what the UUID of our subnet is and pass it in as a string
subnet, err := subnets.Get(client, "id").Extract()
```

### Update an existing subnet

You can edit the name, gateway IP address, DNS nameservers, host routes
and "enable DHCP" status.

```go
opts := subnets.UpdateOpts{Name: "new_subnet_name"}
subnet, err = subnets.Update(client, "id", opts).Extract()
```

### Delete a subnet

```go
result := subnets.Delete(client, "id")
```

## <a name="ports"></a>Ports

Before talking about what ports are, an important concept to define first are
network switches (both the virtual and physical kind). A network switch connects
different network segments together, and a port is the location where devices
connect to the switch. A device in our case is usually a virtual machine. For
more information about these terms, read this [related article](http://www.wisegeek.com/what-is-a-switch-port.htm).

### Create a port

```go
import "github.com/rackspace/gophercloud/openstack/networking/v2/ports"

// You must associate a new port with an existing network - to do this you
// need its UUID. Also notice the "FixedIPs" field; this allows you to specify
// either a specific IP to use for this port, or the subnet ID from which a
// random free IP is selected.
opts := ports.CreateOpts{
  NetworkID:    "network_id",
  Name:         "my_port",
  AdminStateUp: ports.Up,
  FixedIPs:     []ports.IP{ports.IP{SubnetID: "subnet_id"}},
}

// Execute the operation and get back a subnets.Subnet struct
port, err := ports.Create(client, opts).Extract()
```

### List all ports

```go
import "github.com/rackspace/gophercloud/pagination"

// We have the option of filtering ports. For example, we may want to return
// every port that belongs to a specific network. Or filter again by MAC address.
opts := ports.ListOpts{NetworkID: "some_uuid", MACAddress: "some_addr"}

// Retrieve a pager (i.e. a paginated collection)
pager := ports.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  portList, err := ports.ExtractPorts(page)

  for _, s := range portList {
    // "p" will be a ports.Port
  }
})
```

### Get details for an existing port

```go
// We need to know what the UUID of our port is and pass it in as a string
port, err := ports.Get(client, "id").Extract()
```

### Update an existing port

You can edit the name, admin state, fixed IPs, device ID, device owner and
security groups.

```go
opts := ports.UpdateOpts{Name: "new_port_name"}
port, err = ports.Update(client, "id", opts).Extract()
```

### Delete a subnet

```go
result := ports.Delete(client, "id")
```
