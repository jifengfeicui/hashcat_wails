import {createApp} from 'vue'
import naive from 'naive-ui'
import App from './App.vue'
import router from "./router";


const app = createApp(App)

app.use(naive).use(router)

app.mount('#app')