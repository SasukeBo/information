package device

import (
	// "fmt"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/resolvers"
	// "github.com/SasukeBo/information/utils"
)

// Get is a gql resolver, get a device
func Get(params graphql.ResolveParams) (interface{}, error) {
	currentUserUUID := params.Info.RootValue.(map[string]interface{})["currentUserUUID"].(string)
	user := models.User{UUID: currentUserUUID}
	if err := user.GetByUUID(); err != nil {
		return nil, err
	}

	if err := resolvers.ValidateUserPrivilege(&user, "device_r"); err != nil {
		return nil, err
	}

	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.GetByUUID(); err != nil {
		return nil, err
	}

	return device, nil
}

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	// TODO:
	return nil, nil
}

// Delete _
func Delete(params graphql.ResolveParams) (interface{}, error) {
	// TODO:
	return nil, nil
}

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	// TODO:
	return nil, nil
}
