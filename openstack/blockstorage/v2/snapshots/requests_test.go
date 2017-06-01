package snapshots

import (
	"testing"

	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &CreateOpts{VolumeID: "32d8295e-17ef-4ea6-9179-eb71f6827f20",
		Name:        "snap-001",
		Description: "test-snapshot",
		Force:       false}
	n, err := Create(client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Size, 1)
	th.AssertEquals(t, n.ID, "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af")
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	s, err := Get(client.ServiceClient(), "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, s.Name, "snap-001")
	th.AssertEquals(t, s.ID, "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	res := Delete(client.ServiceClient(), "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af")
	th.AssertNoErr(t, res.Err)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	count := 0

	List(client.ServiceClient(), &ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractSnapshots(page)
		if err != nil {
			t.Errorf("Failed to extract snapshots: %v", err)
			return false, err
		}

		expected := []Snapshot{
			{
				ID:          "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af",
				Name:        "",
				CreatedAt:   "2017-05-31T14:18:35.000000",
				UpdatedAt:   "2017-05-31T14:18:36.000000",
				Description: "",
				Metadata:    map[string]string{"foo": "bar"},
				Size:        1,
				SourceVolID: "32d8295e-17ef-4ea6-9179-eb71f6827f20",
				Status:      "available",
			},
			{
				ID:          "c970ff21-3c2b-4a4c-b6a0-731808a81776",
				Name:        "test-snap",
				CreatedAt:   "2017-05-31T14:10:12.000000",
				UpdatedAt:   "2017-05-31T14:10:13.000000",
				Description: "this is only a test snapshot",
				Metadata:    map[string]string{},
				Size:        1,
				SourceVolID: "32d8295e-17ef-4ea6-9179-eb71f6827f20",
				Status:      "available",
			},
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allPages, err := List(client.ServiceClient(), &ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ExtractSnapshots(allPages)
	th.AssertNoErr(t, err)

	expected := []Snapshot{
		{
			ID:          "4ee8a3f6-d1c8-4541-ad09-06b7e84a68af",
			Name:        "",
			CreatedAt:   "2017-05-31T14:18:35.000000",
			UpdatedAt:   "2017-05-31T14:18:36.000000",
			Description: "",
			Size:        1,
			SourceVolID: "32d8295e-17ef-4ea6-9179-eb71f6827f20",
			Status:      "available",
			Metadata:    map[string]string{"foo": "bar"},
		},
		{
			ID:          "c970ff21-3c2b-4a4c-b6a0-731808a81776",
			Name:        "test-snap",
			CreatedAt:   "2017-05-31T14:10:12.000000",
			UpdatedAt:   "2017-05-31T14:10:13.000000",
			Description: "this is only a test snapshot",
			Size:        1,
			SourceVolID: "32d8295e-17ef-4ea6-9179-eb71f6827f20",
			Status:      "available",
			Metadata:    map[string]string{},
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}
