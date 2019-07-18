import { load } from './utils'

function defaultRoutes() {
  return [
    {
      path: '/',
      redirect: { name: 'login' }
    },
    {
      path: '/passport',
      component: load('passport/index.vue'),
      children: [
        {
          path: 'register',
          alias: '/register',
          name: 'register',
          component: load('passport/register.vue')
        },
        {
          path: 'login',
          alias: '/login',
          name: 'login',
          component: load('passport/login.vue')
        },
        {
          path: 'forget_password',
          alias: '/forget_password',
          name: 'forget_password',
          component: load('passport/forget.vue')
        }
      ]
    }
  ]
}

export {
  defaultRoutes
}
