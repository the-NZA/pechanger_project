import { createStore } from 'vuex'
import HTTP from "./../HTTP";

function getNormalizedDayAndMonth(day, month) {
	if (day < 10) {
		day = `0${day}`;
	}

	if (month < 10) {
		month = `0${month}`
	}

	return { day, month }
}

function getDayAndMonth(date) {
	const month = date.getMonth() + 1;
	const day = date.getDate();
	const norm = getNormalizedDayAndMonth(day, month);

	return { ...norm }
}

export default createStore({
	state: {
		curFrom: "RUB",
		curTo: "EUR",
		amountFrom: 0,
		amountTo: 0,
		// currencies: ["RUB", "USD", "EUR", "GBP"],
		currencies: ["RUB", "USD", "EUR", "BYR", "JPY", "CNY", "CAD", "TMТ", "UAH", "GBP"],
		oneAmountFromInTo: 0,
		oneAmountToInFrom: 0,
		selectedDate: null,
		currencyAtDate: 0,
		rangeSlice: [],
		dateSlice: [],
		chartColor: "#ff33ff",
	},
	getters: {
		GetField: state => field => state[field],
	},
	mutations: {
		SetField: (state, data) => state[data.field] = data.value,
		SetSlice: (state, data) => state[data.field] = [...data.value],
		SwapItems: state => {
			[state.curTo, state.curFrom] = [state.curFrom, state.curTo];
			[state.amountFrom, state.amountTo] = [state.amountTo, state.amountFrom];
			[state.oneAmountFromInTo, state.oneAmountToInFrom] = [state.oneAmountToInFrom, state.oneAmountFromInTo];
			state.selectedDate = null;
		}
	},
	actions: {
		async LoadRangeSlice({ state, commit }) {
			let slice = [];
			let curDate = new Date();
			let year, month, day, tmp;

			year = curDate.getFullYear();
			tmp = getDayAndMonth(curDate);

			day = tmp.day;
			month = tmp.month;
			const endDate = `${year}-${month}-${day}`;

			for (let i = 0; i < 30; ++i) {
				tmp = getDayAndMonth(curDate)
				day = tmp.day;
				month = tmp.month;

				slice.push(`${day}.${month}`);
				curDate.setDate(curDate.getDate() - 1);
			}

			tmp = getDayAndMonth(curDate)
			day = tmp.day;
			month = tmp.month;
			year = curDate.getFullYear();

			slice.push(`${day}.${month}`);

			for (let i = 0, j = slice.length - 1; i < j; i++, j--) {
				[slice[i], slice[j]] = [slice[j], slice[i]];
			}

			const startDate = `${year}-${month}-${day}`;

			commit("SetSlice", { value: slice, field: "dateSlice" });

			try {
				const res = await HTTP.get(`/api/exchange_between?from=${state.curTo}&to=${state.curFrom}&start_date=${startDate}&end_date=${endDate}`)
				commit("SetSlice", { value: res.data.res, field: "rangeSlice" });
			} catch (e) {
				console.error(e.response.status, e.response.data);
			}
		},
		async LoadCurrencyAtDate({ getters, commit }) {
			if (getters.GetField("selectedDate") === null) {
				return;
			}

			if (getters.GetField("amountFrom") === 0) {
				commit("SetField", {
					field: "currencyAtDate",
					value: 0,
				});
				return;
			}

			const year = getters.GetField("selectedDate").getFullYear();
			let month = getters.GetField("selectedDate").getMonth() + 1,
				day = getters.GetField("selectedDate").getDate();

			if (month < 10) {
				month = `0${month}`;
			}

			if (day < 10) {
				day = `0${day}`;
			}
			const dt = `${year}-${month}-${day}`;

			try {
				const res = await HTTP.get(
					`/api/exchange_at?from=${getters.GetField(
						"curFrom"
					)}&to=${getters.GetField(
						"curTo"
					)}&amount=${getters.GetField(
						"amountFrom"
					)}&date=${dt}`
				);

				commit("SetField", {
					field: "currencyAtDate",
					value: JSON.parse(res.data.res)
				})
			} catch (e) {
				console.log(e.response.status, e.response.data);
			}
		}
	},
	modules: {
	}
})
