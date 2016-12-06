package hypervisors


type Hypervisor struct {
	HypervisorHostname string	`mapstructure:"hypervisor_hostname"`
	Id int
	State string
	Status string
}

type Service struct {
	Host string
	DisabledReason string		`mapstructure:"disabled_reason"`
	Id int
}

type HypervisorDetail struct {
	HypervisorHostname string	`mapstructure:"hypervisor_hostname"`
	Id int
	State string
	Status string
	Service Service
	VcpuUsed int16					`mapstructure:"vcpus_used"`
	HypervisorType string			`mapstructure:"hypervisor_type"`
	LocalGBUsed	int16				`mapstructure:"local_gb_used"`
	Vcpus int16
	MemoryMBUsed int32				`mapstructure:"memory_mb_used"`
	MemoryMB int32					`mapstructure:"memory_mb"`
	CurrentWorkload int16			`mapstructure:"current_workload"`
	HostIP string 					`mapstructure:"host_ip"`
	CPUInfo map[string]interface{}	`mapstructure:"cpu_info"`
	RunningVMs int					`mapstructure:"running_vms"`
	FreeDiskGB int16				`mapstructure:"free_disk_gb"`
	HypervisorVersion int32			`mapstructure:"hypervisor_version"`
	DistAvailableLeast int16		`mapstructure:"disk_available_least"`
	LocalGB int16					`mapstructure:"local_gb"`
	FreeRamMB int32					`mapstructure:"free_ram_mb"`
}

type ServerBriefInfo struct {
	UUID string
	Name string
}

type HypervisorServersInfo struct {
	HypervisorHostname string	`mapstructure:"hypervisor_hostname"`
	Id int
	State string
	Status string
	Servers []ServerBriefInfo
}

type HypervisorUptimeInfo struct {
	HypervisorHostname string	`mapstructure:"hypervisor_hostname"`
	Id int
	State string
	Status string
	Uptime string
}
