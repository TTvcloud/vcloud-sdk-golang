package iam

import (
	"encoding/json"
	"net/url"
)

// helper functions
func (p *Iam) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Iam) ListAccessKeys(query url.Values) (*AccessKeyListResp, int, error) {
	respBody, status, err := p.Query("ListAccessKeys", query)
	if err != nil {
		return nil, status, err
	}

	output := new(AccessKeyListResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *Iam) CreateAccessKey(query url.Values) (*AccessKeyResp, int, error) {
	respBody, status, err := p.Query("CreateAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(AccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *Iam) DeleteAccessKey(query url.Values) (*AccessKeyResp, int, error) {
	respBody, status, err := p.Query("DeleteAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(AccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *Iam) UpdateAccessKey(query url.Values) (*AccessKeyResp, int, error) {
	respBody, status, err := p.Query("UpdateAccessKey", query)
	if err != nil {
		return nil, status, err
	}
	output := new(AccessKeyResp)
	if err := json.Unmarshal(respBody, output); err != nil {
		return nil, status, err
	} else {
		output.ResponseMetadata.Service = "iam"
		return output, status, nil
	}
}

func (p *Iam) CreateService(query url.Values) (*ServiceResp, int, error) {
	resp := new(ServiceResp)
	statusCode, err := p.commonHandler("CreateService", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *Iam) ListAccessKeysForService(query url.Values) (*AccessKeyListResp, int, error) {
	resp := new(AccessKeyListResp)
	statusCode, err := p.commonHandler("ListAccessKeysForService", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *Iam) CreateAccessKeyForService(query url.Values) (*AccessKeyResp, int, error) {
	resp := new(AccessKeyResp)
	statusCode, err := p.commonHandler("CreateAccessKeyForService", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *Iam) DeleteAccessKeyForService(query url.Values) (*AccessKeyResp, int, error) {
	resp := new(AccessKeyResp)
	statusCode, err := p.commonHandler("DeleteAccessKeyForService", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

func (p *Iam) UpdateAccessKeyForService(query url.Values) (*AccessKeyResp, int, error) {
	resp := new(AccessKeyResp)
	statusCode, err := p.commonHandler("UpdateAccessKeyForService", query, resp)
	if err != nil {
		return nil, statusCode, err
	}

	return resp, statusCode, nil
}

// federation
func (p *Iam) AddAppIDToOAuthProvider(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AddAppIDToOAuthProvider", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) RemoveAppIDFromOAuthProvider(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("RemoveAppIDFromOAuthProvider", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) GetAppIDofOAuthProvider(query url.Values) (*AppResp, int, error) {
	resp := new(AppResp)
	statusCode, err := p.commonHandler("GetAppIDofOAuthProvider", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) UpdateAppIDName(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("UpdateAppIDName", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListAppIDsofOAuthProvider(query url.Values) (*AppListResp, int, error) {
	resp := new(AppListResp)
	statusCode, err := p.commonHandler("ListAppIDsofOAuthProvider", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListRolesAfterActorFilter(query url.Values) (*RoleListResp, int, error) {
	resp := new(RoleListResp)
	statusCode, err := p.commonHandler("ListRolesAfterActorFilter", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) UpdateActorFilter(query url.Values, body string) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	respBody, statusCode, err := p.Json("UpdateActorFilter", query, body)
	if err != nil {
		return nil, statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListIdentityProviders(query url.Values) (*IdentityProviderListResp, int, error) {
	resp := new(IdentityProviderListResp)
	statusCode, err := p.commonHandler("ListIdentityProviders", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// role
func (p *Iam) CreateRole(query url.Values) (*RoleResp, int, error) {
	resp := new(RoleResp)
	statusCode, err := p.commonHandler("CreateRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) GetRole(query url.Values) (*RoleResp, int, error) {
	resp := new(RoleResp)
	statusCode, err := p.commonHandler("GetRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) DeleteRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeleteRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListRoles(query url.Values) (*RoleListResp, int, error) {
	resp := new(RoleListResp)
	statusCode, err := p.commonHandler("ListRoles", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) UpdateRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("UpdateRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) AttachRolePolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AttachRolePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) DetachRolePolicy(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DetachRolePolicy", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListAttachedRolePolicies(query url.Values) (*AttachedPolicyListResp, int, error) {
	resp := new(AttachedPolicyListResp)
	statusCode, err := p.commonHandler("ListAttachedRolePolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListUsersForRole(query url.Values) (*AddedUserListResp, int, error) {
	resp := new(AddedUserListResp)
	statusCode, err := p.commonHandler("ListUsersForRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) AddIdpToRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AddIdpToRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) RemoveIDPFromRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("RemoveIDPFromRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListIDPsForRole(query url.Values) (*IdentityProviderListResp, int, error) {
	resp := new(IdentityProviderListResp)
	statusCode, err := p.commonHandler("ListIDPsForRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// user
func (p *Iam) CreateUser(query url.Values) (*UserResp, int, error) {
	resp := new(UserResp)
	statusCode, err := p.commonHandler("CreateUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) GetUser(query url.Values) (*UserResp, int, error) {
	resp := new(UserResp)
	statusCode, err := p.commonHandler("GetUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) DeleteUser(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("DeleteUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListUsers(query url.Values) (*UserListResp, int, error) {
	resp := new(UserListResp)
	statusCode, err := p.commonHandler("ListUsers", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) AddUserToRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("AddUserToRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) RemoveUserFromRole(query url.Values) (*NullResultResp, int, error) {
	resp := new(NullResultResp)
	statusCode, err := p.commonHandler("RemoveUserFromRole", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Iam) ListRolesForUser(query url.Values) (*RoleListResp, int, error) {
	resp := new(RoleListResp)
	statusCode, err := p.commonHandler("ListRolesForUser", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// policy
func (p *Iam) ListPolicies(query url.Values) (*PolicyListResp, int, error) {
	resp := new(PolicyListResp)
	statusCode, err := p.commonHandler("ListPolicies", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
