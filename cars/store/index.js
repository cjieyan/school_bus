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
		apihost: "https://sc.kenashua.com",
		// apihost: "http://localhost:8000",
		ak: "EtRc2ku5rpOb4aNTskgBr1yb",
		client_id:"oLPVBkl3gkURkuZPdN13XefG",
		client_secret: "qOsoDoVAkvotzLn4ismk4dMmDoNaUrim",
		lat: "",
		lng:"",
		sitename:"",
		lineid:"",
		token:"",
		siteinfo: {},
		isfinish: false,
    },
    mutations: {
		changecantakephone(state, iscan){
			state.cantakephone = iscan
		},
		isfinish(state, is){
			state.isfinish = is
		},
		setToken(state, token){
			state.token = token
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
		setSiteinfo(state, data){
			state.siteinfo = data
		},
		setTeacher(state, data){
			state.teacher = data
		},
		setLineid(state, lineid){
			state.lineid = lineid
		},
	},
    actions: {}
})
export default store