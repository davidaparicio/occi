package util

import (
	"fmt"
	"log"
	"time"

	"github.com/davidaparicio/ovh-inventory/api/types"
	"github.com/ovh/go-ovh/ovh"
)

// PrintService prints all available services on /service
func PrintService(name string, client *ovh.Client) {
	var services []int
	url := "/service"
	err := client.Get(url, &services)
	if err != nil {
		log.Fatal("client.Get("+url+"):", err)
	}
	var service types.Service
	for _, s := range services {
		err = client.Get(url+"/"+fmt.Sprint(s), &service)
		if err != nil {
			log.Fatal("client.Get("+url+"/"+fmt.Sprint(s)+"):", err)
		}
		PrintIfNotEmptyS(name+":", service.Resource.Name)
	}
	time.Sleep(100 * time.Millisecond)
}

// PrintServices prints all available services on /services
func PrintServices(name string, client *ovh.Client) {
	var services []int
	url := "/services"
	err := client.Get(url, &services)
	if err != nil {
		log.Fatal("client.Get("+url+"):", err)
	}
	var service types.Services
	for _, s := range services {
		err = client.Get(url+"/"+fmt.Sprint(s), &service)
		if err != nil {
			log.Fatal("client.Get("+url+"/"+fmt.Sprint(s)+"):", err)
		}
		PrintIfNotEmptyS(name+":", service.Resource.Name)
	}
	time.Sleep(100 * time.Millisecond)
}
