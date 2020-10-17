<template>
	<view>
		<view class="top">
			<u-navbar :is-back="isback" title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="image-content">
			<image src="../../static/banner.jpg" mode="aspectFill" style="width: 100%;" @error="imageError"></image>
		</view>
		<view class="u-demo-wrap">
			<view class="u-demo-area">
				<u-toast ref="uToast"></u-toast>
				<u-grid :col="col" :border="border">
					<u-grid-item name="item1" @tap="gotocars">
						<u-icon name="scan" :size="46"></u-icon>
						<view class="grid-text">刷脸</view>
					</u-grid-item>
					<u-grid-item :index="1" @tap="carinfo">
						<u-icon name="file-text" :size="46"></u-icon>
						<view class="grid-text">跟车记录</view>
					</u-grid-item>
					<u-grid-item :index="2" @tap="student">
						<u-icon name="account" :size="46"></u-icon>
						<view class="grid-text">学生</view>
					</u-grid-item>
				</u-grid>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isback: false,
				background: {
					backgroundColor: '#12c497',
				},
				list: ['integral', 'kefu-ermai', 'coupon', 'gift', 'scan', 'pause-circle', 'wifi', 'email', 'list'],
				isSwiper: false,
				current: 0,
				border: true,
				col: 3
			};
		},
		methods: {
			back() {
				uni.navigateBack({
					success: function() {
						beforePage.onLoad();
					}
				})
			},
			gotocars() {
				uni.showLoading({

				});
				uni.switchTab({
					url: "../location/index",
					success: function() {
						uni.hideLoading({

						})
					},
					fail: function(e) {
						console.log(e)
					}
				})
			},
			carinfo() {
				uni.redirectTo({
					url: "../my/follow"
				})
			},
			student() {
				uni.switchTab({
					url: "../student/index"
				})
			},
			regist() {
				uni.showLoading({
					title: "加载中"
				})
				uni.redirectTo({
					url: "../face/regist"
				})
				uni.hideLoading()
			},
			//设置相关信息 学生人数、上车人数、车辆信息、、、
			setInfo() {
				var token = uni.getStorageSync('token')
				uni.request({
					url: this.$store.state.apihost+"/xcx/auth/line-info",
					method: "GET",
					header: {
						'token': token,
					},
					data: {},
					success: (res) => {
						if (res.data.code == 200) {
							console.log(res)
							this.$store.commit('setcarinfo', res.data.data.car)
							this.$store.commit('setTeacher', res.data.data.teacher)
							this.$store.commit('setLineinfo', res.data.data.line)
							this.$store.commit('setstudent', {
								"studentCount": res.data.data.studentCount,
								"studentGetOnCount": res.data.data.studentGetOnCount
							})
						}
					},
					fail: (err) => {
						console.log(err)
					}
				})
			}
		},
		onShow() {
			var token = uni.getStorageSync("token")
			if (token == "" || token == "undefinded") {
				uni.showLoading({
					title: "会话已过期"
				})
				uni.redirectTo({
					url: "../my/login"
				})
				uni.hideLoading()
			}
		},
		onLoad() {
			
			var token = uni.getStorageSync("token")
			if (token == "" || token == "undefinded") {
				uni.showLoading({
					title: "会话已过期"
				})
				uni.redirectTo({
					url: "../my/login"
				})
				uni.hideLoading()
			}
			this.setInfo()
		}
		
	};
</script>

<style scoped lang="scss">
	.grid-text {
		font-size: 28rpx;
		margin-top: 4rpx;
		color: $u-type-info;
	}

	.image-content image {
		height: 150px;
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
</style>
