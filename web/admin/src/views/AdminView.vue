<template>
  <a-layout class="container">
    <a-layout-sider class="sider"><NavView></NavView>
    </a-layout-sider>
    <a-layout>
      <a-layout-header>
        <HeaderView></HeaderView>
      </a-layout-header>
      <a-layout-content>
        <IndexView></IndexView>
      </a-layout-content>
      <a-layout-footer>
        <FooterView></FooterView>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import FooterView from '../components/admin/FooterView.vue'
import HeaderView from '../components/admin/HeaderView.vue'
import IndexView from '../components/admin/IndexView.vue'
import NavView from '../components/admin/NavView.vue'
export default {
  components: {
    FooterView,
    HeaderView,
    IndexView,
    NavView
  }
}
</script>
<style scoped>
.container {
  height: 100%;
}
</style>
