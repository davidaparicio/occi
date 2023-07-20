package types

import "time"

type PCI_Project struct {
	ProjectID    string    `json:"project_id,omitempty"`
	ProjectName  string    `json:"projectName,omitempty"`
	Description  string    `json:"description,omitempty"`
	PlanCode     string    `json:"planCode,omitempty"`
	Unleash      bool      `json:"unleash,omitempty"`
	Expiration   any       `json:"expiration,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
	OrderID      int       `json:"orderId,omitempty"`
	Access       string    `json:"access,omitempty"`
	Status       string    `json:"status,omitempty"`
	ManualQuota  bool      `json:"manualQuota,omitempty"`
}

type PCI_Instance []struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	IPAddresses []struct {
		IP        string `json:"ip,omitempty"`
		Type      string `json:"type,omitempty"`
		Version   int    `json:"version,omitempty"`
		NetworkID string `json:"networkId,omitempty"`
		GatewayIP string `json:"gatewayIp,omitempty"`
	} `json:"ipAddresses,omitempty"`
	FlavorID       string    `json:"flavorId,omitempty"`
	ImageID        string    `json:"imageId,omitempty"`
	SSHKeyID       string    `json:"sshKeyId,omitempty"`
	Created        time.Time `json:"created,omitempty"`
	Region         string    `json:"region,omitempty"`
	MonthlyBilling struct {
		Since  time.Time `json:"since,omitempty"`
		Status string    `json:"status,omitempty"`
	} `json:"monthlyBilling,omitempty"`
	Status                      string `json:"status,omitempty"`
	PlanCode                    string `json:"planCode,omitempty"`
	OperationIds                []any  `json:"operationIds,omitempty"`
	CurrentMonthOutgoingTraffic any    `json:"currentMonthOutgoingTraffic,omitempty"`
}

type PCI_Volumes []struct {
	ID           string    `json:"id,omitempty"`
	AttachedTo   []string  `json:"attachedTo,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Size         int       `json:"size,omitempty"`
	Status       string    `json:"status,omitempty"`
	Region       string    `json:"region,omitempty"`
	Bootable     bool      `json:"bootable,omitempty"`
	PlanCode     string    `json:"planCode,omitempty"`
	Type         string    `json:"type,omitempty"`
}

type PCI_Swift []struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Archive       any    `json:"archive,omitempty"`
	ContainerType any    `json:"containerType,omitempty"`
	StoredObjects int    `json:"storedObjects,omitempty"`
	StoredBytes   int    `json:"storedBytes,omitempty"`
	Region        string `json:"region,omitempty"`
}
