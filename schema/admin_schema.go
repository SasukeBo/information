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
		/* user */
		// TODO: "adminUserList"

		/*    role    */
		"adminRoleGet":       fields.RoleGetField,
		"adminRoleGetByName": fields.RoleGetByNameField,
		// TODO: "roleList"
	},
})

// AdminMutateRoot is mutation root
var AdminMutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/* user */
		// TODO: "adminUserDelete"
		"adminUserUpdate": fields.UserUpdateField,

		/*    role    */
		"adminRoleCreate": fields.RoleCreateField,
		"adminRoleUpdate": fields.RoleUpdateField,
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
