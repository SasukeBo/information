package schemaadmin

/*
	管理员接口权限逻辑不同，只用于后台调用。
	禁止在前台中调用管理员接口。
*/

/*
import (
	"github.com/graphql-go/graphql"
	"log"
)

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Description: `
	管理员接口权限逻辑不同，只用于后台调用。
	禁止在前台中调用管理员接口。
	`,
	Fields: graphql.Fields{
		"adminUserList": UserListField,
		"adminRoleGet":       RoleGetField,
		"adminRoleGetByName": RoleGetByNameField,
		"adminRoleList":      RoleListField,
		"adminDeviceGet":  DeviceGetField,
		"adminDeviceList": DeviceListField,
		"adminUserLoginList": UserLoginListField,
	},
})

// MutationRoot is mutation root
var MutationRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Description: `
	管理员接口权限逻辑不同，只用于后台调用。
	禁止在前台中调用管理员接口。
	`,
	Fields: graphql.Fields{
		"adminRoleCreate": RoleCreateField,
		"adminRoleUpdate": RoleUpdateField,
		"adminRoleDelete": RoleDeleteField,
		"adminUserDelete": UserDeleteField,
		"adminUserUpdate": UserUpdateField,
		"adminRolePrivCreate": RolePrivCreateField,
		"adminRolePrivDelete": RolePrivDeleteField,
		"adminDeviceUpdate": DeviceUpdateField,
		"adminDeviceDelete": DeviceDeleteField,
	},
})

// Root is graphql admin schema
var Root graphql.Schema

func init() {
	var err error
	Root, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryRoot,
		Mutation: MutationRoot,
	})

	if err != nil {
		log.Fatal("failed to create admin schema, err: ", err)
	}
}
*/
