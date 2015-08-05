// +build acceptance

package v3

import (
	"testing"

	"github.com/rackspace/gophercloud/openstack/identity/v3/projects"
	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
)

func TestProjectCRUDOperations(t *testing.T) {
	serviceClient := createAuthenticatedClient(t)
	if serviceClient == nil {
		return
	}

	// Create project
	opts := projects.ProjectOpts{
		Enabled:     true,
		Name:        "Test project",
		Description: "This is test project",
	}
	project, err := projects.Create(serviceClient, opts).Extract()
	th.AssertNoErr(t, err)
	defer projects.Delete(serviceClient, project.ID)
	th.AssertEquals(t, project.Enabled, true)
	th.AssertEquals(t, project.Name, "Test project")
	th.AssertEquals(t, project.Description, "This is test project")

	// List projects
	pager := projects.List(serviceClient, projects.ListOpts{})
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		t.Logf("--- Page ---")

		projectList, err := projects.ExtractProjects(page)
		th.AssertNoErr(t, err)

		for _, p := range projectList {
			t.Logf("Project: ID [%s] Name [%s] Is enabled? [%s]",
				p.ID, p.Name, p.Enabled)
		}

		return true, nil
	})
	th.CheckNoErr(t, err)
	projectID := project.ID

	// Get a project
	if projectID == "" {
		t.Fatalf("In order to retrieve a project, the ProjectID must be set")
	}
	project, err = projects.Get(serviceClient, projectID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, project.ID, projectID)
	th.AssertEquals(t, project.DomainID, "")
	th.AssertEquals(t, project.ParentID, "")
	th.AssertEquals(t, project.Enabled, true)
	th.AssertEquals(t, project.Name, "Test project")
	th.AssertEquals(t, project.Description, "This is test project")

	// Update project
	project, err = projects.Update(serviceClient, projectID, projects.ProjectOpts{Name: "New test project name"}).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, project.Name, "New test project name")

	// Delete project
	res := projects.Delete(serviceClient, projectID)
	th.AssertNoErr(t, res.Err)
}
