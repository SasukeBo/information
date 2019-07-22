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
| avatar_url | string | 头像连接 |
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

#### 角色与权限关联表 role_abilities

> 存储角色与权限关联信息

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| role_id | int | 角色ID |
| priv_id | int | 权限ID |
