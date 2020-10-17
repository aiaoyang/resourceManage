<template>
  <div>
    <el-menu
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      :router="true"
    >
      <!-- <el-submenu index="1"> -->
      <!-- 子菜单栏 -->
      <template slot="title">资源</template>
      <el-menu-item
        v-for="menu in allResource"
        :key="menu"
        :index="menu"
        @click="pushRouter(menu)"
      >
        {{ menu }}
      </el-menu-item>
      <!-- </el-submenu> -->
    </el-menu>
  </div>
</template>
<script>
import { GetMenu } from "@/http/menu.js";
export default {
  name: "Resource",
  data: function () {
    return {
      allResource: [],
    };
  },
  created: function () {
    this.allResource = GetMenu();
  },
  methods: {
    flushMenu() {
      this.allResource = GetMenu();
    },
    pushRouter(path) {
      // router.push
      console.log(path);
    },
  },
};
</script>