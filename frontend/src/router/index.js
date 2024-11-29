import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import Add from '../views/Add.vue'
import Task from "../views/Task.vue";
import File from "../views/File.vue";

const routes = [
    {
        path: '/',  // 根路径
        redirect: '/add'  // 重定向到默认页面
    },
    {
        path: '/add',
        name: 'Add',
        component: Add
    },
    {
        path: '/file',
        name: 'File',
        component: File
    },
    {
        path: '/task',
        name: 'Task',
        component: Task
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
