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
				<view class="carinfo">1号车粤12345</view>
				<view class="carinfo">张师傅12345678911</view>
			</view>
			<view class="content-hr"></view>
			<view class="location">
				<view class="location-text">
					<span class="location-img">
						<image src="../../static/location.png" class="location-image"></image>
					</span>
					<span class="location-info">
						梧桐小学--六约新村{{address}}
					</span>
				</view>
				<view class="location-d">
					<u-steps :list="numList" active-color="rgb(255 114 58)" mode="dot" direction="column"></u-steps>
				</view>
			</view>
			<view class="comfirm">
				<button class="comfirm-bottom" @tap="gotoLunBo">开始刷脸打卡</button>
			</view>
		</view>
	</view>

</template>

<script>
	var bmap = require('../../libs/bmap-wx.min.js');
	var wxMarkerData = [];
	export default {
		components: {},
		data() {
			return {
				ak: 'tjSZDrqGonHtUiDfmKpa8WbdEoOtSIcH',
				markers: [],
				latitude: '',
				longitude: '',
				placeData: {},
				background: {
					backgroundColor: '#12c497',
				},
				numList: [{
					name: '六约新村'
				}, {
					name: '六约新村'
				}, {
					name: '六约新村'
				}, {
					name: '六约新村'
				}, ],
				latitudeL: "",
				address: '', //地址  
				cityInfo: {} //城市信息 
			};
		},
		makertap: function(e) {
			var that = this;
			var id = e.markerId;
			that.showSearchInfo(wxMarkerData, id);
			that.changeMarkerColor(wxMarkerData, id);
		},
		methods: {
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
		onLoad() {
			var that = this;
			/* 获取定位地理位置 */

			uni.showLoading({
				title: "正在获取定位"
			})
			uni.getLocation({
						type: 'wgs84',
						success: function(res) {
							console.log('当前位置的经度：' + res.longitude);
							console.log('当前位置的纬度：' + res.latitude);
							
							//转换成百度坐标
							uni.request({
								url:"http://api.map.baidu.com/geoconv/v1/?coords="+res.latitude+','+ res.longitude+"&from=3&to=5&ak=hpMsWg4OXrCajEUgHhuiYHMoFEYcCRL7",
								success: (res) => {
									console.log("res:")
									console.log(res)
								},
								fail: (err) => {
									console.log(err)
								}
							})
							//调用百度地图API
							console.log("https://apis.map.qq.com/ws/geocoder/v1/?location="+res.latitude+","+res.longitude+"&key=OB4BZ-D4W3U-B7VVO-4PJWW-6TKDJ-WPB77&get_poi=1")
							uni.request({
								// url: 'http://api.map.baidu.com/reverse_geocoding/v3/?ak=tjSZDrqGonHtUiDfmKpa8WbdEoOtSIcH=json&coordtype=gcj02&location=' +res.latitude+','+ res.longitude,
								// url:'http://api.map.baidu.com/pano/?x=120.320032&y=31.589666&lc=0&ak=hpMsWg4OXrCajEUgHhuiYHMoFEYcCRL7&src=webapp.baidu.openAPIdemo',
								// url:"http://api.map.baidu.com/reverse_geocoding/v3/?ak=hpMsWg4OXrCajEUgHhuiYHMoFEYcCRL7&output=json&coordtype=wgs84ll&location="+(res.latitude+0.0060)+","+(res.longitude+0.0065),
								// data: {},
								//腾讯地图
								url:"https://apis.map.qq.com/ws/geocoder/v1/?location="+res.latitude+","+res.longitude+"&key=OB4BZ-D4W3U-B7VVO-4PJWW-6TKDJ-WPB77&get_poi=1",
								success: (res) => {
									console.log(res)
									uni.showModal({
										title:res.data.result.formatted_address
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
