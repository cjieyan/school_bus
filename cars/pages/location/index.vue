<template class="container">
	<view>
		<view class="top">
			<u-navbar title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="image-content">
			<image src="../../static/banner.jpg" mode="aspectFill" style="width: 100%;" @error="imageError"></image>
		</view>
		<view class="content">
			<view class="content-header">
				<view class="content-header-top">
					<span class="border"></span><span class="teacher">{{teacher.name}}</span> <span class="school">（{{line.name}}）</span>
				</view>
			</view>
			<view class="content-info">
				<view class="carinfo">{{carinfo.carNumber}}{{carinfo.carNo}}（已上车{{students.studentGetOnCount}}人/共{{students.studentCount}}人）</view>
				<view class="carinfo">{{carinfo.driver}}{{carinfo.phone}}</view>
			</view>
			<view class="content-hr"></view>
			<view class="location">
				<view class="location-text">
					<span class="location-img">
						<image src="../../static/location.png" class="location-image"></image>
					</span>
					<span class="location-info">
						{{line.name}}
					</span>
				</view>
				<view class="location-d">
					<u-steps :list="numList" active-color="rgb(255 114 58)" mode="dot" direction="column" :current="current">
					</u-steps>
				</view>
			</view>
			<view class="comfirm">
				<button class="comfirm-bottom" :loading="loading" :disabled="disabled" @tap="gotoLunBo">开始刷脸打卡</button>
			</view>
		</view>
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
				current:0,
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
						console.log(e)
					},
					complete: () => {}
				})
				uni.hideLoading();
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
						console.log("res-----------------")
						console.log(res)
						// const blat = res.y
						// const blng = res.x
						console.log("this.blat")
						console.log(res.blat)
						console.log("this.blng")
						console.log(res.blng)
						console.log("this.latitude")
						console.log(this.latitude)
						console.log(this.longitude)
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
									console.log("-----百度地图----")
									console.log(res)
									resolve(res)
								},
								fail: (err) => {
									console.log(err)
								}
							})
						}).then((res) => {
							console.log("--------siteinfo-----")
							console.log(this.$store.state.siteinfo)
							var siteList = []
							for (var i = 0; i < this.$store.state.siteinfo.length; i++) {
								var distict = utils.getGreatCircleDistance(this.blat, this.blng, this.$store.state.siteinfo[i].latitude, this.$store.state.siteinfo[i].longitude)
								if(distict <= 100){
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
						})
					})
				})

			},
		},
		onHide() {
			console.log('this.ifOnShow=true')
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
			this.timer = setInterval(() => {
				this.getLocation()
			}, 3000)

			uni.hideLoading()
		},
		onLoad() {
			uni.showLoading({

			})
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
		height: 150px;
	}

	.content {
		padding: 10px 20px 10px 20px;
		background-color: #fff;
	}

	.border {
		border-left: 3px solid rgb(254 194 32);
		height: 5px;
		margin-right: 8px;
	}

	.teacher {
		font-size: 16px;
		font-weight: bold;
	}

	.school {
		font-size: 14px;
		color: rgb(148 148 148);
	}

	.content-info {
		margin-top: 10px;
		color: rgb(111 111 111);
		padding-left: 8px;
		line-height: 25px;
	}

	.content-hr {
		border: 0.5px solid rgb(237 237 237);
		margin: 15px 0;
	}

	.location-image {
		width: 36px;
		height: 31px;
	}

	.carinfo {
		font-size: 12px;
	}

	.location {
		padding: 0 30px;
	}

	.location-info {
		color: rgb(85 85 85);
		margin-left: 10px;
		padding: 5px 0;
	}

	.location-d {
		margin-top: 10px;
		padding: 0 50px;
	}

	/*rgb(255 114 58)*/
	.uni-steps__column-text {
		border-bottom: none !important;
		height: 35px;
	}

	.comfirm {
		padding: 10px 40px;
	}

	.comfirm-bottom {
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 25px;
	}

	.comfirm-bottom::after {
		border: none;
	}

	.uni-navbar__header-btns-left {
		width: 250rpx;
	}
</style>
