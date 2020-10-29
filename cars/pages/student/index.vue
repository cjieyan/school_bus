<template>
	<view>
		<view class="top">
			<u-navbar back-text="返回" title="学生列表" @tap="back" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="student">
			<view class="student-top">
				<span class="location-img">
					<image src="../../static/location.png" style="width: 32px; height: 31px;" class="location-image"></image>
				</span>
				<span v-if="line.name == undefined" class="location-info" @tap="show = true">请选择路线</span>
				<span v-else class="location-info" @tap="show = true">{{line.name}}</span>
			</view>
			<view class="student-list">
				<view class="student-list-info" v-for="(item, index) in students" :key="index">
					<view class="student-img"  @tap="inboard(item)" :value="item.id">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						<view class="inboard">
							<span v-if="item.swipeStatus == '-1'">未上车</span>
							<span v-else-if="item.swipeStatus == '0'">已上车</span>
							<span v-else-if="item.swipeStatus == '1'">已下车</span>
							<span v-else>未上车</span>
							<!-- <fa-icon type="check" size="14" color="white" class="font-icon inboard-check"></fa-icon> -->
							<u-icon name="checkbox-mark" color="#fff" size="14" class="inboard-check"></u-icon>
						</view>
					</view>
					<view class="student-name">
						{{item.name}}
					</view>
				</view>
				<u-action-sheet :list="linelist" @click="setline" v-model="show"></u-action-sheet>
				<!-- <view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						<view class="outboard">
							<span>未上车</span>
						</view>
					</view>
					<view class="student-name">
						张三
					</view>
				</view>
				<view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
					</view>
					<view class="student-name">
						张三
					</view>
				</view>
				<view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
					</view>
					<view class="student-name">
						张三
					</view>
				</view>
				<view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
					</view>
					<view class="student-name">
						张三
					</view>
				</view> -->
				<view style="clear:both;height:0"></view>
			</view>
		</view>
		<view class="comfirm">
			<button class="comfirm-bottom" @tap="finish">确认结束</button>
		</view>
	</view>
</template>

<script>
	// import faIcon from "@/components/kilvn-fa-icon/fa-icon.vue"
	export default {
		components: {
			// faIcon
		},
		data() {
			return {
				students: {},
				line: {},
				linelist: [],
				linename: "",
				// list: [{
				// 	text: '点赞',
				// 	color: 'blue',
				// 	fontSize: 28,
				// 	subText: '感谢您的点赞'
				// }, {
				// 	text: '分享'
				// }, {
				// 	text: '评论'
				// }],
				show: false,
				background: {
					backgroundColor: '#12c497',
				},
			}
		},
		methods: {
			back() {
				uni.switchTab({
					url: "../index/index",
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
			inboard(item){
				uni.showLoading({
					
				})
				uni.redirectTo({
					url:"./inboard?id="+item.id,
					fail: (err) => {
						console.log(err)
					}
				})
				uni.hideLoading()
			},
			finish() {
				uni.showLoading({

				})
				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/line-finish",
					method: "POST",
					header: {
						'token': this.$store.state.token,
					},
					data: {
						"line_id": this.$store.state.lineid
					},
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
				uni.switchTab({
					url: "../index/index",
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
				uni.hideLoading()
			},
			studentList() {
				console.log("this.$store.state.lineid-----")
				console.log(this.$store.state.lineid)
				const token = uni.getStorageSync('token')
				if (this.$store.state.lineid) {
					uni.request({
						url: this.$store.state.apihost + "/xcx/auth/line-students",
						method: "POST",
						header: {
							'token': token,
						},
						data: {
							"line_id": this.$store.state.lineid
						},
						success: (res) => {
							console.log(res)
							this.students = res.data.data.studentsDataRet
						},
						fail: (err) => {
							console.log(err)
						}
					})
				}
			},
			lineInfo() {
				var token = uni.getStorageSync('token')
				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/lines",
					method: "post",
					header: {
						'token': token,
					},
					data: {},
					success: (res) => {
						var data = []
						this.linelist = res.data.data
						for (var i = 0; i < res.data.data.length; i++) {
							var obj = {
								text: res.data.data[i].name,
								color: 'blue',
								fontSize: 28,
								id : res.data.data[i].id,
							}
							data.push(obj)
						}
						console.log(data)
						this.linelist = data
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
			setline(index){
				uni.showLoading({
					
				})
				var token = uni.getStorageSync('token')
				new Promise(resolve => {
					uni.request({
						url:this.$store.state.apihost + "/xcx/auth/line-info",
						method:"POST",
						header: {
							'token': token,
						},
						data:{
							"line_id": this.linelist[index].id
						},
						success: (res) => {
							console.log("--------selectline--------")
							console.log(this.$store.state.lineid)
							console.log(res)
							this.$store.commit('setcarinfo', res.data.data.car)
							this.$store.commit('setTeacher', res.data.data.teacher)
							this.$store.commit('setLineinfo', res.data.data.line)
							this.$store.commit('setSiteinfo', res.data.data.sites)
							this.$store.commit('setstudent', {
								"studentCount": res.data.data.studentCount,
								"studentGetOnCount": res.data.data.studentGetOnCount
							})
							this.line = res.data.data.line
							resolve(this.line)
						},
						fail: (err) => {
							console.log(err)
						}
					})
				}).then((res)=> {
					console.log("studentList--------")
					console.log(res)
					// this.studentList()
					uni.request({
						url: this.$store.state.apihost + "/xcx/auth/line-students",
						method: "POST",
						header: {
							'token': token,
						},
						data: {
							"line_id": res.id
						},
						success: (res) => {
							console.log(res)
							this.students = res.data.data.studentsDataRet
						},
						fail: (err) => {
							console.log(err)
						}
					})
				})
				
				uni.hideLoading()
			}
		},
		onShow() {
			this.lineInfo()
			this.line = this.$store.state.liniInfo
			this.studentList()
		},
		onLoad() {
			this.studentList()
		}
	}
</script>

<style>
	page {
		background-color: #f8f8f8;
	}

	.student {
		margin-top: 20rpx;
		padding: 0 10%;
	}

	.student-top {
		border-bottom: 2rpx solid rgb(238 238 238);
		padding-bottom: 20rpx;
	}

	.location-info {
		margin-left: 50rpx;
		color: rgb(89 89 89);
		font-weight: bold;
	}

	.student-list-info {
		width: 80rpx;
		float: left;
		margin: 30rpx 20rpx;
	}

	.student-img {
		width: 80rpx;
		height: 80rpx;
		border-radius: 40rpx;
		text-align: center;
		position: relative;
	}

	.outboard,
	.inboard {
		background-color: #808080;
		width: 84rpx;
		height: 84rpx;
		position: absolute;
		top: 0;
		left: 0;
		border-radius: 40rpx;
		opacity: 0.9;
	}

	.outboard span,
	.inboard span {
		font-size: 16rpx;
		color: #fff;
		position: absolute;
		top: 32rpx;
		left: 18rpx;
	}

	.inboard-check {
		position: absolute;
		bottom: -12rpx;
		left: 36rpx;
		background-color: #4CD964;
		border-radius: 12rpx;
	}

	.student-name {
		margin-top: 16rpx;
		font-size: 24rpx;
		text-align: center;
	}

	.comfirm {
		margin-bottom: 60rpx;
		padding: 20rpx 80rpx;
	}

	.comfirm-bottom {
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 50rpx;
	}

	.comfirm-bottom::after {
		border: none;
	}
</style>
