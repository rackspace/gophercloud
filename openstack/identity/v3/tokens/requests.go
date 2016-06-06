package tokens

import (
	"net/http"

	"github.com/rackspace/gophercloud"
)

// Scope allows a created token to be limited to a specific domain or project.
type Scope struct {
	ProjectID   string
	ProjectName string
	DomainID    string
	DomainName  string
}

func subjectTokenHeaders(c *gophercloud.ServiceClient, subjectToken string) map[string]string {
	return map[string]string{
		"X-Subject-Token": subjectToken,
	}
}

// Create authenticates and either generates a new token, or changes the Scope of an existing token.
func Create(c *gophercloud.ServiceClient, options gophercloud.AuthOptions, scope *Scope) CreateResult {
	type domainReq struct {
		ID   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}

	type projectReq struct {
		Domain *domainReq `json:"domain,omitempty"`
		Name   *string    `json:"name,omitempty"`
		ID     *string    `json:"id,omitempty"`
	}

	type userReq struct {
		ID       *string    `json:"id,omitempty"`
		Name     *string    `json:"name,omitempty"`
		Password string     `json:"password"`
		Domain   *domainReq `json:"domain,omitempty"`
	}

	type passwordReq struct {
		User userReq `json:"user"`
	}

	type tokenReq struct {
		ID string `json:"id"`
	}

	type identityReq struct {
		Methods  []string     `json:"methods"`
		Password *passwordReq `json:"password,omitempty"`
		Token    *tokenReq    `json:"token,omitempty"`
	}

	type scopeReq struct {
		Domain  *domainReq  `json:"domain,omitempty"`
		Project *projectReq `json:"project,omitempty"`
	}

	type authReq struct {
		Identity identityReq `json:"identity"`
		Scope    *scopeReq   `json:"scope,omitempty"`
	}

	type request struct {
		Auth authReq `json:"auth"`
	}

	// Populate the request structure based on the provided arguments. Create and return an error
	// if insufficient or incompatible information is present.
	var req request

	// Test first for unrecognized arguments.
	if options.APIKey != "" {
		return createErr(ErrAPIKeyProvided)
	}
	if options.TenantID != "" {
		return createErr(ErrTenantIDProvided)
	}
	if options.TenantName != "" {
		return createErr(ErrTenantNameProvided)
	}

	if options.Password == "" {
		if c.TokenID != "" {
			// Because we aren't using password authentication, it's an error to also provide any of the user-based authentication
			// parameters.
			switch {
			case options.Username != "":
				return createErr(ErrUsernameWithToken)
			case options.UserID != "":
				return createErr(ErrUserIDWithToken)
			case options.DomainID != "":
				return createErr(ErrDomainIDWithToken)
			case options.DomainName != "":
				return createErr(ErrDomainNameWithToken)
			default:
				// Configure the request for Token authentication.
				req.Auth.Identity.Methods = []string{"token"}
				req.Auth.Identity.Token = &tokenReq{ID: c.TokenID}
			}
		} else {
			// If no password or token ID are available, authentication can't continue.
			return createErr(ErrMissingPassword)
		}
	} else {
		// Password authentication.
		req.Auth.Identity.Methods = []string{"password"}

		if options.UserID != "" && options.Username != "" {
			// At least one of Username and UserID must be specified.
			return createErr(ErrUsernameOrUserID)
		} else if options.Username != "" {
			switch {
			case options.DomainID != "" && options.DomainName != "":
				return createErr(ErrDomainIDOrDomainName)
			// Configure the request for Username and Password authentication with a DomainID.
			case options.DomainID != "":
				req.Auth.Identity.Password = &passwordReq{
					User: userReq{
						Name:     &options.Username,
						Password: options.Password,
						Domain:   &domainReq{ID: &options.DomainID},
					},
				}
			// Configure the request for Username and Password authentication with a DomainName.
			case options.DomainName != "":
				req.Auth.Identity.Password = &passwordReq{
					User: userReq{
						Name:     &options.Username,
						Password: options.Password,
						Domain:   &domainReq{Name: &options.DomainName},
					},
				}
			// Configure the request for Username and Password authentication with a DefaultDomain.
			case options.DefaultDomain != "":
				req.Auth.Identity.Password = &passwordReq{
					User: userReq{
						Name:     &options.Username,
						Password: options.Password,
						Domain:   &domainReq{Name: &options.DefaultDomain},
					},
				}
			default:
				// Either DomainID or DomainName or DefaultDomain must also be specified.
				return createErr(ErrDomainIDOrDomainNameOrDefaultDomain)
			}
		} else if options.UserID != "" {
			// If UserID is specified, neither UserDomainID nor UserDomainName may be.
			// Note: keeping DomainID and DomainName to not break previous versions of GopherCloud
			switch {
			case options.DomainID != "" || options.UserDomainID != "":
				return createErr(ErrDomainIDWithUserID)
			case options.DomainName != "" || options.UserDomainName != "":
				return createErr(ErrDomainNameWithUserID)
			default:
				// Configure the request for UserID and Password authentication.
				req.Auth.Identity.Password = &passwordReq{
					User: userReq{ID: &options.UserID, Password: options.Password},
				}
			}
		} else {
			// At least one of Username and UserID must be specified.
			return createErr(ErrUsernameOrUserID)
		}
	}

	// Add a "scope" element if a Scope has been provided.
	if scope != nil {
		if scope.ProjectName != "" && scope.ProjectID != "" {
			// ProjectID may not be supplied.
			return createErr(ErrScopeProjectIDOrProjectName)
		} else if scope.ProjectName != "" {
			// Project scoping using the project name
			switch {
			case scope.DomainID != "":
				// ProjectName + DomainID
				req.Auth.Scope = &scopeReq{
					Project: &projectReq{
						Name:   &scope.ProjectName,
						Domain: &domainReq{ID: &scope.DomainID},
					},
				}
			case scope.DomainName != "":
				// ProjectName + DomainName
				req.Auth.Scope = &scopeReq{
					Project: &projectReq{
						Name:   &scope.ProjectName,
						Domain: &domainReq{Name: &scope.DomainName},
					},
				}
			default:
				// Either DomainID or DomainName must be supplied.
				return createErr(ErrScopeDomainIDOrDomainName)
			}
		} else if scope.ProjectID != "" {
			// Project scoping using the project id
			// ProjectID provided. ProjectName, DomainID, and DomainName may not be provided.
			switch {
			case scope.DomainID != "":
				return createErr(ErrScopeProjectIDAlone)
			case scope.DomainName != "":
				return createErr(ErrScopeProjectIDAlone)
			default:
				// ProjectID
				req.Auth.Scope = &scopeReq{
					Project: &projectReq{ID: &scope.ProjectID},
				}
			}
		} else if scope.DomainID != "" {
			// Domain scoping using the domain id
			// DomainID provided. ProjectID, ProjectName, and DomainName may not be provided.
			if scope.DomainName != "" {
				return createErr(ErrScopeDomainIDOrDomainName)
			} else {
				// DomainID scope
				req.Auth.Scope = &scopeReq{
					Domain: &domainReq{ID: &scope.DomainID},
				}
			}
		} else if scope.DomainName != "" {
			// Domain scoping using the domain name
			// DomainName scope
			req.Auth.Scope = &scopeReq{
				Domain: &domainReq{Name: &scope.DomainName},
			}
		} else {
			return createErr(ErrScopeEmpty)
		}
	}

	var result CreateResult
	var response *http.Response
	response, result.Err = c.Post(tokenURL(c), req, &result.Body, nil)
	if result.Err != nil {
		return result
	}
	result.Header = response.Header
	return result
}

// Get validates and retrieves information about another token.
func Get(c *gophercloud.ServiceClient, token string) GetResult {
	var result GetResult
	var response *http.Response
	response, result.Err = c.Get(tokenURL(c), &result.Body, &gophercloud.RequestOpts{
		MoreHeaders: subjectTokenHeaders(c, token),
		OkCodes:     []int{200, 203},
	})
	if result.Err != nil {
		return result
	}
	result.Header = response.Header
	return result
}

// Validate determines if a specified token is valid or not.
func Validate(c *gophercloud.ServiceClient, token string) (bool, error) {
	response, err := c.Request("HEAD", tokenURL(c), gophercloud.RequestOpts{
		MoreHeaders: subjectTokenHeaders(c, token),
		OkCodes:     []int{204, 404},
	})
	if err != nil {
		return false, err
	}

	return response.StatusCode == 204, nil
}

// Revoke immediately makes specified token invalid.
func Revoke(c *gophercloud.ServiceClient, token string) RevokeResult {
	var res RevokeResult
	_, res.Err = c.Delete(tokenURL(c), &gophercloud.RequestOpts{
		MoreHeaders: subjectTokenHeaders(c, token),
	})
	return res
}
