package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*							   types
------------------------------------------ */

// DeviceCharger 设备负责人类型
var DeviceCharger = graphql.NewObject(graphql.ObjectConfig{
	Name:        "DeviceCharger",
	Description: "设备负责人类型",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"name":       &graphql.Field{Type: graphql.String, Description: "负责人姓名"},
		"phone":      &graphql.Field{Type: graphql.String, Description: "负责人手机号"},
		"department": &graphql.Field{Type: graphql.String, Description: "负责人部门名称"},
		"jobNumber":  &graphql.Field{Type: graphql.String, Description: "负责人工号"},
		"createdAt":  &graphql.Field{Type: graphql.DateTime},
		"updatedAt":  &graphql.Field{Type: graphql.DateTime},
	},
})

func init() {
	// circular references fixed by dynamically adding inside init(), see https://github.com/graphql-go/graphql/issues/164
	DeviceCharger.AddFieldConfig("device", &graphql.Field{Type: DeviceType, Description: "设备", Resolve: resolver.LoadDevice})
}

/*							query fields
------------------------------------------ */

var deviceChargerGetField = &graphql.Field{
	Type: DeviceCharger,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "设备负责人ID", false),
	},
	Description: "通过id获取设备负责人",
	Resolve:     resolver.GetDeviceCharger,
}

var deviceChargerListField = &graphql.Field{
	Type: graphql.NewList(DeviceCharger),
	Args: graphql.FieldConfigArgument{
		"deviceUUID": GenArg(graphql.String, "设备uuid", false),
	},
	Description: `
	查询本人负责的设备或创建的设备的设备负责人列表，
	可通过设备uuid指定某台设备，但必须是当前用户可访问的设备
	`,
	Resolve: resolver.ListDeviceCharger,
}

/*						mutation fields
------------------------------------------ */

var deviceChargerCreateField = &graphql.Field{
	Type: DeviceCharger,
	Args: graphql.FieldConfigArgument{
		"deivceUUID": GenArg(graphql.String, "设备UUID", false),
		"name":       GenArg(graphql.String, "负责人姓名", false),
		"phone":      GenArg(graphql.String, "手机号"),
		"department": GenArg(graphql.String, "部门名称"),
		"jobNumber":  GenArg(graphql.String, "工号"),
	},

	Resolve: resolver.CreateDeviceCharger,
}

var deviceChargerDeleteField = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": GenArg(graphql.Int, "设备负责人ID", false),
	},
	Resolve: resolver.DeleteDeviceCharger,
}

var deviceChargerUpdateField = &graphql.Field{
	Type: DeviceCharger,
	Args: graphql.FieldConfigArgument{
		"id":         GenArg(graphql.Int, "设备负责人ID", false),
		"name":       GenArg(graphql.String, "负责人姓名", false),
		"phone":      GenArg(graphql.String, "手机号"),
		"department": GenArg(graphql.String, "部门名称"),
		"jobNumber":  GenArg(graphql.String, "工号"),
	},
	Resolve: resolver.UpdateDeviceCharger,
}
