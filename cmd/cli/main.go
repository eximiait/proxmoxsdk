package main

import (
	"fmt"

	proxmox_installconfig "github.com/eximiait/proxmoxsdk/pkg/asset/installconfig/proxmox"
)

func main() {
	p, err := proxmox_installconfig.Platform()
	fmt.Println(err)
	fmt.Println(p)
}
