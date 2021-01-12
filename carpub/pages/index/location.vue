<template>
	<view>
		<view class="top" @click="gotoindex">
			<u-navbar is-back="true" backurl="./index" title="智慧校车" class="top" @click="gotoindex" :background="background" back-icon-color="#fff"
			 title-color="#fff"></u-navbar>
		</view>
		<view class="page-body">
			<view class="u-demo">
				<view class="u-demo-wrap">
					<u-grid>
						<u-grid-item name="item1">
							<u-button size="mini">李明</u-button>
						</u-grid-item>
						<u-grid-item name="item1">
							<u-button size="mini">李明</u-button>
						</u-grid-item>
						<u-grid-item name="item1">
							<u-button size="mini">李明</u-button>
						</u-grid-item>
					</u-grid>
					<u-form ref="uForm">
						<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240" label="路线" prop="name">
							<text>西丽湖--南头</text>
						</u-form-item>
						<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240" label="上车地点" prop="name">
							<text>西丽湖</text>
						</u-form-item>
						<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240" label="下车地点" prop="name">
							<text>南头</text>
						</u-form-item>
					</u-form>
					<map style="width: 100%; height: 300px;" :latitude="latitude" :longitude="longitude" :markers="covers">
					</map>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isback: true,
				background: {
					backgroundColor: '#12c497',
				},
				id: 0, // 使用 marker点击事件 需要填写id
				title: 'map',
				latitude: 39.909,
				longitude: 116.39742,
				covers: [{
					latitude: 39.909,
					longitude: 116.39742,
					iconPath: '../../../static/location.png'
				}, {
					latitude: 39.90,
					longitude: 116.39,
					iconPath: '../../../static/location.png'
				}]
			}
		},
		methods: {
			gotoindex() {
				uni.redirectTo({
					url: "../index/index",
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
			}
		}
	}
</script>

<style>

</style>
