<template>
	<header class="header">
		<h1 class="header__title">
			<a href="/">pechanger</a>
		</h1>
	</header>
</template>

<script>
	export default {
		name: "Header",
	};
</script>

<style>
.header {
	grid-area: header;

	width: 100%;

	background-color: var(--white_color);
	box-shadow: var(--main-shadow);
	border-radius: var(--border-radius);

	margin: 0 auto;
	padding: var(--offset);
}

.header__title {
	text-transform: uppercase;
	text-align: center;

	margin: 0;
	padding: 0;

	color: var(--logo_color);
}

.header__title a {
	color: var(--logo_color);
	text-decoration: none;
}
</style>
