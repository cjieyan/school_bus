<template>
	<view class="face">
		<view class="top">
			<u-navbar title="人脸识别" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<!-- <u-mask :show="show" @click="show = false"></u-mask> -->

		<view class="camera">
			<camera :device-position="camera" flash="off" binderror="error">
			</camera>
		</view>
		<view :class="tipclass">
<!-- 			<u-icon name="checkbox-mark" color="#fff" size="60" class="inboard-check"></u-icon> -->
			<icon :type="icon" size="26"/>
			<text>{{tip}}</text>
		</view>

		<view class="bottom-btn">
			<view class="iconfont icon-avcameraswitchh camera-switch-left"></view>
			<button :type="btntype" class="start-facing" @tap="stoptakephone">{{btntxt}}</button>
			<view :class="cameraicon" @tap="changeCamera"></view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				background: {
					backgroundColor: '#12c497',
				},
				show: true,
				btntxt: "停止打卡",
				btntype: "error",
				tipclass: "",
				istakephone: true,
				accesstoken: "",
				student: {},
				userid: "",
				userinfo: "",
				usergroup: "",
				facetoken: "",
				timer: "",
				ifOnShow: false,
				stauts: "",
				tip: "",
				icon: "",
				camera: "front",
				cameraicon: "iconfont icon-avcameraswitch camera-switch-right",
			}
		},
		methods: {
			back() {
				uni.navigateBack({
					success: function() {
						beforePage.onLoad();
					}
				})
			},
			stoptakephone() {
				if (this.istakephone) {
					this.istakephone = false
					this.btntype = "primary"
					this.btntxt = "开始打卡"
					this.tip = "已停止识别"
					if (this.timer) {
						clearInterval(this.timer)
						this.timer = null
					}


				} else {
					this.btntype = "error"
					this.istakephone = true
					// btntxt:"停止打卡",
					this.btntxt = "停止打卡"
					this.tip = "正在识别"
					if (this.timer == null) {
						this.timer = setInterval(() => {
							this.takephone()
						}, 3000)
					}
				}

			},
			changeCamera() {
				if (this.camera == "front") {
					this.camera = "back"
					this.cameraicon = "iconfont icon-avcameraswitchh camera-switch-right"
				} else {
					this.camera = "front"
					this.cameraicon = "iconfont icon-avcameraswitch camera-switch-right"
				}
			},
			takephone() {
				this.icon = "waiting"
				this.tip = "正在识别"
				const ctx = wx.createCameraContext()
				var that = this
				ctx.takePhoto({
					quality: 'high',
					success: (res) => {
						const imgbase64Url = wx.getFileSystemManager().readFileSync(res.tempImagePath, "base64");
						var token = uni.getStorageSync('token')
						uni.request({
							url: this.$store.state.apihost + "/xcx/auth/face-swipe",
							method: "POST",
							header: {
								'token': token,
							},
							data: {
								"image": imgbase64Url,
								"line_id": this.$store.state.lineid,
								"car_id": this.$store.state.carInfo.id
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
								if (res.data.code == 200) {
									this.icon = "success"
									if (res.data.isFinished == true) {
										this.$store.commit("isfinish", true)
									}
									uni.redirectTo({
										url: "complete"
									})
								} else {
									this.icon = "warn"
									this.tip = res.data.msg
								}
							},
							fail: (err) => {
								uni.showModal({
									title: "识别失败"
								})
							}
						})
						uni.hideLoading()
					},
					fail: (err) => {
						uni.showModal({
							title: "无法打开摄像头，请重试"
						})
					}
				})
			},
		},
		onLoad() {
			this.tipclass = "tips yellow"
			this.tip = "正在初始化"
			console.log(this.$store.state.cantakephone)
			uni.showLoading({
				title: "加载中..."
			})
			console.log(this.accesstoken)

			console.log(this.$store.state.cantakephone)
			//获取百度access_token

			uni.request({
				url: "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + this.$store.state.client_id +
					"&client_secret=" + this.$store.state.client_secret + "&",
				method: 'POST',
				dataType: "json",
				data: {},
				success: (res) => {
					console.log(res)
					this.accesstoken = res.data.access_token
					console.log(this.accesstoken)
				},
				fail: (error) => {
					console.log(error)
				}
			})
			uni.hideLoading()
			this.timer = setInterval(() => {
				this.takephone()

			}, 3000)
		},
		onUnload() {
			if (this.timer) {
				clearInterval(this.timer)
				this.timer = null
			}
		}
	}
</script>

<style lang="scss">
	@import "../../static/icon/iconfont.css";
	page {
		background-color: #fff;
	}

	.facing-tip image {
		height: 370rpx;
		width: 100%;
	}

	.face .camera {
		position: relative;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.camera camera {
		width: 100%;
		height: 100vh;
		position: absolute;
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: -99;
	}

	.progress_bg {
		position: absolute;
		width: 440rpx;
		height: 440rpx;
	}

	.progress_canvas {
		width: 440rpx;
		height: 440rpx;
	}

	.bottom-btn {
		position: fixed;
		bottom: 50rpx;
		display: flex;
		width: 100%;
		justify-content: center;
	}

	.camera {
		margin-top: 12%;
	}

	.start-facing {
		height: 200rpx;
		width: 200rpx;
		border-radius: 100rpx;
		border: 6rpx solid #fff;
		background-color: rgb(228, 63, 25);
		font-weight: bold;
		font-size: 32rpx;
		padding: 50rpx 0;
		color: #fff;
		margin-top: 60rpx;
	}

	.tips {
		position: fixed;
		bottom: 250rpx;
		display: flex;
		width: 100%;
		justify-content: center;
		height: auto;
		font-weight: bold;
		font-size: 55rpx;
		text-align: center;
		margin: 60rpx 0;
	}

	.inboard-check {
		background-color: #4CD964;
		border-radius: 30rpx;
		// width: 30px;
		// height: 30px;
		margin-top: 12rpx;
	}

	.yellow {
		color: #FF9900
	}

	.red {
		color: #FA3534
	}

	.green {
		color: #18B566
	}

	.btn-left {
		width: 100rpx;
		height: 100rpx;
		margin-left: 15%;
	}

	.btn-left {
		background-color: #fff;
		border: none !important;
	}

	.btn-left:after,
	.btn-right:after {
		display: none;
	}

	.btn-right {
		width: 112rpx;
		height: 100rpx;
		padding: 0 16rpx;
	}

	.camera-border {
		margin-right: 15%;
		width: 112rpx;
		height: 112rpx;
		margin-top: 100rpx;
		border: 6rpx solid #ccc;
		border-radius: 56rpx;
	}
	.camera-switch-right{
		font-size: 100rpx;
		margin-top: 140rpx;
		margin-right: 70rpx;
	}
	.camera-switch-left{
		font-size: 100rpx;
		margin-top: 140rpx;
		margin-left: 70rpx;
	}
	.camera-switch-left:before{
		color:#fff;
		background-color: #fff;
	}
	icon{
		margin-top: 10rpx;
	}
</style>
