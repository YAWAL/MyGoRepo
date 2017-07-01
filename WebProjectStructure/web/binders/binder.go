package binders

/*
Web.Binders
https://habrahabr.ru/post/331456/?utm_source=fb&utm_medium=social&utm_campaign=kak-my-homyaka-yablokami-kormili-ili-effek

В папке binders располагаются биндеры, которые парсят параметры из запросов, конвертируют в удобный формат и закидывают
в контекст для дальнейшей работы.
Пример метода из этого пакета. Он берет параметр из query, конвертирует в bool и кладет в контекст:

func OpenNow(c *gin.Context)  {
    openNow, _ := strconv.ParseBool(c.Query(BindingOpenNow))
    c.Set(BindingOpenNow, openNow)
}

Самый простой вариант без обработки ошибок. Просто для наглядности.
 */
