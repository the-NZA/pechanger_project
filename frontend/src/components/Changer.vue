<template>
	<div class="changer">
		<h2 class="changer__title">Выберите валюты</h2>

		<div class="changer__main">
			<div class="changer__card chcard">
				<div class="chcard__title">У вас есть</div>

				<div class="chcard__body">
					<select
						v-model="curFrom"
						@change="updateValues"
						class="chcard__select"
					>
						<option
							v-for="(cur, idx) in getCurrencies"
							:key="idx"
							:value="cur"
						>
							{{ cur }}
						</option>
					</select>

					<input
						v-model.number="amountFrom"
						class="chcard__input"
						type="number"
						min="0"
					/>
				</div>

				<div class="chcard__rate">
					{{ getFromSingleAmount }}
				</div>
			</div>
			<button
				@click="swapCurrencies"
				class="changer__swap"
			></button>
			<div class="changer__card chcard">
				<div class="chcard__title">Вы получите</div>

				<div class="chcard__body">
					<select
						v-model="curTo"
						@change="updateValues"
						class="chcard__select"
					>
						<option
							v-for="(cur, idx) in getCurrencies"
							:key="idx"
							:value="cur"
						>
							{{ cur }}
						</option>
					</select>
					<div class="chcard__output">
						{{ amountTo }}
					</div>
				</div>

				<div class="chcard__rate">{{ getToSingleAmount }}</div>
			</div>
		</div>
		<share-by-email></share-by-email>
	</div>
</template>

<script>
	import { mapGetters, mapMutations, mapActions } from "vuex";
	import ShareByEmail from "./ShareByEmail.vue";
	import HTTP from "./../HTTP";

	export default {
		name: "Changer",
		components: {
			ShareByEmail,
		},
		data() {
			return {
				flag: true,
				isInputEnd: false,
			};
		},
		computed: {
			getFromSingleAmount() {
				return `1 ${this.GetField("curFrom")} – ${this.roundAmount(
					this.GetField("oneAmountFromInTo")
				)} ${this.GetField("curTo")}`;
			},
			getToSingleAmount() {
				return `1 ${this.GetField("curTo")} – ${this.roundAmount(
					this.GetField("oneAmountToInFrom")
				)} ${this.GetField("curFrom")}`;
			},
			amountTo() {
				return this.roundAmount(this.GetField("amountTo"));
			},
			amountFrom: {
				get() {
					return this.GetField("amountFrom");
				},
				async set(val) {
					if (val === "") {
						this.SetField({
							field: "amountFrom",
							value: 0,
						});
						this.SetField({
							field: "amountTo",
							value: 0,
						});
						return;
					}

					if (this.timeout) clearTimeout(this.timeout);
					this.timeout = setTimeout(async () => {
						this.SetField({
							field: "amountFrom",
							value: val,
						});

						try {
							const r = await HTTP.get(
								`/api/exchange?from=${this.curFrom}&to=${this.curTo}&amount=${val}`
							);

							const newVal = JSON.parse(r.data.res);

							this.SetField({
								field: "amountTo",
								value: newVal,
							});

							this.LoadCurrencyAtDate();
						} catch (e) {
							console.error(
								e.response.status,
								e.response.data
							);
						}
					}, 500);
				},
			},
			getCurrencies() {
				return this.GetField("currencies");
			},
			curFrom: {
				get() {
					return this.GetField("curFrom");
				},
				set(v) {
					this.SetField({
						value: v,
						field: "curFrom",
					});

					this.SetField({
						value: null,
						field: "selectedDate",
					});
				},
			},
			curTo: {
				get() {
					return this.GetField("curTo");
				},
				set(v) {
					this.SetField({
						value: v,
						field: "curTo",
					});

					this.SetField({
						value: null,
						field: "selectedDate",
					});
				},
			},
			...mapGetters(["GetField"]),
		},
		methods: {
			roundAmount(num) {
				return parseInt(num * 1000) / 1000;
			},
			updateValues() {
				this.getSingleAmounts();
				this.LoadRangeSlice();
			},
			async getSingleAmounts() {
				try {
					const res = await HTTP.get(
						`/api/exchange?from=${this.curFrom}&to=${this.curTo}&amount=1`
					);
					const oneFromInTo = JSON.parse(res.data.res);
					this.SetField({
						field: "oneAmountFromInTo",
						value: oneFromInTo,
					});

					const oneToInFrom = 1 / oneFromInTo;
					this.SetField({
						field: "oneAmountToInFrom",
						value: oneToInFrom,
					});
				} catch (e) {
					console.error(e.response.status, e.response.data);
				}
			},
			setCurFrom(v) {
				console.log(v.target.options);
			},
			swapCurrencies() {
				this.SwapItems();
				this.LoadRangeSlice();
			},
			...mapMutations(["SetField", "SwapItems"]),
			...mapActions(["LoadCurrencyAtDate", "LoadRangeSlice"]),
		},
		created() {
			this.getSingleAmounts();
		},
	};
</script>

<style>
.changer {
	grid-area: changer;
}

.changer__title {
	color: var(--black_color);
	text-align: center;
	margin-bottom: var(--offset);
}

.changer__main {
	display: grid;
	grid-template-columns: 1fr 100px 1fr;
	gap: var(--offset);
	margin-bottom: calc(var(--offset) * 2);
}

/* .changer__card {
} */

.changer__swap {
	height: 100px;
	border: none;
	align-self: center;
	cursor: pointer;

	background-color: transparent;

	background-image: url("../assets/swap.svg");
	background-size: 70%;
	background-repeat: no-repeat;
	background-position: center center;
}

.chcard__title {
	margin-bottom: var(--offset-half);
	padding-bottom: var(--offset-half);

	font-weight: bold;
	font-size: 1.1rem;

	color: var(--black_color);

	border-bottom: 2px dashed var(--c2);
}

.chcard__body {
	background-color: var(--white_color);
	padding: var(--offset-half);
	border-radius: var(--border-radius);
	box-shadow: var(--main-shadow);

	margin-bottom: var(--offset);

	display: grid;
	grid-template-columns: 80px auto;
	grid-template-rows: 42px;
}

.chcard__select {
	appearance: none;
	border: 2px solid var(--c2);
	background-color: var(--c2);
	color: var(--white_color);

	border-radius: var(--border-radius) 0 0 var(--border-radius);

	font-weight: bold;
	padding: 0 5px;

	background-image: url("../assets/arrow_down.svg");
	background-position: 50px center;
	background-size: 20px;
	background-repeat: no-repeat;
}

.chcard__select option {
	background-color: var(--white_color);
	color: var(--text-color);
}

.chcard__input {
	padding: 5px 10px;

	border-radius: 0 var(--border-radius) var(--border-radius) 0;
	border: 2px solid var(--c2);
}

.chcard__output {
	padding: 5px 10px;

	color: var(--black_color);

	border-radius: 0 var(--border-radius) var(--border-radius) 0;
	border: 2px solid var(--c2);

	display: flex;
	align-items: center;
}

.chcard__rate {
	background-color: var(--ghost_white_color);
	background-color: var(--c3);
	max-width: max-content;

	color: var(--black_color);

	padding: 10px;
	margin: 0 auto;

	font-weight: bold;
	line-height: 1;

	border-radius: var(--border-radius);
	box-shadow: var(--main-shadow);
}
</style>
