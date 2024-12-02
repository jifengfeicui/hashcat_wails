<script setup>

import {Add, FileTrayFullOutline} from "@vicons/ionicons5";
import {Cat} from "@vicons/fa";
import {AlignHorizontalLeftFilled} from "@vicons/material";
</script>

<template>
  <div class="sidebar">
    <div class="logo">
      <n-icon size="40">
        <Cat/>
      </n-icon>
    </div>
    <div class="nav">
      <ul>
        <li>
          <router-link to="/add">
            <n-icon size="40">
              <Add/>
            </n-icon>
            <span class="nav-text">添加</span>
          </router-link>
        </li>
        <li>
          <router-link to="/file">
            <n-icon size="40">
              <FileTrayFullOutline/>
            </n-icon>
            <span class="nav-text">文件列表</span>
          </router-link>
        </li>
        <li>
          <router-link to="/dict">
            <n-icon size="40">
              <AlignHorizontalLeftFilled/>
            </n-icon>
            <span class="nav-text">密码字典</span>
          </router-link>
        </li>
      </ul>
    </div>
  </div>

  <div class="main-content">
    <div class="header">
      <h1>Welcome to HashCat Desktop</h1>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>
/* Sidebar styling */
.sidebar {
  width: 150px;
  background-color: #2c3e50;
  color: white;
  padding: 20px;
  transition: all 0.3s;
}

.sidebar.collapsed {
  width: 70px; /* When collapsed */
}

.sidebar .logo {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 20px;
}

/* Navigation list */
.sidebar .nav ul {
  list-style: none;
  padding: 0;
}

.sidebar .nav ul li {
  display: flex;
  align-items: center;
  margin: 15px 0;
}

.sidebar .nav ul li a {
  display: flex;
  align-items: center;
  color: white;
  text-decoration: none;
  width: 100%;
}

.sidebar .nav ul li a .nav-text {
  margin-left: 10px; /* Space between icon and text */
}

/* Text for collapsed sidebar */
.sidebar.collapsed .nav-text {
  display: none; /* Hide text when sidebar is collapsed */
}

.main-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  width: 100%;
  //background-color: #ecf0f1;
}

.main-content .header {
  background-color: #34495e;
  color: white;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-content .header h1 {
  margin: 0;
}

.main-content .content {
  padding: 20px;
  flex-grow: 1;
}
</style>