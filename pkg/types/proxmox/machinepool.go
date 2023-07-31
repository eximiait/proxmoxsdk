package proxmox_types

// MachinePool stores the configuration for a machine pool installed
// on ovirt.
type MachinePool struct {
	// InstanceTypeID defines the VM instance type and overrides
	// the hardware parameters of the created VM, including cpu and memory.
	// If InstanceTypeID is passed, all memory and cpu variables will be ignored.
	InstanceTypeID string `json:"instanceTypeID,omitempty"`

	// CPU defines the VM CPU.
	// +optional
	CPU *CPU `json:"cpu,omitempty"`

	// MemoryMB is the size of a VM's memory in MiBs.
	// +optional
	MemoryMB int32 `json:"memoryMB,omitempty"`

	// OSDisk is the the root disk of the node.
	// +optional
	OSDisk *Disk `json:"osDisk,omitempty"`

	// Clone makes sure that the disks are cloned from the template and are not linked.
	// Defaults to true for high performance and server VM types, false for desktop types.
	//
	// Note: this option is not documented in the OpenShift documentation. This is intentional as it has sane defaults
	// that shouldn't be changed unless needed for debugging or resolving issues in cooperation with Red Hat support.
	//
	// +optional
	Clone *bool `json:"clone,omitempty"`

	// Sparse indicates that sparse provisioning should be used and disks should be not preallocated.
	// +optional
	Sparse *bool `json:"sparse,omitempty"`

	// Format is the disk format that the disks are in. Can be "cow" or "raw". "raw" disables several features that
	// may be needed, such as incremental backups.
	// +kubebuilder:validation:Enum="";raw;cow
	// +optional
	Format string `json:"format,omitempty"`
}

// Disk defines a VM disk
type Disk struct {
	// SizeGB size of the bootable disk in GiB.
	SizeGB int64 `json:"sizeGB"`
}

// CPU defines the VM cpu, made of (Sockets * Cores).
type CPU struct {
	// Sockets is the number of sockets for a VM.
	// Total CPUs is (Sockets * Cores)
	Sockets int32 `json:"sockets"`

	// Cores is the number of cores per socket.
	// Total CPUs is (Sockets * Cores)
	Cores int32 `json:"cores"`
}

// Set sets the values from `required` to `p`.
func (p *MachinePool) Set(required *MachinePool) {
	if required == nil || p == nil {
		return
	}

	if required.InstanceTypeID != "" {
		p.InstanceTypeID = required.InstanceTypeID
	}

	if required.CPU != nil {
		p.CPU = required.CPU
	}

	if required.MemoryMB != 0 {
		p.MemoryMB = required.MemoryMB
	}

	if required.OSDisk != nil {
		p.OSDisk = required.OSDisk
	}

	p.Clone = required.Clone
	p.Format = required.Format
	p.Sparse = required.Sparse
}
