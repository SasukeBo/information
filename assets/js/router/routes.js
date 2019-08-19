import { load, denyIfLoggedIn } from './utils'

function defaultRoutes() {
  return [
    {
      path: '/',
      name: 'index',
      component: load('main'),
      redirect: '/home',
      children: [
        {
          path: 'home',
          name: 'home',
          component: load('main/pages/home')
        },
        {
          path: 'user-device',
          name: 'user-device',
          component: load('main/pages/device')
        }
      ]
    },
    {
      path: '/auth',
      component: load('authenticate'),
      beforeEnter: denyIfLoggedIn(),
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
