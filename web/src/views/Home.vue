<template>
  <div>
    <!-- 菜单栏 -->
    <DiaplayMenu></DiaplayMenu>
    <!-- 卡片栏 -->
    <DisplayCard></DisplayCard>
  </div>
</template>
<script>
import DiaplayMenu from "@/views/DisplayMenu.vue";
import DisplayCard from "@/views/DisplayCard.vue";
export default {
  name: "Home",
  components: {
    DiaplayMenu,
    DisplayCard,
  },
};
</script>

<style lang="less" scoped>
.DisplayCard {
  padding-top: 10px;
}
</style>