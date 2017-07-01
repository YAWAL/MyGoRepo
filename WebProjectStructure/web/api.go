package web

/*
Web

https://habrahabr.ru/post/331456/?utm_source=fb&utm_medium=social&utm_campaign=kak-my-homyaka-yablokami-kormili-ili-effek

В вебе хранится все, что отвечает за обработку запросов: байндеры, фильтры и контроллеры — а вся их спайка происходит
в api.go. Пример такого склеивания:

regions := r.Group("/regions")
regions.GET("/list", Cache.Gin, rc.List)
regions.GET("/list/active", Cache.Gin, regionController.ListActive)
regions.GET("", binders.Coordinates, regionController.RegionByCoord)

Там же происходит инициализация контроллеров и инъекция зависимостей. По сути весь api.go файл состоит из метода Run,
 где формируется и стартуется роутер, и кучи вспомогательных методов по созданию контроллеров со всеми зависимостями
 и их групп.
 */
