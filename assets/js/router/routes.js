import { load } from './utils'

function defaultRoutes() {
  return [
    {
      path: '/',
      redirect: { name: 'login' }
    },
    {
      path: '/login',
      name: 'login',
      component: load('passport/login.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: load('passport/register.vue')
    }
  ]
}

export {
  defaultRoutes
}
