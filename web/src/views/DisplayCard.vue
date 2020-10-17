<template>
  <div>
    <el-row>
      <el-col :span="8">
        <!-- <div class="grid-content bg-purple"> -->
        <div class="my-card">
          <!-- 即将过期资源展示 -->
          <DelayCard />
          <!-- 即将过期资源展示 -->
          <AlarmCard />
        </div>
      </el-col>
      <el-col :span="16">
        <div class="my-card">
          <!-- 资源显示组件 -->
          <router-view></router-view>
          <!-- <ResourceCard /> -->
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
// import ResourceCard from "@/components/ResourceCard.vue";
import DelayCard from "@/components/DelayCard.vue";
import AlarmCard from "@/components/AlarmCard.vue";
export default {
  name: "DisplayCard",
  components: { DelayCard, AlarmCard },
};
</script>
<style lang="less" scoped>
.el-col {
  padding: 1px;
  padding-top: 10px;
}
.my-card {
  height: 1000px;
}
</style>