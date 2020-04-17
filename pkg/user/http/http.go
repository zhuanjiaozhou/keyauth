package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	api = &handler{}
)

type handler struct {
	service user.Service
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	prmaryRouter := router.ResourceRouter("primary_account")
	prmaryRouter.BasePath("users")
	prmaryRouter.Permission(true)
	prmaryRouter.Handle("POST", "/", h.CreatePrimayAccount)
	prmaryRouter.Handle("DELETE", "/", h.DestroyPrimaryAccount)

	ramRouter := router.ResourceRouter("ram_account")
	ramRouter.Permission(true)
	ramRouter.BasePath("domains/:did/users")
	ramRouter.Handle("POST", "/", h.CreateSubAccount)
}

func (h *handler) Config() error {
	if pkg.User == nil {
		return errors.New("denpence user service is nil")
	}

	h.service = pkg.User
	return nil
}

func init() {
	pkg.RegistryHTTPV1("user", api)
}
