import Vue from 'vue'
import App from './App'

Vue.config.productionTip = false

App.mpType = 'app'

// 引入全局uView
import uView from 'uview-ui'
Vue.use(uView);

import store from './store'
Vue.prototype.$store = store

Vue.prototype.checkLogin = function(backpage, backtype) { //定义一个全局函数
	var token = uni.getStorageSync('token'); //用户 id, 
	console.log("login-token")
	console.log(token)
	if (token == '') {
		uni.showToast({
			icon: 'none',
			title: '请登陆',
			duration: 1000
		});
		setTimeout(function() {
			uni.redirectTo({
				// url: '/my/login?backpage=' + backpage + '&backtype=' + backtype,
				url: '/pages/my/login?backpage=' + backpage + '&backtype=' + backtype,
				fail: (err) => {
					console.log(err)
				}
			});
		}, 1500);
		return false
	}else{
		this.$store.commit("setToken", token)
		return true
	}
}

const app = new Vue({
	...App
})
app.$mount()
