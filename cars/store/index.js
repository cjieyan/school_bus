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
		apihost:"http://127.0.0.1:8000",
		//apihost: "https://sc.kenashua.com",
		ak: "EtRc2ku5rpOb4aNTskgBr1yb",
		client_id:"oLPVBkl3gkURkuZPdN13XefG",
		client_secret: "qOsoDoVAkvotzLn4ismk4dMmDoNaUrim"
    },
    mutations: {
		changecantakephone(state, iscan){
			state.cantakephone = iscan
		},
		setstudent(state, data){
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