<template>
	<view>
		<u-navbar title="打卡成功" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		<view class="completeFetch">
			<image src="../../static/complete.png"></image>
		</view>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				"isback":false,
				background: {
					backgroundColor: '#12c497',
				},
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
			finish() {
				uni.showLoading({
					
				})
				var token = uni.getStorageSync('token')
				uni.request({
					url: this.$store.state.apihost+"/xcx/auth/line-finish",
					method: "POST",
					header: {
						'token': token,
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
						setTimeout(() => {
							uni.switchTab({
								url: "../student/index",
								success: (res) => {

								},
								fail: (err) => {
									console.log("-------err------------")
									console.log(err)
								}
							})
						}, 3000)
						uni.hideLoading()
					},
					fail: (err) => {
						console.log(err)
					}
				})
			}
		},
		onLoad() {
			if ( this.$store.state.isfinish ) {
				this.finish()
				
			} else {
				setTimeout(() => {
					uni.redirectTo({
						url: "./index"
					})
				}, 2000)
			}
		}
	}
</script>

<style>
	.completeFetch {
		height: 100vh;
		width: 100%;
		overflow: hidden;
		background: #1B1C19;
		box-sizing: border-box;
		padding: 20%;
	}

	image {
		height: 550rpx;
		width: 450rpx;
		margin: 0 auto;
		/* margin-top: 200rpx; */
	}
</style>
