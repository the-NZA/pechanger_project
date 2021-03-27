module.exports = {
	pages: {
		index: {
			entry: 'src/main.js',
			template: 'public/index.html',
			filename: 'index.html',
			title: 'PECHANGER – Ваш любимый конвертер валют',
			chunks: ['chunk-vendors', 'chunk-common', 'index']
		},
	}
}