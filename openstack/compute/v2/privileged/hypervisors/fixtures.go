package hypervisors

import (
	"testing"
	"net/http"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
	"fmt"
)


//	nova hypervisor-list (List hypervisors.)
const HypervisorListBody = `
{
  "hypervisors": [
    {
      "status": "enabled",
      "state": "up",
      "id": 2,
      "hypervisor_hostname": "compute-0-3.domain.tld"
    },
    {
      "status": "enabled",
      "state": "up",
      "id": 8,
      "hypervisor_hostname": "compute-0-1.domain.tld"
    },
    {
      "status": "enabled",
      "state": "up",
      "id": 14,
      "hypervisor_hostname": "compute-0-2.domain.tld"
    }
  ]
}
`
// If we call "nova hypervisor-show compute-2.domain.tld"
// (hypervisor hostname), Nova will return details of all hypervisors.
const HypervisorsDetailsBody = `
{
  "hypervisors": [
    {
      "status": "enabled",
      "service": {
        "host": "compute-0-3.domain.tld",
        "disabled_reason": null,
        "id": 33
      },
      "vcpus_used": 0,
      "hypervisor_type": "QEMU",
      "local_gb_used": 240,
      "vcpus": 24,
      "hypervisor_hostname": "compute-0-3.domain.tld",
      "memory_mb_used": 59016,
      "memory_mb": 64136,
      "current_workload": 0,
      "state": "up",
      "host_ip": "192.168.2.28",
      "cpu_info": "{\"vendor\": \"Intel\", \"model\": \"SandyBridge\", \"arch\": \"x86_64\", \"features\": [\"pge\", \"avx\", \"clflush\", \"sep\", \"syscall\", \"vme\", \"dtes64\", \"msr\", \"fsgsbase\", \"xsave\", \"vmx\", \"erms\", \"xtpr\", \"cmov\", \"smep\", \"ssse3\", \"est\", \"pat\", \"monitor\", \"smx\", \"pbe\", \"lm\", \"tsc\", \"nx\", \"fxsr\", \"tm\", \"sse4.1\", \"pae\", \"sse4.2\", \"pclmuldq\", \"acpi\", \"tsc-deadline\", \"mmx\", \"osxsave\", \"cx8\", \"mce\", \"de\", \"tm2\", \"ht\", \"dca\", \"lahf_lm\", \"popcnt\", \"mca\", \"pdpe1gb\", \"apic\", \"sse\", \"f16c\", \"pse\", \"ds\", \"invtsc\", \"pni\", \"rdtscp\", \"aes\", \"sse2\", \"ss\", \"ds_cpl\", \"pcid\", \"fpu\", \"cx16\", \"pse36\", \"mtrr\", \"pdcm\", \"rdrand\", \"x2apic\"], \"topology\": {\"cores\": 10, \"threads\": 2, \"sockets\": 1}}",
      "running_vms": 0,
      "free_disk_gb": 414,
      "hypervisor_version": 2002000,
      "disk_available_least": 633,
      "local_gb": 654,
      "free_ram_mb": 5120,
      "id": 2
    },
    {
      "status": "enabled",
      "service": {
        "host": "compute-0-1.domain.tld",
        "disabled_reason": null,
        "id": 38
      },
      "vcpus_used": 0,
      "hypervisor_type": "QEMU",
      "local_gb_used": 290,
      "vcpus": 22,
      "hypervisor_hostname": "compute-0-1.domain.tld",
      "memory_mb_used": 62088,
      "memory_mb": 64136,
      "current_workload": 0,
      "state": "up",
      "host_ip": "192.168.2.25",
      "cpu_info": "{\"vendor\": \"Intel\", \"model\": \"SandyBridge\", \"arch\": \"x86_64\", \"features\": [\"pge\", \"avx\", \"clflush\", \"sep\", \"syscall\", \"vme\", \"dtes64\", \"msr\", \"fsgsbase\", \"xsave\", \"vmx\", \"erms\", \"xtpr\", \"cmov\", \"smep\", \"ssse3\", \"est\", \"pat\", \"monitor\", \"smx\", \"pbe\", \"lm\", \"tsc\", \"nx\", \"fxsr\", \"tm\", \"sse4.1\", \"pae\", \"sse4.2\", \"pclmuldq\", \"acpi\", \"tsc-deadline\", \"mmx\", \"osxsave\", \"cx8\", \"mce\", \"de\", \"tm2\", \"ht\", \"dca\", \"lahf_lm\", \"popcnt\", \"mca\", \"pdpe1gb\", \"apic\", \"sse\", \"f16c\", \"pse\", \"ds\", \"invtsc\", \"pni\", \"rdtscp\", \"aes\", \"sse2\", \"ss\", \"ds_cpl\", \"pcid\", \"fpu\", \"cx16\", \"pse36\", \"mtrr\", \"pdcm\", \"rdrand\", \"x2apic\"], \"topology\": {\"cores\": 10, \"threads\": 2, \"sockets\": 1}}",
      "running_vms": 0,
      "free_disk_gb": 364,
      "hypervisor_version": 2002000,
      "disk_available_least": 601,
      "local_gb": 654,
      "free_ram_mb": 2048,
      "id": 8
    },
    {
      "status": "enabled",
      "service": {
        "host": "compute-0-2.domain.tld",
        "disabled_reason": null,
        "id": 42
      },
      "vcpus_used": 2,
      "hypervisor_type": "QEMU",
      "local_gb_used": 130,
      "vcpus": 32,
      "hypervisor_hostname": "compute-0-2.domain.tld",
      "memory_mb_used": 20104,
      "memory_mb": 64136,
      "current_workload": 0,
      "state": "up",
      "host_ip": "192.168.2.23",
      "cpu_info": "{\"vendor\": \"Intel\", \"model\": \"SandyBridge\", \"arch\": \"x86_64\", \"features\": [\"pge\", \"avx\", \"clflush\", \"sep\", \"syscall\", \"vme\", \"dtes64\", \"msr\", \"fsgsbase\", \"xsave\", \"vmx\", \"erms\", \"xtpr\", \"cmov\", \"smep\", \"ssse3\", \"est\", \"pat\", \"monitor\", \"smx\", \"pbe\", \"lm\", \"tsc\", \"nx\", \"fxsr\", \"tm\", \"sse4.1\", \"pae\", \"sse4.2\", \"pclmuldq\", \"acpi\", \"tsc-deadline\", \"mmx\", \"osxsave\", \"cx8\", \"mce\", \"de\", \"tm2\", \"ht\", \"dca\", \"lahf_lm\", \"popcnt\", \"mca\", \"pdpe1gb\", \"apic\", \"sse\", \"f16c\", \"pse\", \"ds\", \"invtsc\", \"pni\", \"rdtscp\", \"aes\", \"sse2\", \"ss\", \"ds_cpl\", \"pcid\", \"fpu\", \"cx16\", \"pse36\", \"mtrr\", \"pdcm\", \"rdrand\", \"x2apic\"], \"topology\": {\"cores\": 10, \"threads\": 2, \"sockets\": 1}}",
      "running_vms": 1,
      "free_disk_gb": 524,
      "hypervisor_version": 2002000,
      "disk_available_least": 521,
      "local_gb": 654,
      "free_ram_mb": 44032,
      "id": 14
    }
  ]
}

`

// nova hypervisor-show 14
const HypervisorDetailsBody = `
{"hypervisor": {"status": "enabled",
  "service": {
    "host": "compute-0-2.domain.tld",
    "disabled_reason": null,
    "id": 42
  },
  "vcpus_used": 2,
  "hypervisor_type": "QEMU",
  "local_gb_used": 130,
  "vcpus": 32,
  "hypervisor_hostname": "compute-0-2.domain.tld",
  "memory_mb_used": 20104,
  "memory_mb": 64136,
  "current_workload": 0,
  "state": "up",
  "host_ip": "192.168.2.23",
  "cpu_info": "{\"vendor\": \"Intel\", \"model\": \"SandyBridge\", \"arch\": \"x86_64\", \"features\": [\"pge\", \"avx\", \"clflush\", \"sep\", \"syscall\", \"vme\", \"dtes64\", \"msr\", \"fsgsbase\", \"xsave\", \"vmx\", \"erms\", \"xtpr\", \"cmov\", \"smep\", \"ssse3\", \"est\", \"pat\", \"monitor\", \"smx\", \"pbe\", \"lm\", \"tsc\", \"nx\", \"fxsr\", \"tm\", \"sse4.1\", \"pae\", \"sse4.2\", \"pclmuldq\", \"acpi\", \"tsc-deadline\", \"mmx\", \"osxsave\", \"cx8\", \"mce\", \"de\", \"tm2\", \"ht\", \"dca\", \"lahf_lm\", \"popcnt\", \"mca\", \"pdpe1gb\", \"apic\", \"sse\", \"f16c\", \"pse\", \"ds\", \"invtsc\", \"pni\", \"rdtscp\", \"aes\", \"sse2\", \"ss\", \"ds_cpl\", \"pcid\", \"fpu\", \"cx16\", \"pse36\", \"mtrr\", \"pdcm\", \"rdrand\", \"x2apic\"], \"topology\": {\"cores\": 10, \"threads\": 2, \"sockets\": 1}}",
  "running_vms": 1,
  "free_disk_gb": 524,
  "hypervisor_version": 2002000,
  "disk_available_least": 521,
  "local_gb": 654,
  "free_ram_mb": 44032,
  "id": 14
}
}
`

// nova hypervisor-servers compute-0-2.domain.tld
const HypervisorServersList = `
{
  "hypervisors": [
    {
      "status": "enabled",
      "state": "up",
      "id": 14,
      "hypervisor_hostname": "compute-0-2.domain.tld",
      "servers": [
        {
          "uuid": "b32aa057-1880-46b8-888a-9855414e4a47",
          "name": "instance-00000030"
        },
        {
          "uuid": "1fd61852-2a6e-4f1f-8b62-c1af7f04e997",
          "name": "instance-0000015c"
        }
      ]
    }
  ]
}
`

//	nova hypervisor-uptime compute-0-2.domain.tld (or id which is faster)
const HypervisorUptimeBody = `
{
  "hypervisor": {
    "status": "enabled",
    "state": "up",
    "id": 14,
    "hypervisor_hostname": "compute-0-2.domain.tld",
    "uptime": " 15:22:34 up 20:27,  4 users,  load average: 2.23, 2.31, 2.35\n"
  }
}
`
var (
	// Hypervisor_3 is a Server struct that should correspond to the first
	// result in HypervisorListBody.
	hyper_1 = Hypervisor{
		HypervisorHostname: "compute-0-1.domain.tld",
		Id: 8,
		State: "up",
		Status: "enabled",
	}
	hyper_2 = Hypervisor{
		HypervisorHostname: "compute-0-2.domain.tld",
		Id: 14,
		State: "up",
		Status: "enabled",
	}
	hyper_3 = Hypervisor{
		HypervisorHostname: "compute-0-3.domain.tld",
		Id: 2,
		State: "up",
		Status: "enabled",
	}
	ListHypervisorsExpected = []Hypervisor{ hyper_3, hyper_1, hyper_2}
)

func HandleHypervisorsListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, HypervisorListBody)
	})
}

var (
	HypervisorDetail_1 = HypervisorDetail{
		HypervisorHostname: "compute-0-1.domain.tld",
		Id: 8,
		State: "up",
		Status: "enabled",
		Service: Service{
			Host: "compute-0-1.domain.tld",
			DisabledReason: "",
			Id: 38,
		},
		VcpuUsed: 0,
		HypervisorType: "QEMU",
		LocalGBUsed: 290,
		Vcpus: 22,
		MemoryMBUsed: 62088,
		MemoryMB: 64136,
		CurrentWorkload: 0,
		HostIP: "192.168.2.25",
		CPUInfo: map[string]interface{}{"vendor": "Intel",
			"model": "SandyBridge",
			"arch": "x86_64",
			"features": []string{"pge", "avx", "clflush", "sep",
				"syscall", "vme", "dtes64", "msr", "fsgsbase",
				"xsave", "vmx", "erms", "xtpr", "cmov", "smep",
				"ssse3", "est", "pat", "monitor", "smx", "pbe",
				"lm", "tsc", "nx", "fxsr", "tm", "sse4.1", "pae",
				"sse4.2", "pclmuldq", "acpi", "tsc-deadline",
				"mmx", "osxsave", "cx8", "mce", "de", "tm2",
				"ht", "dca", "lahf_lm", "popcnt", "mca", "pdpe1gb",
				"apic", "sse", "f16c", "pse", "ds", "invtsc", "pni",
				"rdtscp", "aes", "sse2", "ss", "ds_cpl", "pcid",
				"fpu", "cx16", "pse36", "mtrr", "pdcm", "rdrand", "x2apic"},
			"topology": map[string]int{"cores": 10, "threads": 2, "sockets": 1},
		},
		RunningVMs: 0,
		FreeDiskGB: 364,
		HypervisorVersion: 2002000,
		DistAvailableLeast: 601,
		LocalGB: 654,
		FreeRamMB: 2048,
	}
	HypervisorDetail_2 = HypervisorDetail{
		HypervisorHostname: "compute-0-2.domain.tld",
		Id: 14,
		State: "up",
		Status: "enabled",
		Service: Service{
			Host: "compute-0-2.domain.tld",
			DisabledReason: "",
			Id: 42,
		},
		VcpuUsed: 2,
		HypervisorType: "QEMU",
		LocalGBUsed: 130,
		Vcpus: 32,
		MemoryMBUsed: 20104,
		MemoryMB: 64136,
		CurrentWorkload: 0,
		HostIP: "192.168.2.23",
		CPUInfo: map[string]interface{}{"vendor": "Intel",
			"model": "SandyBridge",
			"arch": "x86_64",
			"features": []string{"pge", "avx", "clflush", "sep",
				"syscall", "vme", "dtes64", "msr", "fsgsbase",
				"xsave", "vmx", "erms", "xtpr", "cmov", "smep",
				"ssse3", "est", "pat", "monitor", "smx", "pbe",
				"lm", "tsc", "nx", "fxsr", "tm", "sse4.1", "pae",
				"sse4.2", "pclmuldq", "acpi", "tsc-deadline",
				"mmx", "osxsave", "cx8", "mce", "de", "tm2",
				"ht", "dca", "lahf_lm", "popcnt", "mca", "pdpe1gb",
				"apic", "sse", "f16c", "pse", "ds", "invtsc", "pni",
				"rdtscp", "aes", "sse2", "ss", "ds_cpl", "pcid",
				"fpu", "cx16", "pse36", "mtrr", "pdcm", "rdrand", "x2apic"},
			"topology": map[string]int{"cores": 10, "threads": 2, "sockets": 1},
		},
		RunningVMs: 1,
		FreeDiskGB: 524,
		HypervisorVersion: 2002000,
		DistAvailableLeast: 521,
		LocalGB: 654,
		FreeRamMB: 44032,
	}
	HypervisorDetail_3 = HypervisorDetail{
		HypervisorHostname: "compute-0-3.domain.tld",
		Id: 2,
		State: "up",
		Status: "enabled",
		Service: Service{
			Host: "compute-0-3.domain.tld",
			DisabledReason: "",
			Id: 33,
		},
		VcpuUsed: 0,
		HypervisorType: "QEMU",
		LocalGBUsed: 240,
		Vcpus: 24,
		MemoryMBUsed: 59016,
		MemoryMB: 64136,
		CurrentWorkload: 0,
		HostIP: "192.168.2.28",
		CPUInfo: map[string]interface{}{"vendor": "Intel",
			"model": "SandyBridge",
			"arch": "x86_64",
			"features": []string{"pge", "avx", "clflush", "sep",
				"syscall", "vme", "dtes64", "msr", "fsgsbase",
				"xsave", "vmx", "erms", "xtpr", "cmov", "smep",
				"ssse3", "est", "pat", "monitor", "smx", "pbe",
				"lm", "tsc", "nx", "fxsr", "tm", "sse4.1", "pae",
				"sse4.2", "pclmuldq", "acpi", "tsc-deadline",
				"mmx", "osxsave", "cx8", "mce", "de", "tm2",
				"ht", "dca", "lahf_lm", "popcnt", "mca", "pdpe1gb",
				"apic", "sse", "f16c", "pse", "ds", "invtsc", "pni",
				"rdtscp", "aes", "sse2", "ss", "ds_cpl", "pcid",
				"fpu", "cx16", "pse36", "mtrr", "pdcm", "rdrand", "x2apic"},
			"topology": map[string]int{"cores": 10, "threads": 2, "sockets": 1},
		},
		RunningVMs: 0,
		FreeDiskGB: 414,
		HypervisorVersion: 2002000,
		DistAvailableLeast: 633,
		LocalGB: 654,
		FreeRamMB: 5120,
	}

	HypervisorsDetailsListExpected = []HypervisorDetail{HypervisorDetail_3,
		HypervisorDetail_1, HypervisorDetail_2}
)

func HandleHypervisorsDetailsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, HypervisorsDetailsBody)
	})
}

func HandleHypervisorDetailSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/14", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, HypervisorDetailsBody)
	})
}

var HypervisorServiersListExpected = []HypervisorServersInfo{
	HypervisorServersInfo{
		HypervisorHostname: "compute-0-2.domain.tld",
		Id: 14,
		State: "up",
		Status: "enabled",
		Servers: []ServerBriefInfo {
			ServerBriefInfo{
				UUID: "b32aa057-1880-46b8-888a-9855414e4a47",
				Name: "instance-00000030",
			},
			ServerBriefInfo{
				UUID: "1fd61852-2a6e-4f1f-8b62-c1af7f04e997",
				Name: "instance-0000015c",
			},
		},
	},
}

func HandleHypervisorServerListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/compute-0-2.domain.tld/servers",
		func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, HypervisorServersList)
	})
}

var HypervisorUptimeExpected = HypervisorUptimeInfo{
	HypervisorHostname: "compute-0-2.domain.tld",
	Id: 14,
	State: "up",
	Status: "enabled",
	Uptime: " 15:22:34 up 20:27,  4 users,  load average: 2.23, 2.31, 2.35\n",
}

func HandleHypervisorUptimeSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/14/uptime",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, HypervisorUptimeBody)
		})
}
