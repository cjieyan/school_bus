<template class="container">
	<view>
		<view class="top">
			<u-navbar title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="image-content">
			<image src="../../static/banner.jpg" mode="aspectFill" style="width: 100%;" @error="imageError"></image>
		</view>
		<view class="content">
			<view v-if="line.name == undefined">
			</view>
			<view v-else>
				<view class="content-header">
					<view class="content-header-top">
						<span class="border"></span><span class="teacher">{{teacher.name}}</span> <span class="school">（{{line.name}}）</span>
					</view>
				</view>
				<view class="content-info">
					<view class="carinfo">{{carinfo.carNumber}}{{carinfo.carNo}}</view>
					<view class="carinfo">{{carinfo.driver}}{{carinfo.phone}}</view>
				</view>
			</view>
			<view class="content-hr"></view>
			<view class="location">
				<view class="location-text">
					<span class="location-img">
						<image src="../../static/location.png" class="location-image"></image>
					</span>
					<span class="location-info" v-if="line.name == undefined" @tap="show = true">
						请选择路线
					</span>
					<span v-else class="location-info" @tap="show = true">
						{{line.name}}
					</span>
				</view>

				<view class="location-d"  v-if="numList.length == 0">
					<u-loading mode="circle"></u-loading>
				</view>
				<view class="location-d"  v-else>
					<u-steps :list="numList" active-color="rgb(255 114 58)" mode="dot" direction="column" :current="current">
					</u-steps>
				</view>
			</view>
			<view v-if="line.name == undefined">
			</view>
			<view v-else class="comfirm">
				<button class="comfirm-bottom" :loading="loading" :disabled="disabled" @tap="gotoLunBo">开始刷脸打卡</button>
			</view>
		</view>
		<u-action-sheet :list="linelist" @click="setline" v-model="show"></u-action-sheet>
	</view>

</template>

<script>
	import utils from '../../common/utils.js'
	export default {
		components: {},
		data() {
			return {
				timer: "",
				ifOnShow: false,
				ak: 'tjSZDrqGonHtUiDfmKpa8WbdEoOtSIcH',
				latitude: '',
				longitude: '',
				blat: "",
				blng: "",
				background: {
					backgroundColor: '#12c497',
				},
				numList: [],
				pk: 180 / 3.14169,
				line: {},
				carinfo: {},
				teacher: {},
				students: {},
				disabled: true,
				loading: true,
				current: 0,
				linelist: [],
				show: false,
			};
		},
		methods: {
			back() {
				uni.navigateBack({
					success: function() {
						beforePage.onLoad();
					}
				})
			},
			gotoLunBo() {
				uni.showLoading({
					title: '加载中'
				});
				uni.navigateTo({
					url: "/pages/face/index",
					success: res => {},
					fail: (e) => {
						//console.log(e)
					},
					complete: () => {}
				})
				uni.hideLoading();
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
						if(res.data.code == 401){
							uni.showToast({
								icon: 'none',
								title: '会话过期，请重新登录',
								duration: 1500
							});
							uni.redirectTo({
								url:"../my/login"
							})
						}
						var data = []
						// this.linelist = res.data.data
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
							if(res.data.code == 401){
								uni.showToast({
									icon: 'none',
									title: '会话过期，请重新登录',
									duration: 1500
								});
								uni.redirectTo({
									url:"../my/login"
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
							this.line = res.data.data.line
							this.carinfo = res.data.data.car
							this.teacher = res.data.data.teacher
							this.students = {
								"studentCount": res.data.data.studentCount,
								"studentGetOnCount": res.data.data.studentGetOnCount
							}
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
							"line_id": res.id
						},
						success: (res) => {
							if(res.data.code == 401){
								uni.showToast({
									icon: 'none',
									title: '会话过期，请重新登录',
									duration: 1500
								});
								uni.redirectTo({
									url:"../my/login"
								})
							}
							if( res.data.code == 200 ){
								this.students = res.data.data.studentsDataRet
							}
							
						},
						fail: (err) => {
							console.log(err)
						}
					})
				})

				uni.hideLoading()
			},
			getLocation() {
				new Promise(resolve => {
					uni.getLocation({
						type: "wgs84",
						success: (res) => {
							this.latitude = res.latitude
							this.longitude = res.longitude
							resolve(res);
						}
					})
				}).then((res) => {
					const lat = res.latitude
					const lng = res.longitude
					new Promise(resolve => {
						uni.request({
							url: "http://api.map.baidu.com/geoconv/v1/?coords=" + lng + "," + lat +
								"&from=1&to=5&ak=" + this.$store.state.ak,
							method: "GET",
							data: {},
							success: (res) => {
								this.blat = res.data.result[0].y
								this.blng = res.data.result[0].x
								const data = {
									"blat": res.data.result[0].y,
									"blng": res.data.result[0].x
								}
								// resolve(res.data.result[0])
								resolve(data)
							},
							fail: (err) => {
								console.log(err)
							}
						})
					}).then((res) => {
						new Promise(resolve => {
							var token = uni.getStorageSync('token')
							uni.request({
								url: "http://api.map.baidu.com/reverse_geocoding/v3/?ak=" + this.$store.state.ak +
									"&output=json&coordtype=wgs84ll&location=" +
									this.latitude + "," + this.longitude,
								data: {
									pois: 1,
								},
								success: (res) => {
									resolve(res)
								},
								fail: (err) => {
									console.log(err)
								}
							})
						}).then((res) => {
							var siteList = []
							if (this.$store.state.siteinfo.length > 0) {
								for (var i = 0; i < this.$store.state.siteinfo.length; i++) {
									var distict = utils.getGreatCircleDistance(this.blat, this.blng, this.$store.state.siteinfo[i].latitude,
										this.$store.state.siteinfo[i].longitude)
									if (distict <= 100) {
										this.current = i
										this.$store.state.lat = this.$store.state.siteinfo[i].latitude
										this.$store.state.lng = this.$store.state.siteinfo[i].longitude
										this.$store.state.sitename = this.$store.state.siteinfo[i].name
									}
									var resdata = {
										"name": this.$store.state.siteinfo[i].name + "(" + Math.round(distict) + "米)",
										"distict": Math.round(distict)
									}
									siteList.push(resdata)
									this.numList = siteList
									this.disabled = false
									this.loading = false
								}
							}

						})
					})
				})

			},
		},
		onHide() {
			this.ifOnShow = true
			if (this.timer) {
				clearInterval(this.timer)
				this.timer = null
			}
		},

		onShow() {
			uni.showLoading({
				title: "正在加载"
			})
			this.numList = []
			this.lineInfo()
			this.line = this.$store.state.liniInfo
			if (this.line.hasOwnProperty("name")) {
				// this.line = this.$store.state.liniInfo
				this.carinfo = this.$store.state.carInfo
				this.teacher = this.$store.state.teacher
				this.students = this.$store.state.student
			} else {
				uni.showToast({
					icon: 'none',
					title: '请选择路线',
					duration: 5000
				});
			}
			this.timer = setInterval(() => {
				this.getLocation()
			}, 3000)
			uni.hideLoading()
		},
		onLoad() {
			uni.showLoading({

			})
			this.numList = []
			this.getLocation()
			this.line = this.$store.state.liniInfo
			this.carinfo = this.$store.state.carInfo
			this.teacher = this.$store.state.teacher
			this.students = this.$store.state.student
			uni.hideLoading()
		}
	}
</script>

<style>
	.container {
		background-color: #fafafa;
	}

	.image-content image {
		height: 300rpx;
	}

	.content {
		padding: 20rpx 40rpx 20rpx 40rpx;
		background-color: #fff;
	}

	.border {
		border-left: 6rpx solid rgb(254 194 32);
		height: 10rpx;
		margin-right: 16rpx;
	}

	.teacher {
		font-size: 32rpx;
		font-weight: bold;
	}

	.school {
		font-size: 28rpx;
		color: rgb(148 148 148);
	}

	.content-info {
		margin-top: 20rpx;
		color: rgb(111 111 111);
		padding-left: 16rpx;
		line-height: 30rpx;
	}

	.content-hr {
		border: 1rpx solid rgb(237 237 237);
		margin: 30rpx 0;
	}

	.location-image {
		width: 72rpx;
		height: 62rpx;
	}

	.carinfo {
		font-size: 24rpx;
	}

	.location {
		padding: 0 60rpx;
	}

	.location-info {
		color: rgb(85 85 85);
		margin-left: 20rpx;
		padding: 10rpx 0;
	}

	.location-d {
		margin-top: 20rpx;
		padding: 0 100rpx;
	}

	/*rgb(255 114 58)*/
	.uni-steps__column-text {
		border-bottom: none !important;
		height: 70rpx;
	}

	.comfirm {
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

	.uni-navbar__header-btns-left {
		width: 250rpx;
	}
</style>
