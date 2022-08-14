<template>
  <div class="container">
    <div class="loginbox">
      <a-form-model class="loginForm"
        ref="ruleForm"
        :model="ruleForm"
        :rules="rules"
        v-bind="layout"
      >
        <a-form-model-item>
          <a-input></a-input>
        </a-form-model-item>
        <a-form-model-item></a-form-model-item>
        <a-input></a-input>
        <a-form-model-item></a-form-model-item>
        <a-button type="primary">登录</a-button>
        <a-button type="info">取消</a-button>
      </a-form-model>
    </div>
  </div>
</template>
<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}
.loginbox {
  width: 450px;
  height: 300px;
  background-color: rgb(22, 67, 150);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 20px;
}
.loginForm{
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}
</style>
