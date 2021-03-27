import "./assets/normalize.css"
import { createApp, h } from 'vue'
import App from './App.vue'
import store from './store'

import * as echarts from 'echarts';
import { plugin } from 'echarts-for-vue';


createApp(App).use(store).use(plugin, { echarts, h }).mount('#app');
