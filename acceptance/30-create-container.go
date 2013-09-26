package main

import (
	"flag"
	"fmt"
	"github.com/rackspace/gophercloud"
)

var quiet = flag.Bool("quiet", false, "Quiet mode for acceptance testing.  $? non-zero on error though.")
var rgn = flag.String("r", "", "Datacenter region to interrogate.  Leave blank for provider-default region.")

func main() {
	flag.Parse()
	withIdentity(false, func(auth gophercloud.AccessProvider) {
		withObjectStoreApi(auth, func(osp gophercloud.ObjectStoreProvider) {
			log("Generating random container name")
			containerName := randomString("container-", 16)

			log("Creating container " + containerName)
			err := osp.CreateContainer(containerName)
			if err != nil {
				panic(err)
			}

			log("Deleting container " + containerName)
			err = osp.DeleteContainer(containerName)
			if err != nil {
				panic(err)
			}

			log("Done.")
		})
	})
}

func log(s string) {
	if !*quiet {
		fmt.Println(s)
	}
}
