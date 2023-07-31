package proxmox_types

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	PoolID      string `json:"proxmox_pool_id"`
	GroupID     string `json:"proxmox_group_id"`
	StorageID   string `json:"proxmox_storage_id"`
	NetworkName string `json:"proxmox_network_name,omitempty"`
	// APIVIPs contains the VIP(s) which will be served by bootstrap and then
	// pivoted masters, using keepalived. In dual stack clusters it contains an
	// IPv4 and IPv6 address, otherwise only one VIP
	//
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:UniqueItems=true
	// +kubebuilder:validation:Format=ip
	// +optional
	APIVIPs []string `json:"api_vips,omitempty"`
	// IngressVIPs are external IP(s) which route to the default ingress
	// controller. The VIPs are suitable targets of wildcard DNS records used to
	// resolve default route host names. In dual stack clusters it contains an
	// IPv4 and IPv6 address, otherwise only one VIP
	//
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:UniqueItems=true
	// +kubebuilder:validation:Format=ip
	// +optional
	IngressVIPs []string `json:"ingress_vips,omitempty"`
}
