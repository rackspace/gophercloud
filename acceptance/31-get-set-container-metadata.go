package main

import (
	"flag"
	"fmt"
	"github.com/rackspace/gophercloud"
	"strings"
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
			container, err := osp.CreateContainer(containerName)
			if err != nil {
				panic(err)
			}
			defer container.Delete()

			log("We just created the container; it should have no custom keys")
			metadata, err := container.Metadata()
			if err != nil {
				panic(err)
			}

			settings, err := metadata.CustomValues()
			if err != nil {
				panic(err)
			}
			if len(settings) != 0 {
				panic("Expected no custom attributes at this time.")
			}

			key := randomString("keykey", 16)
			val := randomString("valval", 16)
			log("Attempting to set " + key + " to value " + val)
			err = metadata.SetCustomValue(key, val)
			if err != nil {
				panic(err)
			}

			log("Checking to make sure the custom settings stick")
			val2, err := metadata.CustomValue(key)
			if err != nil {
				panic(err)
			}
			if strings.ToLower(val2) != strings.ToLower(val) {
				panic("Expected value (" + val2 + ") to match what we set before (" + val + ")")
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
