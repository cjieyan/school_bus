<template class="container">
	<view>
		<view class="top">
			<u-navbar title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="image-content">
			<image src="../../static/banner.jpg" @error="imageError"></image>
		</view>
		<view class="content">
			<view class="content-header">
				<view class="content-header-top">
					<span class="border"></span><span class="teacher">王老师</span> <span class="school">（梧桐小学）</span>
				</view>
			</view>
			<view class="content-info">
				<view class="carinfo">1号车粤12345（已上车15人/共20人）</view>
				<view class="carinfo">张师傅12345678911</view>
			</view>
			<view class="content-hr"></view>
			<view class="location">
				<view class="location-text">
					<span class="location-img">
						<image src="../../static/location.png" class="location-image"></image>
					</span>
					<span class="location-info">
						梧桐小学--六约新村
					</span>
				</view>
				<view class="location-d">
					<u-steps :list="numList" active-color="rgb(255 114 58)" mode="dot" direction="column">
						
					</u-steps>
				</view>
			</view>
			<view class="comfirm">
				<button class="comfirm-bottom" @tap="gotoLunBo">开始刷脸打卡</button>
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
				pk: 180 / 3.14169
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
			}
		},
		onHide() {
			console.log('this.ifOnShow=true')
			this.ifOnShow = true
		},

		onShow() {
			if (this.ifOnShow) {
				uni.showLoading({
					title: "正在定位"
				})
				var that = this;
				/* 获取定位地理位置 */
				uni.getLocation({
					type: 'wgs84',
					success: function(res) {
						console.log('当前位置的经度：' + res.longitude);
						console.log('当前位置的纬度：' + res.latitude);
						uni.request({
							url: "http://api.map.baidu.com/reverse_geocoding/v3/?ak=hpMsWg4OXrCajEUgHhuiYHMoFEYcCRL7&output=json&coordtype=wgs84ll&location=" +
								res.latitude + "," + res.longitude,
							data: {},
							success: (res) => {
								// console.log(res)
								// uni.showModal({
								// 	title: res.data.result.formatted_address
								// })
								//获取最近的站点

							},
							fail: (err) => {
								console.log(err)
							}
						})
						console.log("eeeeeeeeeeeeeeeeeee")
						uni.request({
							url: "http://localhost:8000/xcx/sites",
							method: "POST",
							data: {},
							success: (res) => {
								console.log("------计算最近站点-------")
								console.log(res.data.data.length)
								var t = 0
								for (var i = 0; i < res.data.data.length; i++) {

								}
							},
							fail: (err) => {
								console.log("------计算最近站点err-------")
								console.log(err)
							}
						})
						console.log("fffffffffffffffffffffffffff")
					},
					fail: (err) => {
						console.log(err)
						uni.showModal({
							title: "定位失败"
						})
					}
				});
				uni.hideLoading()
			}
		},
		onLoad() {
			uni.showLoading({
				title: "正在获取定位"
			})
			var that = this;
			/* 获取定位地理位置 */
			
			uni.getLocation({
				// type: 'wgs84',
				type:"wgs84",
				success: function(res) {
					console.log('当前位置的经度：' + res.longitude);
					console.log('当前位置的纬度：' + res.latitude);
					that.longitude = res.longitude
					that.latitude = res.latitude
					//转换成百度坐标系
					uni.request({
						url:"http://api.map.baidu.com/geoconv/v1/?coords=114.52152,22.741887&from=1&to=5&ak=hpMsWg4OXrCajEUgHhuiYHMoFEYcCRL7",
						method:"GET",
						data:{},
						success: (res) => {
							console.log("---------百度坐标系--------")
							console.log(res)
							that.blat = res.data.result[0].y
							that.blng = res.data.result[0].x
						},
						fail: (err) => {
							console.log(err)
						}
					})
					var a1 = that.latitude / that.pk
					var a2 = that.longitude / that.pk
					uni.request({
						url: "http://api.map.baidu.com/reverse_geocoding/v3/?ak=hpMsWg4OXrCajEUgHhuiYHMoFEYcCRL7&output=json&coordtype=wgs84ll&location=" +
							res.latitude + "," + res.longitude,
						data: {pois: 1,},
						success: (res) => {
							console.log("-----百度地图----")
							console.log(res)
							uni.request({
								url: "http://localhost:8000/xcx/sites",
								method: "POST",
								data: {},
								success: (res) => {
									console.log("------计算最近站点-------")
									console.log("------计算最近站点-------")
									console.log(res.data)
									var t = 0
									var siteList = []
									//22.744943742471953,114.53285132220785,22.744943742471953
									for (var i = 0; i < res.data.data.length; i++) {
										var distict = utils.getGreatCircleDistance(that.blat,that.blng, res.data.data[i].latitude, res.data.data[i].longitude)
										console.log("-------"+ res.data.data[i].name +"-----")
										console.log(distict)
										console.log("------------")
										var resdata = {
											"name": res.data.data[i].name+"("+Math.round(distict)+"米)",
											"distict": Math.round(distict)
										}
										siteList.push(resdata)
										that.numList = siteList
									}
									console.log(siteList)
								},
								fail: (err) => {
									console.log(err)
								}
							})
						},
						fail: (err) => {
							console.log(err)
						}
					})


				},
				fail: (err) => {
					console.log(err)
					uni.showModal({
						title: "定位失败"
					})
				}
			});
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
