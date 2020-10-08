<template>
	<view class="face">
		<view class="top">
			<u-navbar title="人脸注册" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
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
		<button type="error" class="start-facing" @tap="takephone">开始刷脸</button>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				background: {
					backgroundColor: '#12c497',
				},
				accesstoken: ""
			}
		},
		methods: {
			back() {
				uni.navigateBack({
					success:function(){
						beforePage.onLoad();
					}
				})
			},
			takephone() {
				console.log("ee")
				const ctx = wx.createCameraContext()
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
						console.log(imgbase64Url)
						uni.showLoading({
							title: "正在识别中"
						})
						console.log("https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add?access_token=" + this.accesstoken)
						//请求百度人脸注册接口
						uni.request({
							url: "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add?access_token=" + this.accesstoken,
							method: "POST",
							header: {
								'Content-Type':'application/json',
							},
							data: {
								"image": imgbase64Url,
								"image_type": "BASE64",
								"group_id": "school_bus",
								"user_id": "user1",
								"user_info": "abc",
								"quality_control": "LOW",
								"liveness_control": "NORMAL"
							},
							success: (res) => {
								console.log(res)
								if(res.data.error_code == 0){
									uni.showModal({
										title:"注册成功"
									})
								}else{
									uni.showModal({
										title:"注册失败，请联系管理员"
									})
								}
								uni.hideLoading()
							},
							fail: (err) => {
								console.log(err)
							}
						})
						//本地记录token和url
						uni.hideLoading()
					}
				})
			}
		},
		onLoad() {
			uni.showLoading({
				title: "加载中..."
			})
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
</style>
