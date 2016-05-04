package launch

import (
	"encoding/json"
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
		"loadBalancerId": "123456",
	}

	lb, err := mapToLoadBalancer(
		reflect.TypeOf(lbMap),
		reflect.TypeOf(CloudLB),
		lbMap,
	)

	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CloudLB, lb)
}

func TestServerMarshalJSON(t *testing.T) {
	serverJSON, err := json.MarshalIndent(&ExampleServer, "", "  ")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, ExampleServerJSON, string(serverJSON))
}

func BenchmarkServerMarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(ExampleServer)
	}
}

func TestLoadBalancerMarshalJSON(t *testing.T) {
	lbJSON, err := json.MarshalIndent(&CloudLB, "", "  ")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, CloudLBJSON, string(lbJSON))
}
