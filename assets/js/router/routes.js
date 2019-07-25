import { load } from './utils'

function defaultRoutes() {
  return [
    {
      path: '/',
      redirect: { path: '/login' }
    },
    {
      path: '/auth',
      component: load('authenticate/index'),
      children: [
        {
          path: 'register',
          alias: '/register',
          name: 'register',
          component: load('authenticate/register')
        },
        {
          path: 'login',
          alias: '/login',
          name: 'login',
          component: load('authenticate/login')
        },
        {
          path: 'forget_password',
          alias: '/forget_password',
          name: 'forget_password',
          component: load('authenticate/forget')
        }
      ]
    }
  ]
}

export {
  defaultRoutes
}
