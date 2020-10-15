<template class="container">
	<view class="page-content">
		<u-navbar title="跟车记录" @tap="back" class="top" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		<view class="infolist" @click="popo">
			<view class="info-content">
				<view class="car-info">
					<span>2020/9/13</span> <span>1号车</span> <span>(粤B1234)</span>
				</view>
				<view class="people-info">
					<span>乘车人数</span> <span>14人</span>
				</view>
				<view class="people-detail">
					<span>已上车14人</span> <span>未下车14人</span><span>
						<u-icon name="error" color="red" class="error-icon"></u-icon>未上车14人
					</span>
				</view>
			</view>
		</view>
		<view class="infolist">
			<view class="info-content">
				<view class="car-info">
					<span>2020/9/13</span> <span>1号车</span> <span>(粤B1234)</span>
				</view>
				<view class="people-info">
					<span>乘车人数</span> <span>14人</span>
				</view>
				<view class="people-detail">
					<span>已上车14人</span> <span>未下车14人</span><span>
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
	</view>
</template>

<script>
	// import faIcon from "@/components/kilvn-fa-icon/fa-icon.vue"
	export default {
		components: {
			// faIcon
		},
		data() {
			return {
				background: {
					backgroundColor: '#12c497',
				},
				show: false
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
					url:"../index/index",
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
		}
	}
</script>

<style lang="css">
	page {
		background-color: #f8f8f8;
	}

	.popo-title {
		margin-top: 10px;
		padding: 10px 0;
		font-weight: bold;
		text-align: center;
		font-size: 14px;
	}

	.popo-content {
		padding-bottom: 20px;
		margin: 20px auto;
		text-align: center;
	}

	.popo-content span {
		display: block;
		text-align: center;
		line-height: 25px;
	}

	.popo-content span i {
		font-size: 16px;
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
		border-radius: 10px;
		margin: 0 15px;
		background-color: #fff;
		box-shadow: 0 -1px 10px rgba(0, 0, 0, 0.1);
	}

	.info-content {
		margin: 15px 10px;
		padding: 5px 10px;
	}

	.car-info,
	.people-info,
	.people-detail {
		margin-top: 10px;
		font-size: 12px;
		margin-bottom: 10px;
	}

	.car-info span,
	.people-info span,
	.people-detail span {
		margin-right: 8px;
	}

	.error-icon {
		background: #f5baba;
		padding: 2px;
		border-radius: 8px;
		margin-right: 5px;
	}

	.btn {
		margin-top: 20px;
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
	}
	.popo-content span i{
		display: inline-block;
	}
</style>
