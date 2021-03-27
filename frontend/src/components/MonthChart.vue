<template>
	<div class="monthlyChart">
		<h3 class="monthlyChart__title">
			График валютной пары {{ GetField("curFrom") }} -
			{{ GetField("curTo") }}
		</h3>
		<div class="monthlyChart__content">
			<ECharts
				class="monthlyChart__cont"
				ref="monthlyChart__content"
				:option="chartOptions"
			/>
		</div>
	</div>
</template>

<script>
	import { mapActions, mapGetters } from "vuex";
	export default {
		name: "MonthChart",
		component: {},
		mounted() {
			this.LoadRangeSlice();
		},
		methods: {
			getColor() {
				return window
					.getComputedStyle(window.document.documentElement)
					.getPropertyValue("--c3");
			},
			...mapActions(["LoadRangeSlice"]),
		},
		computed: {
			getChartColor() {
				const c = window
					.getComputedStyle(window.document.documentElement)
					.getPropertyValue("--c3");
				return c;
			},
			chartOptions() {
				const opt = {
					color: {
						type: "linear",
						colorStops: [
							{
								offset: 0,
								color: this.GetField("chartColor"),
							},
							{
								offset: 1,
								color: this.GetField("chartColor"),
							},
						],
						global: false, // false by default
					},
					tooltip: {
						trigger: "axis",
						axisPointer: {
							type: "cross",
							label: {
								backgroundColor: "#6a7985",
							},
						},
					},
					grid: {
						left: "0%",
						right: "3%",
						bottom: "3%",
						// width: "390",
						height: "200",
						containLabel: true,
						show: true,
					},
					xAxis: [
						{
							type: "category",
							boundaryGap: false,
							data: [...this.GetField("dateSlice")],
						},
					],
					yAxis: [
						{
							// type: "value",
							scale: true,
							splitArea: {
								show: true,
							},
						},
					],
					// yAxis: [
					// 	{
					// 		type: "value",
					// 	},
					// ],
					series: [
						{
							type: "line",
							areaStyle: {},
							data: [...this.GetField("rangeSlice")],
						},
					],
				};

				return opt;
			},
			...mapGetters(["GetField"]),
		},
		data() {
			return {};
		},
	};
</script>

<style lang="postcss">
.monthlyChart {
	grid-area: month_chart;
}

.monthlyChart__title {
	color: var(--black_color);
	padding-bottom: var(--offset-half);

	/* margin-bottom: var(--offset-half); */
	border-bottom: 2px dashed var(--c2);
	max-width: 390px;
}

.monthlyChart__cont {
	min-height: 230px;
	/* max-width: 390px; */
}
</style>>