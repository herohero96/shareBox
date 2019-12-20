import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store/store';
import './registerServiceWorker';
import iView from 'iview';
import 'iview/dist/styles/iview.css';
import '@/assets/icons/iconfont.css';
import config from '@/config';

Vue.use(iView);

Vue.config.productionTip = false;

/**
 * @description 全局注册应用配置
 */
Vue.prototype.$config = config;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
