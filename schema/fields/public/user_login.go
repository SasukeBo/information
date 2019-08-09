package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/userlogin"
	"github.com/SasukeBo/information/schema/types"
)

// UserLoginListField _
var UserLoginListField = &graphql.Field{
	Type: graphql.NewList(types.UserLogin),
	Args: graphql.FieldConfigArgument{
		"todo": fields.GenArg(graphql.String, "todo"),
	},
	Resolve: userlogin.List,
}
