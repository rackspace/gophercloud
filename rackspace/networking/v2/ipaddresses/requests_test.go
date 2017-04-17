package ipaddresses

import (
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	count := 0
	err := List(client.ServiceClient(), &ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractIPAddresses(page)
		if err != nil {
			t.Errorf("Failed to extract IP addresses: %v", err)
			return false, err
		}

		expected := []IPAddress{
			IPAddress{
				ID:        "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
				NetworkID: "6870304a-7212-443f-bd0c-089c886b44df",
				Address:   "192.168.10.1",
				PortIDs:   []string{"2f693cca-7383-45da-8bae-d26b6c2d6718"},
				SubnetID:  "f11687e8-ef0d-4207-8e22-c60e737e473b",
				TenantID:  "2345678",
				Version:   4,
				Type:      "fixed",
			},
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allPages, err := List(client.ServiceClient(), &ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ExtractIPAddresses(allPages)
	th.AssertNoErr(t, err)

	expected := []IPAddress{
		IPAddress{
			ID:        "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
			NetworkID: "6870304a-7212-443f-bd0c-089c886b44df",
			Address:   "192.168.10.1",
			PortIDs:   []string{"2f693cca-7383-45da-8bae-d26b6c2d6718"},
			SubnetID:  "f11687e8-ef0d-4207-8e22-c60e737e473b",
			TenantID:  "2345678",
			Version:   4,
			Type:      "fixed",
		},
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	actual, err := Get(client.ServiceClient(), "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737").Extract()
	th.AssertNoErr(t, err)

	expected := &IPAddress{
		ID:        "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
		NetworkID: "fda61e0b-a410-49e8-ad3a-64c595618c7e",
		Address:   "192.168.10.1",
		PortIDs: []string{"6200d533-a42b-4c04-82a1-cc14dbdbf2de",
			"9d0db2d7-62df-4c99-80cb-6f140a5260e8"},
		SubnetID: "f11687e8-ef0d-4207-8e22-c60e737e473b",
		TenantID: "2345678",
		Version:  4,
		Type:     "shared",
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &CreateOpts{
		NetworkID: "00000000-0000-0000-0000-000000000000",
		Version:   4,
		PortIDs: []string{
			"6200d533-a42b-4c04-82a1-cc14dbdbf2de",
			"9d0db2d7-62df-4c99-80cb-6f140a5260e8",
		},
		TenantID: "2345678",
	}
	actual, err := Create(client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	expected := &IPAddress{
		ID:        "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
		NetworkID: "fda61e0b-a410-49e8-ad3a-64c595618c7e",
		Address:   "192.168.10.1",
		PortIDs: []string{"6200d533-a42b-4c04-82a1-cc14dbdbf2de",
			"9d0db2d7-62df-4c99-80cb-6f140a5260e8"},
		SubnetID: "f11687e8-ef0d-4207-8e22-c60e737e473b",
		TenantID: "2345678",
		Version:  4,
		Type:     "shared",
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	res := Delete(client.ServiceClient(), "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737", DeleteOpts{})
	th.AssertNoErr(t, res.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	options := &UpdateOpts{
		PortIDs: []string{"275b0516-206f-4421-8e42-1d3d1e4e9fb2", "66811c0a-fdbd-49d4-b1dd-f0f15a329744"},
	}
	actual, err := Update(client.ServiceClient(), "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737", options).Extract()
	th.AssertNoErr(t, err)

	expected := &IPAddress{
		ID:        "4cacd68e-d7aa-4ff2-96f4-5c6f57dba737",
		NetworkID: "6870304a-7212-443f-bd0c-089c886b44df",
		Address:   "192.168.10.1",
		PortIDs: []string{"275b0516-206f-4421-8e42-1d3d1e4e9fb2",
			"66811c0a-fdbd-49d4-b1dd-f0f15a329744"},
		SubnetID: "f11687e8-ef0d-4207-8e22-c60e737e473b",
		TenantID: "2345678",
		Version:  4,
		Type:     "shared",
	}
	th.CheckDeepEquals(t, expected, actual)
}

func TestListByServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListByServerResponse(t)

	count := 0
	err := ListByServer(client.ServiceClient(), "123456", &ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractIPAssociations(page)
		if err != nil {
			t.Errorf("Failed to extract IP associations: %v", err)
			return false, err
		}

		expected := []IPAssociation{
			IPAssociation{
				ID:      "1",
				Address: "10.1.1.1",
			},
			IPAssociation{
				ID:      "2",
				Address: "10.1.1.2",
			},
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestGetByServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetByServerResponse(t)

	actual, err := GetByServer(client.ServiceClient(), "123456", "1").Extract()
	th.AssertNoErr(t, err)

	expected := &IPAssociation{
		ID:      "1",
		Address: "10.1.1.1",
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestAssociate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockAssociateResponse(t)

	actual, err := Associate(client.ServiceClient(), "123456", "2").Extract()
	th.AssertNoErr(t, err)

	expected := &IPAssociation{
		ID:      "2",
		Address: "10.1.1.2",
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestDisassociate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDisassociateResponse(t)

	res := Disassociate(client.ServiceClient(), "123456", "2")
	th.AssertNoErr(t, res.Err)
}
