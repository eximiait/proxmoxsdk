package proxmox_wrapper

import (
	proxmoxsdk "github.com/Telmate/proxmox-api-go/proxmox"
)

func GetNetworkListFromNode(node string, c proxmoxsdk.Client) (metricServers map[string]interface{}, err error) {
	return c.GetItemList("/nodes/" + node + "/network")
}
