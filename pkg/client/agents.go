package client

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrAgentsGet = errors.New("agents get failed")
)

type Connector struct {
	Name                  string  `json:"name,omitempty"`
	UUIDURL               string  `json:"uuid_url,omitempty"`
	ActivationCode        *string `json:"activation_code"`
	AgentInfraType        int     `json:"agent_infra_type"`
	AgentType             int     `json:"agent_type"`
	AgentVersion          *string `json:"agent_version"`
	CPU                   *string `json:"cpu"`
	DataService           bool    `json:"data_service"`
	DebugChannelPermitted bool    `json:"debug_channel_permitted"`
	Description           *string `json:"description"`
	DHCP                  string  `json:"dhcp"`
	DiskSize              *string `json:"disk_size"`
	DNSServer             *string `json:"dns_server"`
	DownAppsCount         int     `json:"down_apps_count"`
	DownDirCount          int     `json:"down_dir_count"`
	DownloadURL           *string `json:"download_url"`
	Gateway               *string `json:"gateway"`
	GeoLocation           *string `json:"geo_location"`
	Hostname              *string `json:"hostname"`
	IPAddr                *string `json:"ip_addr"`
	LastCheckin           *string `json:"last_checkin"`
	LoadStatus            *string `json:"load_status"`
	MAC                   *string `json:"mac"`
	ManualOverride        bool    `json:"manual_override"`
	OSUpgradesUpToDate    bool    `json:"os_upgrades_up_to_date"`
	OSVersion             *string `json:"os_version"`
	Package               int     `json:"package"`
	Policy                string  `json:"policy"`
	PrivateIP             *string `json:"private_ip"`
	PublicIP              *string `json:"public_ip"`
	RAMSize               *string `json:"ram_size"`
	Reach                 int     `json:"reach"`
	Region                *string `json:"region"`
	State                 int     `json:"state"`
	Status                int     `json:"status"`
	Subnet                *string `json:"subnet"`
	Timezone              *string `json:"tz"`
	UnificationStatus     int     `json:"unification_status"`
	UpAppsCount           int     `json:"up_apps_count"`
	UpDirCount            int     `json:"up_dir_count"`
	UUID                  string  `json:"uuid"`
}

type ConnectorResponse struct {
	Meta struct {
		Limit      int         `json:"limit,omitempty"`
		Next       interface{} `json:"next,omitempty"`
		Offset     int         `json:"offset,omitempty"`
		Previous   interface{} `json:"previous,omitempty"`
		TotalCount int         `json:"total_count,omitempty"`
	} `json:"meta,omitempty"`
	Connectors []Connector `json:"objects,omitempty"`
}

func GetAgents(ec *EaaClient) ([]Connector, error) {
	apiURL := fmt.Sprintf("%s://%s/%s", URL_SCHEME, ec.Host, AGENTS_URL)
	agentsResponse := ConnectorResponse{}

	getResp, err := ec.SendAPIRequest(apiURL, "GET", nil, &agentsResponse, false)
	if err != nil {
		return nil, err
	}
	if !(getResp.StatusCode >= http.StatusOK && getResp.StatusCode < http.StatusMultipleChoices) {
		desc, _ := FormatErrorResponse(getResp)
		updErrMsg := fmt.Errorf("%w: %s", ErrAgentsGet, desc)

		return nil, updErrMsg
	}

	var agents []Connector
	for _, conn := range agentsResponse.Connectors {
		if conn.Name == "" || conn.UUIDURL == "" {
			continue
		}
		agents = append(agents, conn)
	}

	return agents, nil
}

func GetAgentUUIDs(ec *EaaClient, agentNames []string) ([]string, error) {
	agents, err := GetAgents(ec)
	if err != nil {
		return nil, ErrAgentsGet
	}

	agentUUIDs := make([]string, 0)
	for _, agentName := range agentNames {
		for _, agentData := range agents {
			if agentName == agentData.Name {
				agentUUIDs = append(agentUUIDs, agentData.UUIDURL)
				break
			}
		}
	}

	return agentUUIDs, nil
}
