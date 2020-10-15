import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
    state: {
        cantakephone: false,
		student: {},
		carInfo: {},
		liniInfo: {},
		teacher: {},
		apihost: "http://localhost:8000",
    },
    mutations: {
		changecantakephone(state, iscan){
			state.cantakephone = iscan
		},
		setstudent(state, data){
			console.log("-------setstudent-------")
			console.log(data)
			state.student = data
		},
		setcarinfo(state, data){
			state.carInfo = data
		},
		setLineinfo(state, data){
			state.liniInfo = data
		},
		setTeacher(state, data){
			state.teacher = data
		}
	},
    actions: {}
})
export default store