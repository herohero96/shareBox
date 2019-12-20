import axios from 'axios';
import {catchError} from '@/config/utils';
import config from '@/config';
import Vue from 'vue';
const baseUrl = config.baseUrl;


const instance = axios.create({
  // timeout: 10000,
  baseURL: baseUrl,
});

instance.defaults.headers.post['Content-Type'] = 'application/json';

// request拦截器,增加token等
instance.interceptors.request.use(
  (res) => {
    res.headers.token = '412432';
    return res;
  },
  (error) => {
    console.log(error);
    Promise.reject(error);
  },
);

// response 拦截器
instance.interceptors.response.use((response) => {
  if (response.data.status && response.data.status !== 200) {
    Vue.prototype.$Message.error(response.data.msg);
  }
  return response.data;
}, catchError);

export default instance;
