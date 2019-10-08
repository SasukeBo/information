package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// ListPrivilege get list of privilege
func ListPrivilege(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("privilege")

	if privType := params.Args["privType"]; privType != nil {
		qs = qs.Filter("priv_type", privType)
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	var privs []*models.Privilege
	if _, err := qs.All(&privs); err != nil {
		return nil, models.Error{Message: "list privilege failed.", OriErr: err}
	}

	return privs, nil
}

// LoadPrivilege _
func LoadPrivilege(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.RolePriv:
		return v.LoadPrivilege()
	case *models.RolePriv:
		return v.LoadPrivilege()
	default:
		return nil, models.Error{Message: "load related privilege failed."}
	}
}
