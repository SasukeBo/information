package schema

import (
	"github.com/SasukeBo/information/schema/admintypes"
	"github.com/graphql-go/graphql"
	"log"
)

// AdminQueryRoot is query root
var AdminQueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/* user */
		// TODO: "adminUserList"

		/*    role    */
		"adminRoleGet":       admintypes.RoleGetType,
		"adminRoleGetByName": admintypes.RoleGetByNameType,
		// TODO: "roleList"
	},
})

// AdminMutateRoot is mutation root
var AdminMutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/* user */
		// TODO: "adminUserDelete"
		// TODO: "adminUserUpdate"

		/*    role    */
		"adminRoleCreate": admintypes.RoleCreateType,
		"adminRoleUpdate": admintypes.RoleUpdateType,
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
