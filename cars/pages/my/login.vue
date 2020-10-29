<template>
	<view>
		<u-navbar :is-back="isback" title="登录"></u-navbar>
		<image mode="aspectFill" style="width: 100%;" src="../../static/login-banner.jpg"></image>
		<view class="page-content">
			<!-- <u-input class="login-input" v-model="phone" type="number" value="17620323840" placeholder="请输入手机号码" /> -->
			<!-- <u-input class="login-input" v-model="password" type="password" value="123456" placeholder="请输入密码" /> -->
			<u-field icon="phone" label-width="40" maxlength="20" class="login-input" placeholder="请输入手机号码" v-model="phone" type="number"
			 border="true"></u-field>
			<u-field icon="lock" label-width="40" maxlength="20" class="login-input" type="password" password-icon="true"
			 placeholder="请输入密码" v-model="password" passwordIcon="true" border="true"></u-field>
			<!-- <u-button type="primary" class="btn" shape="square" size="medium" open-type="getUserInfo" lang="zh_CN" @getuserinfo="getuserinfo">登录</u-button> -->
			<!-- <u-button type="primary" size="medium" class="btn" @click="login">登录</u-button>
			 -->
			<button class="btn" @tap="login">登录</button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isback: false,
				phone: "17620323840",
				password: "123456",
				backpage: '',
				backtype: '',
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
			getuserinfo: function() {
				uni.showLoading({
					title: "正在登录",
					icon: "none"
				})
				uni.getProvider({
					service: 'oauth',
					success: function(res) {
						console.log(res)
						if (~res.provider.indexOf('weixin')) {
							uni.login({
								provider: 'weixin',
								success: (res2) => {

									uni.getUserInfo({
										provider: 'weixin',
										success: (info) => { //这里请求接口
											console.log(res2);
											console.log(info);
											uni.switchTab({
												url: "../index/index"
											})
											uni.hideLoading()
										},
										fail: (e) => {
											uni.showToast({
												title: "微信登录授权失败",
												icon: "none"
											});
											uni.hideLoading()
										}
									})

								},
								fail: (err) => {
									uni.showToast({
										title: "微信登录授权失败",
										icon: "none"
									});
								}
							})

						} else {
							uni.showToast({
								title: '请先安装微信或升级版本',
								icon: "none"
							});
						}
					}
				});
			},
			login: function(e) {
				if (this.phone == "") {
					uni.showModal({
						title: "请输入手机号码"
					})
					return
				}
				if (this.password == "") {
					uni.showModal({
						title: "请输入密码"
					})
					return
				}
				if (this.phone != "" && this.password != "") {
					console.log("this.backtype")
					console.log(this.backtype)
					uni.request({
						url: this.$store.state.apihost + "/xcx/login",
						method: "POST",
						data: {
							"phone": this.phone,
							"password": this.password
						},
						success: (res) => {
							console.log(res)
							if (res.data.code == "200") {
								//保存token
								uni.setStorageSync("token", res.data.data.token)
								this.$store.commit("setToken", res.data.data.token)
								if (this.backtype == '1') {
									uni.redirectTo({
										url: this.backpage
									});
								} else{
									console.log("aaaaaaaaaaaaaaaaaa")
									console.log(this.backpage)
									if(this.backpage == undefined){
										this.backpage = "../index/index"
									}
									uni.switchTab({
										url: this.backpage,
										fail: (err) => {
											console.log(err)
										}
									});
								}

								// uni.switchTab({
								// 	url: "../index/index"
								// })
								uni.hideLoading()
								return
							} else if (res.data.code == "-1") {
								uni.showModal({
									title: res.data.msg
								})
								return
							} else {
								uni.showModal({
									title: "登录失败"
								})
								return
							}
						},
						fail: (err) => {
							console.log(err)
							uni.showModal({
								title: "网络异常，请重试"
							})
						}
					})
				} else {
					uni.showModal({
						title: "请输入手机号码/密码"
					})
					return false
				}
			}
		},
		onLoad(options) {
			this.backpage = options.backpage;
			this.backtype = options.backtype;
		}
	}
</script>

<style lang="css">
	page {
		background-color: #fff;
	}

	.page-content {
		padding: 40rpx 16%;
		margin: 0 auto;
		text-align: center;
	}

	.btn {
		margin-top: 20rpx;
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 50rpx;
	}

	.u-border-bottom:after {
		border: none !important
	}

	.u-field {
		border: 2rpx solid #ccc;
		border-radius: 50rpx;
		margin-bottom: 40rpx;
	}

	.btn {
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 50rpx;
	}
	.btn:after{
		display: none;
	}
</style>
