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
          path: 'devices',
          name: 'device-list',
          component: load('main/pages/devices')
        },
        {
          path: 'device/:uuid',
          name: 'device-show',
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
