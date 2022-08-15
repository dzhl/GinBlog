<template>
  <div class="container">
    <div class="loginbox">
      <a-form-model ref="formdata" class="loginForm" :model="formdata" :rules="rules">
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="输入用户名" style="border-radius:5px">
            <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input v-model="formdata.password" placeholder="输入密码" type="password">
            <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
          </a-input>
        </a-form-model-item>
        <a-form-model-item class="loginBtn">
          <a-button type="primary" style="margin:20px" @click="onSubmit">登录</a-button>
          <a-button type="info" @click="resetForm">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 20, message: '用户名长度必须至少3位，且不得超过20位', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度必须至少6位，且不得超过20位', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    onSubmit () {
      this.$refs.formdata.validate(async valid => {
        if (valid) {
          const { data: res } = await this.$http.post('user/login', this.formdata)
          if (res.status !== 200) {
            this.$message.error(res.msg)
            console.log(res)
            return
          }
          window.localStorage.setItem('token', res.token)
          this.$message.success(res.msg)
          this.$router.push('admin')
        } else {
          console.log('无效格式,请重新输入')
          this.$message.error('无效格式,请重新输入')
          return false
        }
      })
    },
    resetForm () {
      this.$refs.formdata.resetFields()
    }
  }
}
</script>

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

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 5%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>
