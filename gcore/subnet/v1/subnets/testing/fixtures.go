package testing

import (
	"fmt"
	"net"
	"time"

	gcorecloud "github.com/G-Core/gcorelabscloud-go"
	"github.com/G-Core/gcorelabscloud-go/gcore/subnet/v1/subnets"
	"github.com/G-Core/gcorelabscloud-go/gcore/task/v1/tasks"
	fake "github.com/G-Core/gcorelabscloud-go/testhelper/client"
)

var ListResponse = fmt.Sprintf(`
{
  "count": 1,
  "results": [
    {
      "id": "e7944e55-f957-413d-aa56-fdc876543113",
      "name": "subnet",
      "ip_version": 4,
      "enable_dhcp": true,
      "cidr": "192.168.10.0/24",
      "created_at": "2020-03-05T12:03:24+0000",
      "updated_at": "2020-03-05T12:03:25+0000",
	  "network_id": "ee2402d0-f0cd-4503-9b75-69be1d11c5f1",
	  "task_id": "50f53a35-42ed-40c4-82b2-5a37fb3e00bc",
	  "creator_task_id": "50f53a35-42ed-40c4-82b2-5a37fb3e00bc",
	  "region": "RegionOne",
	  "available_ips": %d,
	  "total_ips": %d,
      "project_id": 1,
      "region_id": 1,
	  "dns_nameservers": [
		"10.0.0.13"
	  ],
	  "has_router": true,
	  "host_routes": [
		{
		  "destination": "10.0.3.0/24",
		  "nexthop": "10.0.0.13"
		}
	  ]	
    }
  ]
}
`, availableIps, totalIps)

var GetResponse = fmt.Sprintf(`
{
  "id": "e7944e55-f957-413d-aa56-fdc876543113",
  "name": "subnet",
  "ip_version": 4,
  "enable_dhcp": true,
  "cidr": "192.168.10.0/24",
  "created_at": "2020-03-05T12:03:24+0000",
  "updated_at": "2020-03-05T12:03:25+0000",
  "network_id": "ee2402d0-f0cd-4503-9b75-69be1d11c5f1",
  "task_id": "50f53a35-42ed-40c4-82b2-5a37fb3e00bc",
  "creator_task_id": "50f53a35-42ed-40c4-82b2-5a37fb3e00bc",
  "region": "RegionOne",
  "project_id": 1,
  "region_id": 1,
  "available_ips": %d,
  "total_ips": %d,
  "dns_nameservers": [
	"10.0.0.13"
  ],
  "has_router": true,
  "host_routes": [
    {
      "destination": "10.0.3.0/24",
      "nexthop": "10.0.0.13"
    }
  ]	
}
`, availableIps, totalIps)

const CreateRequest = `
{
  "name": "subnet",
  "enable_dhcp": true,
  "cidr": "192.168.10.0/24",
  "network_id": "ee2402d0-f0cd-4503-9b75-69be1d11c5f1",
  "connect_to_network_router": true
}
`

const UpdateRequest = `
{
	"name": "subnet"
}	
`

const CreateResponse = `
{
  "tasks": [
    "50f53a35-42ed-40c4-82b2-5a37fb3e00bc"
  ]
}
`

const DeleteResponse = `
{
  "tasks": [
    "50f53a35-42ed-40c4-82b2-5a37fb3e00bc"
  ]
}
`

var createdTimeString = "2020-03-05T12:03:24+0000"
var updatedTimeString = "2020-03-05T12:03:25+0000"
var createdTimeParsed, _ = time.Parse(gcorecloud.RFC3339Z, createdTimeString)
var createdTime = gcorecloud.JSONRFC3339Z{Time: createdTimeParsed}
var updatedTimeParsed, _ = time.Parse(gcorecloud.RFC3339Z, updatedTimeString)
var updatedTime = gcorecloud.JSONRFC3339Z{Time: updatedTimeParsed}
var cidr, _ = gcorecloud.ParseCIDRString("192.168.10.0/24")
var taskID = "50f53a35-42ed-40c4-82b2-5a37fb3e00bc"
var availableIps = 241
var totalIps = 243
var ip = net.ParseIP("10.0.0.13")
var routeCidr, _ = gcorecloud.ParseCIDRString("10.0.3.0/24")

var (
	Subnet1 = subnets.Subnet{
		ID:             "e7944e55-f957-413d-aa56-fdc876543113",
		Name:           "subnet",
		IPVersion:      4,
		EnableDHCP:     true,
		CIDR:           *cidr,
		CreatedAt:      createdTime,
		UpdatedAt:      updatedTime,
		NetworkID:      "ee2402d0-f0cd-4503-9b75-69be1d11c5f1",
		TaskID:         taskID,
		CreatorTaskID:  taskID,
		Region:         "RegionOne",
		ProjectID:      fake.ProjectID,
		RegionID:       fake.RegionID,
		AvailableIps:   availableIps,
		TotalIps:       totalIps,
		HasRouter:      true,
		DNSNameservers: []net.IP{ip},
		HostRoutes: []subnets.HostRoute{
			{DestinationCIDR: *routeCidr, NextHop: ip},
		},
	}
	Tasks1 = tasks.TaskResults{
		Tasks: []tasks.TaskID{"50f53a35-42ed-40c4-82b2-5a37fb3e00bc"},
	}

	ExpectedSubnetSlice = []subnets.Subnet{Subnet1}
)
