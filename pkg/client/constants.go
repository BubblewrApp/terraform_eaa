package client

import (
	"errors"
)

const (
	STR_TRUE  = "true"
	STR_FALSE = "false"
)

const (
	MGMT_POP_URL       = "crux/v1/mgmt-pop"
	APPS_URL           = "crux/v1/mgmt-pop/apps"
	POPS_URL           = "crux/v1/mgmt-pop/pops"
	APPIDP_URL         = "crux/v1/mgmt-pop/appidp"
	APPDIRECTORIES_URL = "crux/v1/mgmt-pop/appdirectories"
	APPGROUPS_URL      = "crux/v1/mgmt-pop/appgroups"
	AGENTS_URL         = "crux/v1/mgmt-pop/agents"
	APP_CATEGORIES_URL = "crux/v1/mgmt-pop/appcategories"
	IDP_URL            = "crux/v1/mgmt-pop/idp"
	URL_SCHEME         = "https"
)

var (
	ErrInvalidArgument = errors.New("invalid arguments provided")
	ErrMarshaling      = errors.New("marshaling input")
	ErrUnmarshaling    = errors.New("unmarshaling output")

	ErrAppCreate = errors.New("app creation failed")
	ErrAppUpdate = errors.New("app update failed")
	ErrAppDelete = errors.New("app delete failed")

	ErrAssignAgentsFailure    = errors.New("assigning agents to the app failed")
	ErrAssignIdpFailure       = errors.New("assigning IDP to the app failed")
	ErrAssignDirectoryFailure = errors.New("assigning directory to the app failed")
	ErrDeploy                 = errors.New("app deploy failed")
	ErrAssignGroupFailure     = errors.New("assigning groups to the app failed")
	ErrGetApp                 = errors.New("app deploy failed")

	ErrInvalidType  = errors.New("value must be of the specified type")
	ErrInvalidValue = errors.New("invalid value for a key")
)

type Domain string

const (
	AppDomainCustom Domain = "custom"
	AppDomainWapp   Domain = "wapp"
)

func (d Domain) ToInt() (int, error) {
	switch d {
	case AppDomainCustom:
		return int(APP_DOMAIN_CUSTOM), nil
	case AppDomainWapp:
		return int(APP_DOMAIN_WAPP), nil
	default:
		return 0, errors.New("Unknown domain value")
	}
}

type DomainInt int

const (
	APP_DOMAIN_CUSTOM DomainInt = 1 + iota
	APP_DOMAIN_WAPP
)

func (cam DomainInt) String() (string, error) {
	switch cam {
	case APP_DOMAIN_CUSTOM:
		return string(AppDomainCustom), nil
	case APP_DOMAIN_WAPP:
		return string(AppDomainWapp), nil
	default:
		return "", errors.New("Unknown domain value")
	}
}

type AppProfile string

const (
	AppProfileHTTP       AppProfile = "http"
	AppProfileSharePoint AppProfile = "sharepoint"
	AppProfileJira       AppProfile = "jira"
	AppProfileRDP        AppProfile = "rdp"
	AppProfileVNC        AppProfile = "vnc"
	AppProfileSSH        AppProfile = "ssh"
	AppProfileJenkins    AppProfile = "jenkins"
	AppProfileConfluence AppProfile = "confluence"
	AppProfileTCP        AppProfile = "tcp"
)

func (ap AppProfile) ToInt() (int, error) {
	switch ap {
	case AppProfileHTTP:
		return int(APP_PROFILE_HTTP), nil
	case AppProfileSharePoint:
		return int(APP_PROFILE_SHAREPOINT), nil
	case AppProfileJira:
		return int(APP_PROFILE_JIRA), nil
	case AppProfileRDP:
		return int(APP_PROFILE_RDP), nil
	case AppProfileVNC:
		return int(APP_PROFILE_VNC), nil
	case AppProfileSSH:
		return int(APP_PROFILE_SSH), nil
	case AppProfileJenkins:
		return int(APP_PROFILE_JENKINS), nil
	case AppProfileConfluence:
		return int(APP_PROFILE_CONFLUENCE), nil
	case AppProfileTCP:
		return int(APP_PROFILE_TCP), nil
	default:
		return 0, errors.New("Unknown App_Profile value")
	}
}

type AppProfileInt int

const (
	APP_PROFILE_HTTP AppProfileInt = 1 + iota
	APP_PROFILE_SHAREPOINT
	APP_PROFILE_JIRA
	APP_PROFILE_RDP
	APP_PROFILE_VNC
	APP_PROFILE_SSH
	APP_PROFILE_JENKINS
	APP_PROFILE_CONFLUENCE
	APP_PROFILE_TCP
)

func (cam AppProfileInt) String() (string, error) {
	switch cam {
	case APP_PROFILE_HTTP:
		return string(AppProfileHTTP), nil
	case APP_PROFILE_SHAREPOINT:
		return string(AppProfileSharePoint), nil
	case APP_PROFILE_JIRA:
		return string(AppProfileJira), nil
	case APP_PROFILE_RDP:
		return string(AppProfileRDP), nil
	case APP_PROFILE_VNC:
		return string(AppProfileVNC), nil
	case APP_PROFILE_SSH:
		return string(AppProfileSSH), nil
	case APP_PROFILE_JENKINS:
		return string(AppProfileJenkins), nil
	case APP_PROFILE_CONFLUENCE:
		return string(AppProfileConfluence), nil
	case APP_PROFILE_TCP:
		return string(AppProfileTCP), nil
	default:
		return "", errors.New("Unknown app_profile value")
	}
}

type ClientAppMode string

const (
	ClientAppModeTCP    ClientAppMode = "tcp"
	ClientAppModeTunnel ClientAppMode = "tunnel"
)

func (cam ClientAppMode) ToInt() (int, error) {
	switch cam {
	case ClientAppModeTCP:
		return int(CLIENT_APP_MODE_TCP), nil
	case ClientAppModeTunnel:
		return int(CLIENT_APP_MODE_TUNNEL), nil
	default:
		return 0, errors.New("Unknown ClientAppMode value")
	}
}

type ClientAppModeInt int

const (
	CLIENT_APP_MODE_TCP ClientAppModeInt = 1 + iota
	CLIENT_APP_MODE_TUNNEL
)

func (cam ClientAppModeInt) String() (string, error) {
	switch cam {
	case CLIENT_APP_MODE_TCP:
		return string(ClientAppModeTCP), nil
	case CLIENT_APP_MODE_TUNNEL:
		return string(ClientAppModeTunnel), nil
	default:
		return "", errors.New("Unknown ClientAppMode value")
	}
}

type ClientAppType string

const (
	ClientAppTypeEnterprise ClientAppType = "enterprise"
	ClientAppTypeSaaS       ClientAppType = "saas"
	ClientAppTypeBookmark   ClientAppType = "bookmark"
	ClientAppTypeTunnel     ClientAppType = "tunnel"
)

func (cat ClientAppType) ToInt() (int, error) {
	switch cat {
	case ClientAppTypeEnterprise:
		return int(APP_TYPE_ENTERPRISE_HOSTED), nil
	case ClientAppTypeSaaS:
		return int(APP_TYPE_SAAS), nil
	case ClientAppTypeBookmark:
		return int(APP_TYPE_BOOKMARK), nil
	case ClientAppTypeTunnel:
		return int(APP_TYPE_TUNNEL), nil
	default:
		return 0, errors.New("Unknown ClientAppType value")
	}
}

type ClientAppTypeInt int

const (
	APP_TYPE_ENTERPRISE_HOSTED ClientAppTypeInt = 1 + iota
	APP_TYPE_SAAS
	APP_TYPE_BOOKMARK
	APP_TYPE_TUNNEL
)

func (cat ClientAppTypeInt) String() (string, error) {
	switch cat {
	case APP_TYPE_ENTERPRISE_HOSTED:
		return string(ClientAppTypeEnterprise), nil
	case APP_TYPE_SAAS:
		return string(ClientAppTypeSaaS), nil
	case APP_TYPE_BOOKMARK:
		return string(ClientAppTypeBookmark), nil
	case APP_TYPE_TUNNEL:
		return string(ClientAppTypeTunnel), nil
	default:
		return "", errors.New("Unknown ClientAppType value")
	}
}
