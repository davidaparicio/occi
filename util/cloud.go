package util

import (
	"fmt"
	"log"
	"time"

	"github.com/davidaparicio/ovh-inventory/api/types"
	"github.com/ovh/go-ovh/ovh"
)

// PrintPCI prints all available instances/volumes/storage on each PCI project
func PrintPCI(client *ovh.Client) {
	var projects_pci []string
	url := "/cloud/project"
	err := client.Get(url, &projects_pci)
	if err != nil {
		log.Fatal("client.Get("+url+"):", err)
	}

	var pci_project types.PCI_Project
	var pci_instances types.PCI_Instance
	var pci_volumes types.PCI_Volumes
	var pci_swifts types.PCI_Swift

	for _, project := range projects_pci {

		err = client.Get(url+"/"+fmt.Sprint(project), &pci_project)
		if err != nil {
			log.Fatal("client.Get("+url+"/"+fmt.Sprint(project)+"):", err)
		}

		fmt.Printf("==== Project %s (%s) ====\n", pci_project.ProjectID, pci_project.Description)
		time.Sleep(100 * time.Millisecond)

		err = client.Get(url+"/"+fmt.Sprint(project)+"/instance", &pci_instances)
		if err != nil {
			log.Fatal("client.Get("+url+"/"+fmt.Sprint(project)+"/instance"+"):", err)
		}

		if len(pci_instances) > 0 {
			fmt.Println("===== VMs =====")
			for _, instance := range pci_instances {
				fmt.Println(instance.ID)
			}
			fmt.Println("==========")
		}
		time.Sleep(100 * time.Millisecond)

		err = client.Get(url+"/"+fmt.Sprint(project)+"/volume", &pci_volumes)
		if err != nil {
			log.Fatal("client.Get("+url+"/"+fmt.Sprint(project)+"/volume"+"):", err)
		}

		if len(pci_volumes) > 0 {
			fmt.Println("===== Volumes =====")
			for _, volume := range pci_volumes {
				fmt.Printf("Volume: %s attached to VM: %v\n", volume.ID, volume.AttachedTo)
			}
			fmt.Println("==========")
		}
		time.Sleep(100 * time.Millisecond)

		err = client.Get(url+"/"+fmt.Sprint(project)+"/storage", &pci_swifts)
		if err != nil {
			log.Fatal("client.Get("+url+"/"+fmt.Sprint(project)+"/storage"+"):", err)
		}

		if len(pci_swifts) > 0 {
			fmt.Println("===== Swift / S3 =====")
			for _, swift := range pci_swifts {
				fmt.Printf("Swift: %s (%s)\n", swift.ID, swift.Name)
			}
			fmt.Println("==========")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
