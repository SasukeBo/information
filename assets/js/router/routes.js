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
          component: load('main/home')
        },
        {
          path: 'product',
          component: load('main/product'),
          children: [
            {
              path: 'new',
              name: 'product-new',
              component: load('main/product/form')
            },
            {
              path: ':id/show',
              name: 'product-show',
              props: true,
              component: load('main/product/show')
            },
            {
              path: ':id/edit',
              name: 'product-edit',
              props: true,
              component: load('main/product/form')
            },
            {
              path: 'list',
              name: 'product-list',
              component: load('main/product/list')
            }
          ]
        },
        {
          path: 'devices',
          name: 'device-list',
          component: load('main/devices')
        },
        {
          path: 'device/new',
          name: 'device-new',
          component: load('main/device/new')
        },
        {
          path: 'device/:uuid',
          name: 'device-show',
          props: true,
          component: load('main/device'),
          redirect: { name: 'device-details' },
          children: [
            {
              path: 'realtime',
              name: 'device-realtime',
              props: true,
              component: load('main/device/realtime')
            },
            {
              path: 'charges',
              name: 'device-charges',
              props: true,
              component: load('main/device/charge')
            },
            {
              path: 'params',
              name: 'device-params',
              props: true,
              component: load('main/device/params')
            },
            {
              path: 'details',
              name: 'device-details',
              props: true,
              component: load('main/device/details')
            },
            {
              path: 'status-log',
              name: 'device-status-log',
              props: true,
              component: load('main/device/status-log')
            },
            {
              path: 'config',
              name: 'device-config',
              props: true,
              component: load('main/device/_config')
            }
          ]
        },
        {
          path: 'device/:uuid/charge',
          name: 'device-charge',
          props: true,
          component: load('main/charge'),
          children: [
            {
              path: ':id/show',
              name: 'charge-show',
              props: true,
              component: load('main/charge/_show.vue')
            },
            {
              path: 'new',
              name: 'charge-new',
              props: true,
              component: load('main/charge/_new.vue')
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
    },
    {
      path: '/*',
      name: 'error-page',
      component: load('errors/404.vue')
    }
  ]
}

export {
  defaultRoutes
}
