<template>
  <div class="login">
    <div class="login-pannel">
      <div class="login-img">
        <img
          style="height: 90px; width: 320px"
          src="@/assets/image/login-head-img.png"
          alt="edoc2"
        />
      </div>
      <div class="login-title">
        <div class="login-welcome-text">欢迎登录</div>
      </div>
      <div class="login-form">
        <el-form :rules="rules" :model="accountLogin" ref="formRef">
          <el-form-item class="formItem" prop="username">
            <template #label>
              <el-icon style="margin: 0 -20px" :size="20">
                <user></user>
              </el-icon>
            </template>
            <el-input
              class="login-form-input"
              v-model="accountLogin.username"
            ></el-input>
          </el-form-item>
          <el-form-item class="formItem" prop="password">
            <template #label>
              <el-icon style="margin: 0 -20px" :size="20"><lock /></el-icon>
            </template>
            <el-input
              class="login-form-input"
              type="password"
              :show-password="false"
              v-model="accountLogin.password"
            ></el-input>
          </el-form-item>
        </el-form>
      </div>
      <div>
        <el-button @click="loginAction" class="login-btn" type="primary"
          >登录</el-button
        >
      </div>
    </div>

    <footer>
      <a href="https://xxxx"
        >Copyright © 2007-2022 xxxxxx</a
      >
    </footer>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useStore } from 'vuex'
import { rules } from './account-config'
import { User, Lock } from '@element-plus/icons-vue'

const store = useStore()
const accountLogin = reactive({
  username: '',
  password: ''
})
const formRef = ref()

const loginAction = () => {
  formRef.value?.validate((valid: any) => {
    if (valid) {
      store.dispatch('login/accountLoginAction', { ...accountLogin })
    }
  })
}
</script>

<style lang="less">
.login {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  background: url('@/assets/image/login-bg.png');

  .login-pannel {
    width: 320px;
    height: 420px;
    border-radius: 20px;
    box-shadow: 5px 5px 5px rgba(95, 92, 92, 0.1),
      -5px -5px 5px rgba(95, 92, 92, 0.1);
    background-color: #fff;
    padding: 0 40px;

    .login-welcome-text {
      font-family: 'Microsoft YaHei', arial;
      font-size: 28px;
      margin-bottom: 35px;
      margin-top: 6px;
      color: #4e4e4e;
    }

    .login-btn {
      width: 100%;
      border-radius: 44px;
      margin-top: 27px;
      height: 47px;
    }

    .el-input__inner {
      border: none;
      border-bottom: 1px solid #eeeeee;
      outline: none;
      background-color: transparent;
      outline: medium;
      box-shadow: 0 0 0 !important;
      border-radius: 0;
      padding-left: 27px !important;
    }

    .el-input__inner.is-error {
      box-shadow: 0 0 0 !important;
    }

    .login-img {
      display: block;
      border: 0;
      margin: 0;
      padding: 0;
      max-width: 100%;
      margin-top: 24px;
    }

    .el-form-item__label {
      margin-top: 5px;
      padding: 0;
    }

    .el-form-item__label::before {
      display: none;
    }
  }
}
footer {
  clear: both;
  display: block;
  text-align: center;
  margin: 0px auto;
  position: absolute;
  bottom: 50px;
  width: 100%;

  a {
    color: #737679;
    text-decoration: none;
    font-family: 'Microsoft YaHei', arial;
    font-size: 12px;
  }
}
</style>
