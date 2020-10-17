<template>
  <div>
    <el-card class="box-card">
      <div>
        <!-- <h3>
          告警数
          <el-button type="primary" @click="flushAlarm">刷新</el-button>
        </h3>
        <h1>{{ this.alarmCount }}</h1> -->
        <el-button type="primary" @click="flushAlarm">刷新</el-button>

        <el-table :data="alarmInfo">
          <el-table-column label="账号名" prop="account"> </el-table-column>
          <el-table-column label="告警数" prop="alarmCount"> </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>
<script>
import { GetAlarm } from "@/http/alarm.js";
export default {
  name: "AlarmCard",
  data() {
    return {
      alarmCount: 0,
      alarmInfo: [
        {
          account: "yongshiwl",
          alarmCount: "1",
        },
      ],
    };
  },
  methods: {
    flushAlarm: function () {
      GetAlarm().then((res) => {
        if (res) {
          let content = JSON.parse(res.data.msg);

          console.log(content);
          this.alarmInfo = content;
        }
      });
    },
  },
  created() {
    this.flushAlarm();
  },
};
</script>
<style lang="less" scoped>
.el-button {
  float: right;
}
.box-card {
  height: 400px;
}
</style>