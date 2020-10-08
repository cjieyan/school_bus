<template>
	<view>
		<u-navbar back-text="返回" title="用户信息"></u-navbar>

		<view class="u-flex">
			<u-image width="100%" height="300rpx" src="../../static/login-banner.jpg"></u-image>
		</view>
		<view class="page-content">
			<view class="u-m-t-20">
				<u-cell-group class="page-item">
					<u-cell-item icon="rmb-circle" @tap="follow" title="跟车记录"></u-cell-item>
				</u-cell-group>
			</view>
			<view class="u-m-t-20">
				<u-cell-group class="page-item">
					<u-cell-item icon="star" @tap="setting" title="设置"></u-cell-item>
				</u-cell-group>
			</view>
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
				uni.navigateBack({
					success:function(){
						beforePage.onLoad();
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
