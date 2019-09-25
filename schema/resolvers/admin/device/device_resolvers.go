package device

import (
	// "fmt"

	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// Get is a gql resolver, get a device
func Get(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_device_r", models.PrivType.Admin); err != nil {
		return nil, err
	}

	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	return device, nil
}

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_device_r", models.PrivType.Admin); err != nil {
		return nil, err
	}

	qs := models.Repo.QueryTable("device").OrderBy("-created_at")

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	if dType := params.Args["type"]; dType != nil {
		qs = qs.Filter("type", dType)
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	if status := params.Args["status"]; status != nil {
		qs = qs.Filter("status", status)
	}

	if userUUID := params.Args["userUUID"]; userUUID != nil {
		user := models.User{UUID: userUUID.(string)}
		if err := user.GetBy("uuid"); err != nil {
			return nil, err
		}

		qs = qs.Filter("user_id", user.ID)
	}

	var devices []*models.Device
	if _, err := qs.All(&devices); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get device list error",
			OriErr:  err,
		}
	}

	return devices, nil
}

func listIDInQS(qs orm.QuerySeter, r orm.RawSeter) orm.QuerySeter {
	var lists []orm.ParamsList
	if _, err := r.ValuesList(&lists); err != nil {
		return nil
	}

	ids := utils.EnumMap(lists, func(item interface{}) interface{} {
		return item.(orm.ParamsList)[0]
	})

	if len(ids) > 0 {
		return qs.Filter("id__in", ids)
	}

	return nil
}

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_device_u", models.PrivType.Admin); err != nil {
		return nil, err
	}

	uuid := params.Args["uuid"].(string)
	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if name := params.Args["name"]; name != nil {
		if err := utils.ValidateStringEmpty(name.(string), "name"); err != nil {
			return nil, err
		}
		device.Name = name.(string)
	}

	if status := params.Args["status"]; status != nil {
		device.Status = status.(int)
	}

	if description := params.Args["description"]; description != nil {
		device.Description = description.(string)
	}

	if err := device.Update("name", "status", "description"); err != nil {
		return nil, err
	}

	return device, nil
}

// Delete _
func Delete(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_device_d", models.PrivType.Admin); err != nil {
		return nil, err
	}

	device := models.Device{UUID: params.Args["uuid"].(string)}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if err := device.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}
