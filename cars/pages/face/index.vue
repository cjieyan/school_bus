<template>
	<view class="face">
		<view class="top">
			<u-navbar title="人脸识别" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<!-- <view class="facing-tip">
			<image src="../../static/facing-tip.png" />
		</view> -->
		<view :class="tipclass">{{tip}}</view>
		<view class="camera">
			<canvas type="2d" class="progress_bg" canvas-id="canvasProgressbg"></canvas>
			<canvas type="2d" class="progress_canvas" canvas-id="canvasProgress"></canvas>
			<camera device-position="front" flash="off" binderror="error">
				<cover-view class='cameraGai'>
					<cover-image src="../../static/circle.png" class='cover-image'></cover-image>
				</cover-view>
			</camera>
		</view>
		<view class="bottom-btn">
			<button type="error" class="start-facing" @tap="stoptakephone">停止打卡</button>
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
				tipclass: "",
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
				faceError: [{
						"code": "222001",
						"msg": "必要参数未传入"
					},
					{
						"code": "222002",
						"msg": "参数格式错误"
					},
					{
						"code": "222003",
						"msg": "参数格式错误"
					},
					{
						"code": "222004",
						"msg": "参数格式错误"
					},
					{
						"code": "222005",
						"msg": "参数格式错误"
					},
					{
						"code": "222006",
						"msg": "参数格式错误"
					},
					{
						"code": "222007",
						"msg": "参数格式错误"
					},
					{
						"code": "222008",
						"msg": "参数格式错误"
					},
					{
						"code": "222009",
						"msg": "参数格式错误"
					},
					{
						"code": "222010",
						"msg": "参数格式错误"
					},
					{
						"code": "222011",
						"msg": "参数格式错误"
					},
					{
						"code": "222012",
						"msg": "参数格式错误"
					},
					{
						"code": "222013",
						"msg": "参数格式错误"
					},
					{
						"code": "222014",
						"msg": "参数格式错误"
					},
					{
						"code": "222015",
						"msg": "参数格式错误"
					},
					{
						"code": "222016",
						"msg": "参数格式错误"
					},
					{
						"code": "222017",
						"msg": "参数格式错误"
					},
					{
						"code": "222018",
						"msg": "参数格式错误"
					},
					{
						"code": "222019",
						"msg": "参数格式错误"
					},
					{
						"code": "222020",
						"msg": "参数格式错误"
					},
					{
						"code": "222021",
						"msg": "参数格式错误"
					},
					{
						"code": "222022",
						"msg": "参数格式错误"
					},
					{
						"code": "222023",
						"msg": "参数格式错误"
					},
					{
						"code": "222024",
						"msg": "参数格式错误"
					},
					{
						"code": "222025",
						"msg": "参数格式错误"
					},
					{
						"code": "222026",
						"msg": "参数格式错误"
					},
					{
						"code": "222027",
						"msg": "参数格式错误"
					},
					{
						"code": "222028",
						"msg": "参数格式错误"
					},
					{
						"code": "222029",
						"msg": "参数格式错误"
					},
					{
						"code": "222030",
						"msg": "参数格式错误"
					},
					{
						"code": "222200",
						"msg": "参数格式错误"
					},
					{
						"code": "222201",
						"msg": "服务端请求失败	"
					},
					{
						"code": "222202",
						"msg": "图片中没有人脸	"
					},
					{
						"code": "222203",
						"msg": "无法解析人脸"
					},
					{
						"code": "222204",
						"msg": "参数格式错误"
					},
					{
						"code": "222205",
						"msg": "服务端请求失败	"
					},
					{
						"code": "222206",
						"msg": "服务端请求失败	"
					},
					{
						"code": "222207",
						"msg": "识别失败"
					},
					{
						"code": "222208",
						"msg": "参数格式错误"
					},
					{
						"code": "222209",
						"msg": "参数格式错误"
					},
					{
						"code": "222210",
						"msg": "参数格式错误"
					},
					{
						"code": "222300",
						"msg": "参数格式错误"
					},
					{
						"code": "222301",
						"msg": "参数格式错误"
					},
					{
						"code": "222302",
						"msg": "参数格式错误"
					},
					{
						"code": "222303",
						"msg": "获取人脸图片失败"
					},
					{
						"code": "223100",
						"msg": "参数格式错误"
					},
					{
						"code": "223101",
						"msg": "参数格式错误"
					},
					{
						"code": "223102",
						"msg": "参数格式错误"
					},
					{
						"code": "223103",
						"msg": "参数格式错误"
					},
					{
						"code": "223104",
						"msg": "参数格式错误"
					},
					{
						"code": "223105",
						"msg": "参数格式错误"
					},
					{
						"code": "223106",
						"msg": "识别失败，请重试"
					},
					{
						"code": "223107",
						"msg": "参数格式错误"
					},
					{
						"code": "223108",
						"msg": "参数格式错误"
					},
					{
						"code": "223109",
						"msg": "识别失败"
					},
					{
						"code": "223110",
						"msg": "目标用户组不存在"
					},
					{
						"code": "223111",
						"msg": "参数格式错误"
					},
					{
						"code": "223112",
						"msg": "人脸有被遮挡"
					},
					{
						"code": "223113",
						"msg": "参数格式错误"
					},
					{
						"code": "223114",
						"msg": "人脸模糊"
					},
					{
						"code": "223115",
						"msg": "人脸光照不好"
					},
					{
						"code": "223116",
						"msg": "人脸不完整"
					},
					{
						"code": "223117",
						"msg": "app_list包含app数量过多"
					},
					{
						"code": "223118",
						"msg": "质量控制项错误	"
					},
					{
						"code": "223119",
						"msg": "活体控制项错误	"
					},
					{
						"code": "223120",
						"msg": "活体检测未通过	"
					},
					{
						"code": "223121",
						"msg": "请勿遮挡左眼"
					},
					{
						"code": "223122",
						"msg": "请勿遮挡右眼"
					},
					{
						"code": "223123",
						"msg": "请勿遮挡左脸颊"
					},
					{
						"code": "223124",
						"msg": "请勿遮挡右脸颊"
					},
					{
						"code": "223125",
						"msg": "请勿遮挡下巴"
					},
					{
						"code": "223126",
						"msg": "请勿遮挡鼻子"
					},
					{
						"code": "223127",
						"msg": "请勿遮挡嘴巴"
					},
					{
						"code": "222901",
						"msg": "请重新尝试"
					},
					{
						"code": "222902",
						"msg": "请重新尝试"
					},
					{
						"code": "222903",
						"msg": "请重新尝试"
					},
					{
						"code": "222904",
						"msg": "请重新尝试"
					},
					{
						"code": "222905",
						"msg": "请重新尝试"
					},
					{
						"code": "222906",
						"msg": "请重新尝试"
					},
					{
						"code": "222907",
						"msg": "请重新尝试"
					},
					{
						"code": "222908",
						"msg": "请重新尝试"
					},
					{
						"code": "222909",
						"msg": "请重新尝试"
					},
					{
						"code": "222910",
						"msg": "请重新尝试"
					},
					{
						"code": "222911",
						"msg": "请重新尝试"
					},
					{
						"code": "222912",
						"msg": "请重新尝试"
					},
					{
						"code": "222913",
						"msg": "请重新尝试"
					},
					{
						"code": "222914",
						"msg": "请重新尝试"
					},
					{
						"code": "222915",
						"msg": "请重新尝试"
					},
					{
						"code": "222916",
						"msg": "请重新尝试"
					},
					{
						"code": "222304",
						"msg": "图片尺寸太大"
					},
					{
						"code": "222305",
						"msg": "当前版本不支持图片存储"
					},
					{
						"code": "223128",
						"msg": "正在清理该用户组的数据"
					},
					{
						"code": "222361",
						"msg": "请重新尝试"
					},
					{
						"code": "222046",
						"msg": "参数格式错误"
					},
					{
						"code": "222101",
						"msg": "参数格式错误"
					},
					{
						"code": "222102",
						"msg": "参数格式错误"
					},
					{
						"code": "222307",
						"msg": "图片非法 鉴黄未通过"
					},
					{
						"code": "222308",
						"msg": "图片非法，含有政治敏感人物"
					},
					{
						"code": "222211",
						"msg": "模板图质量不合格"
					},
					{
						"code": "222212",
						"msg": "人脸融合失败"
					},
					{
						"code": "223129",
						"msg": "请使用面向正前方的人脸图片"
					},
					{
						"code": "223130",
						"msg": "参数格式错误"
					},
					{
						"code": "223131",
						"msg": "合成图检测未通过"
					},
				]
			}
		},
		methods: {
			back() {
				uni.navigateBack({
					success: function() {
						beforePage.onLoad();
						clearInterval(this.timer)
					}
				})
			},
			stoptakephone() {
				if (this.timer) {
					clearInterval(this.timer)
					this.timer = null
				}
			},
			takephone() {
				this.tip = "正在识别"
				this.tipclass = "tips yellow"
				console.log("ee")
				const ctx = wx.createCameraContext()
				console.log(ctx)
				var that = this
				ctx.takePhoto({
					quality: 'high',
					success: (res) => {
						const imgbase64Url = wx.getFileSystemManager().readFileSync(res.tempImagePath, "base64");
						//请求百度接口
						// var userid = ""
						uni.request({
							url: "https://aip.baidubce.com/rest/2.0/face/v3/search?access_token=" + this.accesstoken,
							method: "POST",
							header: {
								'Content-Type': 'application/json',
							},
							data: {
								"image": imgbase64Url,
								"image_type": "BASE64",
								"group_id_list": "classa",
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
										url: this.$store.state.apihost + "/xcx/auth/faceinfo",
										method: "POST",
										header: {
											'token': token,
										},
										data: {
											"user_id": this.userid
										},
										success: (res) => {
											console.log("--------学生信息start------")
											console.log(res)
											console.log(res.data.data)
											//
											if (res.data.code == 200) {
												this.tipclass = "tips green"
												this.tip = "识别成功"
												this.student = res.data
												this.userid = res.data.data.id
											}
											console.log("this.student")
											console.log(res.data.data.id)
											console.log("----------学生信息 end--------")
											console.log(res.data)
											var token = uni.getStorageSync('token')
											uni.request({
												url: this.$store.state.apihost + "/xcx/auth/swipe",
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
													if ( 200 == res.data.code){
														this.tipclass = "tips green"
													}else{
														this.tipclass = "tips red"
													}
													this.tip = res.data.msg
													this.$store.commit('changecantakephone', true)
													this.$store.commit('setstudent', {
														"studentCount": this.$store.state.student.studentCount,
														"studentGetOnCount": this.$store.state.student.studentGetOnCount + 1
													}, )
													uni.redirectTo({
														url: "./complete"
													})
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
									this.tipclass = "tips red"
									for(var i = 0; i< this.faceError.length; i++){
										if(this.faceError[i].code == res.data.error_code){
											this.tip = this.faceError[i].msg
											break
										}
									}
									if (this.tip == "") {
										this.tip = "识别失败"
									}

								}
							},
							fail: (err) => {
								console.log(err)
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
			test() {
				console.log("eeeeeeeeeeeeeeeeeee")
			}
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
				url: "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id="+this.$store.state.client_id+"&client_secret="+this.$store.state.client_secret+"&",
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
		width: 300px;
		height: 300px;
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

	.bottom-btn {
		position: fixed;
		bottom: 50rpx;
		display: flex;
		width: 100%;
		justify-content: center;
	}

	.start-facing {
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

	.tips {
		height: auto;
		font-weight: bold;
		font-size: 55rpx;
		text-align: center;
		margin: 60rpx 0;
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
</style>
