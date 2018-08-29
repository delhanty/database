package auth

import "github.com/linkai-io/am/am"

// Authorizer interfaces between roles and policies to determine if a user is a member of a role allowed to access a resource
// note it does not determine authorization of organization data, that is done at the datastore access level.
type Authorizer interface {
	IsAllowed(subject, resource, action string) error
	IsUserAllowed(orgID, userID int, resource, action string) error
	GetRoles(orgID, userID int) ([]*am.Role, error)
}
