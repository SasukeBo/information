package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// GetLogStopReasons 获取某条日志的停机原因
func GetLogStopReasons(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	logID := params.Args["logID"].(int)
	dsl := models.DeviceStatusLog{ID: logID}
	if err := o.Read(&dsl); err != nil {
		return nil, models.Error{Message: "log not found.", OriErr: err}
	}

	sql := `
	SELECT sr.id, sr.bit_pos, sr.content, sr.device_id, sr.word_index FROM stop_reason sr
	JOIN device_status_log_stop_reasons ship ON sr.id = ship.stop_reason_id
	WHERE ship.device_status_log_id = ?
	`
	var reasons []*models.StopReason
	if _, err := o.Raw(sql, logID).QueryRows(&reasons); err != nil {
		return nil, models.Error{Message: "get log stop_reasons failed.", OriErr: err}
	}

	return reasons, nil
}

// DeleteStopReason 删除停机原因
func DeleteStopReason(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	id := params.Args["id"].(int)
	sr := models.StopReason{ID: id}
	if _, err := o.Delete(&sr); err != nil {
		return nil, models.Error{Message: "delete stop_reason failed.", OriErr: err}
	}

	return "ok", nil
}

// UpdateStopReason 更新停机原因
func UpdateStopReason(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	id := params.Args["id"].(int)
	stopReason := models.StopReason{ID: id}
	if err := o.Read(&stopReason); err != nil {
		return nil, models.Error{Message: "device_stop_reason not found.", OriErr: err}
	}
	updates := []string{}

	if v, ok := params.Args["content"].(string); ok && v != "" {
		updates = append(updates, "content")
		stopReason.Content = v
	}

	if v := params.Args["wordIndex"]; v != nil {
		updates = append(updates, "word_index")
		stopReason.WordIndex = v.(int)
	}

	if v := params.Args["bitPos"]; v != nil {
		updates = append(updates, "bit_pos")
		stopReason.BitPos = v.(int)
	}

	if _, err := o.Update(&stopReason, updates...); err != nil {
		return nil, models.Error{Message: "update device_stop_reason failed.", OriErr: err}
	}

	return stopReason, nil
}

// CreateStopReason 创建停机原因
func CreateStopReason(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()

	deviceID := params.Args["deviceID"].(int)
	device := models.Device{ID: deviceID}
	if err := o.Read(&device); err != nil {
		return nil, models.Error{Message: "device not found.", OriErr: err}
	}

	content := params.Args["content"].(string)
	wordIndex := params.Args["wordIndex"].(int)
	bitPos := params.Args["bitPos"].(int)

	reason := models.StopReason{
		Device:    &device,
		Content:   content,
		WordIndex: wordIndex,
		BitPos:    bitPos,
	}

	if _, err := o.Insert(&reason); err != nil {
		return nil, models.Error{Message: "insert stop_reason failed.", OriErr: err}
	}

	return reason, nil
}
