import Vue from 'vue';


// http错误处理
export const catchError = (error: any) => {
  // console.log('错误', error);

  if (error.response) {
    switch (error.response.status) {
      // case 400:
      //   Vue.prototype.$Message.error(error.response.data.message || '请求参数异常');
      //   break;
      // case 401:
      //   // Vue.prototype.$Message.error('登陆已失效，请重新登陆');
      //   Vue.prototype.$Modal.error({
      //     title: '登陆已失效，请重新登陆',
      //     onOk: () => {
      //       setCookie('token', '', -1); // cookies
      //       parent.location.href = '/login';
      //     },
      //   });
      //   break;
      // case 403:
      //   Vue.prototype.$Message.error(error.response.data.message || '无访问权限，请联系管理员');
      //   break;
      default:
        Vue.prototype.$Message.error(error.response.data.message || '服务端异常，请联系技术支持');
    }
  }
  return Promise.reject(error);
};
