---
layout: page
title: Getting Started with Block Storage
---

## Volume types

A Volume type is... well... the type of a block storage volume you want. You
can define whatever types work best for you, such as SATA, SCSCI, SSD, etc.
These can be customized or defined by the OpenStack admin.

### List volume types

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumetypes"
)

// Retrieve a pager (i.e. a paginated collection)
pager := volumetypes.List(client)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  vtList, err := volumetypes.ExtractVolumeTypes(page)

  for _, vt := range vtList {
    // "vt" will be a volumetypes.VolumeType
  }
})
{% endhighlight %}

### Create volume type

In order to create a new volume type, you must specify a name.

You can also define `ExtraSpecs` associated with your volume types. For
instance, you could have a `SATA` VolumeType with these extra specs: RPM=10000
and RAID-Level=5.

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumetypes"

// Optional
specs := map[string]interface{}{RAID-Level: 5, RPM: 10000}

opts := volumetypes.CreateOpts{Name: "new_type", ExtraSpecs: specs}

vt, err := volumetypes.Create(client, opts).Extract()
{% endhighlight %}

### Show volume type details

{% highlight go %}
vt, err := volumetypes.Get(client, "volume_type_id").Extract()
{% endhighlight %}

### Delete volume type

{% highlight go %}
err := volumetypes.Delete(client, "volume_type_id")
{% endhighlight %}

## Volumes

A volume is a detachable block storage device (you can think of it as a USB
hard drive). It can only be attached to one instance at a time.

### Create volume

The only required attribute when creating a new volume is its size. All other
attributes are optional.

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumes"

opts := volumes.CreateOpts{Size: 100, Name: "foo_volume", VolumeType: "volume_type_id"}

vol, err := volumes.Create(client, opts).Extract()
{% endhighlight %}

### List volumes

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumetypes"
)

// We can filter by status
opts := volumes.ListOpts{Status: "IN-USE"}

// Retrieve a pager (i.e. a paginated collection)
pager := volumes.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  vList, err := volumes.ExtractVolumes(page)

  for _, v := range vList {
    // "v" will be a volumes.Volume
  }
})
{% endhighlight %}

### Show volume details

{% highlight go %}
vol, err := volumes.Get(client, "volume_id").Extract()
{% endhighlight %}

### Update volume

{% highlight go %}
opts := volumes.UpdateOpts{Name: "new_name"}

vol, err := volumes.Update(client, "volume_id", opts).Extract()
{% endhighlight %}

### Delete volume

{% highlight go %}
err := volumes.Delete(client, "volume_id")
{% endhighlight %}

## Snapshots

A snapshot is point-in-time copy of the data contained in a volume.

### Create snapshot

The only required attribute when creating a new snapshot is the ID of the
volume you're backing up.

{% highlight go %}
import "github.com/rackspace/gophercloud/openstack/blockstorage/v1/snapshots"

opts := snapshots.CreateOpts{Name: "2014_oct", VolumeID: "volume_id"}

snap, err := snapshots.Create(client, opts).Extract()
{% endhighlight %}

### List snapshots

{% highlight go %}
import (
  "github.com/rackspace/gophercloud/pagination"
  "github.com/rackspace/gophercloud/openstack/blockstorage/v1/snapshots"
)

// We can filter by status
opts := snapshots.ListOpts{Status: "ERROR"}

// Retrieve a pager (i.e. a paginated collection)
pager := snapshots.List(client, opts)

// Define an anonymous function to be executed on each page's iteration
err := pager.EachPage(func(page pagination.Page) (bool, error) {
  sList, err := snapshots.ExtractSnapshots(page)

  for _, s := range sList {
    // "s" will be a snapshots.Snapshot
  }
})
{% endhighlight %}

### Show snapshot details

{% highlight go %}
snap, err := snapshots.Get(client, "snapshot_id").Extract()
{% endhighlight %}

### Delete snapshot

{% highlight go %}
err := snapshots.Delete(client, "snapshot_id")
{% endhighlight %}

### Update snapshot metadata

{% highlight go %}
opts := snapshots.UpdateMetadataOpts{
  Metadata: map[string]interface{}{
    Foo: "bar",
    Baz: "foo",
  }
}

res := snapshots.UpdateMetadata(client, "snapshot_id", opts)

// To extract snapshot out
snap, err := res.Extract()

// To extract a metadata map
metadata, err := res.ExtractMetadata()
{% endhighlight %}
