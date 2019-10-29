package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
	// "fmt"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"
)

// ReasonTypeLoadReasons 停机类型获取停机原因
func ReasonTypeLoadReasons(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

// ReasonLoadType 停机原因获取停机类型
func ReasonLoadType(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	parent, ok := params.Source.(models.DeviceStopReason)
	if !ok {
		return nil, models.Error{Message: "parent is not device_stop_reason."}
	}

	rt := parent.Type

	if err := o.Read(rt); err != nil {
		return nil, models.Error{Message: "reason_type not found.", OriErr: err}
	}

	return rt, nil
}

// UpdateStopReason 更新停机原因
func UpdateStopReason(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	id := params.Args["id"].(int)
	stopReason := models.DeviceStopReason{ID: id}
	if err := o.Read(&stopReason); err != nil {
		return nil, models.Error{Message: "device_stop_reason not found.", OriErr: err}
	}
	updates := []string{}

	if v := params.Args["content"]; v != nil {
		content := v.(string)
		if err := utils.ValidateStringEmpty(content, "content"); err != nil {
			return nil, err
		}
		updates = append(updates, "content")
		stopReason.Content = content
	}

	if v := params.Args["code"]; v != nil {
		code := v.(string)
		if err := utils.ValidateStringEmpty(code, "code"); err != nil {
			return nil, err
		}
		updates = append(updates, "code")
		stopReason.Code = code
	}

	if _, err := o.Update(&stopReason, updates...); err != nil {
		return nil, models.Error{Message: "update device_stop_reason failed.", OriErr: err}
	}

	return stopReason, nil
}

// CreateStopReason 创建停机原因
func CreateStopReason(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	content := params.Args["content"].(string)
	code := params.Args["code"].(string)
	typeID := params.Args["typeID"].(int)

	rt := models.ReasonType{ID: typeID}
	if err := o.Read(&rt); err != nil {
		return nil, models.Error{Message: "reason type not found.", OriErr: err}
	}

	reason := models.DeviceStopReason{Type: &rt, Content: content, Code: code}
	if _, err := o.Insert(&reason); err != nil {
		return nil, models.Error{Message: "insert stop reason failed.", OriErr: err}
	}

	return reason, nil
}

// CreateReasonType 创建停机类型
func CreateReasonType(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)
	reasonType := models.ReasonType{Name: name}
	o := orm.NewOrm()
	if _, err := o.Insert(&reasonType); err != nil {
		return nil, models.Error{Message: "create reason type failed.", OriErr: err}
	}

	return reasonType, nil
}

// DeleteReasonType 删除停机类型
func DeleteReasonType(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	reasonType := models.ReasonType{ID: id}
	o := orm.NewOrm()
	if _, err := o.Delete(&reasonType); err != nil {
		return nil, models.Error{Message: "delete reason type failed.", OriErr: err}
	}

	return "ok", nil
}

// UpdateReasonType 更新停机类型
func UpdateReasonType(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	name := params.Args["name"].(string)
	reasonType := models.ReasonType{ID: id, Name: name}
	o := orm.NewOrm()
	if _, err := o.Update(&reasonType, "name"); err != nil {
		return nil, models.Error{Message: "update reason type failed.", OriErr: err}
	}

	return reasonType, nil
}
