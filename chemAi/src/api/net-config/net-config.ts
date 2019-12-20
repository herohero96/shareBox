import instance from '../axios';

/**
 * 无线
 */

// 查询 wifi 状态
export const queryWifiStateApi = () => {
  return instance.post('/wifi/queryWifiState');
};

// 打开 wifi
export const openWifiApi = () => {
  return instance.post('/wifi/openWifi');
};

// 关闭 wifi
export const closeWifiApi = () => {
  return instance.post('/wifi/closeWifi');
};

// 开始扫描 wifi
export const scanWifiApi = () => {
  return instance.post('/wifi/scanWifi');
};

// 查询周边 wifi 列表
export const areaWifiSpotApi = () => {
  return instance.post('/wifi/areaWifiSpot');
};

// 查询当前连接的 wifi
export const queryWifiConnectionApi = () => {
  return instance.post('/wifi/queryWifiConnection');
};

// 连接 wifi
interface ConnectWifiParam {
  password: any;
  ssid: string;
}
export const connectWifiApi = (param: ConnectWifiParam) => {
  return instance.post('/wifi/connectWifi', param);
};



/**
 * 有线
 */

// 查询当前LAN的配置
export const queryLANConfigApi = () => {
  return instance.post('/network/queryLANConfig');
};

// 查询所有LAN的配置（配置下拉列表）
export const queryAllLANConfigApi = () => {
  return instance.post('/network/queryAllLANConfig');
};

// 激活某一配置
interface ActiveLANConfigParam {
  connectionName: string;
  ifname: string;
}
export const activeLANConfigApi = (param: ActiveLANConfigParam) => {
  return instance.post('/network/activeLANConfig', param);
};

// 新增 LAN
interface AddLANConfigParam {
  autoConnect: boolean;
  dhcp: string;
  dns: any;
  gateway: string;
  ip: string;
  mask: string;
  connectionName: string;
  netName: string;
}
export const addLANConfigApi = (param: AddLANConfigParam) => {
  return instance.post('/network/addLANConfig', param);
};

// 修改 LAN
interface ModifyLANConfigParam {
  autoConnect: boolean;
  dhcp: string;
  dns: any;
  gateway: string;
  ip: string;
  mask: string;
  connectionName: string;
  netName: string;
}
export const modifyLANConfigApi = (param: ModifyLANConfigParam) => {
  return instance.post('/network/modifyLANConfig', param);
};

// 删除 LAN
export const deleteLANConfigApi = (param: any) => {
  return instance.post('/network/deleteLANConfig', param);
};











