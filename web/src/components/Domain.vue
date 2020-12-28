<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="flushDomain" style="text-align: left"
        >刷新</el-button
      >
      <el-table :data="resources" style="width: 100%">
        <el-table-column label="ID" prop="index"> </el-table-column>
        <el-table-column label="资源类型" prop="type"> </el-table-column>
        <el-table-column label="资源名称" prop="name"> </el-table-column>
        <el-table-column label="到期日" prop="end">
          <template slot-scope="scope">
            <span v-if="scope.row.status == '0'" class="ok-span">
              {{ scope.row.end }}
            </span>
            <span v-if="scope.row.status == '1'" class="warning-span">
              {{ scope.row.end }}
            </span>
            <span v-if="scope.row.status == '2'" class="danger-span">
              {{ scope.row.end }}
            </span>
            <span v-if="scope.row.status == '3'" class="fatal-span">
              {{ scope.row.end }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="所属账号" prop="account"> </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import { GetDomain } from "@/http/resources.js";
export default {
  name: "Domain",
  data() {
    return {
      resourceLabel: [],
      resources: [],
    };
  },
  methods: {
    flushDomain: function () {
      let tmp = "";
      GetDomain().then((res) => {
        console.log(res.data);
        let js = JSON.parse(res.data.msg);
        console.log(js);
        this.resources = js;
      });
      console.log(tmp.toString());
    },
  },
  created() {
    this.flushDomain();
  },
};
</script>
<style lang="less" scoped>
.box-card {
  height: 100%;
}
.warning-span {
  background-color: yellow;
  color: black;
}
.ok-span {
  background-color: rgb(1, 255, 1);
  color: black;
}
.danger-span {
  background-color: orange;
  color: black;
}
.fatal-span {
  background-color: red;
  color: black;
}
</style>