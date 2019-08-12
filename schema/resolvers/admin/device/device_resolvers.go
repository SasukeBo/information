package device

import (
	// "fmt"

	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/errors"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// Get is a gql resolver, get a device
func Get(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "device_r"); err != nil {
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
	if err := utils.ValidateAccess(&params, "device_r"); err != nil {
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

	if userName := params.Args["userName"]; userName != nil {
		userExtend := models.UserExtend{Name: userName.(string)}
		if err := userExtend.GetBy("name"); err != nil {
			return nil, err
		}
		if user, err := userExtend.LoadUser(); err != nil {
			return nil, err
		} else {
			qs = qs.Filter("user_id", user.ID)
		}
	}

	if userUUID := params.Args["userUUID"]; userUUID != nil {
		user := models.User{UUID: userUUID.(string)}
		if err := user.GetBy("uuid"); err != nil {
			return nil, err
		}

		qs = qs.Filter("user_id", user.ID)
	}

	if chargerName := params.Args["chargerName"]; chargerName != nil {
		r := models.Repo.Raw(`
		SELECT device_id FROM device_charge
		WHERE user_id = (
			SELECT u.id FROM public.user AS u INNER JOIN user_extend AS ue
			ON u.user_extend_id = ue.id WHERE ue.name = ?
		);`, chargerName)

		if qs = listIDInQS(qs, r); qs == nil {
			return []interface{}{}, nil
		}
	}

	if chargerUUID := params.Args["chargerUUID"]; chargerUUID != nil {
		r := models.Repo.Raw(`
		SELECT device_id FROM device_charge
		WHERE user_id = (
			SELECT u.id FROM public.user as u WHERE u.uuid = ?
		);`, chargerUUID)

		if qs = listIDInQS(qs, r); qs == nil {
			return []interface{}{}, nil
		}
	}

	var devices []*models.Device
	if _, err := qs.All(&devices); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Device",
			Message: "List() error",
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
	if err := utils.ValidateAccess(&params, "device_w"); err != nil {
		return nil, err
	}

	uuid := params.Args["uuid"].(string)
	device := models.Device{UUID: uuid}
	if err := device.GetBy("uuid"); err != nil {
		return nil, err
	}

	if name := params.Args["name"]; name != nil {
		device.Name = name.(string)
	}

	if mac := params.Args["mac"]; mac != nil {
		device.Mac = mac.(string)
	}

	if status := params.Args["status"]; status != nil {
		device.Status = status.(int)
	}

	if description := params.Args["description"]; description != nil {
		device.Description = description.(string)
	}

	if err := device.Update("name", "mac", "status", "description"); err != nil {
		return nil, err
	}

	return device, nil
}

// Delete _
func Delete(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "device_w"); err != nil {
		return nil, err
	}

	uuid := params.Args["uuid"].(string)

	device := models.Device{UUID: uuid}
	if err := device.DeleteByUUID(); err != nil {
		return nil, err
	}

	return "ok", nil
}
