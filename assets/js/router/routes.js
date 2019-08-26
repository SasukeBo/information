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
          path: 'device/new',
          name: 'device-new',
          component: load('main/pages/device-new')
        },
        {
          path: 'device/:uuid',
          name: 'device-show',
          props: true,
          component: load('main/pages/device'),
          redirect: { name: 'device-details' },
          children: [
            {
              path: 'details',
              name: 'device-details',
              props: true,
              component: load('main/pages/device/details')
            },
            {
              path: 'charges',
              name: 'device-charges',
              props: true,
              component: load('main/pages/device/_charge')
            },
            {
              path: 'params',
              name: 'device-params',
              props: true,
              component: load('main/pages/device/_params')
            },
            {
              path: 'params-values',
              name: 'device-params-values',
              props: true,
              component: load('main/pages/device/_params-values')
            },
            {
              path: 'status-log',
              name: 'device-status-log',
              props: true,
              component: load('main/pages/device/_status-log')
            },
            {
              path: 'config',
              name: 'device-config',
              props: true,
              component: load('main/pages/device/_config')
            }
          ]
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
