package main

import (
	"flag"
	"fmt"
	"github.com/rackspace/gophercloud"
	"io"
	"strings"
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

			reader := strings.NewReader("Hello, Pythoncloud!")

			_, err = io.Copy(writer, reader)
			if err != nil {
				panic(err)
			}

			log("writer before WriteAt: ", writer)

			_, err = writer.WriteAt([]byte("Gopher"), 7)

			if err != nil {
				panic(err)
			}

			log("writer after WriteAt: ", writer)

			err = writer.Commit(gophercloud.ObjectOpts{
				Name:      "gophercloud_upload_test.txt",
				Container: container,
			})
			if err != nil {
				panic(err)
			}

			log("Deleting object gophercloud_upload_test.txt")
			err = container.DeleteObject("gophercloud_upload_test.txt")
			if err != nil {
				panic(err)
			}

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
