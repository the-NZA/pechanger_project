<template>
	<div class="container">
		<app-header></app-header>
		<app-changer></app-changer>
		<app-change-at-date> </app-change-at-date>
		<app-month-chart></app-month-chart>

		<app-theme-changer v-if="true"></app-theme-changer>
	</div>
</template>

<script>
	import Header from "./components/Header.vue";
	import ThemeChanger from "./components/ThemeChanger.vue";
	import Changer from "./components/Changer.vue";
	import ChangeAtDate from "./components/ChangeAtDate.vue";
	import MonthChart from "./components/MonthChart";

	// import { onMounted } from "vue";
	// import HTTP from "@/HTTP";

	// import { mapGetters } from "vuex";

	export default {
		name: "App",
		components: {
			"app-header": Header,
			"app-changer": Changer,
			"app-change-at-date": ChangeAtDate,
			"app-theme-changer": ThemeChanger,
			"app-month-chart": MonthChart,
		},
		setup() {
			// onMounted(async () => {
			// 	const r = await HTTP.get("/api");
			// 	console.log(r);
			// });

			// const { GetField } = mapGetters(["GetField"]);

			return {};
		},
	};
</script>

<style lang="postcss" >
:root {
	/* Main colors */

	--black_color: #002a29;
	--white_color: #ffffff;
	--orange_color: #ffc145;
	--light_gray_color: #c3c3c3;
	--logo_color: #59585e;
	--ghost_white_color: #f0eff4;

	/* Sizes */
	--site-width: 1000px;

	--offset: 30px;
	--offset-half: calc(var(--offset) / 2);

	/* Shadows */
	--main-shadow: 0 1px 3px rgba(0, 0, 0, 0.12),
		0 1px 2px rgba(0, 0, 0, 0.24);
	--hover-shadow: 0 2px 5px rgba(0, 0, 0, 0.16),
		0 2px 5px rgba(0, 0, 0, 0.23);

	/* Border radious */
	--border-radius: 5px;

	/* Anim func */
	--btn-cub: cubic-bezier(0.25, 0.8, 0.25, 1);
	--main-transition: all 0.3s var(--btn-cub);
}

* {
	box-sizing: border-box;
	padding: 0;
	margin: 0;
	outline: none;
}

body {
	background-color: var(--bg-color);
	color: var(--text-color);

	max-width: var(--site-width);
	margin: 0 auto;

	font-family: Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
}

.container {
	display: grid;
	grid-template-areas:
		"header"
		"changer"
		"change_at_date"
		"month_chart"
		"theme_changer";

	/* grid-auto-columns: auto;
	grid-auto-rows: minmax(100px, 1fr); */

	gap: calc(var(--offset) * 3);
	padding: var(--offset);
}

#app {
}
</style>
