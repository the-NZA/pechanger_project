<template>
	<div class="theme_changer">
		<button
			v-for="(theme, idx) in themes"
			:key="idx"
			@click="setColorTheme(theme.name)"
			class="theme_changer__item"
			:style="{ 'background-color': theme.color }"
		></button>
	</div>
</template>

<script>
	import { reactive } from "vue";
	import { useStore } from "vuex";
	export default {
		name: "ThemeChanger",
		setup() {
			const themes = reactive([
				// { name: "normal", color: "#c7cedf" },
				{ name: "normal", color: "#c4ffff" },
				{ name: "second", color: "#c4c4ff" },
				{ name: "third", color: "#ffc4ff" },
				{ name: "fifth", color: "#ffc4c4" },
				{ name: "sixth", color: "#c4ffc4" },
				{ name: "seventh", color: "#ffffc4" },
				// { name: "dark", color: "#dcdcdd" },
			]);

			const store = useStore();

			const setColorTheme = (theme_name) => {
				document.querySelector("html").dataset.theme = theme_name;
				//TODO: Add store to localstorage with autoload when loaded

				const curColor = window
					.getComputedStyle(window.document.documentElement)
					.getPropertyValue("--c3");

				store.commit("SetField", {
					field: "chartColor",
					value: curColor,
				});
			};

			return {
				themes,
				setColorTheme,
			};
		},
	};
</script>

<style >
/* Default Theme */
:root {
	/* Theme colors */
	/* --text-color: #002a29;
	--bg-color: #c7cedf;

	--c1: #007a79;
	--c2: #976f4f;
	--c3: #4d2a16; */
	--text-color: #66b3b3;
	--bg-color: #c4ffff;

	--c1: #abffff;
	--c2: #b37754;
	--c3: #ffcaab;
}

/* Dark Theme */
html[data-theme="dark"] {
	--text-color: #1d2424;
	--bg-color: #dcdcdd;

	--c1: #69849b;
	--c2: #d0c3bd;
	--c3: #3d6a66;
}

html[data-theme="second"] {
	--text-color: #6666b3;
	--bg-color: #c4c4ff;

	--c1: #ababff;
	--c2: #b3a154;
	--c3: #fff0ab;
}

html[data-theme="third"] {
	--text-color: #b366b3;
	--bg-color: #ffc4ff;

	--c1: #ffabff;
	--c2: #76b354;
	--c3: #c9ffab;
}

html[data-theme="fourth"] {
	--text-color: #66b3b3;
	--bg-color: #c4ffff;

	--c1: #abffff;
	--c2: #b37754;
	--c3: #ffcaab;
}

html[data-theme="fifth"] {
	--text-color: #b36666;
	--bg-color: #ffc4c4;

	--c1: #ffabab;
	--c2: #54b36e;
	--c3: #abffc2;
}

html[data-theme="sixth"] {
	--text-color: #66b366;
	--bg-color: #c4ffc4;

	--c1: #abffab;
	--c2: #b35484;
	--c3: #ffabd6;
}

html[data-theme="seventh"] {
	--text-color: #b3b366;
	--bg-color: #ffffc4;

	--c1: #ffffab;
	--c2: #8254b3;
	--c3: #d4abff;
}

.theme_changer {
	grid-area: theme_changer;
	border: 1px solid var(--orange_color);
	padding: 10px;
	margin: 0 auto;

	max-width: fit-content;

	background-color: var(--white_color);

	display: flex;
}

.theme_changer__item {
	/* width: 20px;
	height: 20px; */
	width: var(--offset);
	height: var(--offset);

	margin-right: 10px;

	border: none;
	box-shadow: var(--main-shadow);
	border-radius: var(--border-radius);

	cursor: pointer;
	transition: var(--main-transition);
}

.theme_changer__item:last-of-type {
	margin-right: 0;
}

.theme_changer__item:hover {
	box-shadow: var(--hover-shadow);
}
</style>
