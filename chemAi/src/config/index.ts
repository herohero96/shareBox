
const url = {
  dev: '/api/',
  pro: '/apis/',
};

const baseUrl = process.env.NODE_ENV === 'development' ? url.dev : url.pro;


export default {
  baseUrl,
  /**
   * @description 默认打开的首页的路由name值，默认为home
   */
  homeName: 'home',
};
