package cloudcone

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	ID      = "cloudcone"
	ApiKey  = "api_key"
	ApiHash = "api_hash"
)

type InstancesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Instances []struct {
			Id       int    `json:"id,omitempty"`
			Hostname string `json:"hostname,omitempty"`
			Created  string `json:"created,omitempty"`
			Pvt      int    `json:"pvt,omitempty"`
			NodeName string `json:"node_name,omitempty"`
			State    struct {
				Name  string `json:"name,omitempty"`
				Color string `json:"color,omitempty"`
				Id    int    `json:"id,omitempty" `
			} `json:"state"`
			Xid    string `json:"xid,omitempty"`
			Ips    string `json:"ips,omitempty"`
			Distro string `json:"distro,omitempty"`
			Ram    int    `json:"ram,omitempty"`
			Cpu    int    `json:"cpu,omitempty"`
			Disk   int    `json:"disk,omitempty"`
		} `json:"instances"`
		TotalRam       int `json:"total_ram"`
		TotalDisk      int `json:"total_disk"`
		TotalCpu       int `json:"total_cpu"`
		TotalInstances int `json:"total_instances"`
	} `json:"__data"`
}

type InstanceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Instances struct {
			Hostname string `json:"hostname"`
			Label    string `json:"label"`
			Status   string `json:"status"`
			Template string `json:"template"`
			Mainip   string `json:"mainip"`
			Ips      []struct {
				ID      int    `json:"id"`
				Address string `json:"address"`
				Prefix  int    `json:"prefix"`
				Gateway string `json:"gateway"`
				Ipv4    bool   `json:"ipv4"`
			} `json:"ips"`
			Disk struct {
				Total int  `json:"total"`
				Used  bool `json:"used"`
				Free  bool `json:"free"`
				Usage bool `json:"usage"`
			} `json:"disk"`
			RAM                 int    `json:"ram"`
			CPU                 int    `json:"cpu"`
			PrivateIP           bool   `json:"private_ip"`
			Ipv6                int    `json:"ipv6"`
			Node                int    `json:"node"`
			RecoveryMode        bool   `json:"recovery_mode"`
			InitialRootPassword string `json:"initial_root_password"`
			AttachedDisks       []struct {
				ID         int    `json:"id"`
				Size       int    `json:"size"`
				IsSwap     bool   `json:"is_swap"`
				Built      bool   `json:"built"`
				Locked     bool   `json:"locked"`
				Primary    bool   `json:"primary"`
				MountPoint string `json:"mount_point"`
				Mounted    bool   `json:"mounted"`
				Label      string `json:"label"`
				FileSystem string `json:"file_system"`
			} `json:"attached_disks"`
			AttachedNics struct {
				Num225929 struct {
					ID                  int    `json:"id"`
					Label               string `json:"label"`
					Primary             bool   `json:"primary"`
					MacAddress          string `json:"mac_address"`
					DefaultFirewallRule string `json:"default_firewall_rule"`
				} `json:"225929"`
			} `json:"attached_nics"`
			Backups      bool   `json:"backups"`
			HypervisorID int    `json:"hypervisor_id"`
			Created      string `json:"created"`
			Billing      struct {
				Hourly struct {
					CPU float64 `json:"cpu"`
					RAM float64 `json:"ram"`
					Hdd float64 `json:"hdd"`
					Ssd int     `json:"ssd"`
					Ips float64 `json:"ips"`
				} `json:"hourly"`
				Monthly struct {
					CPU       float64 `json:"cpu"`
					RAM       float64 `json:"ram"`
					Hdd       float64 `json:"hdd"`
					Ssd       int     `json:"ssd"`
					Ips       float64 `json:"ips"`
					Backups   int     `json:"backups"`
					Snapshots float64 `json:"snapshots"`
				} `json:"monthly"`
				OriginalPlanID  int         `json:"original_plan_id"`
				Ssd             bool        `json:"ssd"`
				MinCPU          int         `json:"min_cpu"`
				MaxCPU          int         `json:"max_cpu"`
				MinRAM          int         `json:"min_ram"`
				MaxRAM          int         `json:"max_ram"`
				RAMStep         int         `json:"ram_step"`
				CheckRAM        int         `json:"check_ram"`
				PriceMultiplier float64     `json:"price_multiplier"`
				MaxIP           int         `json:"max_ip"`
				MinDisk         int         `json:"min_disk"`
				MaxDisk         int         `json:"max_disk"`
				DiskShare       float64     `json:"disk_share"`
				Stock           int         `json:"stock"`
				MaxOrders       int         `json:"max_orders"`
				Bandwidth       int         `json:"bandwidth"`
				Plan            string      `json:"plan"`
				WindowsPlan     int         `json:"windows_plan"`
				Virt            string      `json:"virt"`
				Upgradeable     int         `json:"upgradeable"`
				OfflineBilling  int         `json:"offline_billing"`
				Whitelist       string      `json:"whitelist"`
				PlanID          int         `json:"plan_id"`
				Contract        string      `json:"contract"`
				Node            interface{} `json:"node"`
				SsdEnabled      int         `json:"ssd_enabled"`
				NewClientsOnly  int         `json:"new_clients_only"`
				Joins           interface{} `json:"joins"`
				Hidden          int         `json:"hidden"`
				Notes           bool        `json:"notes"`
				Hook            interface{} `json:"hook"`
			} `json:"billing"`
			Ssd             int         `json:"ssd"`
			VirtType        string      `json:"virt_type"`
			Distro          string      `json:"distro"`
			Pvt             int         `json:"pvt"`
			Rescue          bool        `json:"rescue"`
			AdvancedStatsID int         `json:"advanced_stats_id"`
			Ipv4Count       int         `json:"ipv4_count"`
			Ipv6Count       int         `json:"ipv6_count"`
			TplID           interface{} `json:"tpl_id"`
			TplType         bool        `json:"tpl_type"`
			TplInfo         bool        `json:"tpl_info"`
			Nodename        string      `json:"nodename"`
			Bandwidth       struct {
				Used  int `json:"used"`
				Total int `json:"total"`
				Free  int `json:"free"`
				Usage int `json:"usage"`
			} `json:"bandwidth"`
			State struct {
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"state"`
			Vid   int `json:"vid"`
			Price struct {
				Due    float64 `json:"due"`
				Online struct {
					Monthly float64 `json:"monthly"`
					Hourly  float64 `json:"hourly"`
				} `json:"online"`
				Offline struct {
					Monthly float64 `json:"monthly"`
					Hourly  float64 `json:"hourly"`
				} `json:"offline"`
			} `json:"price"`
			OnlineHours  int    `json:"online_hours"`
			OfflineHours int    `json:"offline_hours"`
			Contract     string `json:"contract"`
			NextDue      string `json:"next_due"`
			Addons       struct {
				Billing bool          `json:"billing"`
				Items   []interface{} `json:"items"`
			} `json:"addons"`
			BackupSchedules bool `json:"backup_schedules"`
			BackupFiles     bool `json:"backup_files"`
		} `json:"instances"`
	} `json:"__data"`
}

type CloudCone struct {
	ApiKey  string
	ApiHash string
}

func NewCloudCone(apiKey string, apiHash string) *CloudCone {
	return &CloudCone{ApiKey: apiKey, ApiHash: apiHash}
}

func (v *CloudCone) GetInstances() (*InstancesResponse, error) {
	c := resty.New()
	resp, err := c.R().
		SetResult(&InstancesResponse{}).
		SetHeader("App-Secret", v.ApiKey).
		SetHeader("Hash", v.ApiHash).
		Get("https://api.cloudcone.com/api/v2/compute/instances")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*InstancesResponse)
		return result, nil
	} else {
		return nil, fmt.Errorf("cloudcone api error %d", resp.StatusCode())
	}
}

func (v *CloudCone) GetInstance(id int) (*InstanceResponse, error) {
	c := resty.New()
	resp, err := c.R().
		SetResult(&InstanceResponse{}).
		SetHeader("App-Secret", v.ApiKey).
		SetHeader("Hash", v.ApiHash).
		Get(fmt.Sprintf("https://api.cloudcone.com/api/v2/compute/%d/information", id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*InstanceResponse)
		return result, nil
	} else {
		return nil, fmt.Errorf("cloudcone api error %d", resp.StatusCode())
	}
}
