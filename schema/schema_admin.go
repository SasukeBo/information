package schema

import (
	"log"

	"github.com/graphql-go/graphql"

	fields "github.com/SasukeBo/information/schema/fields/admin"
)

// AdminQueryRoot is query root
var AdminQueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
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
