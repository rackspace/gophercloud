package members

import (
	"fmt"

	"github.com/rackspace/gophercloud"
)

// Create member for specific image
//
// Preconditions
//    The specified images must exist.
//    You can only add a new member to an image which 'visibility' attribute is private.
//    You must be the owner of the specified image.
// Synchronous Postconditions
//    With correct permissions, you can see the member status of the image as pending through API calls.
//
// More details here: http://developer.openstack.org/api-ref-image-v2.html#createImageMember-v2
func Create(client *gophercloud.ServiceClient, id string, member string) CreateMemberResult {
	var res CreateMemberResult
	body := map[string]interface{}{}
	body["member"] = member

	response, err := client.Post(imageMembersURL(client, id), body, &res.Body,
		&gophercloud.RequestOpts{OkCodes: []int{200, 409, 403}})

	//some problems in http stack or lower
	if err != nil {
		res.Err = err
		return res
	}

	// membership conflict
	if response.StatusCode == 409 {
		res.Err = fmt.Errorf("Given tenant '%s' is already member for image '%s'.", member, id)
		return res
	}

	// visibility conflict
	if response.StatusCode == 403 {
		res.Err = fmt.Errorf("You can only add a new member to an image "+
			"which 'visibility' attribute is private (image '%s')", id)
		return res
	}

	return res
}

// List members returns list of members for specifed image id
// More details: http://developer.openstack.org/api-ref-image-v2.html#listImageMembers-v2
func List(client *gophercloud.ServiceClient, id string) ListMembersResult {
	var res ListMembersResult
	_, res.Err = client.Get(listMembersURL(client, id), &res.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return res
}

// Get image member details.
// More details: http://developer.openstack.org/api-ref-image-v2.html#getImageMember-v2
func Get(client *gophercloud.ServiceClient, imageID string, memberID string) MemberDetailsResult {
	var res MemberDetailsResult
	_, res.Err = client.Get(imageMemberURL(client, imageID, memberID), &res.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return res
}

// Delete membership for given image.
// Callee should be image owner
// More details: http://developer.openstack.org/api-ref-image-v2.html#deleteImageMember-v2
func Delete(client *gophercloud.ServiceClient, imageID string, memberID string) MemberDeleteResult {
	var res MemberDeleteResult
	response, err := client.Delete(imageMemberURL(client, imageID, memberID), &gophercloud.RequestOpts{OkCodes: []int{204, 403}})

	//some problems in http stack or lower
	if err != nil {
		res.Err = err
		return res
	}

	// Callee is not owner of specified image
	if response.StatusCode == 403 {
		res.Err = fmt.Errorf("You must be the owner of the specified image. "+
			"(image '%s')", imageID)
		return res
	}
	return res
}

// Update fuction updates member
// More details: http://developer.openstack.org/api-ref-image-v2.html#updateImageMember-v2
func Update(client *gophercloud.ServiceClient, imageID string, memberID string, status string) MemberUpdateResult {
	var res MemberUpdateResult
	body := map[string]interface{}{}
	body["status"] = status
	_, res.Err = client.Put(imageMemberURL(client, imageID, memberID), body, &res.Body,
		&gophercloud.RequestOpts{OkCodes: []int{200}})
	return res
}
