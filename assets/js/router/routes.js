import { load } from './utils'

function defaultRoutes() {
  return [
    {
      path: '/',
      redirect: { path: '/login' }
    },
    {
      path: '/home',
      name: 'home',
      component: load('home')
    },
    {
      path: '/auth',
      component: load('authenticate'),
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
          path: 'reset_password',
          alias: '/reset_password',
          name: 'reset_password',
          component: load('authenticate/reset')
        }
      ]
    }
  ]
}

export {
  defaultRoutes
}
