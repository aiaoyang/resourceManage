<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="flushRDS" style="text-align: left"
        >刷新</el-button
      >
      <el-table :data="resources" style="width: 100%">
        <el-table-column label="ID" prop="index"> </el-table-column>
        <el-table-column label="资源类型" prop="type"> </el-table-column>
        <el-table-column label="资源名称" prop="name"> </el-table-column>
        <el-table-column label="资源规格" prop="size"> </el-table-column>
        <el-table-column label="到期日" prop="end"> </el-table-column>
        <el-table-column label="所属账号" prop="account"> </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script>
import { GetRDS } from "@/http/resources.js";
export default {
  name: "RDS",
  data() {
    return {
      resourceLabel: [],
      resources: [],
    };
  },
  methods: {
    flushRDS: function () {
      let tmp = "";
      GetRDS().then((res) => {
        console.log(res.data);
        let js = JSON.parse(res.data.msg);
        console.log(js);
        this.resources = js;
      });
      console.log(tmp.toString());
    },
  },
  created() {
    this.flushRDS();
  },
};
</script>
<style lang="less" scoped>
.box-card {
  height: 100%;
}
</style>