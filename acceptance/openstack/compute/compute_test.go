// +build acceptance

package compute

import (
	"fmt"
	"github.com/rackspace/gophercloud/openstack/compute/flavors"
	"github.com/rackspace/gophercloud/openstack/compute/images"
	"github.com/rackspace/gophercloud/openstack/compute/servers"
	"os"
	"testing"
	"github.com/rackspace/gophercloud/acceptance/tools"
)

var service = "compute"

func TestListServers(t *testing.T) {
	ts, err := tools.SetupForList(service)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Fprintln(ts.W, "ID\tRegion\tName\tStatus\tIPv4\tIPv6\t")

	region := os.Getenv("OS_REGION_NAME")
	n := 0
	for _, ep := range ts.EPs {
		if (region != "") && (region != ep.Region) {
			continue
		}

		client := servers.NewClient(ep.PublicURL, ts.A, ts.O)

		listResults, err := servers.List(client)
		if err != nil {
			t.Error(err)
			return
		}

		svrs, err := servers.GetServers(listResults)
		if err != nil {
			t.Error(err)
			return
		}

		n = n + len(svrs)

		for _, s := range svrs {
			fmt.Fprintf(ts.W, "%s\t%s\t%s\t%s\t%s\t%s\t\n", s.Id, s.Name, ep.Region, s.Status, s.AccessIPv4, s.AccessIPv6)
		}
	}
	ts.W.Flush()
	fmt.Printf("--------\n%d servers listed.\n", n)
}

func TestListImages(t *testing.T) {
	ts, err := tools.SetupForList(service)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Fprintln(ts.W, "ID\tRegion\tName\tStatus\tCreated\t")

	region := os.Getenv("OS_REGION_NAME")
	n := 0
	for _, ep := range ts.EPs {
		if (region != "") && (region != ep.Region) {
			continue
		}

		client := images.NewClient(ep.PublicURL, ts.A, ts.O)

		listResults, err := images.List(client)
		if err != nil {
			t.Error(err)
			return
		}

		imgs, err := images.GetImages(listResults)
		if err != nil {
			t.Error(err)
			return
		}

		n = n + len(imgs)

		for _, i := range imgs {
			fmt.Fprintf(ts.W, "%s\t%s\t%s\t%s\t%s\t\n", i.Id, ep.Region, i.Name, i.Status, i.Created)
		}
	}
	ts.W.Flush()
	fmt.Printf("--------\n%d images listed.\n", n)
}

func TestListFlavors(t *testing.T) {
	ts, err := tools.SetupForList(service)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Fprintln(ts.W, "ID\tRegion\tName\tRAM\tDisk\tVCPUs\t")

	region := os.Getenv("OS_REGION_NAME")
	n := 0
	for _, ep := range ts.EPs {
		if (region != "") && (region != ep.Region) {
			continue
		}

		client := flavors.NewClient(ep.PublicURL, ts.A, ts.O)

		listResults, err := flavors.List(client, flavors.ListFilterOptions{})
		if err != nil {
			t.Error(err)
			return
		}

		flavs, err := flavors.GetFlavors(listResults)
		if err != nil {
			t.Error(err)
			return
		}

		n = n + len(flavs)

		for _, f := range flavs {
			fmt.Fprintf(ts.W, "%s\t%s\t%s\t%d\t%d\t%d\t\n", f.Id, ep.Region, f.Name, f.Ram, f.Disk, f.VCpus)
		}
	}
	ts.W.Flush()
	fmt.Printf("--------\n%d flavors listed.\n", n)
}

func TestGetFlavor(t *testing.T) {
	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Fatal(err)
	}

	region := os.Getenv("OS_REGION_NAME")
	for _, ep := range ts.EPs {
		if (region != "") && (region != ep.Region) {
			continue
		}
		client := flavors.NewClient(ep.PublicURL, ts.A, ts.O)

		getResults, err := flavors.Get(client, ts.FlavorId)
		if err != nil {
			t.Fatal(err)
		}
		flav, err := flavors.GetFlavor(getResults)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%#v\n", flav)
	}
}

func TestCreateDestroyServer(t *testing.T) {
	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Error(err)
		return
	}

	err = tools.CreateServer(ts)
	if err != nil {
		t.Error(err)
		return
	}

	// We put this in a defer so that it gets executed even in the face of errors or panics.
	defer func() {
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()

	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateServer(t *testing.T) {
	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Error(err)
		return
	}

	err = tools.CreateServer(ts)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()

	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Error(err)
		return
	}

	err = tools.ChangeServerName(ts)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestActionChangeAdminPassword(t *testing.T) {
	t.Parallel()

	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Fatal(err)
	}

	err = tools.CreateServer(ts)
	if err != nil {
		t.Fatal(err)
	}

	defer func(){
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()

	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Fatal(err)
	}

	err = tools.ChangeAdminPassword(ts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestActionReboot(t *testing.T) {
	t.Parallel()

	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Fatal(err)
	}

	err = tools.CreateServer(ts)
	if err != nil {
		t.Fatal(err)
	}

	defer func(){
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()

	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Fatal(err)
	}

	err = servers.Reboot(ts.Client, ts.CreatedServer.Id, "aldhjflaskhjf")
	if err == nil {
		t.Fatal("Expected the SDK to provide an ArgumentError here")
	}

	err = tools.RebootServer(ts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestActionRebuild(t *testing.T) {
	t.Parallel()

	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Fatal(err)
	}

	err = tools.CreateServer(ts)
	if err != nil {
		t.Fatal(err)
	}

	defer func(){
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()

	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Fatal(err)
	}

	err = tools.RebuildServer(ts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestActionResizeConfirm(t *testing.T) {
	t.Parallel()
	
	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Fatal(err)
	}
	
	err = tools.CreateServer(ts)
	if err != nil {
		t.Fatal(err)
	}
	
	defer func(){
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()
	
	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Fatal(err)
	}
	
	err = tools.ResizeServer(ts)
	if err != nil {
		t.Fatal(err)
	}
	
	err = tools.ConfirmResize(ts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestActionResizeRevert(t *testing.T) {
	t.Parallel()
	
	ts, err := tools.SetupForCRUD()
	if err != nil {
		t.Fatal(err)
	}
	
	err = tools.CreateServer(ts)
	if err != nil {
		t.Fatal(err)
	}
	
	defer func(){
		servers.Delete(ts.Client, ts.CreatedServer.Id)
	}()
	
	err = tools.WaitForStatus(ts, "ACTIVE")
	if err != nil {
		t.Fatal(err)
	}
	
	err = tools.ResizeServer(ts)
	if err != nil {
		t.Fatal(err)
	}
	
	err = tools.RevertResize(ts)
	if err != nil {
		t.Fatal(err)
	}
}