import instance from '../axios';

// 导入配置
interface ImportSystemConfigParam {
  importSystemConfig: any;
}
export const importSystemConfigApi = (param: any) => {
  const config = {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  };
  return instance.post('/system/importSystemConfig', param, config);
};

// 导出配置
export const exportSystemConfigApi = () => {
  return instance.post('/system/exportSystemConfig');
};


// 查询端口信息
export const queryPortInfoApi = () => {
  return instance.post('/device/queryPortInfo');
};

// 查询端口配置
export const queryPortConfigApi = (param: any) => {
  return instance.post('/device/queryPortConfig', param);
};

// 修改端口配置信息
interface ModifyPortConfigParam {
  baudrate: any;
  dataBit: any;
  parity: any;
  portName: any;
  portType: any;
  stopBit: any;
}
export const modifyPortConfigApi = (param: ModifyPortConfigParam) => {
  return instance.post('/device/modifyPortConfig', param);
};

// 查询所有设备
export const queryAllDeviceApi = () => {
  return instance.post('/device/queryAllDevice');
};

// 添加设备
interface AddDeviceParam {
  connectionPort: string;
  createTime: any;
  deviceName: string;
  importFix: boolean;
  installLocation: string;
  mqttnorthboundList: any;
  serverPort: string;
  sevrerIp: string;
  southProtocol: string;
  southboundList: any;
  tcpType: string;
  template: boolean;
  tibocoRvnorthboundList: any;
  type: string;
  uuid: any;
}
export const addDeviceApi = (param: AddDeviceParam) => {
  return instance.post('/device/addDevice', param);
};

// 查询设备
export const queryDeviceApi = (param: any) => {
  return instance.post('/device/queryDevice', param);
};

// 编辑设备
interface EditDeviceParam {
  connectionPort: string;
  createTime: any;
  deviceName: string;
  importFix: boolean;
  installLocation: string;
  mqttnorthboundList: any;
  serverPort: string;
  sevrerIp: string;
  southProtocol: string;
  southboundList: any;
  tcpType: string;
  template: boolean;
  tibocoRvnorthboundList: any;
  type: string;
  uuid: any;
}
export const editDeviceApi = (param: EditDeviceParam) => {
  return instance.post('/device/editDevice', param);
};

// 删除设备
export const deleteDeviceApi = (param: any) => {
  return instance.post('/device/deleteDevice', param);
};

// 添加平田南向协议
interface AddHirataProtocolParam {
  deviceId: string;
  scanTime: string;
  uuid: string;
}
export const addHirataProtocolApi = (param: AddHirataProtocolParam) => {
  return instance.post('/device/addHirataProtocol', param);
};

// 编辑平田南向协议
interface EditHirataProtocolParam {
  deviceId: string;
  scanTime: string;
  uuid: string;
}
export const editHirataProtocolApi = (param: EditHirataProtocolParam) => {
  return instance.post('/device/editHirataProtocol', param);
};

// 查询平田南向协议
export const queryHirataProtocolApi = (param: any) => {
  return instance.post('/device/queryHirataProtocol', param);
};

// 添加南向三协协议
interface AddSankyoProtocolParam {
  deviceId: string;
  uuid: string;
}
export const addSankyoProtocolApi = (param: AddSankyoProtocolParam) => {
  return instance.post('/device/addSankyoProtocol', param);
};

// 查询三协协议
interface QuerySankyoProtocolParam {
  uuid: string;
}
export const querySankyoProtocolApi = (param: QuerySankyoProtocolParam) => {
  return instance.post('/device/querySankyoProtocol', param);
};

// 编辑南向三协协议
interface EditSankyoProtocolParam {
  deviceId: string;
  uuid: string;
}
export const editSankyoProtocolApi = (param: EditSankyoProtocolParam) => {
  return instance.post('/device/editSankyoProtocol', param);
};

// 删除南向三协协议
interface DeleteSankyoProtocolParam {
  uuid: string;
}
export const deleteSankyoProtocolApi = (param: DeleteSankyoProtocolParam) => {
  return instance.post('/device/deleteSankyoProtocol', param);
};


// 添加 MQTT 北向协议
interface AddMQTTNorthboundProtocolParam {
  deviceId: any;
  mqttServerClientID: any;
  mqttServerIp: string;
  mqttServerPort: any;
  password: string;
  username: string;
  mqttServerTopic: string;
  template: boolean;
  uuid: string;
}
export const addMQTTNorthboundProtocolApi = (param: AddMQTTNorthboundProtocolParam) => {
  return instance.post('/device/addMQTTNorthboundProtocol', param);
};

// 查询 MQTT 北向协议
export const queryMQTTNorthboundProtocolApi = (param: any) => {
  return instance.post('/device/queryMQTTNorthboundProtocol', param);
};

// 编辑 MQTT 北向协议
interface EditMQTTNorthboundProtocolParam {
  deviceId: any;
  mqttServerClientID: any;
  mqttServerIp: string;
  mqttServerPort: any;
  password: string;
  username: string;
  mqttServerTopic: string;
  template: boolean;
  uuid: string;
}
export const editMQTTNorthboundProtocolApi = (param: EditMQTTNorthboundProtocolParam) => {
  return instance.post('/device/editMQTTNorthboundProtocol', param);
};

// 删除 MQTT 北向协议
export const deleteMQTTNorthboundProtocolApi = (param: any) => {
  return instance.post('/device/deleteMQTTNorthboundProtocol', param);
};

// 添加 TibcoRv 北向协议
interface AddTibcoRvNorthboundProtocolParam {
  deviceId: any;
  service: string;
  network: string;
  daemon: string;
  subject: string;
  uuid: any;
}
export const addTibcoRvNorthboundProtocolApi = (param: AddTibcoRvNorthboundProtocolParam) => {
  return instance.post('/device/addTibcoRvNorthboundProtocol', param);
};

// 编辑 TibcoRv 北向协议
interface EditTibcoRvNorthboundProtocolParam {
  deviceId: any;
  service: string;
  network: string;
  daemon: string;
  subject: string;
  uuid: any;
}
export const editTibcoRvNorthboundProtocolApi = (param: EditTibcoRvNorthboundProtocolParam) => {
  return instance.post('/device/editTibcoRvNorthboundProtocol', param);
};

// 删除 TibcoRv 北向协议
export const deleteTibcoRvNorthboundProtocolApi = (param: any) => {
  return instance.post('/device/deleteTibcoRvNorthboundProtocol', param);
};

// 查询 TibcoRv 北向协议
export const queryTibcoRvNorthboundProtocolApi = (param: any) => {
  return instance.post('/device/queryTibcoRvNorthboundProtocol', param);
};

// 测试 MQTT 的连通性
export const testMqttConnectApi = (param: any) => {
  return instance.post('/device/testMqttConnect', param);
};

// 开始通讯
export const beginCommApi = (param: any) => {
  return instance.post('/device/beginComm', param);
};

// 查询相关设备通讯状态
export const queryCommStateApi = (param: any) => {
  return instance.post('/device/queryCommState', param);
};
