package groups

import (
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGroupListSuccessfully(t)

	pages := 0
	err := List(client.ServiceClient(), ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		groups, err := ExtractGroups(page)

		if err != nil {
			return false, err
		}

		if len(groups) != 3 {
			t.Fatalf("Expected 3 groups, got %d", len(groups))
		}

		th.CheckDeepEquals(t, FirstGroup, groups[0])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestGetState(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGroupGetStateSuccessfully(t)

	client := client.ServiceClient()
	groupID := "10eb3219-1b12-4b34-b1e4-e10ee4f24c65"

	state, err := GetState(client, groupID).Extract()

	if err != nil {
		t.Fatalf("Unexpected GetState error: %v", err)
	}

	th.CheckDeepEquals(t, FirstGroupState, *state)
}

func TestGetConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGroupGetConfigSuccessfully(t)

	client := client.ServiceClient()
	groupID := "10eb3219-1b12-4b34-b1e4-e10ee4f24c65"

	state, err := GetConfig(client, groupID).Extract()

	if err != nil {
		t.Fatalf("Unexpected GetState error: %v", err)
	}

	th.CheckDeepEquals(t, GroupConfiguration, *state)
}
