package endpoint

import (
	"github.com/go-godin/permission-service/internal/permission"
)

type SetRequest struct {
	Identifier string
	Permission permission.Permission
}
type SetResponse struct {
	Err error
}

type GetRequest struct {
	Identifier string
}
type GetResponse struct {
	Permissions []permission.Permission
	Err         error
}

type HasRequest struct {
	Identifier string
	Permission permission.Permission
}
type HasResponse struct {
	HasPermission bool
	Err           error
}

func (r GetResponse) Failed() error { return r.Err }
func (r SetResponse) Failed() error { return r.Err }
func (r HasResponse) Failed() error { return r.Err }
