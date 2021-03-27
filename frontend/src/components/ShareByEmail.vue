<template>
	<div class="shareByEmail">
		<h4 v-if="isActive" class="shareByEmail__title">Введите email</h4>
		<div v-if="isActive" class="shareByEmail__form">
			<input
				v-model.trim="email"
				@input="isShowMsg = false"
				type="email"
				class="shareByEmail__input"
				placeholder="Место для email"
			/>
			<button @click="sendEmail" class="shareByEmail__sendbtn">
				Отправить
			</button>
			<div
				v-show="isShowMsg"
				:style="{ 'background-color': bgcColor }"
				class="shareByEmail__msg"
			>
				{{ showMsg }}
			</div>
		</div>

		<button
			class="shareByEmail__visbtn"
			@click="isActive = true"
			v-else
		>
			<img src="../assets/share.png" alt="Поделиться" />
		</button>
	</div>
</template>

<script>
	import HTTP from "../HTTP";
	import { mapGetters } from "vuex";
	export default {
		name: "ShareByEmail",
		data() {
			return {
				isActive: false,
				isShowMsg: false,
				email: "",
				showMsg: "",
				bgcColor: "red",
			};
		},
		computed: {
			...mapGetters(["GetField"]),
		},
		methods: {
			async sendEmail() {
				this.isShowMsg = false;
				const am_from = this.GetField("amountFrom"),
					am_to = this.GetField("amountTo"),
					from = this.GetField("curFrom"),
					to = this.GetField("curTo");

				if (!this.email || am_from <= 0) {
					this.showMsg = "Проверьте правильность полей";
					this.isShowMsg = true;
					this.bgcColor = "coral";
					return;
				}

				try {
					const res = await HTTP.post("/api/email_share", {
						amount_from: `${am_from}`,
						amount_to: `${am_to}`,
						from: from,
						to: to,
						email: this.email,
					});

					console.log(res.data.res);
					this.isShowMsg = true;
					this.showMsg = "Email успешно отправлен";
					this.bgcColor = "lightgreen";

					setTimeout(() => {
						this.isActive = false;
						this.email = "";
						this.isShowMsg = false;
					}, 1000);
				} catch (e) {
					this.isShowMsg = true;
					this.bgcColor = "coral";
					this.showMsg = e.response.data.error;
					console.error(e.response.status, e.response.data);
				}
			},
		},
	};
</script>

<style lang="postcss">
.shareByEmail {
}

.shareByEmail__form {
	display: flex;
	align-items: center;
}

.shareByEmail__title {
	margin-bottom: 10px;
	color: var(--black_color);
}

.shareByEmail__input {
	margin-right: var(--offset-half);
	box-shadow: var(--main-shadow);
	padding: 10px;
	border: none;
	border-radius: var(--border-radius);
}

.shareByEmail__sendbtn {
	cursor: pointer;
	padding: 10px;
	margin-right: var(--offset-half);
	border: none;
	box-shadow: var(--main-shadow);
	background-color: var(--c1);
	border-radius: var(--border-radius);
	transition: var(--main-transition);
}

.shareByEmail__sendbtn:hover {
	background-color: var(--c3);
}

.shareByEmail__msg {
	transition: var(--main-transition);
	padding: 10px;
	color: var(--black_color);
	font-weight: bold;
	border-radius: var(--border-radius);
	box-shadow: var(--main-shadow);
}

.shareByEmail__visbtn {
	cursor: pointer;
	border: none;
	background-color: transparent;
}

.shareByEmail__visbtn img {
	display: block;
	max-width: 24px;
	height: auto;
}
</style>
