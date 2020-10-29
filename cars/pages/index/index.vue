<template>
	<view>
		<view class="top">
			<u-navbar :is-back="isback" title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff"
			 title-color="#fff"></u-navbar>
		</view>
		<view class="image-content">
			<image src="../../static/banner.jpg" mode="aspectFill" style="width: 100%;" @error="imageError"></image>
		</view>
		<view class="wrap">
			<view v-for="(item, index) in lines" :key="index">
				<u-card :title="item.carIds+'号线'" :sub-title="'发车 '+item.departed_at+'——到达 '+item.arrivedAt">
					<view class="" slot="body"  @tap="selectline(item)" :value="item.id" >
						<view class="u-body-item u-flex u-col-between u-p-t-0">
							<view class="u-body-item-title">
								<u-tag class="u-tag" text="出发" mode="dark" type="success" />{{item.startSite.name}}</view>
						</view>
						<view class="u-body-item u-flex u-col-between u-p-t-0">
							<view class="u-body-item-title">
								<u-tag class="u-tag" text="到达" mode="dark" type="warning" />{{item.endSite.name}}</view>
						</view>
					</view>
				</u-card>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				checked: false,
				isback: false,
				background: {
					backgroundColor: '#12c497',
				},
				isSwiper: false,
				current: 0,
				border: true,
				col: 3,
				lines: {}
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
			selectline(item){
				console.log("----item--------")
				console.log(item)
				const lineid = item.id
				console.log(lineid)
				this.$store.commit('setLineid', lineid)
				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/line-info",
					method:"POST",
					header: {
						'token': this.$store.state.token,
					},
					data:{
						"line_id": lineid
					},
					success: (res) => {
						console.log("--------selectline--------")
						console.log(this.$store.state.lineid)
						console.log(res)
						if(res.data.code == 401){
							uni.showToast({
								icon: 'none',
								title: '会话过期，请重新登录',
								duration: 1500
							});
							uni.redirectTo({
								url:"../my/login"
							})
						}else{
							this.$store.commit('setcarinfo', res.data.data.car)
							this.$store.commit('setTeacher', res.data.data.teacher)
							this.$store.commit('setLineinfo', res.data.data.line)
							this.$store.commit('setSiteinfo', res.data.data.sites)
							this.$store.commit('setstudent', {
								"studentCount": res.data.data.studentCount,
								"studentGetOnCount": res.data.data.studentGetOnCount
							})
							uni.switchTab({
								url:"../location/index",
								fail: (err) => {
									console.log(err)
								}
							})
						}
						
					},
					fail: (err) => {
						console.log(err)
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
					url: this.$store.state.apihost + "/xcx/auth/line-info",
					method: "POST",
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
						}else if(res.data.code == 401){
							uni.showToast({
								icon: 'none',
								title: '会话过期，请重新登录',
								duration: 1500
							});
							uni.redirectTo({
								url:"../my/login"
							})
						}
					},
					fail: (err) => {
						uni.showModal({
							title:"网络异常,请重试"
						})
						return
					}
				})
			},
			lineInfo() {
				var token = uni.getStorageSync('token')
				console.log("lineinfo")
				console.log(token)
				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/lines",
					method: "post",
					header: {
						'token': token,
					},
					data: {},
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
						}else{
							this.lines = res.data.data
						}
					},
					fail: (err) => {
						console.log(err)
					}
				})
			}
		},
		onLoad() {
			uni.showLoading({
				title:"正在加载"
			})
			var login = this.checkLogin('/pages/index/index',2);
			if(login != false){
				this.lineInfo()
			}
			
			uni.hideLoading()
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

	.u-body-item {
		font-size: 26rpx;
		color: #333;
		padding: 10rpx 10rpx;
	}

	u-tag {
		margin-right: 30rpx;
	}

	.indicator-dots {
		margin-top: 40rpx;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.indicator-dots-item {
		background-color: $u-tips-color;
		height: 12rpx;
		width: 12rpx;
		border-radius: 20rpx;
		margin: 0 6rpx;
	}

	.indicator-dots-active {
		background-color: $u-type-primary;
	}
</style>
