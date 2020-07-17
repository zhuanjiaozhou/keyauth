package http

import (
	"net/http"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/provider"
	"github.com/infraboard/keyauth/pkg/provider/ldap"
	"github.com/infraboard/keyauth/pkg/user/types"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	page := request.NewPageRequestFromHTTP(r)
	req := provider.NewQueryLDAPConfigRequest(page)
	req.WithToken(tk)

	apps, err := h.service.QueryConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, apps)
	return
}

// CreateApplication 创建主账号
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := ldap.NewDefaultConfig()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	if !tk.UserType.Is(types.SupperAccount, types.PrimaryAccount) {
		response.Failed(w, exception.NewPermissionDeny("只有域管理员可以设置域的LDAP"))
		return
	}

	d, err := h.service.SaveConfig(tk, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	req := provider.NewDescribeLDAPConfigWithID(rctx.PS.ByName("id"))
	d, err := h.service.DescribeConfig(req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
	return
}
