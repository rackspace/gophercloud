package launch

import (
	"reflect"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
)

func TestMapToFile(t *testing.T) {
	fileMap := map[string]interface{}{
		"path":     "/etc/motd",
		"contents": "Z29waGVyY2xvdWQ=",
	}

	file, err := mapToFile(
		reflect.TypeOf(fileMap),
		reflect.TypeOf(&MOTDFile),
		fileMap,
	)

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &MOTDFile, file)
}

func TestStringToBytes(t *testing.T) {
	data := "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQ="

	bytes, err := stringToBytes(
		reflect.TypeOf(data),
		reflect.TypeOf(ExampleUserData),
		data,
	)

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ExampleUserData, bytes)
}

func TestMapToLoadBalancer(t *testing.T) {
	lbMap := map[string]interface{}{
		"port":           float64(443),
		"loadBalancerId": float64(123456),
	}

	lb, err := mapToLoadBalancer(
		reflect.TypeOf(lbMap),
		reflect.TypeOf(CloudLB),
		lbMap,
	)

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CloudLB, lb)
}
