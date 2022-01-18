import instance from '../axios';

// ping
export const pingTestApi = (param: string) => {
  return instance.post('/netDiagnosis/pingTest', param);
};

// traceroute
export const tracertTestApi = (param: string) => {
  return instance.post('/netDiagnosis/tracertTest', param);
};

// nslookupTest
export const nslookupTestApi = (param: string) => {
  return instance.post('/netDiagnosis/nslookupTest', param);
};

// 立即重启
export const rebootApi = () => {
  return instance.post('/system/reboot');
};

// 查看是否重启成功
export const getRebootStateApi = () => {
  return instance.post('/heart/heart');
};

// 恢复出厂设置
interface ResetApiParam {
  password: any;
  username: any;
}
export const resetApi = (param: ResetApiParam) => {
  return instance.post('/system/reset', param);
};



