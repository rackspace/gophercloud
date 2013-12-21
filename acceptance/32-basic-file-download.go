package main

// To run this file, call it from the command-line like the following:
// go run 32-basic-file-download.go libargs.go -container=jclouds-example-publish -file=createObjectFromFile.html

import (
	"flag"
	"fmt"
	"github.com/rackspace/gophercloud"
	"io"
	"os"
)

var containerName = flag.String("container", "", "The name of the container from which to download the file")
var fileName = flag.String("file", "", "The name of the file to download")

type testCase struct {
	o int
	l int
}

func main() {
	flag.Parse()

	withIdentity(false, func(auth gophercloud.AccessProvider) {
		withObjectStoreApi(auth, func(osp gophercloud.ObjectStoreProvider) {
			container := osp.GetContainer(*containerName)
			// You can define more interesting test cases here. The default test {0,0} downloads the whole file.
			testCases := [...]testCase{{0, 0}}
			for ind, test := range testCases {
				fmt.Println("\n")
				fmt.Printf("offset:%d, length:%d\n", test.o, test.l)
				reader, err := container.BasicObjectDownloader(gophercloud.ObjectOpts{
					Name:   *fileName,
					Offset: test.o,
					Length: test.l,
				})
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					continue
				}
				defer reader.Close()
				cloud_file, err := os.Create(fmt.Sprintf("cloud_file_%d.html", ind))
				if err != nil {
					panic(err)
				}
				defer cloud_file.Close()
				if reader != nil {
					n, err := io.Copy(cloud_file, reader)
					if err != nil {
						panic(err)
					}
					fmt.Printf("Bytes read: %d\n", n)
				}
			}
		})
	})
}
