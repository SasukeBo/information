package schema

/*
	管理员接口权限逻辑不同，只用于后台调用。
	禁止在前台中调用管理员接口。
*/

import (
	"log"

	"github.com/graphql-go/graphql"

	fields "github.com/SasukeBo/information/schema/fields/admin"
)

// AdminQueryRoot is query root
var AdminQueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Description: `
	管理员接口权限逻辑不同，只用于后台调用。
	禁止在前台中调用管理员接口。
	`,
	Fields: graphql.Fields{
		/* user */
		"adminUserList": fields.UserListField,
		/*    role    */
		"adminRoleGet":       fields.RoleGetField,
		"adminRoleGetByName": fields.RoleGetByNameField,
		"adminRoleList":      fields.RoleListField,

		/* device */
		"adminDeviceGet":  fields.DeviceGetField,
		"adminDeviceList": fields.DeviceListField,

		/* userLogin */
		"adminUserLoginList": fields.UserLoginListField,
	},
})

// AdminMutateRoot is mutation root
var AdminMutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Description: `
	管理员接口权限逻辑不同，只用于后台调用。
	禁止在前台中调用管理员接口。
	`,
	Fields: graphql.Fields{
		/*    role    */
		"adminRoleCreate": fields.RoleCreateField,
		"adminRoleUpdate": fields.RoleUpdateField,
		"adminRoleDelete": fields.RoleDeleteField,

		/* userExtend */
		"adminUserExtendUpdate": fields.UserExtendUpdateField,

		/* user */
		"adminUserDelete": fields.UserDeleteField,
		"adminUserUpdate": fields.UserUpdateField,

		/* rolePriv */
		"adminRolePrivCreate": fields.RolePrivCreateField,
		"adminRolePrivDelete": fields.RolePrivDeleteField,

		/* device */
		"adminDeviceUpdate": fields.DeviceUpdateField,
		"adminDeviceDelete": fields.DeviceDeleteField,
	},
})

// AdminSchema is graphql schema
var AdminSchema graphql.Schema

func init() {
	var err error
	AdminSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    AdminQueryRoot,
		Mutation: AdminMutateRoot,
	})

	if err != nil {
		log.Fatal("failed to create admin schema, err: ", err)
	}
}
