package main

import (
	"fmt"
	"flag"
	"github.com/rackspace/gophercloud"
)

var quiet = flag.Bool("quiet", false, "Quiet mode for acceptance testing.  $? non-zero on error though.")
var rgn = flag.String("r", "", "Datacenter region to interrogate.  Leave blank for provider-default region.")

func main() {
	flag.Parse()
	withIdentity(false, func(auth gophercloud.AccessProvider) {
		withObjectStoreApi(auth, func(_ gophercloud.ObjectStoreProvider) {
			fmt.Printf("Hello world\n")
		})
	})
}