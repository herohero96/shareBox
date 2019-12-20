<template>
  <div class="gateway-status">
    <div class="gateway-status-con">
      <div class="con-section system">
        <title-bar :titleName="'系统信息'"></title-bar>
        <div class="con system-con">
          <div class="item">
            <span class="item-key" title="GeWeb19v1.0.1">软件版本号：</span>
            <span class="item-value">{{system.software_version}}</span>
          </div>
          <div class="item">
            <span class="item-key">网关序列号：</span>
            <span class="item-value">{{system.gateway_no}}</span>
          </div>
          <div class="item">
            <span class="item-key">网关型号：</span>
            <span class="item-value">{{system.gateway_type}}</span>
          </div>
          <div class="item">
            <span class="item-key">网关备注：</span>
            <span class="item-value item-input-wrap item-textarea-wrap">
              <Input v-model="remark" type="textarea" :rows="4" />
              <g-button @click="saveRemark" class="item-input-btn">保存</g-button>
            </span>
          </div>
          <div class="item item-input gateway-heart-input">
            <span class="item-key">网关心跳：</span>
            <span class="item-value item-input-wrap" style="width: 105px;">
              <Input
                v-model="heart"
                @on-blur="checkItem('heart')"
                :class="{'error': heartError}"
                type="text"
                style="width: 90px;" /> s
              <g-button @click="setHeart" class="item-input-btn">确定</g-button>
            </span>
            <p v-if="heartError" class="error-msg">网关心跳应为大于等于30的整数</p>
          </div>
        </div>
      </div>
      <div class="con-section wifi">
        <title-bar :titleName="'无线网络状态'"></title-bar>
        <div class="con wifi-con" v-if="wifiState">
          <div class="item">
            <span class="item-key">无线网络状态：</span>
            <span class="item-value">{{wifiState ? '已连接' : '已断开'}}</span>
          </div>
          <div class="item">
            <span class="item-key">SSID：</span>
            <span class="item-value">{{wifi.ssid}}</span>
          </div>
          <div class="item">
            <span class="item-key">IPv4 地址：</span>
            <span class="item-value">{{wifi.ip}}</span>
          </div>
          <div class="item">
            <span class="item-key">IPv4DNS：</span>
            <span class="item-value">{{wifi.dns ? wifi.dns.join(' ; ') : ''}}</span>
          </div>
          <div class="item">
            <span class="item-key">网关：</span>
            <span class="item-value">{{wifi.gateway}}</span>
          </div>
          <div class="item">
            <span class="item-key">MAC 地址：</span>
            <span class="item-value">{{wifi.mac}}</span>
          </div>
        </div>
        <div v-else class="con wifi-con" >
          <div class="item empty">
            <span class="item-key">无线网络状态：</span>
            <span class="item-value">未连接任何无线网络</span>
            <router-link to="/net_config/page" class="item-btn">去连接</router-link>
          </div>
        </div>
      </div>
      <div class="con-section network">
        <title-bar :titleName="'有线网络状态'"></title-bar>
        <div class="con network-con">
          <div class="network-list">
            <div
              v-for="(net, index) in lanNames"
              :key="index"
              @click="selectNet(net)"
              class="network-list-item"
              :class="{'selected': selectLanName === net}"
            >
              <p class="left">
                <i class="iconfont icon-network"></i>
                <span class="list-name" :title="net">{{net}}</span>
              </p>
              <p class="right">
                {{lanNetwordk[net].connect ? '已连接' : '已断开'}}
              </p>
            </div>
          </div>
          <div class="item empty" v-if="lanNames.length==0">
            <span class="item-key">有线网络状态：</span>
            <span class="item-value">未连接任何有线网络</span>
            <router-link class="item-btn" to="/net_config/page">去连接</router-link>
          </div>
          <div class="divider-line"></div>
          <div class="network-detail" v-if="lanNames.length>0">
            <div class="item">
              <span class="item-key">当前配置：</span>
              <span class="item-value">{{selectLanDetail.connectionName}}</span>
            </div>
            <div class="item">
              <span class="item-key">网卡名称：</span>
              <span class="item-value">{{selectLanDetail.netName}}</span>
            </div>
            <div class="item">
              <span class="item-key">IPv4 地址：</span>
              <span class="item-value">{{selectLanDetail.ip}}</span>
            </div>
            <div class="item">
              <span class="item-key">IPv4DNS：</span>
              <span class="item-value">{{selectLanDetail.dns ? selectLanDetail.dns.join(' ; ') : ''}}</span>
            </div>
            <div class="item">
              <span class="item-key">网关：</span>
              <span class="item-value">{{selectLanDetail.gateway}}</span>
            </div>
            <div class="item">
              <span class="item-key">MAC 地址：</span>
              <span class="item-value">{{selectLanDetail.mac}}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="modify-password" v-if="modalShow">
      <modify-password-modal
        :show="modalShow"
        @cancel-modify-password="modalShow=false"
      >
      </modify-password-modal>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TitleBar from '@/components/title-bar/title-bar.vue';
import GButton from '@/components/button/g-button.vue';
import ModifyPasswordModal from '@/components/modify-password/modify-password-modal.vue';
import {
  getGatewayApi,
  queryWifiStateApi,
  queryWifiConnectionApi,
  getNetworkApi,
  getGatewayRemarkApi,
  changeGatewayRemarkApi,
  queryHeartApi,
  setHeartApi,
} from '@/api/gateway-status/gateway-status';

@Component({
  components: {
    TitleBar,
    GButton,
    ModifyPasswordModal,
  },
})
export default class GatewayStatus extends Vue {
  private modalShow: boolean = false; // 是否显示修改密码弹框
  private system: object = {}; // 系统信息
  private remark: string = ''; // 备注
  private wifiState: boolean = false; // wifi 连接状态
  private wifi: object = {}; // 无线网络
  private selectLanName: string = ''; // 有线网选中项 name
  private selectLanDetail: object = {}; // 有线网选中项详情
  private lanNames: any[] = []; // 有线网 name 列表
  private lanNetwordk: any = {}; // 有线网详情 obj
  private heart: number = 30; // 网关心跳
  private heartError: boolean = false; // 网关心跳校验

  // life-cycle
  private created(): void {
    // 判断是否显示修改密码弹窗
    const hasShowModifyPasswordModal = sessionStorage.getItem('hasShowModifyPasswordModal');
    this.modalShow = hasShowModifyPasswordModal !== 'yes' &&
      (this as any).$route.query.show_modal === 'true' ? true : false;
    if (this.modalShow) {
      sessionStorage.setItem('hasShowModifyPasswordModal', 'yes');
    }
  }

  private mounted() {
    this.getGateway();
    this.getRemark();
    this.getNetwork();
    this.queryWifiState();
    this.queryHeart();
  }

  // 系统信息
  private getGateway() {
    getGatewayApi().then((res: any) => {
      this.system = res && res.data;
    });
  }

  // wifi 连接状态
  private queryWifiState() {
    queryWifiStateApi().then((res: any) => {
      this.wifiState = res.data || false;
      if (this.wifiState) {
        this.queryWifiConnection();
      }
    });
  }

  // 获取当前连接的 wifi
  private queryWifiConnection() {
    queryWifiConnectionApi().then((res: any) => {
      this.wifi = res && res.data;
    });
  }

  // 有线网
  private getNetwork() {
    getNetworkApi().then((res: any) => {
      if (res.data) {
        this.lanNames = res.data.lanNames;
        this.lanNetwordk = res.data.lanNetwordk;
        this.selectLanName = this.lanNames[0];
        this.selectLanDetail = this.lanNetwordk[this.selectLanName];
      }
    });
  }

  // 切换有线网选中项
  private selectNet(net: string) {
    this.selectLanName = net;
    this.selectLanDetail = this.lanNetwordk[net];
  }

  // 获取备注
  private getRemark() {
    getGatewayRemarkApi().then((res: any) => {
      this.remark = res && res.data;
    });
  }

  // 保存备注
  private saveRemark() {
    if (this.remark) {
      changeGatewayRemarkApi(this.remark).then((res: any) => {
        if (res.status === 200) {
          this.$Message.info('保存成功');
          this.getRemark();
        }
      });
    }
  }

  // 获取网关心跳
  private queryHeart(): void {
    queryHeartApi().then((res: any) => {
      this.heart = res.data || 0;
    });
  }

  private checkItem(type: string) {
    let isValid: boolean = true;

    if (type === 'heart') {
      if (!isNaN(this.heart) && Math.floor(this.heart) === Number(this.heart) && this.heart >= 30) {
        this.heartError = false;
      } else {
        this.heartError = true;
        isValid = false;
      }
    }

    return isValid;
  }

  // 设置网关心跳
  private setHeart() {
    if (this.checkItem('heart')) {
      setHeartApi(this.heart).then((res: any) => {
        if (res.status === 200) {
          this.$Message.info('保存成功');
          this.queryHeart();
        }
      });
    }
  }
}
</script>
<style lang="less" scoped>
@blueColor: #4D9DED;
@grayColor: #888888;
@blackColor: #333333;
.gateway-status {
  .gateway-status-con {
    .con-section {
      margin-top: 30px;
      font-size: 12px;
      &:not(:first-child) {
        margin-top: 13px;
      }
      .title {
        font-size: 16px;
        color: @blueColor;
        margin-bottom: 8px;
        font-weight: 400;
      }
      .divider-line {
        width: 100%;
        height: 1px;
        background: @grayColor;
      }
      .con {
        width: 100%;
        padding: 30px 100px;
        .item {
          font-weight: 400;
          line-height: 36px;
          .error-msg {
            position: absolute;
            left: 0;
            bottom: -31px;
            padding-left: 60px;
            color: #F4333C;
          }
          &.gateway-heart-input {
            position: relative;
          }
          &.item-input {
            line-height: 32px;
          }
          .item-key {
            color: @blackColor;
          }
          .item-value {
            color: #888;
            &.item-input-wrap {
              position: relative;
              display: inline-block;
              margin-top: 30px;
              width: 90px;
              .item-input-btn {
                position: absolute;
                bottom: 0;
                right: -80px;
              }
            }
            &.item-textarea-wrap {
              width: 400px;
              vertical-align: text-top;
              margin-top:0;
            }
          }
        }
        .network-list {
          width: 100%;
          display: flex;
          justify-content: flex-start;
          align-items: center;
          flex-wrap: wrap;
          .network-list-item {
            padding: 10px;
            border-radius: 3px;
            width: 140px;
            height: 58px;
            margin-bottom: 30px;
            background: #fff;
            border: 1px solid transparent;
            cursor: pointer;
            &.selected {
              border:1px solid rgba(77, 157, 237, 1);
              box-shadow:0px 1px 3px 0px rgba(77,157,237,0.59);
            }
            &:not(:last-child) {
              margin-right: 20px;
            }
            .right {
              text-align: right;
            }
            .left {
              .icon-network {
                margin-right: 5px;
              }
            }
          }
        }
        .network-detail {
          margin: 38px auto 0 auto;
          .item {
            line-height: 36px;
          }
        }
        .empty {
          padding-bottom: 76px;
          .item-btn {
            color: #4D9DED;
            padding-left: 24px;
            cursor: pointer;
          }
        }
      }
    }
  }
}
</style>
