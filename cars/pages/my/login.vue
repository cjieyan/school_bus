<template>
	<view>
		<u-navbar back-text="返回" title="登录"></u-navbar>
		<image width="100%" height="300rpx" src="../../static/login-banner.jpg"></image>
		<view class="page-content">
			<u-button type="primary" class="btn" shape="square" size="medium" open-type="getUserInfo" lang="zh_CN" @getuserinfo="getuserinfo">授权登录</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
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
			getuserinfo: function() {
				uni.showLoading({
					title:"正在登录",
					icon:"none"
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
												url:"../index/index"
											})
											uni.hideLoading()
										},
										fail: (e) => {
											console.log(e)
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
			}
		}
	}
</script>

<style lang="css" >
	page{
		background-color: #fff;
	}
	.page-content{
		padding: 20px 20px;
		margin: 0 auto;
		text-align: center;
	}
	.btn{
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 25px;
	}
</style>
