package pools

import (
	"fmt"
	"net"

	"github.com/G-Core/gcorelabscloud-go/gcore/k8s/v1/types"

	gcorecloud "github.com/G-Core/gcorelabscloud-go"
	"github.com/G-Core/gcorelabscloud-go/gcore/task/v1/tasks"
	"github.com/G-Core/gcorelabscloud-go/pagination"
)

type commonResult struct {
	gcorecloud.Result
}

// Extract is a function that accepts a result and extracts a cluster Pool resource.
func (r commonResult) Extract() (*ClusterPool, error) {
	var s ClusterPool
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Pool.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Pool.
type UpdateResult struct {
	commonResult
}

// ClusterPool represents a cluster Pool.
type ClusterPool struct {
	ClusterID        string            `json:"cluster_id"`
	ProjectID        string            `json:"project_id"`
	Labels           map[string]string `json:"labels"`
	NodeAddresses    []net.IP          `json:"node_addresses"`
	StatusReason     string            `json:"status_reason"`
	DockerVolumeSize int               `json:"docker_volume_size"`
	*ClusterListPool
}

// ClusterListPool represents a cluster Pool in the list response.
type ClusterListPool struct {
	UUID         string         `json:"uuid"`
	Name         string         `json:"name"`
	FlavorID     string         `json:"flavor_id"`
	ImageID      string         `json:"image_id"`
	NodeCount    int            `json:"node_count"`
	MinNodeCount int            `json:"min_node_count"`
	MaxNodeCount *int           `json:"max_node_count"`
	IsDefault    bool           `json:"is_default"`
	StackID      string         `json:"stack_id"`
	Status       string         `json:"status"`
	Role         types.PoolRole `json:"role"`
}

// ClusterPoolPage is the page returned by a pager when traversing over a
// collection of networks.
type ClusterPoolPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of cluster Pools has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r ClusterPoolPage) NextPageURL() (string, error) {
	var s struct {
		Links []gcorecloud.Link `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gcorecloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a ClusterPool struct is empty.
func (r ClusterPoolPage) IsEmpty() (bool, error) {
	is, err := ExtractClusterPools(r)
	return len(is) == 0, err
}

// ExtractClusterPools accepts a Page struct, specifically a ClusterPoolPage struct,
// and extracts the elements into a slice of ClusterPool structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractClusterPools(r pagination.Page) ([]ClusterListPool, error) {
	var s []ClusterListPool
	err := ExtractClusterPoolsInto(r, &s)
	return s, err
}

func ExtractClusterPoolsInto(r pagination.Page, v interface{}) error {
	return r.(ClusterPoolPage).Result.ExtractIntoSlicePtr(v, "results")
}

type ClusterTaskResult struct {
	Pools []string `json:"pools"`
}

func ExtractClusterPoolIDFromTask(task *tasks.Task) (string, error) {
	var result ClusterTaskResult
	err := gcorecloud.NativeMapToStruct(task.CreatedResources, &result)
	if err != nil {
		return "", fmt.Errorf("cannot decode cluster information in task structure: %w", err)
	}
	if len(result.Pools) == 0 {
		return "", fmt.Errorf("cannot decode cluster information in task structure: %w", err)
	}
	return result.Pools[0], nil
}