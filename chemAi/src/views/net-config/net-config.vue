<template>
  <div class="net-config">
    <div class="net-config-wifi">
      <title-bar :titleName='"无线网络配置"'>
        <Switch v-model="isOpenWifi" size="large">
          <span slot="open">ON</span>
          <span slot="close">OFF</span>
        </Switch>
      </title-bar>
      <div class="con">
        <p class="con-row">附近的WLAN：</p>
        <div class="con-scan">
          <!-- <Select
            v-model="currentWifi.value"
            style="width:200px"
            :disabled="!isOpenWifi"
            class="wifi-select"
            @on-change="changeWifiSelect"
          >
            <Option v-for="item in wifiList" :value="item.value" :key="item.value">
              <span v-if="item.isSelected" class="selected-icon"><i class="iconfont icon-check"></i></span>
              {{item.label}}
              <span class="right-icon">
                <i class="iconfont icon-privac-open"></i>
                <i class="iconfont icon-wifi"></i>
              </span>
            </Option>
          </Select> -->
          <Dropdown :disabled="!isOpenWifi" trigger="click" class="wifi-dropdown">
            <a href="javascript:void(0)" class="wifi-dropdown-rel">
                {{currentWifi.ssid}}
                <Icon type="ios-arrow-down" class="icon"></Icon>
            </a>
            <DropdownMenu slot="list" class="wifi-select-dropdown">
                <DropdownItem
                  v-for="(item, index) in wifiList"
                  :key="index"
                  @click.native="changeWifiSelect(item)"
                >
                  <span v-if="item.ssid === currentWifi.ssid" class="selected-icon-wrap"><i class="iconfont icon-check"></i></span>
                  <span :class="{'ssid-text': true, 'text-selected': item.ssid === currentWifi.ssid}">{{item.ssid}}</span>
                  <span :class="{'right-icon': true, 'selected': item.ssid === currentWifi.ssid}">
                    <i v-if="item.security" class="iconfont icon-privac-open"></i>
                    <i class="iconfont icon-wifi"></i>
                  </span>
                </DropdownItem>
            </DropdownMenu>
          </Dropdown>       
          <g-button @click="scanWifi" class="btn fr" type="primary" :disabled='!isOpenWifi' :loading="isScanWifi">{{isScanWifi ? '扫描中...' : '重新扫描'}}</g-button>
        </div>
        <div class='con-detail'>
          <template v-if="isOpenWifi">
            <p>
              <span class="item-key">SSID：</span>
              <span class="item-value" :title="currentWifi.ssid">{{currentWifi.ssid}}</span>
            </p>
            <p>
              <span class="item-key">密码：</span>
              <span v-if="!isShowPassword" class="item-value">********</span>
              <span v-if="isShowPassword" class="item-value">{{currentWifi.password}}</span>
              <Checkbox class="fr" v-model="isShowPassword">显示密码</Checkbox>
            </p>
            <p>
              <span class="item-key">IPv4地址：</span>
              <span class="item-value">{{currentWifi.ip}}</span>
            </p>
            <p>
              <span class="item-key">子网掩码：</span>
              <span class="item-value">{{currentWifi.mask}}</span>
            </p>
            <p>
              <span class="item-key">IPv4DNS：</span>
              <span class="item-value">{{currentWifi.dns ? currentWifi.dns.join(' ') : ''}}</span>
            </p>
            <p>
              <span class="item-key">网关：</span>
              <span class="item-value">{{currentWifi.gateway}}</span>
            </p>
          </template>
        </div>
      </div>
    </div>
    <div class="net-config-wired">
      <title-bar :titleName='"有线网络配置"'></title-bar>
      <div class="con">
        <div class="con-card">
          <div
            v-for="(item, index) in lanNames"
            :key="index" class="con-card-item"
            :class="{'disabled': !lanNetwordk[item].connect, 'selected': selectLanName === item}"
            @click="selectLanName = item"
          >
            <p class="link">{{lanNetwordk[item].connect? '已连接': '未连接'}}</p>
            <p class="no"><i class="iconfont icon-network"></i>{{item}}</p>
          </div>
        </div>    
        <div class="con-select">
          <p class="title">选择配置方案：</p>
          <div>
            <Select v-model="connectionName" style="width:200px" class="lan-select">
              <Option v-for="(item, index) in allLanList" :label="item.connectionName" :value="item.connectionName" :key="index">
                <span v-if="connectionName === item.connectionName" class="selected-icon"><i class="iconfont icon-check"></i></span>
                <span>{{ item.connectionName }}</span>
                <span class="right-icon">
                  <i class="iconfont icon-edit" @click="editLanSet(item.connectionName)" ></i>
                </span>
              </Option>
            </Select>
            <g-button @click="editLanSet" class="select-btn select-edit">新增</g-button>
            <!-- :disabled="!connectionName || selectLanDetail.connectionName === connectionName || selectLanDetail.netName === 'console'" -->
            <g-button
              :disabled="!connectionName || selectLanDetail.netName === 'console'"
              @click="activateSelectSet()"
              type="primary"
              class="select-btn select-activate"
            >
              激活所选配置
            </g-button>
          </div>
        </div>
        <div class="con-detail" v-if="lanNames && lanNames.length>0">
          <div class="item">
            <span class="item-key">当前配置：</span>
            <span class="item-value">{{selectLanDetail.connectionName}}</span>
          </div>
          <div class="item">
            <span class="item-key">网卡名称：</span>
            <span class="item-value">{{selectLanDetail.netName}}</span>
          </div>
          <div class="item">
            <span class="item-key">IPv4地址：</span>
            <span class="item-value">{{selectLanDetail.ip}}</span>
          </div>
          <div class="item">
            <span class="item-key">IPv4DNS：</span>
            <span class="item-value">{{selectLanDetail.dns ? selectLanDetail.dns.join(' ') : ''}}</span>
          </div>
          <div class="item">
            <span class="item-key">网关：</span>
            <span class="item-value">{{selectLanDetail.gateway}}</span>
          </div>
        </div>
      </div>
    </div>

    <win-common v-if="isShowConnectModal" class="password-win" :title="connectForm.ssid">
      <div class="password-win-con">
        <p class="input">
          <span>密码</span>
          <Input
            v-model="connectForm.password"
            @on-blur="checkPassword"
            :type="connectForm.isShowConnectPassword ? 'text' : 'password'"
            style="width: 250px"
          />
        </p>
        <p class="Checkbox">
          <Checkbox v-model="connectForm.isShowConnectPassword">显示密码</Checkbox>
        </p>
        <p class="Checkbox">
          <Checkbox v-model="connectForm.isremenberNet">记住该网络</Checkbox>
        </p>
        <div class="btn">
          <g-button class="btn-cancel" type="default" @click="closeConnectModal()"> 取消</g-button>
          <g-button type="primary" :disabled = "isDisabledBtn" :unusable = "isUnusableBtn"  @click="connectWifi(true)">确定</g-button>
        </div>
        <div class="error">
          <p class="error-con" v-if="showPasswordError">
            <i class="iconfont icon-warning"></i>
            <span>{{passwordError}}</span>
          </p>
        </div>
      </div>
    </win-common>

    <win-common v-if="isShowLanModal" class="lan-win" :title="'编辑有线网络配置方案'" width="800">
      <div class="lan-set-header">
        <ul class="table-wrap">
          <li style="width: 100px; padding-left: 15px;">名称</li>
          <li style="width: 40px;">自动获取</li>
          <li style="width: 126px;">Ipv4地址</li>
          <li style="width: 240px;">Ipv4DNS</li>
          <!-- <li style="width: 116px;">网关</li> -->
          <li style="width: 126px;">子网掩码</li>
          <li style="width: 80px;">操作</li>
        </ul>
      </div>
      <div class="lan-set-con" ref="lanSetCon">
        <ul v-for="(item, index) of allLanList" :key="index" :class="{'table-wrap': true, 'table-wrap-readonly': !item.currentEdit}">
          <li style="width: 100px; padding-left: 15px;">
            <span v-if="!item.isAdd">{{item.connectionName}}</span>
            <Input
              v-else
              v-model="item.connectionName"
              @on-blur="checkLanItem(item, 'connectionName', index)"
              size="small"
              :class="{'set-input': true, 'set-input-error' : item.connectionNameStatus === 'error', 'set-input-right' : item.connectionNameStatus === 'right' }"
              style="width: 80px;"
            />
          </li>
          <li style="width: 40px;"><Checkbox v-model="item.dhcp" :disabled="!item.currentEdit"></Checkbox></li>
          <li style="width: 126px;">
            <Input
              v-model="item.ip"
              @on-blur="checkLanItem(item, 'ip', index)"
              v-if="!item.dhcp"
              :readonly="!item.currentEdit"
              size="small"
              :class="{'set-input': true, 'set-input-error' : item.ipStatus === 'error', 'set-input-right' : item.ipStatus === 'right' }"
            />
          </li>
          <li style="width: 240px;">
            <Input
              v-model="item.dns[0]" 
              @on-blur="checkLanItem(item, 'dns0', index)"
              v-if="!item.dhcp" :readonly="!item.currentEdit"
              size="small"
              :class="{'set-input': true, 'set-input-error' : item.dns0Status === 'error', 'set-input-right' : item.dns0Status === 'right' }"
            />
            <span v-if="!item.dhcp">;</span>
            <Input
              v-model="item.dns[1]"
              @on-blur="checkLanItem(item, 'dns1', index)"
              v-if="!item.dhcp"
              :readonly="!item.currentEdit"
              size="small"
              :class="{'set-input': true, 'set-input-error' : item.dns1Status === 'error', 'set-input-right' : item.dns1Status === 'right' }"
            />
          </li>
          <!-- <li style="width: 116px;">
            <Input
              v-model="item.gateway"
              @on-blur="checkLanItem(item, 'gateway', index)"
              v-if="!item.dhcp"
              :readonly="!item.currentEdit"
              size="small"
              :class="{'set-input': true, 'set-input-error' : item.gatewayStatus === 'error', 'set-input-right' : item.gatewayStatus === 'right' }"
            />
          </li> -->
          <li style="width: 126px;">
            <Input
              v-model="item.mask"
              @on-blur="checkLanItem(item, 'mask', index)"
              v-if="!item.dhcp"
              :readonly="!item.currentEdit"
              size="small"
              :class="{'set-input': true, 'set-input-error' : item.maskStatus === 'error', 'set-input-right' : item.maskStatus === 'right' }"
            />
          </li>
          <li style="width: 80px;">
            <span class="edit-icon-wrap">
              <i class="iconfont icon-edit" v-if="item.connectionName !== 'console' && !item.isLoading && !item.currentEdit" @click="editLanItem(item, index)"></i>
              <i class="iconfont icon-check" v-if="!item.isLoading && item.currentEdit" @click="saveLanItem(item, index)"></i>
              <Icon v-if="item.isLoading" type="ios-loading" size=14 class="demo-spin-icon-load"></Icon>
              <Poptip
                v-if="item.connectionName !== 'console'"
                placement="top-end"
                class="delete-poptip-set"
                confirm
                offset="15"
                @on-ok="deleLanItem(item, index)"
                @on-cancel="cancelDeleLanItem"
                ok-text="Yes"
                cancel-text="No"
              >
                <i class="iconfont icon-delete"></i>
                <div slot="title">
                  <i class="iconfont icon-warning"></i>
                  <span>确认要删除吗？</span>
                </div>
              </Poptip>
            </span>
          </li>
        </ul>
      </div>
      <div class="lan-footer">
        <!-- <p
          :class="[{'net-set-gateway-tip' : true}, {'zero-set' : allLanList.length === 0}, {'one-set' : allLanList.length === 1}, {'two-set' : allLanList.length === 2}, {'three-set' : allLanList.length === 3}, {'four-set' : allLanList.length === 4}]">
          <i class="iconfont icon-warn-fill"></i>
          南向接口不要配置网关
        </p> -->
        <span class="footer-add" @click="addLanItem">新增 +</span>
        <div class="footer-btn">
          <g-button class="btn" @click="cancelAddLan">取消</g-button>
          <g-button class="btn" :disabled="lanConfirmDisabled" @click="cancelAddLan" type="primary">确定</g-button>
        </div>
      </div>
    </win-common>

  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import TitleBar from '@/components/title-bar/title-bar.vue';
import WinCommon from '@/components/windows/win-common.vue';
import GButton from '@/components/button/g-button.vue';
import {
  queryWifiStateApi,
  openWifiApi,
  closeWifiApi,
  scanWifiApi,
  areaWifiSpotApi,
  queryWifiConnectionApi,
  connectWifiApi,
  queryLANConfigApi,
  queryAllLANConfigApi,
  activeLANConfigApi,
  addLANConfigApi,
  modifyLANConfigApi,
  deleteLANConfigApi,
} from '@/api/net-config/net-config';

@Component({
  components: {
    TitleBar,
    WinCommon,
    GButton,
  },
})

export default class NetConfig extends Vue {

  private isOpenWifi: boolean = false; // 无线网络配置 ON/OFF
  private isScanWifi: boolean = false; // 正在扫描 wifi
  private wifiList: object[] = []; // wifi 列表
  private currentWifi: any = {}; // 当前连接的 wifi
  private isShowPassword: boolean = false; // 显示当前 wifi 密码

  private isShowConnectModal: boolean = false; // 显示连接 wifi 弹框
  private isDisabledBtn: boolean = false;
  private isUnusableBtn: boolean = true;
  private passwordError: string = '密码错误';
  private showPasswordError: boolean = false;
  private connectForm: any = { // 弹框
    password: '',
    isShowConnectPassword: false,
    isremenberNet: true,
    ssid: '',
  };

  private selectLanName: string = ''; // 当前选中的有线网络 name
  private connectionName: string = ''; // 当前选中的有线网络的配置 name
  private selectLanDetail: any = {}; // 当前选中有线网络详情
  private lanNames: any = []; // 有线网络名字列表
  private lanNetwordk: any = {}; // 有线网络详情
  private isShowLanModal: boolean = false; // 显示有线网络配置弹框
  private allLanList: any = []; // 所有的 LAN 配置列表
  private lanConfirmDisabled: boolean = true;
  private lanItemModel = {
    autoConnect: true,
    dhcp: false,
    dns: [],
    gateway: '',
    ip: '',
    mask: '',
    name: '',
    connectionName: '',
    netName: '',
    isAdd: true, // 是否为新增
    currentEdit: true, // 当前正在编辑
    connectionNameStatus: 'nomal',
    ipStatus: 'nomal',
    dns0Status: 'nomal',
    dns1Status: 'nomal',
    gatewayStatus: 'nomal',
    maskStatus: 'nomal',
    isLoading: false,
  };

  public data() {
    return {
      allLanList: [],
      connectionName: '',
    };
  }

  public mounted() {
    this.queryWifiState(); // 查询当前 wifi 连接状态
    this.getAllLanList(); // 查询有线网络所有配置列表
    this.queryLANConfig(); // 查询有线网络
  }

  // wifi on/off
  @Watch('isOpenWifi')
  public startOpenWifi(val: boolean, oldVal: boolean) {
    if (val) {
      this.openWifi();
    } else {
      this.wifiList = [];
      this.currentWifi = {};
      this.closeWifi();
    }
  }

  // 校验弹框密码位数
  @Watch('connectForm.password')
  public onPasswordChanged(val: string, oldVal: string) {
    this.isUnusableBtn = val.length < 8 ? true : false;
  }

  // 当前选择的 Lan
  @Watch('selectLanName')
  public onSelectLanNameChanged(val: string, oldVal: string) {
    this.selectLanDetail = this.lanNetwordk[this.selectLanName];
    this.getCurrentLanSetName();
  }

  // 查询 wifi 连接状态
  private queryWifiState(): void {
    queryWifiStateApi().then((res: any) => {
      this.isOpenWifi = res.data || false;
      if (this.isOpenWifi) {
        this.queryWifiConnection();
      }
    });
  }

  // 打开 wifi
  private openWifi(): void {
    this.isScanWifi = true;
    openWifiApi().then((res: any) => {
      console.log(res);
      if (res.status === 200) {
        this.scanWifi();
      }
    });
  }

  // 关闭 wifi
  private closeWifi(): void {
    closeWifiApi().then((res: any) => {
      console.log(res);
      if (res.status === 200) {
        this.$Message.info('WIFI 已关闭');
      }
    });
  }

  // 开始扫描 wifi
  private scanWifi(): void {
    this.isScanWifi = true;
    this.queryWifiConnection();
    scanWifiApi().then((res: any) => {
      console.log(res);
      if (res.status === 200) {
        let intervalRun = 0;
        const timer = setInterval(() => {
          intervalRun ++;
          if (intervalRun > 5) {
            this.isScanWifi = false;
            clearInterval(timer);
          } else {
            this.areaWifiSpot();
            this.queryWifiConnection();
          }
        }, 3 * 1000);
      }
    });
  }

  // 查询周边 wifi 列表
  private areaWifiSpot(): void {
    areaWifiSpotApi().then((res: any) => {
      console.log('wifi-list');
      console.log(res);
      this.wifiList = res.data;
    });
  }

  // 查询当前连接的 wifi 详情
  private queryWifiConnection(): void {
    queryWifiConnectionApi().then((res: any) => {
      this.currentWifi = res.data;
    });
  }

  // 连接 wifi
  private connectWifi(hasPassword: boolean): void {
    const param: any = {
      ssid: this.connectForm.ssid,
    };
    if (hasPassword) {
      if (this.isUnusableBtn || this.isDisabledBtn) {
        return;
      }
      param.password = this.connectForm.password;
      this.isDisabledBtn = true;
    }
    connectWifiApi(param).then((res: any) => {
      console.log(res);
      this.isDisabledBtn = false;
      if (res.status === 200) {
        this.closeConnectModal();
        this.queryWifiConnection();
        this.$Message.info('连接成功');
      }
    });
  }

  // 切换无线网络
  private changeWifiSelect(item: any): void {
    console.log(item);
    if (item.ssid !== this.currentWifi.ssid) {
      this.connectForm.ssid = item.ssid;
      if (item.security) { // 需要输入密码
        this.isShowConnectModal = true;
        } else { // 不需要输入密码
        this.currentWifi = JSON.parse(JSON.stringify(item));
        this.connectWifi(false);
      }
    }
  }

  // 校验密码格式
  private checkPassword(value: any): void {
    if (this.connectForm.password.length < 8) {
      this.showPasswordError = true;
      const timer = setTimeout(() => {
        this.showPasswordError = false;
        clearTimeout(timer);
      }, 3000);
    }
  }

  // 关闭连接 wifi 弹框
  private closeConnectModal(): void {
    this.isShowConnectModal = false;
    this.connectForm.ssid = '';
    this.connectForm.password = '';
  }





  // 校验 LAN
  private checkLanItem(lanItem: any, type: any, index: any) {
    let isValid: boolean = true;
    const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (type === 'dns0') {
      if (!lanItem.dns[0] || (lanItem.dns[0] && reg.test(lanItem.dns[0]))) {
        lanItem.dns0Status = 'right';
      } else {
        lanItem.dns0Status = 'error';
        isValid = false;
      }
    } else if (type === 'dns1') {
      if (!lanItem.dns[1] || (lanItem.dns[1] && reg.test(lanItem.dns[1]))) {
        lanItem.dns1Status = 'right';
      } else {
        lanItem.dns1Status = 'error';
        isValid = false;
      }
    } else if (type === 'connectionName') {
      const regEn = /[ `!@#$%^&*()_+<>?:"{},.\/;'[\]]/im;
      const regCn = /[ ·！#￥（——）：；“”‘、，|《。》？、【】[\]]/im;
      if (lanItem.connectionName && !regEn.test(lanItem.connectionName) && !regCn.test(lanItem.connectionName)) {
        lanItem.connectionNameStatus = 'right';
      } else {
        lanItem.connectionNameStatus = 'error';
        isValid = false;
        this.$Message.error('请输入不包含特殊字符的名称');
      }
    } else if (type === 'gateway') {
      if (!lanItem.gateway || (lanItem.gateway && reg.test(lanItem.gateway))) {
        lanItem.gatewayStatus = 'right';
      } else {
        lanItem.gatewayStatus = 'error';
        isValid = false;
      }
    } else {
      if (lanItem[type] && reg.test(lanItem[type])) {
        lanItem[`${type}` + 'Status'] = 'right';
      } else {
        lanItem[`${type}` + 'Status'] = 'error';
        isValid = false;
      }
    }
    this.$set(this.allLanList, index, lanItem);

    return isValid;
  }

  private checkLanAllItem(lanItem: any, index: any) {
    let isValid: boolean = true;
    const list = ['connectionName', 'ip', 'dns0', 'dns1', 'gateway', 'mask'];
    const dhcpList = ['connectionName'];
    if (lanItem.dhcp) {
      dhcpList.forEach((item: any) => {
        if (!this.checkLanItem(lanItem, item, index)) {
          isValid = false;
        }
      });
    } else {
      list.forEach((item: any) => {
        if (!this.checkLanItem(lanItem, item, index)) {
          isValid = false;
        }
      });
    }

    return isValid;
  }

  // 查询有线网络列表
  private queryLANConfig(): void {
    queryLANConfigApi().then((res: any) => {
      const data = res.data;
      this.lanNames = data.lanNames;
      this.lanNetwordk = data.lanNetwordk;
      if (!this.selectLanName) {
        this.selectLanName = this.lanNames[0];
      }
      this.selectLanDetail = this.lanNetwordk[this.selectLanName];
      queryAllLANConfigApi().then((allRes: any) => {
        this.allLanList = allRes.data;
        this.getCurrentLanSetName();
      });
    });
  }

  // 查询所有配置列表
  private getAllLanList(): void {
    queryAllLANConfigApi().then((res: any) => {
      this.allLanList = res.data;
    }).catch(() => {
      this.allLanList = [];
    });
  }

  // 获取当前有线网络的当前配置
  private getCurrentLanSetName(): void {
    this.connectionName = '';
    this.allLanList.forEach((item: any) => {
      if (item.connectionName === this.selectLanDetail.connectionName) {
        this.connectionName = item.connectionName;
      }
    });
  }

  // 编辑 LAN 配置
  private editLanSet(connectionName?: any): void {
    if (connectionName) {
      this.connectionName = connectionName;
    }
    this.isShowLanModal = true;
    this.lanConfirmDisabled = true;
    this.allLanList.forEach((item: any, index: any) => {
      item.isAdd = false;
      item.currentEdit = false;
      item.connectionNameStatus = 'nomal';
      item.ipStatus = 'nomal';
      item.dns0Status = 'nomal';
      item.dns1Status = 'nomal';
      item.gatewayStatus = 'nomal';
      item.maskStatus = 'nomal';
      if (item.connectionName && item.connectionName === this.connectionName) {
        (this as any).$nextTick(() => {
          (this as any).$refs.lanSetCon.scrollTop = index > 4 ? index * 40 : 0;
        });
        if (connectionName) { // 新增时都不是编辑状态
          item.currentEdit = true;
        }
        // this.checkLanAllItem(item, index);
      }
    });
  }

  // 添加空 LAN 配置项
  private addLanItem(): void {
    this.lanConfirmDisabled = true;
    this.allLanList.push(JSON.parse(JSON.stringify(this.lanItemModel)));
    (this as any).$nextTick(() => {
      (this as any).$refs.lanSetCon.scrollTop = this.allLanList.length * 40;
    });
  }

  // 保存 LAN 配置
  private saveLanItem(item: any, index: any): void {
    if (this.checkLanAllItem(item, index)) {
      item.isLoading = true;
      const param = {
        autoConnect: true,
        dhcp: item.dhcp,
        dns: JSON.parse(JSON.stringify(item.dns)),
        gateway: item.gateway,
        ip: item.ip,
        mask: item.mask,
        connectionName: item.connectionName,
        netName: this.selectLanName,
      };
      if (item.isAdd) { // 新增
        addLANConfigApi(param).then((res: any) => {
          if (res.status === 200) {
            this.$Message.info('保存成功');
            item.currentEdit = false;
            item.isAdd = false;
            item.isLoading = false;
            this.$set(this.allLanList, index, item);
            this.lanConfirmDisabled = false;
          }
        }).catch(() => {
          item.isLoading = false;
        });
      } else { // 编辑
        modifyLANConfigApi(param).then((res: any) => {
          if (res.status === 200) {
            this.$Message.info('修改成功');
            item.currentEdit = false;
            item.isLoading = false;
            this.$set(this.allLanList, index, item);
            this.lanConfirmDisabled = false;
          }
        }).catch(() => {
          item.isLoading = false;
        });
      }
    }
  }

  // 编辑 LAN 配置项
  private editLanItem(item: any, index: any): void {
    this.lanConfirmDisabled = true;
    item.currentEdit = true;
    this.$set(this.allLanList, index, item);
    this.checkLanAllItem(item, index);
  }

  // 取消添加 LAN 配置
  private cancelAddLan(): void {
    this.isShowLanModal = false;
    this.getAllLanList(); // 查询有线网络所有配置列表
    this.queryLANConfig(); // 查询有线网络列表
  }

  // 删除 LAN 配置
  private deleLanItem(param: any, index: any): void {
    if (param.isAdd) {
      this.allLanList.splice(index, 1);
    } else {
      deleteLANConfigApi(param.connectionName).then((res: any) => {
        if (res.status === 200) {
          this.allLanList.splice(index, 1);
          this.$Message.info('删除成功');
          this.lanConfirmDisabled = false;
        }
      });
    }
  }

  // 取消删除 LAN
  private cancelDeleLanItem(): void {
    //
  }

  // 激活所选 LAN 配置
  private activateSelectSet(): void {
    activeLANConfigApi({connectionName: this.connectionName , ifname: this.selectLanName}).then((res: any) => {
      if (res.status === 200) {
        this.$Message.info('激活成功');
        this.selectLanDetail.connectionName = this.connectionName;
        setTimeout(() => {
          this.queryLANConfig();
        }, 3000);
      }
    });
  }
}
</script>

<style lang="less" scoped>
@borderColor: #888;
@titleColor: #333;
@hoverColor: #2d8cf0;

.net-config {
  .ivu-switch {
    margin-left: 10px;
    margin-bottom: 5px;
  }
  .net-config-wifi {
    padding-top: 30px;
    .con {
      .con-row {
        line-height: 40px;
      }
      padding: 10px 0 10px 0;
      padding-left: 100px;
      width: 420px;
      .con-scan {
        // .wifi-select {
        //   .ivu-select-item {
        //     padding: 7px 7px 7px 18px;
        //     .selected-icon {
        //       position: relative;
        //       .iconfont {
        //         position: absolute;
        //         top: 1px;
        //         left: -15px;
        //       }
        //     }
        //     .right-icon {
        //       float: right;
        //       .iconfont {
        //         &:last-child {
        //           padding-left: 4px;
        //         }
        //       }
        //     }
        //     &:hover {
        //       background: #A6CEF6;
        //     }
        //   }
        // }
        .wifi-dropdown {
          .wifi-dropdown-rel {
            border: 1px solid @borderColor;
            background: #fff;
            height: 30px;
            line-height: 30px;
            width: 200px;
            border-radius: 3px;
            padding: 0 6px;
            color: @borderColor;
            display: inline-block;
            overflow: hidden;
            text-overflow: ellipsis;
            .icon {
              float: right;
              margin-top: 8px;
            }
          }
          .wifi-select-dropdown {
            width: 200px;
            max-height: 200px;
            overflow: auto;
            .selected-icon-wrap {
              // position: relative;
              // width: 0;
              // height: 0;
              // display: inline-block;
              color: @hoverColor;
              vertical-align: text-bottom;
              .iconfont {
                position: absolute;
                top: 8px;
                left: 3px;
              }
            }
            .text-selected {
              color: @hoverColor;
            }
            .right-icon {
              float: right;
              .iconfont {
                &:last-child {
                  padding-left: 4px;
                }
              }
              &.selected {
                color: @hoverColor;
              }
            }
            li {
              padding: 7px 7px 7px 18px;
              position: relative;
              &:hover {
                background: #A6CEF6;
              }
              .ssid-text {
                // overflow: hidden;
                // text-overflow: ellipsis;
                // white-space: nowrap;
                width: 120px;
                max-width: 120px;
                display: inline-block;
                word-break: break-all;
                white-space: pre-wrap;
                vertical-align: text-top;
              }
            }
          }
        }
        .btn {
          // margin-top: 4px;
        }
      }
      .con-detail {
        min-height: 240px;
        p {
          line-height: 40px;
          font-weight: 400;
          .item-key {
            color: @titleColor;
          }
          .item-value {
            color: @borderColor;
          }
        }
        .con-row-right {
          float: right;
        }
      }

    }
  }
  .net-config-wired {
    .con {
      padding:  0 100px 0 100px;
      .con-card {
        display: flex;
        flex-flow:row wrap;
        justify-content: flex-start;
        align-items:center;
        border-bottom: 1px solid @borderColor;
        padding-top: 29px;
        padding-bottom: 9px;
        .con-card-item {
          width:140px;
          height:58px;
          padding: 10px;
          margin-right: 20px;
          margin-bottom: 20px;
          background-color: #fff;
          box-shadow:0px 1px 3px 0px rgba(136,136,136,0.59);
          border-radius: 4px;
          cursor: pointer;
          &.disabled{
            color: #888;
            border:1px solid rgba(222, 222, 222, 1);
            box-shadow: none;
          }
          &.selected {
            border:1px solid #4D9DED;            
          }
          .link {
            text-align: right;
          }
          .no {
            i {
              margin-right: 5px;
            }
          }
        }
      }
      .con-select {
        padding: 30px 0;
        border-bottom: 1px solid @borderColor;
        .title {
          font-weight: 400;
          color: @titleColor;
          padding-bottom: 16px;
        }
        .select-edit {
          margin: 0 30px 0 20px;
        }
        .lan-select {
          .ivu-select-item {
            padding: 7px 12px 7px 18px;
            .selected-icon {
              position: relative;
              width: 0;
              height: 0;
              display: inline-block;
              .iconfont {
                position: absolute;
                top: -11px;
                left: -15px;
              }
            }
            .right-icon {
              float: right;
              .iconfont {
                &:last-child {
                  padding-left: 4px;
                }
              }
            }
            &:hover {
              background: #A6CEF6;
            }
          }
        }
      }
      .con-detail {
        padding: 28px 0 56px 0;
        .item {
          font-weight: 400px;
          line-height: 40px;
          .item-key {
            color: @titleColor;
          }
          .item-value {
            color: @borderColor;
          }
        }
      }
    }
  }

  .password-win {
    .password-win-con {
      padding: 30px 89px 20px 76px;
      .error {
        position: relative;
        color: #F4333C;
        .error-con {
          position: absolute;
          left: -36px;
          bottom: -3px;
          .iconfont {
            padding-right: 4px;
          }
        }
      }
      p.input {
        margin-bottom: 8px;
        span {
          margin-right: 10px;
        }
      }
      .Checkbox {
        text-align: left;
        padding-left: 35px;
        line-height: 20px;
      }
      .btn {
        padding-top: 10px;
        text-align: right;
        .btn-cancel {
          margin-right: 14px;
        }
      }
    }
  }

  .lan-win {
    h1 {
      border-bottom: none;
    }
    .table-wrap {
      color: @titleColor;
      display: table;
      width: 100%;
      height: 40px;
      li {
        display: table-cell;
        padding: 0 6px;
        text-align: left;
        vertical-align: middle;
        box-sizing: border-box;
        .set-input {
          width: 100px;
          
        }
      }
    }
    .lan-set-header {
      .table-wrap {
        background: #ECECEC;
        font-weight: 600;
        li {
          border-right: 1px solid #C3C3C3;
          &:last-child {
            border: none;
          }
        }
      }
    }
    .lan-set-con {
      height: 200px;
      overflow: auto;
      .table-wrap {
        transition: all .3s;
        &:hover {
          background: #DBEBFB;
        }
        li {
          border-bottom: 1px solid rgba(136,136,136,0.5);
          .icon-edit{
            padding-right: 20px;
          }
          .icon-check {
            padding-right: 16px;
            color: #4D9DED;
            font-size: 20px;
            vertical-align: sub;
          }
        }
      }
    }
    .lan-footer {
      // position: relative;
      padding: 30px 30px 20px 30px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      // .net-set-gateway-tip {
      //   .iconfont {
      //     color: #FF5B05;
      //     padding-right: 4px;
      //     vertical-align: middle;
      //   }
      //   position: absolute;
      //   top: 10px;
      //   left: 30px;
      //   &.zero-set {
      //     top: -190px;
      //   }
      //   &.one-set {
      //     top: -150px;
      //   }
      //   &.two-set {
      //     top: -110px;
      //   }
      //   &.three-set {
      //     top: -70px;
      //   }
      //   &.four-set {
      //     top: -30px;
      //   }
      // }
      .footer-add {
        color: #4D9DED;
        cursor: pointer;
      }
      .footer-btn {
        .btn {
          margin-left: 14px;
        }
      }
    }
    .edit-icon-wrap {
      display: flex;
      align-items: center;
      justify-content: space-between;
      width: 60px;
    }
    .demo-spin-icon-load {
      animation: ani-demo-spin 1s linear infinite;
      color: #4D9DED;
    }
    @keyframes ani-demo-spin {
      from { transform: rotate(0deg);}
      50%  { transform: rotate(180deg);}
      to   { transform: rotate(360deg);}
    }
  }
}
</style>
<style lang="less">
.lan-select {
  .ivu-select-selection {
    color: #888;
    border-color: #888;
    box-shadow: none;
  }
}
.lan-win {
  .win-con {
    h1 {
      border: none;
    }
    .set-input {
      .ivu-input-small {
        padding: 1px 1px;
      }
      &.set-input-error {
        .ivu-input-small {
          border: 1px solid #F4333C;
        }
      }
      &.set-input-right {
        .ivu-input-small {
          border: 1px solid #4D9DED;
        }
      }
    }
    .table-wrap-readonly {
      .set-input {
        .ivu-input-small {
          border: none;
          background: transparent;
          text-overflow: ellipsis;
          &:focus {
            box-shadow: none;
          }
        }
      }
    }
  }
}
</style>


