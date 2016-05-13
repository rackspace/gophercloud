// +build acceptance

package v2

import (
	"testing"
	"time"

	"github.com/rackspace/gophercloud/openstack/blockstorage/v2/volumetypes"
	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
)

func TestVolumeTypes(t *testing.T) {
	client, err := newClient(t)
	th.AssertNoErr(t, err)

	const volName = "gophercloud-test-volumeType"
	vt, err := volumetypes.Create(client, &volumetypes.CreateOpts{
		ExtraSpecs: map[string]interface{}{
			"capabilities": "gpu",
			"priority":     3,
		},
		Name: volName,
	}).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		time.Sleep(10000 * time.Millisecond)
		err = volumetypes.Delete(client, vt.ID).ExtractErr()
		if err != nil {
			t.Error(err)
			return
		}
	}()
	t.Logf("Created volume type: %+v\n", vt)

	vt, err = volumetypes.Get(client, vt.ID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Got volume type: %+v\n", vt)

	found := false
	err = volumetypes.List(client).EachPage(func(page pagination.Page) (bool, error) {
		volTypes, err := volumetypes.ExtractVolumeTypes(page)
		if err != nil {
			return false, err
		}
		for _, volType := range volTypes {
			t.Logf("Listing volume type: %+v\n", volType)
			if volType.Name == volName {
				found = true
			}
		}
		return true, err
	})
	th.AssertNoErr(t, err)
	if !found {
		t.Errorf("Didn't find volume we created: %q", volName)
	}
}
