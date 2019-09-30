import VueRouter from 'vue-router'
import store from '../vuex'
import currentUserQuery from './gql/query.currentUser.gql';

import { defaultRoutes } from './routes.js'

const router = new VueRouter({
  mode: 'history',
  routes: defaultRoutes()
})

router.beforeEach((to, from, next) => {
  NProgress.start()
  var app = router.app

  if (!store.state.user.phone) { // 没有用户信息
    app.$apollo.query({
      query: currentUserQuery,
      fetchPolicy: 'network-only'
    }).then(({ data: { currentUser } }) => { // 获取成功
      app.$store.dispatch('user/setUserData', currentUser)

      if (isAuthPage(from)) {
        // 登录状态下如果是 auth 相关页面则导向 首页
        // 如果有 return_to 则导向 return_to
        var return_to = from.query.return_to
        return_to ? next({ name: return_to, params: from.query }) : next({ name: 'index' })
      } else {
        next()
      }
    }).catch((_) => { // 获取失败
      // 未登录状态下如果是 auth 相关页面则继续
      if (isAuthPage(to)) {
        next()
      } else {
        next({ path: '/login', query: { return_to: to.name, ...to.params } })
      }
    })
  } else {
    // store 中有用户信息时则将导航过滤交给后续
    next()
  }
})

// 如果是 authenticate 相关页面则返回 true
function isAuthPage(toPath) {
  return ['login', 'register', 'reset_password'].indexOf(toPath.name) > -1
}

export default router;
