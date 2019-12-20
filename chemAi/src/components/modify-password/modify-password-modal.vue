<template>
  <!-- 首次登录修改密码 -->
  <Modal
    class-name="modify-password-modal vertical-center-modal"
    width="450" 
    v-model="modalShow"
    :mask="false"
    :closable="false"
    :mask-closable="false"
    title="为了系统安全，请修改默认密码"
  >
    <div class="content">
      <modify-password-form
        @cancel-modify-password="cancelModifyPassword"
        :showLabel="true"
        :type="'modify'"
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
  components: {
    ModifyPasswordForm,
  },
})
export default class ModifyPasswordModal extends Vue {
  @Prop(Boolean) private show!: boolean;
  @Prop(String) private username!: string;
  private passwordShow?: boolean = false;
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
      padding: 30px 43px 0 43px;
    }
    .ivu-modal-footer {
      border-top: none;
      padding: 0;
    }
  }
}
</style>

