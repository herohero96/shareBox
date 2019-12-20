<template>
  <div class="system-set">
    <div class="system-set-config">
      <title-bar :titleName='"网络诊断"'></title-bar>
      <div class="con">
        <div class="network-tomography-item">
          <Input
            v-model="tomography.ping"
            @on-blur="checkTomographyItem('ping')"
            :class="{'tomography-input': true, 'selected': tomographyType==='ping' && !tomography.pingError, 'error': tomography.pingError}"
          />
          <p v-if="tomography.pingError" class="error-text">请输入正确的 IP 或网址</p>
          <g-button :loading="tomography.loading === 'ping'" style="width: 105px;" @click="handleTomography('ping')">Ping</g-button>
        </div>
        <div class="network-tomography-item">
          <Input
            v-model="tomography.traceroute"
            @on-blur="checkTomographyItem('traceroute')"
            :class="{'tomography-input': true, 'selected': tomographyType==='traceroute' && !tomography.tracerouteError, 'error': tomography.tracerouteError}"
          />
          <p v-if="tomography.tracerouteError" class="error-text">请输入正确的 IP 或网址</p>
          <g-button :loading="tomography.loading === 'traceroute'" style="width: 105px;" @click="handleTomography('traceroute')">Traceroute</g-button>
        </div>
        <div class="network-tomography-item">
          <Input
            v-model="tomography.nsloopup"
            @on-blur="checkTomographyItem('nsloopup')"
            :class="{'tomography-input': true, 'selected': tomographyType==='nsloopup' && !tomography.nsloopupError, 'error': tomography.nsloopupError}"
          />
          <p v-if="tomography.nsloopupError" class="error-text">请输入正确的 IP 或网址</p>
          <g-button :loading="tomography.loading === 'nsloopup'" style="width: 105px;" @click="handleTomography('nsloopup')">Nsloopup</g-button>
        </div>
      </div>
      <div class="tomography-result">
        <Input v-model="tomography.result" readonly type="textarea" :autosize="{minRows: 6,maxRows: 6}" />
      </div>
    </div>
    <div class="system-set-update">
      <title-bar :titleName='"系统升级"'></title-bar>
      <div class="con">
          <p class="con-tip">仅支持bin,tar,gz,zip类型的文件。</p>
          <div class="con-file">
            <span class="con-file-name"></span>
            <input type="file" @change="fileChange($event)" class="con-file-input" accept=".bin,.tar,.gz,.zip" />
            <g-button class="con-file-btn" type="default">上传固件</g-button>
            <g-button type="primary" :unusable = 'isUpdateUnusableBtn' class="con-file-btn" @click=" !isUpdateUnusableBtn? isShowUpdateWin = true : null">
              升级
            </g-button>
          </div>
      </div>
    </div>
    <div class="system-set-reboot">
      <title-bar :titleName='"重启网关"'></title-bar>
      <div class="con">
        <p class="con-tip">点击立即重启后将中断当前所有连接，大约需要1分钟恢复。</p>
        <g-button class="btn" @click="isShowRebootWin = true" type="default">立即重启</g-button>
      </div>
    </div>
    <div class="system-set-recover">
      <title-bar :titleName='"恢复出厂设置"'></title-bar>
      <div class="con">
        <p class="con-tip">将网关恢复到出厂时的默认状态，所有配置将会消失。</p>
        <g-button class="btn" type="default" @click='isShowResetWin = true'>恢复出厂设置</g-button>
      </div>
    </div>
    <div class="system-set-password">
      <title-bar :titleName='"修改登录密码"'></title-bar>
      <div class="con">
        <modify-password-form :showLabel="false" :type="'page'"></modify-password-form>
      </div>
    </div>

    <win-tip v-if="isRebootNow" :title="'网关正在重启...'" :content="'网关重启成功后，将会自动连接，请不要刷新页面。'"></win-tip>
    <win-tip-second v-if="isRebootSuccess" :content="'网关重启成功'"></win-tip-second>
    <win-tip v-if="isResetNow" :title="'网关正在恢复出厂设置...'" :content="'网关恢复出厂设置成功后，需要原始密码重新登录。'"></win-tip>
    <win-tip-second v-if="isResetSuccess" :content="'网关恢复出厂设置成功'"></win-tip-second>
    

    <win-common class='rebootWin' v-if='isShowRebootWin' :title="'重启网关确认'">
      <p class='rebootWin-tip'>重启网关将中断当前所有连接，大约需要一分钟恢复。</p>
      <div class='rebootWin-con clearfix'>
        <g-button class="btn fr" type="primary" @click="reboot">确定</g-button>
        <g-button class="btn btn-cancel fr" type="default" @click="isShowRebootWin = false"> 取消</g-button>
      </div>
    </win-common>

    <win-common class='recoverWin' v-if='isShowResetWin' :title="'恢复出厂设置'">
      <p class='recoverWin-tip'>将网关恢复到出厂时的默认状态，所有配置将会失效。</p>
      <div class='recoverWin-con clearfix'>
        <p class='recoverWin-con-password'><span>密码</span> <Input class="input" v-model="resetpassword" :type="isShowPassword? 'text': 'password'"/></p>   
        <p class="Checkbox">
          <Checkbox v-model="isShowPassword">显示密码</Checkbox>
        </p>
        <g-button class="fr" type="primary" :unusable = "isResetUnusableBtn" @click="reset">确定</g-button>
        <g-button class="btn-cancel fr" type="default" @click="isShowResetWin = false"> 取消</g-button>
      </div>
    </win-common>

    <win-common class="updateWin" :width = 600 v-if='isShowUpdateWin' :title="'升级固件'">
      <p class="updateWin-tip">
        对当前网关进行更新操作，操作后，系统将选择的固件版本升级到当前网关中。
      </p>
      <div class="updateWin-con">
        <span class="updateWin-con-text">当前版本</span> 
        <span class="updateWin-con-input">GeLinkIII19v1.0.0</span>
        <span class="updateWin-con-text">更新为</span> 
        <span class="updateWin-con-input">GeLinkIII19v1.0.0</span>
      </div>
      <div class="updateWin-btn">
        <g-button class="fr" type="primary"  @click="reset">确认更新</g-button>
        <g-button class="btn-cancel fr" type="default" @click="isShowUpdateWin = false"> 取消</g-button>
      </div>
    </win-common>

  </div>
</template>

<script lang="ts">
import { Component, Watch, Vue } from 'vue-property-decorator';
import TitleBar from '@/components/title-bar/title-bar.vue';
import WinTip from '@/components/windows/win-tip.vue';
import WinTipSecond from '@/components/windows/win-tip-second.vue';
import WinCommon from '@/components/windows/win-common.vue';
import GButton from '@/components/button/g-button.vue';

import {
  pingTestApi,
  tracertTestApi,
  nslookupTestApi,
  rebootApi,
  getRebootStateApi,
  resetApi,
} from '@/api/system-set/system-set';
import ModifyPasswordForm from '@/components/modify-password/modify-password-form.vue';

@Component({
  components: {
    TitleBar,
    WinTip,
    WinTipSecond,
    WinCommon,
    GButton,
    ModifyPasswordForm,
  },
})
export default class SystemSet extends Vue {
  private tomographyType: string = '';
  private tomography: any = {
    ping: '',
    traceroute: '',
    nsloopup: '',
    pingError: false,
    tracerouteError: false,
    nsloopupError: false,
    result: '',
    loading: '',
  };
  private isRebootNow: boolean = false; // 正在重启
  private isRebootSuccess: boolean = false; // 重启成功
  private isResetNow: boolean = false; // 正在恢复出厂设置
  private isResetSuccess: boolean = false; // 恢复出厂设置成功
  private resetpassword: string = ''; // 恢复出厂设置-密码
  private isShowResetWin: boolean = false; // 恢复出厂设置弹窗
  private isShowPassword: boolean = false; // 恢复出厂设置显示密码
  private isResetUnusableBtn: boolean = true; // 恢复出厂设置不可点
  private isShowRebootWin: boolean = false; // 重启二次确认弹框
  private isShowUpdateWin: boolean = true;
  private isUpdateUnusableBtn: boolean = true;



  @Watch('resetpassword')
  public onResetPasswordChanged(val: string, oldVal: string) {
    this.isResetUnusableBtn = val.length < 5 ? true : false;
  }

  // 网络诊断
  private handleTomography(type: string): void {
    if (this.checkTomographyItem(type)) {
      this.tomography.result = '';
      const apiObj: any = {
        ping: pingTestApi,
        traceroute: tracertTestApi,
        nsloopup: nslookupTestApi,
      };

      this.tomographyType = type;
      this.tomography.loading = type;
      apiObj[type](this.tomography[type]).then((res: any) => {
        this.tomography.result = res.data;
        this.tomography.loading = '';
      });
    }
  }

  // 校验网络诊断
  private checkTomographyItem(type: string) {
    let isValid: boolean = true;
    if (!this.tomography[type]) {
      this.tomography[type + 'Error'] = true;
      isValid = false;
    } else {
      this.tomography[type + 'Error'] = false;
      isValid = true;
    }

    return isValid;
  }

  // 重启
  private reboot(): void {
    this.isShowRebootWin = false;
    this.isRebootNow = true;
    rebootApi().then((res: any) => {
      if (res.status === 200) {
        const timer = setInterval(() => {
          getRebootStateApi().then((rebootRes: any) => {
            if (rebootRes.status === 200) {
              this.isRebootNow = false;
              clearInterval(timer);
              if (!this.isRebootSuccess) {
                this.isRebootSuccess = true;
              }
              setTimeout(() => {
                this.isRebootSuccess = false;
              }, 2000);
            }
          });
        }, 3000);
      } else {
        this.isRebootNow = false;
        this.$Message.error('重启失败');
      }
    });
  }


  // 导入配置文件
  private fileChange(event: any) {
    const file = event.target.files[0];
    if (file) {
      const fileName = file.name;
      const formData = new FormData();
      formData.append('systemConfig', file);
      // importSystemConfigApi(formData).then((res: any) => {
      //   if (res.status === 200) {
      //     this.currentFileName = fileName;
      //     console.log(this.currentFileName);
      //     this.isImportSuccess = true;
      //     const timer = setTimeout(() => {
      //       this.isImportSuccess = false;
      //       this.queryPortInfo();
      //     }, 3000);
      //   } else {
      //     this.isImportFail = true;
      //     const timer = setTimeout(() => {
      //       this.isImportFail = false;
      //     }, 3000);
      //   }
      // });
    }
  }

  // 恢复出厂设置
  private reset(): void {
    if (!this.isResetUnusableBtn) {
      this.isShowResetWin = false;
      this.isResetNow = true;
      resetApi({password: this.resetpassword, username: sessionStorage.getItem('userName')}).then((res: any) => {
        this.isResetNow = false;
        if (res.status === 200) {
          this.isResetSuccess = true;
          setTimeout(() => {
            sessionStorage.removeItem('userName');
            sessionStorage.removeItem('hasShowModifyPasswordModal');
            this.isResetSuccess = false;
            this.$router.push({
              name: 'login',
            });
          }, 3000);
        }
      });
    }
  }
}
</script>

<style lang="less" scoped>
@font-color: #333333;
.system-set {
  color: @font-color;
  .con {
    padding: 0 100px;
  }
  .system-set-config {
    padding-top: 30px;
    .con {
      display: flex;
      justify-content: space-between;
      padding-top: 30px;
      padding-bottom: 53px;
      .network-tomography-item {
        display: flex;
        position: relative;
        flex-direction: column;
        justify-content: flex-start;
        .tomography-input {
          width: 220px;
          margin-bottom: 20px;
          &.selected {
            .ivu-input {
              border-color: #57a3f3;
            }
          }
        }
        .error-text {
          position: absolute;
          left: 1px;
          top: 32px;
          color: #F4333C;
        }
      }
    }
    .tomography-result {
      padding: 0 100px 30px 100px;
    }
  }

  .system-set-update {
    .con {
      .con-tip {
        height: 74px;
        line-height: 74px;
        border-bottom: 1px solid #888888;
      }
      .con-file {
        position: relative;
        .con-file-name {
          position: relative;
          top: 12px;
          display: inline-block;
          width: 180px;
          height: 32px;
          line-height: 30px;
          background: #fff;
          padding: 0 15px;
          margin-right: 10px;
          border: 1px solid #888888;
          border-radius: 3px;
        }
        .con-file-input {
          position: absolute;
          left: 202px;
          top: 15px;  
          width: 80px;
          height: 30px;
          border: 1px solid red;        
          opacity: 0;
        }
        .con-file-btn {
          margin-right: 20px;
          margin-left: 10px;
        }
      }
      .btn {
        margin: 30px 0;
      }
    }
  }

  .system-set-reboot {
    .con {
      .con-tip {
        height: 74px;
        line-height: 74px;
        border-bottom: 1px solid #888888;
      }
      .btn {
        margin: 30px 0;
      }
    }
  }
  .system-set-recover {
    .con {
      .con-tip {
        height: 74px;
        line-height: 74px;
        border-bottom: 1px solid #888888;
      }
      .btn {
        margin: 30px 0;
      }
    }
  }

  .system-set-password {
    .con {
      padding-top: 20px;
    }
  }

  .recoverWin {
    .recoverWin-tip {
      padding: 10px 0px 0px 76px;
      text-align: left;
    }
    .recoverWin-con {
      padding: 10px 89px 20px 76px;
      .recoverWin-con-password {
        span {
          margin-right: 10px;
        }
        .input {
          width:250px;
          height: 30px;
        }
      }
      .Checkbox {
        text-align: left;
        padding-left: 35px;
        padding-top: 10px;
        line-height: 20px;
      }

      .btn-cancel {
        margin-right: 14px;
      }
    }
  }

  .rebootWin {
    .rebootWin-tip {
      padding: 30px 0px 34px 76px;
      text-align: left;
    }
    .rebootWin-con {
      padding: 10px 89px 20px 76px;
      .btn {
        width: 70px;
      }
      .btn-cancel {
        margin-right: 14px;
      }
    }
  }

  .updateWin {
  
    .updateWin-tip {
      padding: 20px 0;
      text-align: center;
    }
    .updateWin-con {
      padding: 20px 0;  
      .updateWin-con-text {
        padding-right: 10px;
      }    
      .updateWin-con-input {
        display: inline-block;
        width: 180px;
        height: 32px;
        line-height: 30px;
        background: #fff;
        padding: 0 15px;
        margin-right: 10px;
        border: 1px solid #888888;
        border-radius: 3px;        
      }
    }
    .updateWin-btn {
      padding-right: 30px;
      height: 60px;
      .btn-cancel {
        margin-right: 20px;
      }
    }
  }
}
</style>





