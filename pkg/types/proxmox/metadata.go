package proxmox_types

// Metadata contains ovirt metadata (e.g. for uninstalling the cluster).
type Metadata struct {
	RemoveTemplate bool `json:"remove_template"`
}
