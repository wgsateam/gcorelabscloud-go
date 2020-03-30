package instances

import (
	"net"

	"bitbucket.gcore.lu/gcloud/gcorecloud-go/gcore/instance/v1/types"

	"bitbucket.gcore.lu/gcloud/gcorecloud-go"
	"bitbucket.gcore.lu/gcloud/gcorecloud-go/gcore/flavor/v1/flavors"
	"bitbucket.gcore.lu/gcloud/gcorecloud-go/pagination"
)

type commonResult struct {
	gcorecloud.Result
}

// Extract is a function that accepts a result and extracts a instance resource.
func (r commonResult) Extract() (*Instance, error) {
	var s Instance
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Instance.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Instance.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Instance.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation
type DeleteResult struct {
	gcorecloud.ErrResult
}

type InstanceVolume struct {
	ID                  string `json:"id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type InstanceAddress struct {
	Type       types.AddressType `json:"type"`
	Address    net.IP            `json:"addr"`
	SubnetID   *string           `json:"subnet_id,omitempty"`
	SubnetName *string           `json:"subnet_name,omitempty"`
}

// Instance represents a instance structure.
type Instance struct {
	ID             string                       `json:"instance_id"`
	Name           string                       `json:"instance_name"`
	Description    string                       `json:"instance_description"`
	CreatedAt      gcorecloud.JSONRFC3339ZZ     `json:"instance_created"`
	Status         string                       `json:"status"`
	VMState        string                       `json:"vm_state"`
	TaskState      *string                      `json:"task_state"`
	Flavor         flavors.Flavor               `json:"flavor"`
	Metadata       map[string]interface{}       `json:"metadata"`
	Volumes        []InstanceVolume             `json:"volumes"`
	Addresses      map[string][]InstanceAddress `json:"addresses"`
	SecurityGroups []types.ItemName             `json:"security_groups"`
	CreatorTaskID  *string                      `json:"creator_task_id"`
	TaskID         *string                      `json:"task_id"`
	ProjectID      int                          `json:"project_id"`
	RegionID       int                          `json:"region_id"`
	Region         string                       `json:"region"`
}

// Interface represents a instance port interface.
type Interface struct {
	PortID            string         `json:"port_id"`
	MacAddress        gcorecloud.MAC `json:"mac_address"`
	NetworkID         string         `json:"network_id"`
	IPAssignments     []PortIP       `json:"ip_assignments"`
	NetworkDetails    NetworkDetail  `json:"network_details"`
	FloatingIPDetails []FloatingIP   `json:"floatingip_details"`
}

// PortIP
type PortIP struct {
	IPAddress net.IP `json:"ip_address"`
	SubnetID  string `json:"subnet_id"`
}

// FloatingIP represents a floating ip of instance port.
type FloatingIP struct {
	FloatingIPAddress net.IP                   `json:"floating_ip_address"`
	RouterID          string                   `json:"router_id"`
	SubnetID          string                   `json:"subnet_id"`
	Status            string                   `json:"status"`
	ID                string                   `json:"id"`
	PortID            string                   `json:"port_id"`
	DNSDomain         string                   `json:"dns_domain"`
	DNSName           string                   `json:"dns_name"`
	FixedIPAddress    net.IP                   `json:"fixed_ip_address"`
	UpdatedAt         *gcorecloud.JSONRFC3339Z `json:"updated_at"`
	CreatedAt         gcorecloud.JSONRFC3339Z  `json:"created_at"`
	CreatorTaskID     *string                  `json:"creator_task_id"`
	ProjectID         int                      `json:"project_id"`
	RegionID          int                      `json:"region_id"`
	Region            string                   `json:"region"`
}

// Subnet port subnet
type Subnet struct {
	ID            string                   `json:"id"`
	Name          string                   `json:"name"`
	IPVersion     gcorecloud.IPVersion     `json:"ip_version"`
	EnableDHCP    bool                     `json:"enable_dhcp"`
	Cidr          gcorecloud.CIDR          `json:"cidr"`
	CreatedAt     gcorecloud.JSONRFC3339Z  `json:"created_at"`
	UpdatedAt     *gcorecloud.JSONRFC3339Z `json:"updated_at"`
	NetworkID     string                   `json:"network_id"`
	TaskID        *string                  `json:"task_id"`
	CreatorTaskID *string                  `json:"creator_task_id"`
	ProjectID     int                      `json:"project_id"`
	RegionID      int                      `json:"region_id"`
	Region        string                   `json:"region"`
}

// NetworkDetail represents a NetworkDetails of instance port.
type NetworkDetail struct {
	Mtu           int                      `json:"mtu"`
	UpdatedAt     *gcorecloud.JSONRFC3339Z `json:"updated_at"`
	CreatedAt     gcorecloud.JSONRFC3339Z  `json:"created_at"`
	ID            string                   `json:"id"`
	External      bool                     `json:"external"`
	Default       bool                     `json:"default"`
	Name          string                   `json:"name"`
	Shared        bool                     `json:"shared"`
	Subnets       []Subnet                 `json:"subnets"`
	ProjectID     int                      `json:"project_id"`
	RegionID      int                      `json:"region_id"`
	Region        string                   `json:"region"`
	TaskID        *string                  `json:"task_id"`
	CreatorTaskID *string                  `json:"creator_task_id"`
}

// InstancePage is the page returned by a pager when traversing over a
// collection of instances.
type InstancePage struct {
	pagination.LinkedPageBase
}

// InstanceInterfacePage is the page returned by a pager when traversing over a
// collection of instance interfaces.
type InstanceInterfacePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of instances has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r InstancePage) NextPageURL() (string, error) {
	var s struct {
		Links []gcorecloud.Link `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gcorecloud.ExtractNextURL(s.Links)
}

// InstanceInterfacePage is invoked when a paginated collection of instance interfaces has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r InstanceInterfacePage) NextPageURL() (string, error) {
	var s struct {
		Links []gcorecloud.Link `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gcorecloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a InstancePage struct is empty.
func (r InstancePage) IsEmpty() (bool, error) {
	is, err := ExtractInstances(r)
	return len(is) == 0, err
}

// IsEmpty checks whether a InstancePage struct is empty.
func (r InstanceInterfacePage) IsEmpty() (bool, error) {
	is, err := ExtractInstanceInterfaces(r)
	return len(is) == 0, err
}

// ExtractInstances accepts a Page struct, specifically a InstancePage struct,
// and extracts the elements into a slice of Instance structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractInstances(r pagination.Page) ([]Instance, error) {
	var s []Instance
	err := ExtractInstancesInto(r, &s)
	return s, err
}

// ExtractInstanceInterfaces accepts a Page struct, specifically a InstancePage struct,
// and extracts the elements into a slice of Instance structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractInstanceInterfaces(r pagination.Page) ([]Interface, error) {
	var s []Interface
	err := ExtractInstancesInto(r, &s)
	return s, err
}

func ExtractInstancesInto(r pagination.Page, v interface{}) error {
	return r.(InstancePage).Result.ExtractIntoSlicePtr(v, "results")
}