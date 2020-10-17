<template>
  <div>
    <el-card class="box-card">
      <el-table style="width: 100%" :data="delays">
        <el-table-column label="7日内过期" prop="seven"> </el-table-column>
        <el-table-column label="30日内过期" prop="thirty"> </el-table-column>
        <el-table-column label="已过期" prop="delay"> </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script>
import { GetDelay } from "@/http/delay.js";
export default {
  name: "DelayCard",
  data() {
    return {
      delays: [],
    };
  },
  methods: {},
  created() {
    this.delays = GetDelay();
  },
};
</script>
<style lang="less" scoped>
.box-card {
  height: 200px;
}
</style>