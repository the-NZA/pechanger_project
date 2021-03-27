<template>
	<div class="changeAtDate">
		<div class="picker">
			<h3 class="picker__title changeAtDate__title">
				Выберите дату:
			</h3>
			<datepicker
				v-model="selectedDate"
				:lowerLimit="lowerLimit"
				:upperLimit="upperLimit"
				placeholder="Нажмите для выбора"
			></datepicker>
		</div>

		<div v-if="isDateSelected" class="shower">
			<h3 class="shower__title changeAtDate__title">
				{{ getDescription }}
			</h3>
			<p class="shower__content">{{ currencyAtDate }}</p>
		</div>
	</div>
</template>

<script>
	import Datepicker from "vue3-datepicker";
	import { mapGetters, mapMutations, mapActions } from "vuex";

	export default {
		name: "ChangeAtDate",
		components: {
			datepicker: Datepicker,
		},
		data() {
			return {
				lowerLimit: new Date(2006, 0, 1),
				upperLimit: new Date(),
			};
		},
		computed: {
			selectedDate: {
				get() {
					return this.GetField("selectedDate");
				},
				async set(val) {
					this.SetField({
						field: "selectedDate",
						value: val,
					});

					this.LoadCurrencyAtDate();
				},
			},
			getDescription() {
				const dt = this.GetField("selectedDate").toLocaleDateString(
					"ru"
				);

				return `На ${dt} курс равен:`;
			},
			currencyAtDate() {
				if (this.GetField("selectedDate") === null) {
					return "";
				}

				return `${this.GetField("amountFrom")} ${this.GetField(
					"curFrom"
				)} – ${this.roundAmount(
					this.GetField("currencyAtDate")
				)} ${this.GetField("curTo")}`;
			},
			isDateSelected() {
				return this.selectedDate !== null;
			},
			...mapGetters(["GetField"]),
		},
		methods: {
			roundAmount(num) {
				return parseInt(num * 1000) / 1000;
			},
			...mapMutations(["SetField"]),
			...mapActions(["LoadCurrencyAtDate"]),
		},
	};
</script>

<style lang="postcss">
.changeAtDate {
	grid-area: change_at_date;

	display: flex;
	justify-content: space-between;
	align-items: center;
	align-items: flex-start;
}

.changeAtDate__title {
	margin-bottom: var(--offset-half);
	border-bottom: 2px dashed var(--c2);
	padding-bottom: var(--offset-half);
	width: 100%;
}

.picker {
	flex-basis: 390px;
	box-sizing: border-box;
}

.picker__title {
	color: var(--black_color);
	margin-right: 15px;
}

.v3dp__datepicker {
	max-width: max-content;
}

.v3dp__datepicker input {
	padding: 15px;

	width: 100%;
	border: none;
	border-radius: var(--border-radius);
	box-shadow: var(--main-shadow);
}

.shower {
	color: var(--black_color);
	flex-basis: 390px;
	box-sizing: border-box;
}

.shower__content {
	padding: 15px;
	border-radius: var(--border-radius);
	box-shadow: var(--main-shadow);
	background-color: var(--white_color);
}
</style>
