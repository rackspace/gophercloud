---
layout: page
title: Getting Started with Object Storage
---

* [Containers](#containers)
* [Objects](#objects)
* [Account](#account)

## Setup

```go
import "github.com/rackspace/gophercloud/openstack"

authOpts, err := utils.AuthOptions()

provider, err := openstack.AuthenticatedClient(authOpts)

client, err := openstack.NewStorageV1(provider, gophercloud.EndpointOpts{
  Region: "RegionOne",
})
```

## Containers

A container is a storage compartment that provides a way for you to organize
your objects. It is analogous to a Linux directory or Windows folder, with the
exception that you cannot nest containers in other containers like a filesystem.

### Create a new container

```go
import "github.com/rackspace/gophercloud/openstack/objectstorage/v1/containers"

// We have the option of passing in configuration options for our new container
opts := containers.CreateOpts{
  ContainerSyncTo: "backup_container",
  Metadata:        map[string]string{"author": "emily dickinson"},
}

res := containers.Create(client, "container_name", opts)

// If we want to extract information out from the response headers, we can.
// The first return value will be http.Header (alias of map[string][]string).
headers, err := res.ExtractHeaders()
```

### List containers

```go
import "github.com/rackspace/gophercloud/pagination"

// We have the option of filtering containers by their attributes
opts := &containers.ListOpts{Full: true, Prefix: "backup_"}

// Retrieve a pager (i.e. a paginated collection)
pager := containers.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {

  // Get a slice of containers.Container structs
	containerList, err := containers.ExtractInfo(page)
	for _, c := range containerList {
    // ...
	}

  // Get a slice of strings, i.e. container names
  containerNames, err := containers.ExtractNames(page)
  for _, n := range containerNames {
    // ...
  }

	return true, nil
})
```

### View and modify container metadata

To retrieve a container's metadata:

```go
metadata, err := containers.Get(client, "container_name").ExtractMetadata()

// Iterate over the map[string]string
for key, val := range metadata {
  // ...
}
```

To update a container's metadata:

```go
// We need to specify the new metadata. Keys that do not exist will be added,
// keys that already exist will be overriden. Keys that are not included in
// this struct will be deleted.
opts := &containers.UpdateOpts{
  Metadata: map[string]string{"new_key": "new_value"},
}

result := containers.Update(client, "container_name", opts)
```

### Delete an existing container

```go
response := containers.Delete(client, "container_name")

// Like most operations, we can extract headers values too
headers, err := response.ExtractHeaders()
```

## Objects

An object stores data content, such as documents, images, and so on. Another way
to think about it is that it serves like a traditional file on a local
filesystem but with lots of additional functionality. For example, you can  
store custom metadata on an object, compress files, manage access with CORS
and temporary URLs, schedule deletions, and execute batch operations (like
deleting 10,000 objects at a time).

### Upload objects

```go

```

### List objects

```go

```

### Copy to new location

```go

```

### Download object

```go

```

### Retrieve and update metadata

```go

```

### Delete object

```go

```

## Account

An account represents the very top-level namespace of the resource hierarchy -
containers belong to accounts, and objects belong to containers. Normally your
service provider creates your account and you then own and can control all the
resources in that account. The account defines a namespace for containers. In
the OpenStack environment, account is synonymous with a project or a tenant.

### Retrieve metadata

```go

```

### Update metadata

```go

```
