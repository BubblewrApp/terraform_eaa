package client

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrAppIdpMembershipGet       = errors.New("unable to get app idp membership")
	ErrAppDirectoryMembershipGet = errors.New("unable to get app directory membership")
	ErrAppGroupMembershipGet     = errors.New("unable to get app group membership")
)

type IDPMembership struct {
	IDPUUIDURL string `json:"idp_uuid_url"`
	Name       string `json:"name"`
}

type AppMembership struct {
	AppUUIDURL string `json:"app_uuid_url"`
	Name       string `json:"name"`
}

type AppIdpMembership struct {
	App         AppMembership `json:"app"`
	EnableMFA   string        `json:"enable_mfa"`
	IDP         IDPMembership `json:"idp"`
	Resource    string        `json:"resource"`
	ResourceURI struct {
		Apps string `json:"apps"`
		Href string `json:"href"`
		IDP  string `json:"idp"`
	} `json:"resource_uri"`
	UUIDURL string `json:"uuid_url"`
}

type AppIdpMembershipResponse struct {
	Meta              Meta               `json:"meta"`
	AppIdpMemberships []AppIdpMembership `json:"objects"`
}

func (app *Application) GetAppIdpMembership(ec *EaaClient) (*AppIdpMembership, error) {
	apiURL := fmt.Sprintf("%s://%s/%s/%s/idp_membership", URL_SCHEME, ec.Host, APPS_URL, app.UUIDURL)
	appidpMembershipResponse := AppIdpMembershipResponse{}

	getResp, err := ec.SendAPIRequest(apiURL, "GET", nil, &appidpMembershipResponse, false)
	if err != nil {
		return nil, err
	}
	if !(getResp.StatusCode >= http.StatusOK && getResp.StatusCode < http.StatusMultipleChoices) {
		desc, _ := FormatErrorResponse(getResp)
		appIdpErrMsg := fmt.Errorf("%w: %s", ErrAppIdpMembershipGet, desc)
		return nil, appIdpErrMsg
	}
	if len(appidpMembershipResponse.AppIdpMemberships) > 0 {
		appIdpmem := appidpMembershipResponse.AppIdpMemberships[0]
		return &appIdpmem, nil
	}
	return nil, ErrAppIdpMembershipGet
}

type DirectoryMembership struct {
	DirectoryUUIDURL string `json:"directory_uuid_url"`
	Name             string `json:"name"`
}

type AppDirectoryMembership struct {
	App       AppMembership       `json:"app"`
	Directory DirectoryMembership `json:"directory"`
	EnableMFA string              `json:"enable_mfa"`
	Resource  string              `json:"resource"`
	UUIDURL   string              `json:"uuid_url"`
}

type AppDirectoryMembershipResponse struct {
	Meta                    Meta                     `json:"meta"`
	AppDirectoryMemberships []AppDirectoryMembership `json:"objects"`
}

func (app *Application) GetAppDirectoryMembership(ec *EaaClient) ([]AppDirectoryMembership, error) {
	apiURL := fmt.Sprintf("%s://%s/%s/%s/directories_membership", URL_SCHEME, ec.Host, APPS_URL, app.UUIDURL)
	appdirectoryMembershipResponse := AppDirectoryMembershipResponse{}

	getResp, err := ec.SendAPIRequest(apiURL, "GET", nil, &appdirectoryMembershipResponse, false)
	if err != nil {
		return nil, err
	}
	if !(getResp.StatusCode >= http.StatusOK && getResp.StatusCode < http.StatusMultipleChoices) {
		desc, _ := FormatErrorResponse(getResp)
		appDirErrMsg := fmt.Errorf("%w: %s", ErrAppDirectoryMembershipGet, desc)
		return nil, appDirErrMsg
	}
	if len(appdirectoryMembershipResponse.AppDirectoryMemberships) >= 0 {
		return appdirectoryMembershipResponse.AppDirectoryMemberships, nil
	}
	return nil, ErrAppDirectoryMembershipGet
}

type GroupMembership struct {
	DirName      string `json:"dir_name"`
	DirUUIDURL   string `json:"dir_uuid_url"`
	GroupUUIDURL string `json:"group_uuid_url"`
	GroupName    string `json:"name"`
}

type AppGroupMembership struct {
	App       AppMembership   `json:"app"`
	EnableMFA string          `json:"enable_mfa"`
	Group     GroupMembership `json:"group"`
	UUIDURL   string          `json:"uuid_url"`
}

type AppGroupMembershipResponse struct {
	Meta                Meta                 `json:"meta"`
	AppGroupMemberships []AppGroupMembership `json:"objects"`
}

func (app *Application) GetAppGroupMembership(ec *EaaClient) ([]AppGroupMembership, error) {
	apiURL := fmt.Sprintf("%s://%s/%s/%s/groups", URL_SCHEME, ec.Host, APPS_URL, app.UUIDURL)
	appgroupMembershipResponse := AppGroupMembershipResponse{}

	getResp, err := ec.SendAPIRequest(apiURL, "GET", nil, &appgroupMembershipResponse, false)
	if err != nil {
		return nil, err
	}
	if !(getResp.StatusCode >= http.StatusOK && getResp.StatusCode < http.StatusMultipleChoices) {
		desc, _ := FormatErrorResponse(getResp)
		appGrpErrMsg := fmt.Errorf("%w: %s", ErrAppGroupMembershipGet, desc)
		return nil, appGrpErrMsg
	}
	if len(appgroupMembershipResponse.AppGroupMemberships) >= 0 {
		return appgroupMembershipResponse.AppGroupMemberships, nil
	}
	return nil, ErrAppGroupMembershipGet
}

func (app *Application) CreateAppAuthenticationStruct(ec *EaaClient) ([]interface{}, error) {
	appAuth := make(map[string]interface{})

	// Get the data from the auth membership functions
	appIDPMembership, err := app.GetAppIdpMembership(ec)
	if err != nil {
		return nil, err
	}
	appAuth["app_idp"] = appIDPMembership.IDP.Name
	appDirectoryMemberships, err := app.GetAppDirectoryMembership(ec)
	if err != nil {
		return nil, err
	}

	appGroupMemberships, err := app.GetAppGroupMembership(ec)
	if err != nil {
		return nil, err
	}

	directories := make(map[string]map[string]interface{})
	for _, dir := range appDirectoryMemberships {
		dirName := dir.Directory.Name
		directories[dirName] = make(map[string]interface{})
		directories[dirName]["name"] = dirName
	}

	for _, group := range appGroupMemberships {
		dirName := group.Group.DirName
		groupName := group.Group.GroupName

		if dir, exists := directories[dirName]; exists {
			// If "app_groups" key is not present, add it before appending groups
			if _, hasGroups := dir["app_groups"]; !hasGroups {
				dir["app_groups"] = make([]map[string]interface{}, 0)
			}
			groupInfo := make(map[string]interface{})
			groupInfo["name"] = groupName
			dir["app_groups"] = append(dir["app_groups"].([]map[string]interface{}), groupInfo)
		}
	}

	directoriesData := make([]map[string]interface{}, 0)
	for _, dir := range directories {
		appGroups, ok := dir["app_groups"].([]map[string]interface{})
		if ok && len(appGroups) > 0 {
			directoriesData = append(directoriesData, dir)
		}
	}
	// add "app_directories" key if the list is not empty
	if len(directoriesData) > 0 {
		appAuth["app_directories"] = directoriesData
	}
	return []interface{}{appAuth}, nil

}
