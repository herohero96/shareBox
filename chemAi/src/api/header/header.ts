import instance from '../axios';

// 网关时间
export const getGatewayTimeApi = () => {
  return instance.post('/gateway/time');
};
