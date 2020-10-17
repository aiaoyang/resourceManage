<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="flushRDS" style="text-align: left"
        >RDS</el-button
      >
      <el-button type="primary" @click="flushECS" style="text-align: left"
        >ECS</el-button
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
import {
  TestResource,
  TestResourceLabel,
  GetECS,
  GetRDS,
} from "@/http/resources.js";
export default {
  name: "ResourceCard",
  data() {
    return {
      resourceLabel: [],
      resources: [],
    };
  },
  methods: {
    flushECS: function () {
      // TODO: 更换为从后端获取得api函数
      let tmp = "";
      GetECS().then((res) => {
        console.log(res.data);
        let js = JSON.parse(res.data.msg);
        console.log(js);
        this.resources = js;
      });
      console.log(tmp.toString());
    },
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
    // TODO: 更换为从后端获取得api函数
    this.resourceLabel = TestResourceLabel();
    // TODO: 更换为从后端获取得api函数
    this.resources = TestResource();
  },
};
</script>
<style lang="less" scoped>
.box-card {
  height: 100%;
}
</style>