<template>
	<view class="u-demo">
		<view class="top">
			<u-navbar :is-back="isback" title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="u-demo-wrap">
			<view class="u-demo-area">
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
					<u-grid-item @tap="unbind">
						<u-icon name="hourglass" :size="46"></u-icon>
						<view class="grid-text">解绑</view>
					</u-grid-item>
				</u-grid>
				<u-button class="btn" type="success">返回</u-button>
			</view>
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
			isSwiperChange(index) {
				this.isSwiper = index == 0 ? true : false;
			},
			borderChange(index) {
				this.border = index == 0 ? true : false;
			},
			colChange(index) {
				this.col = index == 0 ? 3 : 4;
			},
			click(index) {
				this.$refs.uToast.show({
					title: `点击了第${index + 1}宫格`,
					type: 'warning'
				})
			},
			change(e) {
				this.current = e.detail.current;
			},
			// 针对单个grid-item的事件
			itemClick(index) {
				// console.log(index);
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
</style>
