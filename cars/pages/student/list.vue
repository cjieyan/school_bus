<template>
	<view>
		<view class="top">
			<u-navbar back-text="返回" title="学生列表"></u-navbar>
		</view>
		<uni-card mode="basic">
			<view class="student">
				<view class="student-top">
					<view class="location-info">
						梧桐小学--六约新村路线
					</view>
					<view class="datetime">
						日期:2020/09/15
					</view>
				</view>
				<view class="student-list">
					<view class="student-list-info">
						<view class="student-img">
							<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
							<view class="inboard">
								<span>已上车</span>
								<fa-icon type="check" size="14" color="white" class="font-icon inboard-check"></fa-icon>
							</view>
						</view>
						<view class="student-name">
							张三
						</view>
					</view>
					<view class="student-list-info">
						<view class="student-img">
							<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
							<view class="outboard">
								<span>未上车</span>
							</view>
						</view>
						<view class="student-name">
							张三
						</view>
					</view>
					<view class="student-list-info">
						<view class="student-img">
							<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						</view>
						<view class="student-name">
							张三
						</view>
					</view>
					<view class="student-list-info">
						<view class="student-img">
							<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						</view>
						<view class="student-name">
							张三
						</view>
					</view>
					<view class="student-list-info">
						<view class="student-img">
							<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						</view>
						<view class="student-name">
							张三
						</view>
					</view>
					<view style="clear:both;height:0"></view>
				</view>
			</view>
		</uni-card>
		<uni-card mode="basic" class="commnet">
			<view class="comment-top">
				评价内容:
			</view>
			<view class="comment-list">
				<span>靠谱开车稳</span>
				<span>靠谱开车稳</span>
				<span>靠谱开车稳</span>
				<span>靠谱开车稳</span>
				<span>靠谱开车稳</span>
			</view>
		</uni-card>
	</view>
</template>

<script>
	// import faIcon from "@/components/kilvn-fa-icon/fa-icon.vue"
	// import uniCard from '@/components/uni-card/uni-card.vue'
	export default {
		components: {
			// faIcon,
			// uniCard,
		},
		data() {
			return {

			}
		},
		methods: {

		}
	}
</script>

<style>
	page {
		background-color: #f8f8f8;
	}

	.student {
		margin-top: 10px;
	}

	.student-top {
		border-bottom: 1px solid rgb(238 238 238);
		padding-bottom: 10px;
	}

	.location-info {
		font-size: 14px;
	}

	.datetime {
		margin-top: 5px;
		font-size: 14px;
	}

	.student-list-info {
		width: 40px;
		float: left;
		margin: 15px 10px;
	}

	.student-img {
		width: 40px;
		height: 40px;
		border-radius: 20px;
		text-align: center;
		position: relative;
	}

	.outboard,
	.inboard {
		background-color: #808080;
		width: 42px;
		height: 42px;
		position: absolute;
		top: 0;
		left: 0;
		border-radius: 20px;
		opacity: 0.9;
	}

	.outboard span,
	.inboard span {
		font-size: 8px;
		color: #fff;
		position: absolute;
		top: 12px;
		left: 3px;
	}

	.inboard-check {
		position: absolute;
		bottom: -6px;
		left: 15px;
		background-color: #4CD964;
		border-radius: 6px;
	}

	.student-name {
		margin-top: 8px;
		font-size: 12px;
		text-align: center;
	}
	
	.commnet{
		font-size: 12px;
		background-color: #f8f8f8;
	}
	.comment-list span{
		border: 1px solid #f8f8f8;
		border-radius: 5px;
		padding: 5px 8px;
		background-color: #fff;
		margin-right: 8px;
		margin-top: 10px;
		display: inline-block;
	}
</style>
