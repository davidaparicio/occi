package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davidaparicio/ovh-inventory/util"
	"github.com/ovh/go-ovh/ovh"
)

const (
	EXIT_NOCONF_FILE = iota + 1
)

// PartialMe holds the first name of the currently logged-in user.
// Visit https://api.ovh.com/console/#/me#GET for the full definition
type PartialMe struct {
	Firstname string `json:"firstname"`
	Nic       string `json:"nichandle"`
}

// Instantiate an OVH client and get the firstname of the currently logged-in user.
// Visit https://api.ovh.com/createToken/index.cgi?GET=/me to get your credentials.
func main() {
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
		os.Exit(EXIT_NOCONF_FILE)
	}

	c, _ := ovh.NewClient(
		conf.Endpoint,
		conf.Application_Key,
		conf.Application_Secret,
		conf.Consumer_Key,
	)

	var me PartialMe
	err = c.Get("/me", &me)
	if err != nil {
		log.Fatal("client.Get(/me):", err)
	}
	fmt.Printf("Welcome %s (%s)!\n", me.Firstname, me.Nic)

	util.PrintAssets("/allDom", "All Doms", c)
	util.PrintAssets("/cdn/dedicated", "CDN Dedicated", c)
	util.PrintAssets("/cloud/project", "PCI projects", c)
	util.PrintAssets("/cluster/hadoop", "Hadoop clusters", c)
	util.PrintAssets("/dbaas/logs", "DBAAS logs", c)
	fmt.Println("=== DEDICATED ===")
	util.PrintAssets("/dedicated/ceph", "-Ceph", c)
	util.PrintAssets("/dedicated/housing", "-Housing", c)
	util.PrintAssets("/dedicated/installationTemplate", "-Template", c)
	util.PrintAssets("/dedicated/nas", "-NAS", c)
	util.PrintAssets("/dedicated/nasha", "-NAS HA", c)
	//util.PrintAssets("/dedicated/nvmeof ", "-NVMeoF clusters", c) //:OVHcloud API error (status code 404): Client::NotFound: "Got an invalid (or empty) URL"
	util.PrintAssets("/dedicated/server", "-Servers", c)
	util.PrintAssets("/dedicatedCloud", "-VMware", c)
	fmt.Println("======")
	util.PrintAssets("/domain", "Domains", c)
	util.PrintAssets("/email/domain", "Email Domains", c)
	util.PrintAssets("/email/exchange", "Email Exchange", c)
	//util.PrintAssets("/email/mxplan ", "Email MX Plans", c) //:OVHcloud API error (status code 404): Client::NotFound: "Got an invalid (or empty) URL"
	util.PrintAssets("/email/pro", "Email Pro", c)
	util.PrintAssets("/freefax", "FreeFax", c)
	util.PrintAssets("/horizonView", "Horizon View as a Service", c)
	util.PrintAssets("/hosting/privateDatabase", "Private DBs", c)
	util.PrintAssets("/hosting/web", "Web Hosting projects", c)
	util.PrintAssets("/ip", "IPs", c)
	//util.PrintAssets("/ip/campus", "IP Campus", c) // client.Get(/ip/campus):json: cannot unmarshal object into Go value of type string
	util.PrintAssets("/ip/service", "IP Services", c)
	util.PrintAssets("/ipLoadbalancing", "IPLBs", c)
	fmt.Println("=== LICENCES ===")
	util.PrintAssets("/license/cloudLinux", "+CloudLinux", c)
	util.PrintAssets("/license/cpanel", "+CPanel", c)
	util.PrintAssets("/license/directadmin", "+DirectAdmin", c)
	util.PrintAssets("/license/office", "+Office", c)
	util.PrintAssets("/license/officePrepaid", "+OfficePrepaid", c)
	util.PrintAssets("/license/plesk", "+Plesk", c)
	util.PrintAssets("/license/redhat", "+RedHat", c)
	util.PrintAssets("/license/sqlserver", "+SQLServer", c)
	util.PrintAssets("/license/virtuozzo", "+Virtuozzo", c)
	util.PrintAssets("/license/windows", "+Windows", c)
	util.PrintAssets("/license/worklight", "+Worklight", c)
	fmt.Println("======")
	util.PrintAssets("/metrics", "Metrics", c)
	util.PrintAssets("/msServices", "msServices", c)
	util.PrintAssets("/nutanix", "Nutanix clusters", c)
	util.PrintAssets("/overTheBox", "OverTheBox", c)
	util.PrintAssets("/ovhCloudConnect", "ovhCloudConnect", c)
	util.PrintAssets("/pack/siptrunk", "siptrunk", c)
	util.PrintAssets("/pack/xdsl", "xdsl", c)
	util.PrintAssets("/saas/csp2", "csp2", c)
	//util.PrintAssets("/service", "?Service ?", c)
	//util.PrintAssets("/services", "?Services ?", c)
	util.PrintAssets("/sms", "SMS", c)
	util.PrintAssets("/ssl", "SSL", c)
	util.PrintAssets("/sslGateway", "SSL Gateway", c)
	util.PrintAssets("/stack/mis", "Stack MIS", c)
	util.PrintAssets("/storage/netapp", "NetAPP", c)
	util.PrintAssets("/veeam/veeamEnterprise", "VEEAM Enterprise", c)
	util.PrintAssets("/veeamCloudConnect", "VEEAM CloudConnect", c)
	util.PrintAssets("/vps", "VPS", c)
	util.PrintAssets("/vrack", "vRack", c)
	util.PrintAssets("/webPaaS/subscription", "WebPaaS", c)
	util.PrintAssets("/xdsl", "xDSL", c)

	fmt.Println("=== SERVICE ===")
	util.PrintService("Service", c)
	fmt.Println("======")

	fmt.Println("=== SERVICES ===")
	util.PrintService("Service", c)
	fmt.Println("======")

	fmt.Println("=== PCI (Public Cloud) ===")
	util.PrintPCI(c)
	fmt.Println("======")

	/*fmt.Println("=== PCC (Private Cloud) ===")
	util.PrintPCC(c)
	fmt.Println("======")*/
}
