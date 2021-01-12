<template>
	<view>
		<view class="top">
			<u-navbar :is-back="isback" title="智慧校车" @tap="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="bind-info">
			您已绑定的手机号:17620323840
		</view>
		<u-card title="颜长越" sub-title="上车地点:坂田" thumb="http://pic2.sc.chinaz.com/Files/pic/pic9/202002/hpic2119_s.jpg">
			<view class="" slot="body">
				<u-icon name="car" size="34" color="" label="北京-上海-天津路线"></u-icon>
			</view>
		</u-card>
		<view class="u-area">
			<u-toast ref="uToast"></u-toast>
			<u-grid :col="col" @click="click" v-if="!isSwiper" :border="border">
				<u-grid-item name="item1" @tap="location">
					<u-icon name="photo" :size="46"></u-icon>
					<view class="grid-text">实时位置</view>
				</u-grid-item>
				<u-grid-item @tap="changeline">
					<u-icon name="lock" :size="46"></u-icon>
					<view class="grid-text">更换路线</view>
				</u-grid-item>
				<u-grid-item @tap="bind">
					<u-icon name="hourglass" :size="46"></u-icon>
					<view class="grid-text">解绑</view>
				</u-grid-item>
			</u-grid>
			<u-button class="btn" type="success" @tap="gotolines">校车线路</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				list: ['integral', 'kefu-ermai', 'coupon', 'gift', 'scan', 'pause-circle', 'wifi', 'email', 'list'],
				isSwiper: false,
				current: 0,
				border: true,
				col: 3,
				isback: false,
				background: {
					backgroundColor: '#12c497',
				},
			};
		},
		methods: {
			gotolines() {
				uni.redirectTo({
					url:"./lines"
				})
			},
			bind() {
				uni.redirectTo({
					url:"./bind"
				})
			},
			changeline() {
				uni.showLoading({

				})
				uni.redirectTo({
					url: "./changeline"
				})
				uni.hideLoading()
			},
			location() {
				uni.showLoading({

				})
				uni.redirectTo({
					url: "./location"
				})
				uni.hideLoading()
			}
		}
	};
</script>

<style scoped lang="scss">
	.grid-text {
		font-size: 28rpx;
		margin-top: 4rpx;
		color: $u-type-info;
	}

	.badge-icon {
		position: absolute;
		width: 40rpx;
		height: 40rpx;
	}

	.swiper {
		height: 480rpx;
	}

	.indicator-dots {
		margin-top: 40rpx;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.indicator-dots-item {
		background-color: $u-tips-color;
		height: 6px;
		width: 6px;
		border-radius: 10px;
		margin: 0 3px;
	}

	.indicator-dots-active {
		background-color: $u-type-primary;
	}
	.btn{
		margin-top: 30rpx;
	}
	.u-area{
		margin:36rpx
	}
	.bind-info{
		border: 1px solid #e4e7ed;
		margin:36rpx;
		padding: 36rpx;
		border-radius: 12rpx;
	}
</style>
