import VueRouter from 'vue-router'
import tag from 'graphql-tag'
import store from '../vuex'

import { defaultRoutes } from './routes.js'
import { parseGQLError } from '../utils'

const router = new VueRouter({
  mode: 'history',
  routes: defaultRoutes()
})

router.beforeEach((to, from, next) => {
  var app = router.app

  if (!store.state.user.uuid) { // 没有用户信息
    app.$apollo.query({
      query: tag`query { currentUser { uuid phone status avatarURL userExtend { name email } role { roleName isAdmin } } }`
    }).then(({ data: { currentUser } }) => { // 获取成功
      app.$store.dispatch('user/setUserData', currentUser)

      if (isAuthPage(to)) {
        // 登录状态下如果是 auth 相关页面则导向 首页
        // 如果有 return_to 则导向 return_to
        var return_to = to.query.return_to
        return_to ? next({ name: return_to }) : next({ name: 'index' })
      } else {
        next()
      }
    }).catch((e) => { // 获取失败
      var error = parseGQLError(e)
      // 未登录状态下如果是 auth 相关页面则继续
      if (isAuthPage(to)) {
        next()
      } else {
        // 否则前往登录页，加上 return_to
        app.$message({ message: error.message, type: 'warning' })
        next({ path: '/login', query: { return_to: to.name } })
      }
    })
  } else {
    // store 中有用户信息时则将导航过滤交给后续
    next()
  }
})

// 如果是 authenticate 相关页面则返回 true
function isAuthPage(toPath) {
  if (['login', 'register', 'reset_password'].indexOf(toPath.name) == -1) return false
  else return true
}

export default router;
