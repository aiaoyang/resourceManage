<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="flush">刷新</el-button>
      <el-table style="width: 100%" :data="delays">
        <el-table-column label="7日内过期" prop="seven"> </el-table-column>
        <el-table-column label="30日内过期" prop="thirty"> </el-table-column>
        <el-table-column label="已过期" prop="expired"> </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script>
import { GetECS, GetRDS } from "@/http/resources.js";
export default {
  name: "DelayCard",
  data() {
    return {
      delays: [],
    };
  },
  methods: {
    flush() {
      this.delays = [
        {
          seven: 0,
          thirty: 0,
          expired: 0,
        },
      ];
      GetECS().then((res) => {
        console.log(res);
        let content = JSON.parse(res.data.msg);
        let thirty = 0;
        let seven = 0;
        for (let i = 0; i < content.length; i++) {
          if (content[i].status == "1") {
            thirty++;
          }
          if (content[i].status == "2") {
            seven++;
          }
        }
        this.delays[0].seven += seven;
        this.delays[0].thirty += thirty;
      });
      GetRDS().then((res) => {
        console.log(res);
        let content = JSON.parse(res.data.msg);
        let thirty = 0;
        for (let i = 0; i < content.length; i++) {
          if (content[i].status == "1") {
            thirty++;
          }
        }
        this.delays[0].thirty += thirty;
      });
    },
  },
  created() {
    this.flush();
  },
};
</script>
<style lang="less" scoped>
.box-card {
  height: 200px;
}
.el-button {
  float: right;
}
</style>