# 设备资讯化

## 数据库设计

### 用户

#### 用户表 users

> 用于存储基本账号信息

| 字段 | 类型 | 描述 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| uuid | string | 通用唯一标识 |
| phone | string | 手机号 |
| password | string | 密码 |
| avatar_url | string | 头像链接 |
| role_id | int | 角色ID |
| user_extend_id | int | 用户拓展信息ID |
| status | int | 基础状态 |
| createdAt | datetime | 创建时间 |
| updatedAt | datetime | 更新时间 |

#### 用户信息 user_extends

> 用户存储拓展用户信息，例如

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| user_id | int | 用户ID |
| name | string | 真实姓名 |
| email | string | 邮箱 |

#### 第三方用户信息表 oauths

> 未来开发第三方登录可能需要

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| user_id | int | 用户ID |
| oauth_type | int | 第三方登录类型 |
| oauth_id | string | 第三方uid、openid等 |

### 权限

#### 角色表 roles

> 存储角色信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| role_name | string | 角色名称 |
| status | int | 基础状态 |
| inserted_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 权限表 privileges

> 存储权限信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | 唯一标识 |
| priv_name | string | 权限名称 |
| priv_type | int | 权限类型 |
| status | int | 基础状态 |
| inserted_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

**权限**
各项读写权限

#### 角色与权限关联表 role_abilities

> 存储角色与权限关联信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| role_id | int | 角色ID |
| priv_id | int | 权限ID |

### 设备

#### 设备表 devices

> 存储设备信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| uuid | string | 设备UUID |
| type | string | 设备类型 |
| name | string | 设备名称 |
| user_id | int | 注册人id |
| description | string | 设备描述 |
| mac | string | 设备mac地址 |
| token | string | 设备唯一token，用于加密传输内容 |
| status | int | 基础状态 |
| inserted_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 设备负责人表 device_charge

> 存储设备和负责人关系记录

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| user_id | int | 用户ID |
| device_id | int | 设备ID |
| inserted_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

#### 设备负责人权限表 device_charge_abilities

> 存储负责人权限记录

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pk |
| device_user_id | int | 设备负责人关系ID |
| priv_id | int | 权限ID |


#### 设备参数表 device_params

> 存储设备的参数内容

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| name | string | 参数名 |
| sign | string | 标识 |
| type | string | 值类型 |
| author_id | int | 创建人ID |
| inserted_at | datetime | 创建时间 |

#### 设备参数值表 device_param_values

> 存储设备参数的值

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| device_param_id | int | 设备参数ID |
| value | string | 值 |
| inserted_at | datetime | 插入时间 |


#### 设备运行状态表 device_status

> 存储设备运行状态变更记录

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| id | int | pkey |
| device_id | int | 设备id |
| status | int | 运行状态 |
| change_at | datetime | 状态变更时间 |
