<template>
  <div class="header-bar">
    <div class="custom-content-con">
      <div class="item sys-time">{{sysTime}}</div>
      <div class="item help">
        <i class="iconfont icon-help"></i>
        <span class="help-text">帮助</span>
      </div>
      <div class="item logout">
        <span class="user-name">{{userName}}</span>
        <i class="iconfont icon-sign-out" @click="handleLogOut"></i>
        <span class="logout-text" @click="handleLogOut">退出</span>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import { getGatewayTimeApi } from '@/api/header/header';

@Component
export default class HeaderBar extends Vue {
  // data
  private userName?: string;
  private sysTime?: any;
  private timer?: any;
  private data() {
    return {
      userName: '',
      sysTime: '',
    };
  }

  // life-cycle
  private created() {
    this.userName = sessionStorage.getItem('userName') || '';
    this.getDate();
  }
  private destroyed() {
    clearInterval(this.timer);
  }

  // methods
  private handleLogOut() {
    sessionStorage.removeItem('userName');
    sessionStorage.removeItem('hasShowModifyPasswordModal');
    this.$router.push({
      name: 'login',
    });
  }
  private getDate() {
    getGatewayTimeApi().then((res: any) => {
      if (res.data) {
        let currentTime = new Date(res.data.replace(/-/g, '/')).getTime(); // 接口获取网关时间
        this.sysTime = this.dateFormat(currentTime);
        this.timer = setInterval(() => {
          currentTime += 1000;
          this.sysTime = this.dateFormat(currentTime);
          return;
        }, 1000);
      }
    });
  }
  private digitFormat(date: number) {
    return date < 10 ? '0' + date : date;
  }
  private dateFormat(timestamp: any) {
    const dateTime = new Date(timestamp);
    const year = dateTime.getFullYear();
    const month = this.digitFormat(dateTime.getMonth() + 1);
    const date = this.digitFormat(dateTime.getDate());
    const hour = this.digitFormat(dateTime.getHours());
    const minute = this.digitFormat(dateTime.getMinutes());
    const second = this.digitFormat(dateTime.getSeconds());
    return  year + '-' + month + '-' + date + ' ' + hour + ':' + minute + ':' + second;
  }
}
</script>
<style lang="less" scoped>
.header-bar {
  .custom-content-con {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    line-height: 40px;
    .item {
      font-size: 12px;
      color: #333;
      margin-left: 32px;
      cursor: pointer;
      .iconfont {
        font-size: 14px;
        color: #333;
      }
      span {
        display: inline-block;
        vertical-align: top;
      }
      .help-text, .user-name, .logout-text {
        padding-left: 4px;
      }
      .user-name {
        padding-right: 4px;
      }
    }
    .help {
      margin-left: 60px;
    }
  }
}
</style>
