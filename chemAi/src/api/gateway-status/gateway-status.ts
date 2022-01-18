import instance from '../axios';

// 网关基本信息
export const getGatewayApi = () => {
  return instance.post('/gateway/basicInfo');
};

// 查询 WIFI 状态
export const queryWifiStateApi = () => {
  return instance.post('/wifi/queryWifiState');
};

// 查询当前连接 WIFI
export const queryWifiConnectionApi = () => {
  return instance.post('/wifi/queryWifiConnection');
};

// 有线网络状态
export const getNetworkApi = () => {
  return instance.post('/network/queryLANConfig');
};

// 获取网关备注
export const getGatewayRemarkApi = () => {
  return instance.post('/gateway/queryNote');
};

// 修改网关备注
export const changeGatewayRemarkApi = (param: string) => {
  return instance.post('/gateway/setNote', param);
};

// 获取网关心跳
export const queryHeartApi = () => {
  return instance.post('/gateway/queryHeart');
};

// 修改网关心跳
export const setHeartApi = (param: number) => {
  return instance.post('/gateway/setHeart', param);
};
