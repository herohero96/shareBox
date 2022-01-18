<template>
  <!-- 使用超级密码重置密码 -->
  <Modal
    class-name="modify-password-modal vertical-center-modal"
    width="450" 
    v-model="modalShow"
    :closable="false"
    :mask-closable="false"
    title="重置密码"
  >
    <div class="content">
      <p class="sub-title">重置密码需要超级密码，重置后原密码将会失效。</p>
      <modify-password-form
        @cancel-modify-password="cancelModifyPassword"
        :showLabel="true"
        :type="'reset'"
        :username="username"
      >
      </modify-password-form>
    </div>
    <div slot="footer"></div>
  </Modal>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';
import ModifyPasswordForm from './modify-password-form.vue';
@Component({
  components: { ModifyPasswordForm },
})
export default class ResetPasswordModal extends Vue {
  @Prop(Boolean) private show!: boolean;
  @Prop(String) private username!: string;
  private passwordShow?: boolean;
  private modalShow!: boolean;
  private data() {
    return {
      passwordShow: false,
      modalShow: this.show,
    };
  }

  // methods
  private handleSwitchPasswordStatus() {
    this.passwordShow = !this.passwordShow;
  }

  @Emit()
  private cancelModifyPassword() {
    this.modalShow = false;
    // this.$emit('cancel-modify-password')
  }
}
</script>
<style lang="less" scope>
.modify-password-modal {
  .ivu-modal {
    .ivu-modal-content {
      .ivu-modal-header {
        .ivu-modal-header-inner {
          font-size: 16px;
          color: #333;
          text-align: center;
        }
      }
    }
    .ivu-modal-body {
      padding: 10px 43px 0 43px;
      .sub-title {
        padding-bottom: 20px;
        text-align: center;
      }
    }
    .ivu-modal-footer {
      border-top: none;
      padding: 0;
    }
  }
}
</style>

