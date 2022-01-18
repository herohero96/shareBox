<template>
  <Form
    ref="passwordForm"
    :model="form"
    @keydown.enter.native="handleSubmit"
    :label-width="showLabel === true ? 74 : 0"
    label-position="left"
    :footer-hide="true"
    class="modify-password-form"
    :class="showLabel ? 'show-label-form': 'hide-label-form'"
  >
    <FormItem
      :label="showLabel === true && type === 'modify' ? '默认密码'
      : showLabel === true && type === 'reset' ? '超级密码' : ''"
      class="form-item-required"
    >
      <Input
        type="password"
        v-model="form.oldPassword"
        @on-blur="checkItem('oldPassword')"
        :placeholder="type === 'page' ? '当前密码' : ''"
      />
      <p v-show="check.oldPassword.show" :class="{'input-error': true, 'input-error-show': check.oldPassword.show}">{{check.oldPassword.text}}</p>
    </FormItem>
    <FormItem :label="showLabel === true ? '新密码' : ''" class="form-item-required" >
      <Input
        :type="!passwordShow ? 'password' : 'text'"
        v-model="form.newPassword"
        @on-blur="checkItem('newPassword')"
        style="width: 249px;"
        :placeholder="type === 'page' ? '新密码' : ''"
      />
      <p v-show="check.newPassword.show" :class="{'input-error': true, 'input-error-show': check.newPassword.show}">{{check.newPassword.text}}</p>
      <a class="switch-password-status" @click="handleSwitchPasswordStatus">
        <i class="iconfont icon-hide" :class="!passwordShow ? 'icon-hide' : 'icon-display'"></i>
      </a>
    </FormItem>
    <FormItem :label="showLabel === true ? '再次输入' : ''" class="form-item-required" >
      <Input
        type="password"
        v-model="form.againPassword"
        @on-blur="checkItem('againPassword')"
        :placeholder="type === 'page' ? '再次输入' : ''"
      />
      <p v-show="check.againPassword.show" :class="{'input-error': true, 'input-error-show': check.againPassword.show}">{{check.againPassword.text}}</p>
    </FormItem>
    <FormItem :class="type === 'page' ? 'btn-align-left' : 'btn-align-right'">
      <g-button v-if="type !== 'page'" type="default" width="84" @click="cancel">取消</g-button>
      <g-button v-if="type !== 'page'" type="primary" width="84" @click="handleSubmit" style="margin-left: 8px;">{{type === 'modify' ? '更改密码' : '确认'}}</g-button>
      <g-button v-if="type==='page'" type="default" width="84" @click="handleSubmit">确定</g-button>
    </FormItem>
  </Form>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import GButton from '@/components/button/g-button.vue';
import { changePasswordApi } from '@/api/login/login';

@Component({
  components: {
    GButton,
  },
})
export default class ModifyPasswordForm extends Vue {
  /**
   * page 在页面中使用
   * modify 修改密码弹框
   * reset 超级密码重置弹框
   */
  @Prop({ type: String, default: 'modal'}) private type!: string;
  @Prop({ type: Boolean, default: true }) private showLabel!: boolean;
  @Prop({ type: String, default: 'admin' }) private username!: string;
  private timers: any = []; // 计时器
  private passwordShow: boolean = false; // 是否显示密码
  private form: any = { // Form 表单数据
    oldPassword: '',
    newPassword: '',
    againPassword: '',
  };
  private check: any = { // 表单校验
    list: ['oldPassword', 'newPassword', 'againPassword'],
    oldPassword: {
      show: false,
      text: '',
    },
    newPassword: {
      show: false,
      text: '',
    },
    againPassword: {
      show: false,
      text: '',
    },
  };
  private data() {
    return {
      passwordShow: false,
      form: {
        oldPassword: '',
        newPassword: '',
        againPassword: '',
      },
    };
  }

  // methods
  private handleSwitchPasswordStatus() {
    this.passwordShow = !this.passwordShow;
  }
  private handleSubmit() {
    if (this.checkAll()) {
      const param = {
        newPassword: this.form.newPassword,
        oldPassword: this.form.oldPassword,
        repeatPassword: this.form.againPassword,
        username: this.username,
      };
      changePasswordApi(param).then((res: any) => {
        if (res.status === 200) {
          this.$Message.info('修改成功!');
          setTimeout(() => {
            this.cancel();
            sessionStorage.removeItem('userName');
            sessionStorage.removeItem('hasShowModifyPasswordModal');
            this.$router.push({
              name: 'login',
            });
          }, 2000);
        }
      });
    }
  }
  private cancel() {
    this.$emit('cancel-modify-password');
  }
  private checkItem(type: string) { // 校验
    let isValid: boolean = true;
    switch (type) {
      case 'oldPassword':
        if (!this.form.oldPassword) {
          isValid = false;
          this.showError(type, '请输入密码');
        }
        break;
      case 'newPassword':
        if (!this.form.newPassword) {
          isValid = false;
          this.showError(type, '请输入新密码');
        } else {
          const reg = /^[0-9a-zA-Z]{6,}$/;
          if (!reg.test(this.form.newPassword)) {
            isValid = false;
            this.showError(type, '密码由不少于6位的字母和数字组成');
          }
        }
        break;
      case 'againPassword':
        if (!this.form.againPassword) {
          isValid = false;
          this.showError(type, '请再次输入新密码');
        } else {
          if (this.form.newPassword !== this.form.againPassword) {
            isValid = false;
            this.showError(type, '两次输入密码不一致');
          }
        }
        break;
    }
    return isValid;
  }
  private showError(type: string, text: string) { // 显示错误信息
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

  // life-cycle
  private destroyed() {
    this.timers.forEach((item: any) => {
      clearTimeout(item);
    });
  }
}
</script>
<style lang="less" scope>
.modify-password-form {
  &.show-label-form {
    width: 364px;
  }
  &.hide-label-form {
    width: 290px;
  }
  .ivu-form-item {
    margin-bottom: 20px;
    &.form-item-required {
      .ivu-form-item-label {
        &:before {
          content: '*';
          display: inline-block;
          margin-right: 4px;
          line-height: 1;
          font-family: SimSun;
          font-size: 12px;
          color: #ed4014;
        }
      }
    }
    &:last-child {
      padding: 6px 0 21px 0;
      margin-bottom: 0;
    }
    &.btn-align-left {
      text-align: left;
    }
    &.btn-align-right {
      text-align: right;
    }
    .ivu-form-item-content {
      .ivu-input {
        border: 1px solid #888;
        &:hover {
          border-color: #57a3f3;
        }
      }
      .input-error {
        position: absolute;
        text-align: left;
        display: inline-block;
        left: 50%;
        transform: translateX(-50%);
        top: 34px;
        color: #ed4014;
        height: 20px;
        line-height: 20px;
        width: 100%;
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
      .switch-password-status {
        background: #fff;
        width: 30px;
        display: inline-block;
        height: 30px;
        border: 1px solid #888;
        border-radius: 3px;
        cursor: pointer;
        padding: 2px;
        color: #888;
        text-align: center;
        line-height: 26px;
        margin-left: 10px;
        vertical-align: middle;
        transition: all .2s ease-in-out;
        .iconfont {
          font-size: 12px;
        }
        &:hover {
          border-color: #57a3f3;
          color: #57a3f3;
        }
      }
    }
  }
}
</style>

