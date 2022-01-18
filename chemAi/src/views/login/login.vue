<template>
  <div class="login">
    <div class="login-header">
      <img src="@/assets/images/logo.png" alt="">
      <span class="title">格创东智科技有限公司</span>
    </div>
    <div class="login-con">
      <div class="login-form-wrap">
        <div class="form-title">
          <span class="cycle left-cycle"></span>
          <p>欢迎使用格创东智网关</p>
          <span class="cycle right-cycle"></span>
        </div>
        <div class="form-con">
          <Form ref="loginForm" :model="form" @keydown.enter.native="handleSubmit">
            <FormItem prop="userName">
              <i-input v-model="form.userName" placeholder="请输入用户名" @on-blur="checkItem('userName')" />
              <p v-show="check.userName.show" :class="{'input-error': true, 'input-error-show': check.userName.show}">{{check.userName.text}}</p>
            </FormItem>
            <FormItem prop="password">
              <i-input type="password" v-model="form.password" placeholder="请输入登录密码" @on-blur="checkItem('password')" />
              <p v-show="check.password.show" :class="{'input-error': true, 'input-error-show': check.password.show}">{{check.password.text}}</p>
              <p class="forget-password" @click="openSuperPassword">忘记密码</p>
            </FormItem>
            <FormItem>
              <g-button @click="handleSubmit" type="primary" size="large" long :disabled="loginDisabled">登录</g-button>
            </FormItem>
          </Form>
        </div>
      </div>
    </div>
    <!-- 超级密码 -->
    <div class="modify-password" v-if="modalShow">
      <reset-password-modal
        :show="modalShow"
        :username="this.form.userName"
        @cancel-modify-password="modalShow=false"
      >
      </reset-password-modal>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import GButton from '@/components/button/g-button.vue';
import ResetPasswordModal from '@/components/modify-password/reset-password-modal.vue';
import config from '@/config';
import { loginApi } from '@/api/login/login';

@Component({
  components: {
    GButton,
    ResetPasswordModal,
  },
})
export default class Login extends Vue {
  private timers: any = [];
  private form: any = {
    userName: 'admin',
    password: '',
  };
  private check: any = {
    list: ['userName', 'password'],
    userName: {
      show: false,
      text: '',
    },
    password: {
      show: false,
      text: '',
    },
  };
  private loginDisabled: boolean = false;
  private modalShow: boolean = false;
  private errorCnt: number = 0;
  private data() {
    return {
      form: {
        userName: 'admin',
        password: '',
      },
    };
  }
  private checkItem(type: string) {
    let isValid: boolean = true;
    switch (type) {
      case 'userName':
        if (!this.form.userName) {
          isValid = false;
          this.showError(type, '请输入用户名');
        }
        break;
      case 'password':
        if (!this.form.password) {
          isValid = false;
          this.showError(type, '请输入密码');
        } else {
          const reg = /^[0-9a-zA-Z]*$/;
          if (!reg.test(this.form.password)) {
            isValid = false;
            this.showError(type, '密码格式错误');
          }
        }
        break;
    }
    return isValid;
  }
  private showError(type: string, text: string) {
    this.check[type].show = true;
    this.check[type].text = text;
    const timer = setTimeout(() => {
      this.check[type].show = false;
      this.check[type].text = '';
    }, 3000);
    this.timers.push(timer);
  }
  private checkAll() {
    let isValid: boolean = true;
    this.check.list.forEach((item: string) => {
      if (!this.checkItem(item)) {
        isValid = false;
      }
    });
    return isValid;
  }
  private handleSubmit() {
    if (this.checkAll()) {
      this.loginDisabled = true;
      loginApi({username: this.form.userName, password: this.form.password}).then((res: any) => {
        this.loginDisabled = false;
        if (res.status === 200) {
          sessionStorage.setItem('userName', this.form.userName);
          let shouldModifyPassword: boolean = false;
          if (this.form.password === 'admin' && this.form.userName === 'admin') {
            shouldModifyPassword = true;
          }
          this.$router.push({
            name: config.homeName,
            query: {show_modal: shouldModifyPassword ? 'true' : 'false'},
          });
        } else {
          this.errorCnt ++;
          if (this.errorCnt > 5) {
            this.showError('password', '请确定后再输入或者重置密码');
          }
        }
      }, (error) => {
        this.loginDisabled = false;
      });
    }
  }
  private openSuperPassword() {
    this.modalShow = true;
  }

  private mounted() {
    sessionStorage.removeItem('userName');
  }
}
</script>
<style lang="less" scoped>
.login {
  height: 100vh;
  background: #2A2C35;
  .login-header {
      width: 100%;
      height: 60px;
      display: flex;
      justify-content: flex-start;
      padding: 0 30px;
      align-items: flex-end;
    .title {
      font-size: 12px;
      color: #fff;
      font-weight: 400;
      padding-left: 10px;
    }
  }
  .login-con {
    height: ~'calc(100% - 120px)';
    display: flex;
    justify-content: center;
    align-items: center;
    .login-form-wrap {
      width: 400px;
      height: 380px;
      background:rgba(255,255,255,1);
      box-shadow:0px 0px 6px 0px rgba(42,45,55,0.74);
      border-radius:5px;
      .form-title {
        background: #4D9DED;
        height: 140px;
        line-height: 140px;
        text-align: center;
        border-top-left-radius: 5px;
        border-top-right-radius: 5px;
        position: relative;
        overflow: hidden;
        .cycle {
            display: inline-block;
            position: absolute;
            width: 140px;
            height: 140px;
            background: #D7EBFF;
            opacity: 0.2;
            border-radius: 50%;
          &.left-cycle {
            left: -49px;
            bottom: -47px;
          }
          &.right-cycle {
            top: -64px;
            right: -21px;
          }
        }
        p {
          font-size: 24px;
          color: #fff;
          font-weight: 400;
        }
      }
      .form-con {
        padding: 24px 50px 0 50px;
      }
    }
  }
}
</style>
<style lang="less" scope>
.login {
  .form-con {
    .ivu-form-item {
      &:nth-of-type(2) {
        margin-bottom: 20px;
      }
      .ivu-input {
        border: 1px solid transparent;
        border-bottom: 1px solid #888;
        border-radius: 0;
        &:focus {
          box-shadow: none;
          border-bottom-color: #57a3f3;
        }
      }
      .ivu-btn-large {
        padding: 9px 15px 9px 15px;
      }
      .forget-password {
        color: #4D9DED;
        float: right;
        cursor: pointer;
        font-weight: 400;
      }
      .ivu-form-item-content {
        .input-error {
          position: absolute;
          text-align: left;
          display: inline-block;
          left: 0;
          top: 34px;
          color: #ed4014;
          height: 20px;
          line-height: 20px;
          width: 230px;
          overflow: hidden;
          text-overflow: ellipsis;
          &.input-error-show {
            animation: myErrorShow 3s forwards;
          }
          @keyframes myErrorShow {
            0% { opacity: 1; }
            50% { opacity: 1; }
            75% { opacity: 1; }
            100% { opacity: 0; }
          }
        }
      }
    }
  }
}
</style>
