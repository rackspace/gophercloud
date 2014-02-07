package main

import (
	"flag"
	"fmt"
	"github.com/rackspace/gophercloud"
	"io"
	"strconv"
	"strings"
	"time"
)

var quiet = flag.Bool("quiet", false, "Quiet mode, for acceptance testing. $? still indicates errors though.")

func main() {
	withIdentity(false, func(auth gophercloud.AccessProvider) {
		withObjectStoreApi(auth, func(osp gophercloud.ObjectStoreProvider) {
			log("Generating random container name")
			containerName := randomString("container-", 16)

			log("Creating container " + containerName)
			container, err := osp.CreateContainer(containerName)
			if err != nil {
				panic(err)
			}

			writer := container.BasicObjectUploader()
			defer writer.Close()

			// The number of objects to create in the container
			numObjs := 3

			for i := 0; i < numObjs; i++ {
				reader := strings.NewReader(randomString("", 30))

				_, err = io.Copy(writer, reader)
				if err != nil {
					panic(err)
				}

				err = writer.Commit(gophercloud.ObjectOpts{
					Name:      "gophercloud-list-objects_test-object-" + strconv.Itoa(i) + ".txt",
					Container: container,
				})
				if err != nil {
					panic(err)
				}
			}

			objects, err := container.ListObjects(gophercloud.OpenstackListOpts{
				Full: false,
			})
			if err != nil {
				panic(err)
			}
			log("Basic Object Info")
			if !*quiet {
				fmt.Printf("%+v\n", objects)
			}

			objects, err = container.ListObjects(gophercloud.OpenstackListOpts{
				Full: true,
			})
			if err != nil {
				panic(err)
			}
			log("Full Object Info")
			if !*quiet {
				fmt.Printf("%+v\n", objects)
			}

			for i := 0; i < numObjs; i++ {
				log("Deleting object gophercloud-list-objects_test-object-" + strconv.Itoa(i) + ".txt")
				err = container.DeleteObject("gophercloud-list-objects_test-object-" + strconv.Itoa(i) + ".txt")
				if err != nil {
					panic(err)
				}
			}

			// Give the endpoint some time to fully delete all the objects in the container.
			// Removing this pause can result in a 409 Error (Container not empty)
			time.Sleep(1000 * time.Millisecond)

			log("Deleting container " + containerName)
			err = container.Delete()
			if err != nil {
				panic(err)
			}

			log("Done.")
		})
	})
}

func log(s ...interface{}) {
	if !*quiet {
		fmt.Println(s...)
	}
}
