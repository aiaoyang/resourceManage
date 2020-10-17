<template>
  <div>
    <Menu />
  </div>
</template>
<script>
import Menu from "@/components/Menu.vue";
export default {
  name: "DisplayMenu",
  components: {
    Menu,
  },
};
</script>
<style lang="less" scoped>
.el-col {
  padding-top: 10px;
}
</style>