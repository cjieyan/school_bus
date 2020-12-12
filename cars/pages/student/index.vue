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
				<span v-if="line.name == undefined || lineid == ''" class="location-info" @tap="show = true">请选择路线</span>
				<span v-else class="location-info" @tap="show = true">{{line.name}} (切换)</span>
			</view>
			<view class="student-list">
				<view class="student-list-info" v-for="(item, index) in students" :key="index">
					<view class="student-img" @tap="inboard(item)" :value="item.id">
						<image :src="item.headImg" style="width: 100rpx; height: 100rpx;" class="headImg"></image>
						<view v-if="item.swipeStatus == '-1'" class="inboard">
							<span>未上车</span>
						</view>
						<view  v-else-if="item.swipeStatus == '0'" class="">
							<span></span>
							<icon v-if="item.swipeStatus == '1'" type="success" size="13" />
						</view>
						<view  v-else-if="item.swipeStatus == '1'" class="inboard">
							<span >已下车</span>
						</view>
						<view  v-else class="inboard">
							<span>未上车</span>
						</view>
					</view>
					<view class="student-name">
						{{item.name}}
					</view>
				</view>
				<u-action-sheet :list="linelist" @click="setline" v-model="show"></u-action-sheet>
				<view style="clear:both;height:0"></view>
			</view>
		</view>
		<view class="comfirm">
			<button class="comfirm-bottom" @tap="finish" :disabled="disabled">{{btntex}}</button>
		</view>
	</view>
</template>

<script>
	export default {
		components: {
			// faIcon
		},
		data() {
			return {
				btntex:"确认结束",
				students: {},
				line: {},
				linelist: [],
				lineid: "",
				carid: "",
				linename: "",
				show: false,
				disabled: false,
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
						
					},
					fail: (err) => {
						
					}
				})
			},
			inboard(item) {
				uni.showLoading({

				})
				uni.redirectTo({
					url: "./inboard?id=" + item.id,
					fail: (err) => {
						
					}
				})
				uni.hideLoading()
			},
			finish() {
				if (this.lineid == "" || this.lineid == undefined) {
					uni.showToast({
						title: "请选择线路",
						icon: "none"
					})
					return false
				}
				uni.showLoading({

				})
				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/line-finish",
					method: "POST",
					header: {
						'token': this.$store.state.token,
					},
					data: {
						"line_id": this.$store.state.lineid,
						"car_id": this.$store.state.carid,
					},
					success: (res) => {
						if (res.data.code == -1) {
							uni.showModal({
								title: "尚未有人上车，是否结束",
								icon: "none",
								success: (res) => {
									if (res.confirm) {
										uni.request({
											url: this.$store.state.apihost + "/xcx/auth/line-finish",
											method: "POST",
											header: {
												'token': this.$store.state.token,
											},
											data: {
												"line_id": this.$store.state.lineid,
												"car_id": this.$store.state.carid,
												"force": 1,
											},
											success: (res) => {
												if(res.data.code == 200)
												{
													uni.showToast({
														title:"线路已结束",
														icon:"none"
													})
													this.btntex = "已结束"
													this.disabled = true
												}
											},
											fail: (err) => {
												uni.showToast({
													title:"请求失败，请重试",
													icon:"none"
												})
												return false
											}
										})
									} else if (res.cancel) {
										return false;
									}
								}
							})
							
						}
						if (res.data.code == 200) {
							uni.showToast({
								title: "打卡结束",
								icon: "none"
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
						}

					},
					fail: (err) => {
						console.log(err)
					}
				})

				uni.hideLoading()
			},
			studentList() {
				const token = uni.getStorageSync('token')
				if (this.$store.state.lineid) {
					uni.request({
						url: this.$store.state.apihost + "/xcx/auth/line-students",
						method: "POST",
						header: {
							'token': token,
						},
						data: {
							"line_id": this.$store.state.lineid,
							"car_id": this.$store.state.carid,
						},
						success: (res) => {
							if (res.data.code == 401) {
								uni.showToast({
									icon: 'none',
									title: '会话过期，请重新登录',
									duration: 1500
								});
								uni.redirectTo({
									url: "../my/login"
								})
							}
							if (res.data.code == 200 && res.data.data.hasOwnProperty('studentsDataRet')) {
								this.students = res.data.data.studentsDataRet
							} else {
								this.students = {}
								uni.showToast({
									title: "无查询记录",
									icon: "none"
								})
							}
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
					url: this.$store.state.apihost + "/xcx/auth/cars",
					method: "post",
					header: {
						'token': token,
					},
					data: {},
					success: (res) => {
						if (res.data.code == 401) {
							uni.showToast({
								icon: 'none',
								title: '会话过期，请重新登录',
								duration: 1500
							});
							uni.redirectTo({
								url: "../my/login"
							})
						}
						var data = []
						this.linelist = res.data.data
						for (var i = 0; i < res.data.data.length; i++) {
							var obj = {
								text: res.data.data[i].name,
								color: 'blue',
								fontSize: 28,
								id: res.data.data[i].id,
								carid: res.data.data[i].carId,
							}
							data.push(obj)
						}
						this.linelist = data
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
			setline(index) {
				uni.showLoading({

				})
				var token = uni.getStorageSync('token')
				new Promise(resolve => {
					uni.request({
						url: this.$store.state.apihost + "/xcx/auth/line-info",
						method: "POST",
						header: {
							'token': token,
						},
						data: {
							"line_id": this.linelist[index].id,
							"car_id": this.linelist[index].carid
						},
						success: (res) => {
							if (res.data.code == 401) {
								uni.showToast({
									icon: 'none',
									title: '会话过期，请重新登录',
									duration: 1500
								});
								uni.redirectTo({
									url: "../my/login"
								})
							}
							this.$store.commit('setLineid', this.linelist[index].id)
							this.$store.commit('setCarid', this.linelist[index].carid)
							this.$store.commit('setcarinfo', res.data.data.car)
							this.$store.commit('setTeacher', res.data.data.teacher)
							this.$store.commit('setLineinfo', res.data.data.line)
							this.$store.commit('setSiteinfo', res.data.data.sites)
							this.$store.commit('setstudent', {
								"studentCount": res.data.data.studentCount,
								"studentGetOnCount": res.data.data.studentGetOnCount
							})
							this.lineid = this.linelist[index].id
							this.carid = this.linelist[index].carid
							this.line = res.data.data.line
							resolve(this.line)
						},
						fail: (err) => {
							console.log(err)
						}
					})
				}).then((res) => {
					uni.request({
						url: this.$store.state.apihost + "/xcx/auth/line-students",
						method: "POST",
						header: {
							'token': token,
						},
						data: {
							"line_id": res.id,
							"car_id": this.$store.state.carInfo.id
						},
						success: (res) => {
							if (res.data.code == 401) {
								uni.showToast({
									icon: 'none',
									title: '会话过期，请重新登录',
									duration: 1500
								});
								uni.redirectTo({
									url: "../my/login",
									icon: "none"
								})
							}
							if (res.data.code == 200 && res.data.data.hasOwnProperty('studentsDataRet')) {
								this.students = res.data.data.studentsDataRet
							} else {
								this.students = {}
								uni.showToast({
									icon: "none",
									title: "无查询记录"
								})
							}

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
			this.linelist = {}
			this.students = {}
			this.lineid = this.$store.state.lineid
			this.lineInfo()
			this.line = this.$store.state.liniInfo
			this.studentList()
		},
		onLoad() {
			this.linelist = {}
			this.students = {}
			this.lineid = this.$store.state.lineid
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
		width: 100rpx;
		height: 100rpx;
		position: absolute;
		top: 0;
		left: 0;
		border-radius: 50rpx;
		opacity: 0.9;
	}

	.outboard span,
	.inboard span {
		font-size: 20rpx;
		color: #fff;
		position: absolute;
		top: 38rpx;
		left: 24rpx;
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

	image.headImg {
		border-radius: 50rpx;
	}

	icon {
		top: 70rpx;
		position: relative;
	}
</style>
