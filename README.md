# 设备资讯化

## 数据库设计

### 用户表 user

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| UUID | string | 唯一标识 |
| account | string | 账号 |
| password | string | 密码 |
| status | int | 基础状态 |
| createdAt | datetime | 创建时间 |
| updatedAt | datetime | 更新时间 |

### 用户信息 user_profile

| 字段 | 类型 | 说明 |
|:----:|:----:|:----:|
| UUID | string | 唯一标识 |
| realname | string | 真实姓名 |
| phone | string | 手机号 |
| email | string | 邮箱 |
