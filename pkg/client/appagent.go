package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type AppAgentResponse struct {
	Agents []struct {
		Agent struct {
			Name    string `json:"name,omitempty"`
			UUIDURL string `json:"uuid_url,omitempty"`
		} `json:"agent,omitempty"`
		ResourceURI struct {
			Href string `json:"href,omitempty"`
		} `json:"resource_uri,omitempty"`
	} `json:"objects,omitempty"`
}

var (
	ErrAgentsAssign = errors.New("agents assign failed")
)

type AssignAgents struct {
	AppId      string   `json:"app_id"`
	AgentNames []string `json:"agents"`
}

type Agent struct {
	UUIDURL string `json:"uuid_url"`
}

type AssignAgentsRequest struct {
	Agents []Agent `json:"agents"`
}

func (aar *AssignAgents) AssignAgents(ctx context.Context, ec *EaaClient) error {
	var agents AssignAgentsRequest
	agentUUIDs, err := GetAgentUUIDs(ec, aar.AgentNames)
	if err != nil {
		ec.Logger.Error("unable to lookup uuids from agent names")
		return err
	}
	for _, uuid := range agentUUIDs {
		agent := Agent{
			UUIDURL: uuid,
		}
		agents.Agents = append(agents.Agents, agent)
		ec.Logger.Info(uuid)
	}

	apiURL := fmt.Sprintf("%s://%s/%s/%s/agents", URL_SCHEME, ec.Host, APPS_URL, aar.AppId)
	ec.Logger.Info(apiURL)
	agentsResp, err := ec.SendAPIRequest(apiURL, "POST", agents, nil, false)
	if err != nil {
		ec.Logger.Error("assign agents failed StatusCode: ", agentsResp.StatusCode)
		return err
	}
	if !(agentsResp.StatusCode >= http.StatusOK && agentsResp.StatusCode < http.StatusMultipleChoices) {
		desc, _ := FormatErrorResponse(agentsResp)
		assignErrMsg := fmt.Errorf("%w: %s", ErrAgentsAssign, desc)
		ec.Logger.Error("assign agents failed StatusCode: desc: ", agentsResp.StatusCode, desc)
		return assignErrMsg
	}
	return nil
}

func (app *Application) GetAppAgents(ec *EaaClient) ([]string, error) {
	apiURL := fmt.Sprintf("%s://%s/%s/%s/agents", URL_SCHEME, ec.Host, APPS_URL, app.UUIDURL)
	agentsResponse := AppAgentResponse{}

	getResp, err := ec.SendAPIRequest(apiURL, "GET", nil, &agentsResponse, false)
	if err != nil {
		return nil, err
	}
	if !(getResp.StatusCode >= http.StatusOK && getResp.StatusCode < http.StatusMultipleChoices) {
		desc, _ := FormatErrorResponse(getResp)
		updErrMsg := fmt.Errorf("%w: %s", ErrAgentsGet, desc)

		return nil, updErrMsg
	}

	agentNames := make([]string, 0, len(agentsResponse.Agents))
	for _, agent := range agentsResponse.Agents {
		agentNames = append(agentNames, agent.Agent.Name)
	}

	return agentNames, nil
}
