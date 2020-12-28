import App from '@/App.vue'
import Domain from '@/components/Domain'
import ECS from '@/components/ECS'
import RDS from '@/components/RDS'
import Cert from '@/components/Cert'
import Home from '@/views/Home.vue'
import ElementUI from 'element-ui'
import "element-ui/lib/theme-chalk/index.css"
import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.config.productionTip = false

Vue.use(VueRouter)
Vue.use(ElementUI)


const routes = [
	{ path: "/", redirect: "/resources" },
	{
		path: "/resources", component: Home,
		children: [
			{ path: "/ecs", component: ECS },
			{ path: "/rds", component: RDS },
			{ path: "/domain", component: Domain },
			{ path: "/cert", component: Cert }
		]

	}
]

const router = new VueRouter({
	routes: routes,
	mode: "history",
})

router.mode = "history"
new Vue({
	router: router,
	render: h => h(App)
}).$mount('#app')

