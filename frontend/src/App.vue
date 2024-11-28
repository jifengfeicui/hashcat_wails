<script setup>
import {Add} from '@vicons/ionicons5'
import {AlignHorizontalLeftFilled} from '@vicons/material'
import {Cat} from '@vicons/fa'
</script>

<template>
  <div id="app">
    <div class="sidebar" :class="{ collapsed: isSidebarCollapsed }">
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
            <router-link to="/task">
                      <n-icon size="40">
                        <AlignHorizontalLeftFilled />
                      </n-icon>
              <span class="nav-text">任务列表</span>
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
  </div>
</template>

<style>
#app {
  display: flex;
  height: 100vh;
  overflow: hidden;
  width: 100%;
}

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
