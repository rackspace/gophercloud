package hypervisors

import (
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)


func TestListHypervisors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHypervisorsListSuccessfully(t)

	actual, err := List(client.ServiceClient()).Extract()

	th.AssertNoErr(t, err)

	if len(actual) != 3 {
		t.Errorf("Expected 3 hypervisors, saw %d", len(actual))
	}

	th.CheckDeepEquals(t, ListHypervisorsExpected[0], actual[0])
	th.CheckDeepEquals(t, ListHypervisorsExpected[1], actual[1])
	th.CheckDeepEquals(t, ListHypervisorsExpected[2], actual[2])
}

func TestListDetailsHypervisors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHypervisorsDetailsSuccessfully(t)

	actual, err := GetDetailsList(client.ServiceClient()).ExtractDetails()

	th.AssertNoErr(t, err)

	if len(actual) != 3 {
		t.Errorf("Expected 3 hypervisors, saw %d", len(actual))
	}

	th.CheckDeepEquals(t, HypervisorsDetailsListExpected[0], actual[0])
	th.CheckDeepEquals(t, HypervisorsDetailsListExpected[1], actual[1])
	th.CheckDeepEquals(t, HypervisorsDetailsListExpected[2], actual[2])
}

func TestHypervisorDetails(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHypervisorDetailSuccessfully(t)

	actual, err := GetDetail(client.ServiceClient(), "14").ExtractDetail()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, HypervisorDetail_2, *actual)
}


func TestServersList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHypervisorServerListSuccessfully(t)

	actual, err := GetHypervisorServers(client.ServiceClient(), "compute-0-2.domain.tld").ExtractServersInfo()

	th.AssertNoErr(t, err)
	if len(actual) != 1 {
		t.Errorf("Expected 1 hypervisors, saw %d", len(actual))
	}

	th.CheckDeepEquals(t, HypervisorServiersListExpected[0], actual[0])
}

func TestHypervisorUptime(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleHypervisorUptimeSuccessfully(t)

	actual, err := GetHypervisorUptime(client.ServiceClient(), "14").ExtractUptime()

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, HypervisorUptimeExpected, *actual)
}
