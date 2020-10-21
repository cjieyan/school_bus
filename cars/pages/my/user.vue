<template>
	<view>
		<u-navbar back-text="返回"  @tap="back" title="用户信息"></u-navbar>

		<view class="u-flex">
			<image src="../../static/login-banner.jpg" mode="heightFix" @error="imageError"></image>
		</view>
		<view class="page-content">
			<view class="u-m-t-20">
				<u-cell-group class="page-item">
					<u-cell-item icon="rmb-circle" @tap="follow" title="跟车记录"></u-cell-item>
				</u-cell-group>
			</view>
			<!-- <view class="u-m-t-20">
				<u-cell-group class="page-item">
					<u-cell-item icon="star" @tap="setting" title="设置"></u-cell-item>
				</u-cell-group>
			</view> -->
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				pic: 'https://uviewui.com/common/logo.png',
				show: true
			}
		},
		methods: {
			back() {
				uni.switchTab({
					url:"../index/index",
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
		}
	}
</script>

<style lang="scss">
	.camera {
		width: 54px;
		height: 44px;

		&:active {
			background-color: #ededed;
		}
	}
	.page-content{
	}
	.user-box {
		background-color: #fff;
	}
</style>
