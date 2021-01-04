<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="flushECS" style="text-align: left"
        >刷新</el-button
      >
      <el-table :data="resources" style="width: 100%" sort="status">
        <el-table-column label="ID" prop="index" sortable="true">
        </el-table-column>

        <!-- <el-table-column label="状态" prop="status" sortable="true">
          <template scope="scope">
            <span v-if="scope.row.status == '0'" class="ok-span">good</span>
            <span v-if="scope.row.status == '1'" class="warning-span">
              warning
            </span>
            <span v-if="scope.row.status == '2'" class="danger-span">
              danger
            </span>
            <span v-if="scope.row.status == '3'" class="fatal-span">fatal</span>
          </template>
        </el-table-column> -->
        <el-table-column label="资源类型" prop="type"> </el-table-column>
        <el-table-column label="资源名称" prop="name"> </el-table-column>
        <el-table-column label="资源规格" prop="detail"> </el-table-column>
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
import { GetECS } from "@/http/resources.js";
export default {
  name: "ECS",
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

    // SetStatus({ row, column, rowIndex, columnIndex }) {
    //   console.log(row, column, rowIndex, columnIndex);
    //   return "cell-red";
    // },
  },
  created() {
    this.flushECS();
  },
  computed: {},
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