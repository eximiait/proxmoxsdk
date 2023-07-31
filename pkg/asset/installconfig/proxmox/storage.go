package proxmox_installconfig

import (
	"errors"
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	proxmox_types "github.com/eximiait/proxmox_installer/pkg/types/proxmox"
	proxmox_wrapper "github.com/eximiait/proxmox_installer/pkg/wrapper"
)

func askStorage(c *proxmox_wrapper.Connection, p *proxmox_types.Platform) error {
	var storageName string
	domainList, err := c.Client().GetStorageList()
	if err != nil {
		return err
	}
	dataInterfaces, ok := domainList["data"].([]interface{})
	if !ok {
		return errors.New("the key 'data' is not in proxmox api response")
	}
	storageList := make([]string, 0)
	// Recorrer la lista de interfaces y acceder a los mapas
	for _, dataInterface := range dataInterfaces {
		// Verificar si la interfaz es de tipo map[string]interface{}
		if dataMap, ok := dataInterface.(map[string]interface{}); ok {
			// Acceder a los valores de cada mapa
			if storageValue, exists := dataMap["storage"].(string); exists {
				storageList = append(storageList, storageValue)
			}
		}
	}
	if err := survey.AskOne(
		&survey.Select{
			Message: "Storage name",
			Help:    "The storage name will be used to create the disks of all the cluster nodes.",
			Options: storageList,
		},
		&storageName,
		survey.WithValidator(func(ans interface{}) error {
			choice := ans.(core.OptionAnswer).Value
			sort.Strings(storageList)
			i := sort.SearchStrings(storageList, choice)
			if i == len(storageList) || storageList[i] != choice {
				return fmt.Errorf("invalid storage " + choice)
			}
			storage := choice
			p.StorageID = storage
			return nil
		}),
	); err != nil {
		return err
	}
	return err
}
