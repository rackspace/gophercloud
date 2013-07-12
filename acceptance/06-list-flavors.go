package main

import (
	"fmt"
	"flag"
	"github.com/rackspace/gophercloud"
)

var quiet = flag.Bool("quiet", false, "Quiet mode for acceptance testing.  $? non-zero on error though.")
var rgn = flag.String("r", "DFW", "Datacenter region to interrogate.")

func main() {
	provider, username, password := getCredentials()
	flag.Parse()

	auth, err := gophercloud.Authenticate(
		provider,
		gophercloud.AuthOptions{
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		panic(err)
	}

	servers, err := gophercloud.ServersApi(auth, gophercloud.ApiCriteria{
		Name:      "cloudServersOpenStack",
		Region:    *rgn,
		VersionId: "2",
		UrlChoice: gophercloud.PublicURL,
	})
	if err != nil {
		panic(err)
	}

	flavors, err := servers.ListFlavors()
	if err != nil {
		panic(err)
	}

	if !*quiet {
		fmt.Println("ID,Name,MinRam,MinDisk")
		for _, f := range flavors {
			fmt.Printf("%s,\"%s\",%d,%d\n", f.Id, f.Name, f.Ram, f.Disk)
		}
	}
}
