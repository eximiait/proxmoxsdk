package proxmox_machines

import proxmox_types "github.com/eximiait/proxmoxsdk/pkg/types/proxmox"

func defaultProxmoxMachinePoolPlatform() proxmox_types.MachinePool {
	return proxmox_types.MachinePool{
		CPU: &proxmox_types.CPU{
			Cores:   4,
			Sockets: 1,
		},
		MemoryMB: 16348,
		OSDisk: &proxmox_types.Disk{
			SizeGB: decimalRootVolumeSize,
		},
	}
}

const (
	// decimalRootVolumeSize is the size in GB we use for some platforms.
	// See below.
	decimalRootVolumeSize = 120
)
