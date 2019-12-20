import Mian from '@/components/main/main.vue';

export default [
  {
    path: '/',
    name: 'home',
    redirect: '/gateway/status',
    meta: {
      hideInMenu: true,
    },
  },
  {
    path: '/login',
    name: 'login',
    meta: {
      hideInMenu: true,
    },
    component: () => import(/* webpackChunkName: "login" */ '@/views/login/login.vue'),
  },
  {
    path: '*',
    name: 'error',
    meta: {
      hideInMenu: true,
    },
    component: () => import(/* webpackChunkName: "error" */ '@/views/error/error.vue'),
  },

  {
    path: '/gateway',
    name: 'Gateway',
    component: Mian,
    meta: {
      title: '网关状态',
    },
    children: [
      {
        path: 'status',
        name: 'GatewayStatus',
        meta: {
          title: '网关状态',
          icon: 'icon-gateway',
        },
        component: () => import(/* webpackChunkName: "gateway-status" */ '@/views/gateway-status/gateway-status.vue'),
      },
    ],
  },


  {
    path: '/net_config',
    name: 'net_config',
    component: Mian,
    meta: {
      title: '网络设置',
    },
    children: [
      {
        path: 'page',
        name: 'net_config_page',
        meta: {
          title: '网络设置',
          icon: 'icon-network-configuration',
        },
        component: () => import(/* webpackChunkName: "net_config_page" */ '@/views/net-config/net-config.vue'),
      },
    ],
  },
  // {
  //   path: '/data_flow',
  //   name: 'data_flow',
  //   component: Mian,
  //   meta: {
  //     title: '数据流配置',
  //   },
  //   children: [
  //     {
  //       path: 'page',
  //       name: 'data_flow_page',
  //       meta: {
  //         title: '数据流配置',
  //         icon: 'icon-data-stream',
  //       },
  //       component: () => import(/* webpackChunkName: "data_flow_page" */ '@/views/data-flow/data-flow.vue'),
  //     },
  //   ],
  // },
  {
    path: '/child_device',
    name: 'child_device',
    component: Mian,
    meta: {
      title: '设备配置',
    },
    children: [
      {
        path: 'page',
        name: 'child_device_page',
        meta: {
          title: '设备配置',
          icon: 'icon-data-stream',
        },
        component: () => import(/* webpackChunkName: "child_device_page" */ '@/views/child-device/child-device.vue'),
      },
    ],
  },
  {
    path: '/system_set',
    name: 'system_set',
    component: Mian,
    meta: {
      title: '系统设置',
    },
    children: [
      {
        path: 'page',
        name: 'system_set_page',
        meta: {
          title: '系统设置',
          icon: 'icon-system-configuration',
        },
        component: () => import(/* webpackChunkName: "system_set_page" */ '@/views/system-set/system-set.vue'),
      },

    ],
  },

  {
    path: 'test',
    name: 'test',
    meta: {
      hideInMenu: true,
    },
    component: () => import(/* webpackChunkName: "test" */ '@/views/test/test.vue'),
  },

];






