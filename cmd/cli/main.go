package main

import (
	"fmt"

	proxmox_installconfig "github.com/eximiait/proxmox_installer/pkg/asset/installconfig"
)

func main() {
	p, err := proxmox_installconfig.Platform()
	fmt.Println(err)
	fmt.Println(p)
}
