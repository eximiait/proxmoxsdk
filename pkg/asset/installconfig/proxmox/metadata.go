package proxmox_installconfig

import (
	"sync"

	proxmoxsdk "github.com/Telmate/proxmox-api-go/proxmox"
)

//go:generate mockgen -source=./metadata.go -destination=./mock/powervsmetadata_generated.go -package=mock

// Metadata holds additional metadata for InstallConfig resources that
// do not need to be user-supplied (e.g. because it can be retrieved
// from external APIs).
type Metadata struct {
	BaseDomain string
	client     *proxmoxsdk.Client

	mutex sync.Mutex
}

// NewMetadata initializes a new Metadata object.
func NewMetadata(baseDomain string) *Metadata {
	return &Metadata{BaseDomain: baseDomain}
}
