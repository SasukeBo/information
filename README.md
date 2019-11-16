# 设备资讯化

web服务器，设备资讯化管理系统

### 数据库设计

#### 用户表 users

> 用于存储基本账号信息

| 字段 | 类型 | 描述 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| name | string | 真实姓名 |
| phone | string | 手机号 |
| password | string | 密码 |
| avatar_url | string | 头像链接 |
| email | string | 邮箱 |
| role_id | int | 角色ID |
| status | int | 基础状态 |
| createdAt | datetime | 创建时间 |
| updatedAt | datetime | 更新时间 |

#### 角色表 roles

> 存储角色信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| role_name | string | 角色名称 |
| status | int | 基础状态 |
| is_admin | bool | 是否为管理员 |
| inserted_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 权限表 privileges

> 存储权限信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| name | string | 权限名称 |
| priv_type | int | 权限类型 |
| sign | string | 权限标识 |

#### 设备表 devices

> 存储设备信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| type | string | 设备类型 |
| name | string | 设备名称 |
| prod_speed | float | 理论最大生产速率 |
| address | string | 设备地址 |
| number | string | 设备编号 |
| remote_ip | string | 接入IP |
| token | string | 设备唯一token，用于加密传输内容 |
| status | int | 基础状态 |
| user_id | int | 注册人id |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |
| status_change_at | datetime | 状态变更时间 |

#### 用户登录记录表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| encrypted_passwd | string | 加密后的密码 |
| user_agent | string | 用户代理 |
| user_id | int | 用户id |
| remote_ip | string | 登录IP |
| session_id | string | session id |
| remembered | bool | 记住登录 |
| logout | bool | 是否已登出 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 设备状态变更日志表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| status | int | 设备运行状态 |
| device_id | int | 设备id |
| begin_at | datetime | 状态开始时间 |
| finish_at | datetime | 状态结束时间 |

#### 设备停机原因表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| device_id | int | 设备id |
| word_index | int | 约定第n个停机原因字 |
| bit_pos | int | 停机原因字约定的bit位 |
| content | string | 原因 |

#### 产品表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| name | string | 产品名称 |
| token | string | 产品Token |
| productor_contact | string | 生产负责人联系方式 |
| productor | string | 生产负责人名称 |
| register_id | int | 产品注册人ID |
| finish_time | datetime | 产品交期 |
| total | int | 指标生产总量 |
| order_num | string | 订单编号 |
| customer_contact | string | 订货人联系方式 |
| customer | string | 订货人名称 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 产品实例表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| qualified | bool | 是否合格 |
| device_product_ship_id | int | 设备产品关联关系id |
| created_at | datetime | 生产时间 |

#### 产品检测项表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| sign | string | 检测项标识 |
| product_id | int | 产品id |
| upper_limit | float | 检测值上限 |
| lower_limit | float | 检测值下限 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 产品检测值表

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| detect_item_id | int | 检测项id |
| product_ins_id | int | 产品实例id |
| value | float | 检测值 |

### 硬件需求预估

|项次|名称|型号|数量|备注|
|:--:|:--:|:--:|:--:|:--:|
| 1 | 数据库服务器 | 联想至强四核E3-1220v6 | 1台 | 型号暂定，配置根据实际效果弹性更改 |
| 2 | web服务器 | 联想至强四核E3-1220v6 | 1台 | 型号暂定，配置根据实际效果弹性更改 |
| 3 | 交换机 | 华为48口百兆交换机 | 7台 | |
| 4 | 无线路由器 | TP-LINK TL-WAR308 | 15台 | 实际数量需要根据连接状态做增删 |

### 计算公式

#### OEE

> OEE = Availability(可用率) * Performance(表现性) * Quality(质量指数)

- **Availability 可用性** = 运行时间 / 计划生产时间
- **Performance 表现性** = 实际产量 / (运行时间 * 理论最大生产速率)
- **Quality 质量指数** = 合格产品数量 / 总生产产品数
