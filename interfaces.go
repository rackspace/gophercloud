package gophercloud

// AccessProvider instances encapsulate a Keystone authentication interface.
type AccessProvider interface {
	// FirstEndpointUrlByCriteria searches through the service catalog for the first
	// matching entry endpoint fulfilling the provided criteria.  If nothing found,
	// return "".  Otherwise, return either the public or internal URL for the
	// endpoint, depending on both its existence and the setting of the ApiCriteria.UrlChoice
	// field.
	FirstEndpointUrlByCriteria(ApiCriteria) string

	// AuthToken provides a copy of the current authentication token for the user's credentials.
	// Note that AuthToken() will not automatically refresh an expired token.
	AuthToken() string

	// Revoke allows you to terminate any program's access to the OpenStack API by token ID.
	Revoke(string) error

	// Reauthenticate attempts to acquire a new authentication token, if the feature is enabled by
	// AuthOptions.AllowReauth.
	Reauthenticate() error
}

// ServiceCatalogerIdentityV2 interface provides direct access to the service catalog as offered by the Identity V2 API.
// We regret we need to fracture the namespace of what should otherwise be a simple concept; however,
// the OpenStack community saw fit to render V3's service catalog completely incompatible with V2.
type ServiceCatalogerForIdentityV2 interface {
	V2ServiceCatalog() []CatalogEntry
}

// ObjectStorageProvider instances encapsulate a cloud-based storage API, should one exist in the service catalog
// for your provider.
type ObjectStoreProvider interface {
	// CreateContainer attempts to create a container for objects on the remote provider's cloud
	// storage infrastructure.
	CreateContainer(name string) (Container, error)

	// ListContainers returns a slice of ContainerInfo interfaces
	ListContainers(listOpts ListOpts) ([]ContainerInfo, error)

	// GetContainer returns a Container identified by a well-known name. No server interactions occur;
	// it assumes you already know the name of an existing container.
	GetContainer(name string) Container

	// DeleteContainer attempts to delete an empty container.
	// This call WILL fail if the container is not empty.
	DeleteContainer(name string) error
}

// Container instances more or less correspond to directories or volume table of contents structures in
// more traditional filesystems.
type Container interface {
	// Delete() dispenses with the container.
	// NOTE: Upon returning from this method call without error,
	// the reference to the container no longer refers to a container hosted by your provider.
	// Future calls to the container will produce errors.
	Delete() error

	// Metadata() provides access to a container's set of custom metadata settings.
	Metadata() (MetadataProvider, error)

	// BasicObjectUploader allows for uploading an object.
	BasicObjectUploader() *BasicUploader

	// DeleteObject removes an object from the container.
	DeleteObject(name string) error

	// ListObjects will return objects in the container. The listOpts parameter holds the
	// options for retreiving the objects.
	ListObjects(listOpts ListOptions) ([]ObjectInfo, error)
}

// ContainerInfo instances encapsulate information relating to a container. An object implementing the ContainerInfo
//interface must have methods to accessing the container's name, object count, and size.
type ContainerInfo interface {
	// Return the container's name
	Label() string
	// Return the number of objects in the container
	ObjCount() int
	// Return the size of the container (in bytes)
	Size() int
}

// ObjectInfo instances encapsulate information relating to a storage object. The methods associated
// with an instance of ObjectInfo are for specifying fields that must be present in the implemented
// structure
type ObjectInfo interface {
	// Return the name of the object.
	GetName() string
	// Return the object's hash.
	GetHash() string
	// Return the size of the object (in bytes).
	GetSize() int
	// Return the object's content-type.
	GetContentType() string
	// Return (as a string) the most recent date the object was modified.
	GetLastModified() string
}

// ListOptions are options for any query that can be 'paged'.
type ListOptions interface {
	// Return the desired format. Setting Full to true will retrieve all the information available
	// for the listed items, whereas false will only retrieve the item names.
	GetFull() bool
	// Return the limit. Limit is an integer parameter that represents the maximum number of items
	// to return.
	GetLimit() int
	// Return the marker. Marker is a string parameter that tells the endpoint to return items whose name
	// is greater in value than the specified marker.
	GetMarker() string
	// Return the end marker. EndMarker is a string parameter that tells the endpoint to return items
	// whose name is lesser in value than the specified end marker.
	GetEndMarker() string
}

// MetadataProvider grants access to custom metadata on some "thing", whatever that thing may be (e.g., containers,
// files, etc.)
type MetadataProvider interface {
	// CustomValue retrieves the value associated with a single custom attribute, if one exists.
	// Note: It is explicitly not an error if no value is bound to an attribute.
	// In this case, the value returned will be "".
	CustomValue(key string) (string, error)

	// CustomValues provides a complete set of metadata settings, in map form.
	// Try not to use this method unless you need the full batch at once, or unless you find your
	// code is making repeated CustomValue() calls.
	CustomValues() (map[string]string, error)

	// SetCustomValue provides a means by which your application may set a custom attribute on an entity.
	SetCustomValue(key, value string) error
}

// CloudServersProvider instances encapsulate a Cloud Servers API, should one exist in the service catalog
// for your provider.
type CloudServersProvider interface {
	// Servers

	// ListServers provides a complete list of servers hosted by the user
	// in a given region.  This function differs from ListServersLinksOnly()
	// in that it returns all available details for each server returned.
	ListServers() ([]Server, error)

	// ListServers provides a complete list of servers hosted by the user
	// in a given region.  This function differs from ListServers() in that
	// it returns only IDs and links to each server returned.
	//
	// This function should be used only under certain circumstances.
	// It's most useful for checking to see if a server with a given ID exists,
	// or that you have permission to work with that server.  It's also useful
	// when the cost of retrieving the server link list plus the overhead of manually
	// invoking ServerById() for each of the servers you're interested in is less than
	// just calling ListServers() to begin with.  This may be a consideration, for
	// example, with mobile applications.
	//
	// In other cases, you probably should just call ListServers() and cache the
	// results to conserve overall bandwidth and reduce your access rate on the API.
	ListServersLinksOnly() ([]Server, error)

	// ServerById will retrieve a detailed server description given the unique ID
	// of a server.  The ID can be returned by either ListServers() or by ListServersLinksOnly().
	ServerById(id string) (*Server, error)

	// CreateServer requests a new server to be created by the cloud server provider.
	// The user must pass in a pointer to an initialized NewServerContainer structure.
	// Please refer to the NewServerContainer documentation for more details.
	//
	// If the NewServer structure's AdminPass is empty (""), a password will be
	// automatically generated by your OpenStack provider, and returned through the
	// AdminPass field of the result.  Take care, however; this will be the only time
	// this happens.  No other means exists in the public API to acquire a password
	// for a pre-existing server.  If you lose it, you'll need to call SetAdminPassword()
	// to set a new one.
	CreateServer(ns NewServer) (*NewServer, error)

	// DeleteServerById requests that the server with the assigned ID be removed
	// from your account.  The delete happens asynchronously.
	DeleteServerById(id string) error

	// SetAdminPassword requests that the server with the specified ID have its
	// administrative password changed.  For Linux, BSD, or other POSIX-like
	// system, this password corresponds to the root user.  For Windows machines,
	// the Administrator password will be affected instead.
	SetAdminPassword(id string, pw string) error

	// ResizeServer can be a short-hand for RebuildServer where only the size of the server
	// changes.  Note that after the resize operation is requested, you will need to confirm
	// the resize has completed for changes to take effect permanently.  Changes will assume
	// to be confirmed even without an explicit confirmation after 24 hours from the initial
	// request.
	ResizeServer(id, newName, newFlavor, newDiskConfig string) error

	// RevertResize will reject a server's resized configuration, thus
	// rolling back to the original server.
	RevertResize(id string) error

	// ConfirmResizeServer will acknowledge a server's resized configuration.
	ConfirmResize(id string) error

	// RebootServer requests that the server with the specified ID be rebooted.
	// Two reboot mechanisms exist.
	//
	// - Hard.  This will physically power-cycle the unit.
	// - Soft.  This will attempt to use the server's software-based mechanisms to restart
	//           the machine.  E.g., "shutdown -r now" on Linux.
	RebootServer(id string, hard bool) error

	// RescueServer requests that the server with the specified ID be placed into
	// a state of maintenance.  The server instance is replaced with a new instance,
	// of the same flavor and image.  This new image will have the boot volume of the
	// original machine mounted as a secondary device, so that repair and administration
	// may occur.  Use UnrescueServer() to restore the server to its previous state.
	// Note also that many providers will impose a time limit for how long a server may
	// exist in rescue mode!  Consult the API documentation for your provider for
	// details.
	RescueServer(id string) (string, error)

	// UnrescueServer requests that a server in rescue state be placed into its nominal
	// operating state.
	UnrescueServer(id string) error

	// UpdateServer alters one or more fields of the identified server's Server record.
	// However, not all fields may be altered.  Presently, only Name, AccessIPv4, and
	// AccessIPv6 fields may be altered.   If unspecified, or set to an empty or zero
	// value, the corresponding field remains unaltered.
	//
	// This function returns the new set of server details if successful.
	UpdateServer(id string, newValues NewServerSettings) (*Server, error)

	// RebuildServer reprovisions a server to the specifications given by the
	// NewServer structure.  The following fields are guaranteed to be recognized:
	//
	//		Name (required)				AccessIPv4
	//		imageRef (required)			AccessIPv6
	//		AdminPass (required)		Metadata
	//		Personality
	//
	// Other providers may reserve the right to act on additional fields.
	RebuildServer(id string, ns NewServer) (*Server, error)

	// CreateImage will create a new image from the specified server id returning the id of the new image.
	CreateImage(id string, ci CreateImage) (string, error)

	// Addresses

	// ListAddresses yields the list of available addresses for the server.
	// This information is also returned by ServerById() in the Server.Addresses
	// field.  However, if you have a lot of servers and all you need are addresses,
	// this function might be more efficient.
	ListAddresses(id string) (AddressSet, error)

	// Images

	// ListImages yields the list of available operating system images.  This function
	// returns full details for each image, if available.
	ListImages() ([]Image, error)

	// ImageById yields details about a specific image.
	ImageById(id string) (*Image, error)

	// DeleteImageById will delete the specific image.
	DeleteImageById(id string) error

	// Flavors

	// ListFlavors yields the list of available system flavors.  This function
	// returns full details for each flavor, if available.
	ListFlavors() ([]Flavor, error)

	// KeyPairs

	// ListKeyPairs yields the list of available keypairs.
	ListKeyPairs() ([]KeyPair, error)

	// CreateKeyPairs will create or generate a new keypair.
	CreateKeyPair(nkp NewKeyPair) (KeyPair, error)

	// DeleteKeyPair wil delete a keypair.
	DeleteKeyPair(name string) error

	// ShowKeyPair will yield the named keypair.
	ShowKeyPair(name string) (KeyPair, error)
}
