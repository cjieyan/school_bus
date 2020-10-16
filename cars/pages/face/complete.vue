<template>
	<view>
		<u-navbar title="打卡成功" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		<view class="completeFetch">
			<image src="../../static/complete.png"></image>
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
						console.log(res)
						setTimeout(() => {
							uni.switchTab({
								url: "../student/index",
								success: (res) => {

								},
								fail: (err) => {
									console.log(err)
								}
							})
						}, 3000)
						uni.hideLoading()
					},
					fail: (err) => {
					}
				})
			}
		},
		onLoad() {
			if (this.$store.state.student.studentCount > this.$store.state.student.studentGetOnCount) {
				this.$store.commit('changecantakephone', false)
				setTimeout(() => {
					uni.redirectTo({
						url: "../index/index"
					})
				}, 2000)
			} else {
				this.finish()
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
