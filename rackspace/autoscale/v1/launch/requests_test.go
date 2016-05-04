package launch

import (
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestGetLaunch(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLanuchGetSuccessfully(t)

	client := client.ServiceClient()
	groupID := "10eb3219-1b12-4b34-b1e4-e10ee4f24c65"

	config, err := Get(client, groupID).Extract()

	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, LaunchConfig, *config)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLaunchUpdateSuccessfully(t)

	client := client.ServiceClient()
	groupID := "10eb3219-1b12-4b34-b1e4-e10ee4f24c65"

	opts := UpdateOpts{
		Configuration: LaunchConfig,
	}

	err := Update(client, groupID, opts).ExtractErr()

	th.AssertNoErr(t, err)
}
