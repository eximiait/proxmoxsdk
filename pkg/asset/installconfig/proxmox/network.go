package proxmox_installconfig

import (
	"errors"
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	proxmox_types "github.com/eximiait/proxmoxsdk/pkg/types/proxmox"
	proxmox_wrapper "github.com/eximiait/proxmoxsdk/pkg/wrapper"
)

func askNetwork(c *proxmox_wrapper.Connection, p *proxmox_types.Platform) error {
	var selectedNode string
	//var networkName string
	nodeListRaw, err := c.Client().GetNodeList()
	if err != nil {
		return err
	}
	dataInterfaces, ok := nodeListRaw["data"].([]interface{})
	if !ok {
		return errors.New("the key 'data' is not in proxmox api response")
	}
	nodeList := make([]string, 0)
	// Recorrer la lista de interfaces y acceder a los mapas
	for _, dataInterface := range dataInterfaces {
		// Verificar si la interfaz es de tipo map[string]interface{}
		if dataMap, ok := dataInterface.(map[string]interface{}); ok {
			// Acceder a los valores de cada mapa
			if nodeValue, exists := dataMap["node"].(string); exists {
				nodeList = append(nodeList, nodeValue)
			}
		}
	}

	if err := survey.AskOne(
		&survey.Select{
			Message: "Node",
			Help:    "The Node from where obtain the netwroks list",
			Options: nodeList,
		},
		&selectedNode,
		survey.WithValidator(func(ans interface{}) error {
			choice := ans.(core.OptionAnswer).Value
			sort.Strings(nodeList)
			i := sort.SearchStrings(nodeList, choice)
			if i == len(nodeList) || nodeList[i] != choice {
				return fmt.Errorf("cannot find a node by name %s", choice)
			}
			return nil
		}),
	); err != nil {
		return fmt.Errorf("failed UserInput")
	}

	var networkName string
	networkListRaw, err := proxmox_wrapper.GetNetworkListFromNode(selectedNode, *c.Client())
	if err != nil {
		return err
	}
	dataInterfacesNet, ok := networkListRaw["data"].([]interface{})
	if !ok {
		return errors.New("the key 'data' is not in proxmox api response")
	}
	networkList := make([]string, 0)
	// Recorrer la lista de interfaces y acceder a los mapas
	for _, dataInterface := range dataInterfacesNet {
		// Verificar si la interfaz es de tipo map[string]interface{}
		if dataMap, ok := dataInterface.(map[string]interface{}); ok {
			// Acceder a los valores de cada mapa
			if networkValue, exists := dataMap["iface"].(string); exists {
				networkList = append(networkList, networkValue)
			}
		}
	}

	if err := survey.AskOne(
		&survey.Select{
			Message: "Interface",
			Help:    "The network interface",
			Options: networkList,
		},
		&networkName,
		survey.WithValidator(func(ans interface{}) error {
			choice := ans.(core.OptionAnswer).Value
			sort.Strings(networkList)
			i := sort.SearchStrings(networkList, choice)
			if i == len(networkList) || networkList[i] != choice {
				return fmt.Errorf("cannot find a network by name %s", choice)
			}
			p.NetworkName = choice
			return nil
		}),
	); err != nil {
		return fmt.Errorf("failed UserInput")
	}
	return nil
}

func askVIPs(p *proxmox_types.Platform) error {
	var apiVIP, ingressVIP string
	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Internal API virtual IP",
				Help:    "This is the virtual IP address that will be used to address the OpenShift control plane. Make sure the IP address is not in use.",
				Default: "",
			},
		},
	}, &apiVIP)
	if err != nil {
		return fmt.Errorf("failed UserInput")
	}
	p.APIVIPs = []string{apiVIP}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Ingress virtual IP",
				Help:    "This is the virtual IP address that will be used to address the OpenShift ingress routers. Make sure the IP address is not in use.",
				Default: "",
			},
		},
	}, &ingressVIP)
	if err != nil {
		return fmt.Errorf("failed UserInput")
	}
	p.IngressVIPs = []string{ingressVIP}

	return nil
}
