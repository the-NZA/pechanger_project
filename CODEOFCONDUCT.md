# Командное соглашение о чистоте кода
1. Все имена - емкие и однозначные, для временных переменных допускаются короткие имена. 
	+ pointX - координата точки, 
	+ getOneChar() - название не экспортируемой функции,
	+ GetAnotherChar() – название экспортируемой функции.
2. Размер функции - до 75 строк, в силу особенностей синтаксиса языка и метода обработки ошибок.
3. В комментариях (минимум один на функцию/метод) указывается: 
	+ автор каждой функции/метода, 
	+ аннотации к экспортируемым функциям и замечания о важнейших решениях разработчика ("почему так?").
4. Форматирование кода: горизотальные отступы:
	+ 8 символов для backend кода,
	+ 6 символов для frontend кода.
5. Код, не попадающий под правила, отмечается тестировщиком и проходит рефакторинг.
