<template>
	<view>
		<u-navbar back-text="返回" title="用户信息" @tap="back" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		<view class="u-flex">
			<image src="../../static/login-banner.jpg" mode="heightFix" @error="imageError"></image>
		</view>
		<view class="page-content">
			<view class="u-m-t-20">
				<u-cell-group class="page-item">
					<u-cell-item icon="info" @tap="follow" title="跟车记录"></u-cell-item>
				</u-cell-group>
			</view>
			<view class="u-m-t-20">
				<u-cell-group class="page-item">
					<u-cell-item icon="star" @tap="loginout" title="退出登录"></u-cell-item>
				</u-cell-group>
			</view>
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
			}
		},
		methods: {
			back() {
				uni.switchTab({
					url: "../index/index",
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
			follow: () => {
				console.log("userinfo")
				uni.redirectTo({
					url: "./follow"
				})
			},
			setting: () => {
				console.log("picker")
				uni.redirectTo({
					url: "./setting"
				})
			},
			loginout() {
				console.log(this)
				this.$store.commit("setToken", null)
				this.$store.commit("setstudent", null)
				this.$store.commit("setcarinfo", null)
				this.$store.commit("setLineinfo", null)
				this.$store.commit("setSiteinfo", null)
				this.$store.commit("setTeacher", null)
				this.$store.commit("setLineid", null)
				
				uni.setStorageSync("token", null)
				console.log(this)
				uni.redirectTo({
					url: "./login",
					success: (res) => {
						uni.showToast({
							icon: 'none',
							title: '已退出登录',
							duration: 1500
						});
					}
				})
			}
		}
	}
</script>

<style lang="scss">
	.camera {
		width: 108rpx;
		height: 88rpx;

		&:active {
			background-color: #ededed;
		}
	}

	.page-content {}

	.user-box {
		background-color: #fff;
	}
</style>
