<template class="container">
	<view class="page-content">
		<u-navbar title="跟车记录" @tap="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		<view class="infolist" @click="popo" v-for="(item, index) in datalist" :key="index">
			<view class="info-content">
				<view class="car-info">
					<span>2020/9/13  (id：{{item.id}})</span> <span>1号车</span> <span>(粤B1234)</span>
				</view>
				<view class="people-info">
					<span>乘车人数</span> <span>{{item.allCount}}人</span>
				</view>
				<view class="people-detail">
					<span>已上车{{item.getOn}}人</span> <span>未下车{{item.getOff}}人</span><span>
						<u-icon name="error" color="red" class="error-icon"></u-icon>未上车14人
					</span>
				</view>
			</view>
		</view>
		<u-popup v-model="show" mode="bottom" border-radius="18" class="popo">
			<view class="popo-title">
				2020/9/13 1号车 (粤B1234)
			</view>
			<u-line color="gray" length="95%" style="margin: 0 auto;"></u-line>
			<view class="popo-content">
				<span>已上车<i>15</i>人</span>
				<span>已上车<i>15</i>人</span>
				<span>已上车<i>15</i>人</span>
				<u-button type="primary" class="btn" shape="square" size="medium" @click="popo">确定</u-button>
			</view>
		</u-popup>
		<u-loadmore v-if="showLoadMore" :status="status" :icon-type="iconType" :load-text="loadText" />
	</view>
</template>

<script>
	import utils from '../../common/utils.js'
	export default {
		components: {
			// faIcon
		},
		data() {
			return {
				background: {
					backgroundColor: '#12c497',
				},
				datalist: {},
				show: false,
				pagesize: 10,
				currentPage: 1,
				pagetotal: 0,
				pullcount: 0,
				showLoadMore: false,
				status: 'loadmore',
				iconType: 'flower',
				loadText: {
					loadmore: '轻轻上拉',
					loading: '努力加载中',
					nomore: '实在没有了'
				}
			}
		},
		methods: {
			popo() {
				console.log(this.show)
				this.show = this.show == false ? true : false;
				console.log(this.show)
			},
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
			getData() {
				var token = uni.getStorageSync('token')

				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/follow-record",
					method: "POST",
					header: {
						'token': token,
					},
					data: {
						"page_index": 1,
						"page_size": this.pagesize,
					},
					success: (res) => {
						if (res.data.code == 200) {
							console.log('-totalpage-')
							console.log(res.data.data.totalpage)
							this.pagetotal = res.data.data.totalpage
							console.log("totalpage")
							console.log(res.data.data.totalpage)
							if (this.pagetotal > 1) {
								this.showLoadMore = true
								this.currentPage++
							}
							this.datalist = res.data.data.result
						}
						console.log("--res--")
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
			getMoreData() {
				var token = uni.getStorageSync('token')
				console.log("------getmoredata-----")
				console.log(this.currentPage)
				console.log(this.pagesize)
				console.log(this.pagetotal)
				if (this.currentPage <= this.pagetotal) {
					uni.request({
						url: this.$store.state.apihost + "/xcx/auth/follow-record",
						method: "POST",
						header: {
							'token': token,
						},
						data: {
							"page_index": this.currentPage,
							"page_size": this.pagesize,
						},
						success: (res) => {
							console.log("-----getmore ---- res-----")
							if (res.data.code == 200) {
								this.pagetotal = res.data.data.totalpage
								this.showLoadMore = true
								this.currentPage++
								if( this.currentPage >= this.pagetotal)
								{
									this.showLoadMore = false
								}
								this.datalist = this.datalist.concat(res.data.data.result)
							}
						},
						fail: (err) => {
							console.log(err)
						}
					})
				}else{
					return false
				}

			}
		},
		onLoad() {
			uni.showLoading({

			})
			this.getData()
			uni.hideLoading()
		},
		onReachBottom() {
			console.log("-------已到达底部-------")
			if (this.currentPage <= this.pagetotal) {
				utils.throttle(this.getMoreData())
			}
		}
	}
</script>

<style lang="css">
	page {
		background-color: #f8f8f8;
	}

	.popo-title {
		margin-top: 20rpx;
		padding: 20rpx 0;
		font-weight: bold;
		text-align: center;
		font-size: 28rpx;
	}

	.popo-content {
		padding-bottom: 40rpx;
		margin: 40rpx auto;
		text-align: center;
	}

	.popo-content span {
		display: block;
		text-align: center;
		line-height: 50rpx;
	}

	.popo-content span i {
		font-size: 32rpx;
		font-weight: bold;
	}

	.top {
		background-color: #12c497;
		color: #fff;
	}

	.page-content {
		margin: 0 auto;
	}

	.infolist {
		background-image: linear-gradient(-45deg, transparent 80%, rgb(242 237 183), transparent 90%);
		border-radius: 20rpx;
		margin: 0 30rpx;
		background-color: #fff;
		box-shadow: 0 -2rpx 20rpx rgba(0, 0, 0, 0.1);
	}

	.info-content {
		margin: 30rpx 20rpx;
		padding: 10rpx 20rpx;
	}

	.car-info,
	.people-info,
	.people-detail {
		margin-top: 20rpx;
		font-size: 24rpx;
		margin-bottom: 20rpx;
	}

	.car-info span,
	.people-info span,
	.people-detail span {
		margin-right: 16rpx;
	}

	.error-icon {
		background: #f5baba;
		padding: 4rpx;
		border-radius: 16rpx;
		margin-right: 10rpx;
	}

	.btn {
		margin-top: 40rpx;
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
	}

	.popo-content span i {
		display: inline-block;
	}
</style>
