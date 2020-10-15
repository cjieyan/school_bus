<template>
	<view class="face">
		<view class="top">
			<u-navbar title="人脸识别" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="facing-tip">
			<image src="../../static/facing-tip.png" />
		</view>
		<view class="camera">
			<canvas type="2d" class="progress_bg" canvas-id="canvasProgressbg"></canvas>
			<canvas type="2d" class="progress_canvas" canvas-id="canvasProgress"></canvas>
			<camera device-position="front" flash="off" binderror="error">
				<cover-view class='cameraGai'>
					<cover-image src="../../static/circle.png" class='cover-image'></cover-image>
				</cover-view>
			</camera>
		</view>
		<view class="tips">{{tip}}</view>
		<button type="error" class="start-facing" @tap="stoptakephone">停止打卡</button>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				background: {
					backgroundColor: '#12c497',
				},
				accesstoken: "",
				student: {},
				userid: "",
				userinfo: "",
				usergroup: "",
				facetoken: "",
				timer: "",
				ifOnShow: false,
				tip: ""
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
			stoptakephone(){
				uni.redirectTo({
					url:"../index/index"
				})
			},
			takephone() {
				console.log("ee")
				const ctx = wx.createCameraContext()
				console.log(ctx)
				var that = this
				ctx.takePhoto({
					quality: 'high',
					success: (res) => {
						console.log(res.tempImagePath)
						var imgbase64Url = wx.getFileSystemManager().readFileSync(res.tempImagePath, "base64");
						
						// data:image/png;base64,
						// console.log(imgbase64Url)
						// that.uploadImg(encodeURIComponent(imgbase64Url))
						// 识别成功之后跳转
						//请求百度接口
						var userid;
						uni.request({
							url: "https://aip.baidubce.com/rest/2.0/face/v3/search?access_token=" + this.accesstoken,
							method: "POST",
							header: {
								'Content-Type': 'application/json',
							},
							data: {
								"image": imgbase64Url,
								"image_type": "BASE64",
								"group_id_list": "school_bus",
								"quality_control": "LOW",
								"liveness_control": "NORMAL"
							},
							success: (res) => {
								console.log(res)
								if (res.data.error_code == 0) {
									this.facetoken = res.data.face_token
									this.$store.commit('changecantakephone', true)
									const token = uni.getStorageSync('token')
									uni.request({
										url: "http://localhost:8000/xcx/auth/faceinfo",
										method: "POST",
										header: {
											'token': token,
										},
										data: {
											"user_id": this.userid
										},
										success: (res) => {
											console.log("--------学生信息------")
											console.log(res)
											//
											if (res.code == 200) {
												this.student = res.data
												this.userid = res.data.data.id
												userid = res.data.data.id
											}
											console.log("this.student")
											console.log(res.data.data.id)
											console.log("----------学生信息--------")
											console.log(res.data)
											console.log("userid----------")
											console.log(userid)
											var token = uni.getStorageSync('token')
											uni.request({
												url: "http://localhost:8000/xcx/auth/swipe",
												method: "POST",
												header: {
													'token': token,
												},
												data: {
													"student_id": res.data.data.id.toString(),
												},
												success: (res) => {
													console.log("打卡结果-  --------")
													console.log(res)
													this.$store.commit('changecantakephone', true)
													this.$store.commit('setstudent', {
														"studentCount": this.$store.state.student.studentCount,
														"studentGetOnCount": this.$store.state.student.studentGetOnCount + 1
														},
													)
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
								} else {
									// uni.showModal({
									// 	title: "识别失败"
									// })
									this.tip = "识别失败"
									console.log("识别失败")
								}
							},
							fail: (err) => {
								console.log(err)
							}
						})
						uni.hideLoading()
					}
				})
			},
			test() {
				console.log("eeeeeeeeeeeeeeeeeee")
			}
		},
		onLoad() {
			console.log(this.$store.state.cantakephone)
			uni.showLoading({
				title: "加载中..."
			})
			console.log(this.accesstoken)

			console.log(this.$store.state.cantakephone)
			//获取百度access_token

			uni.request({
				url: "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=tTyqXmbuYWkl3nGMmCtkL6pY&client_secret=yoG5QAsLuCp9yU4cmXi8Gf2vapOTZGAi&",
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
				if (!this.$store.state.cantakephone) {
					this.takephone()
				} else {
					uni.redirectTo({
						url: "./complete"
					})
				}

			}, 2000)
		},
		onUnload() {
			console.log("this.timer")
			console.log(this.timer)
			if (this.timer) {
				clearInterval(this.timer)
				this.timer = null
			}
		}
	}
</script>

<style lang="scss">
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
		width: 220px;
		height: 220px;
		position: absolute;
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: -99;
	}

	.progress_bg {
		position: absolute;
		width: 220px;
		height: 220px;
	}

	.progress_canvas {
		width: 220px;
		height: 220px;
	}

	.start-facing {
		bottom: 20px;
		height: 100px;
		width: 100px;
		border-radius: 50px;
		border: 3px solid #fff;
		background-color: rgb(228, 63, 25);
		font-weight: bold;
		font-size: 16px;
		padding: 25px 0;
		color: #fff;
		margin-top: 30px;
	}
	.tips{
		height: 30rpx;
		font-weight: bold;
		font-size: 18px;
		color: #ff0000;
		text-align: center;
	}
</style>
