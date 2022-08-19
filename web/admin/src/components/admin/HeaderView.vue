<template>
    <div class="btn">
        <a-button type="danger" @click="logout">退出</a-button>
    </div>
</template>
<script>
export default {
  methods: {
    logout () {
      window.localStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
<style>
.btn {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    height: 100%;
}
</style>
