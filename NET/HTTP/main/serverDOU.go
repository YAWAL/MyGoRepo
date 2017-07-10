// Объявляем название пакета.
// Все *.go - файлы должны начинаться с названия пакета.
// Пакет с названием "main" имеет специальное назначение - он указывает
// компилятору go, что из этого пакета нужно собрать самодостаточный
// исполняемый файл. Это бинарник, которому для запуска не нужны дополнительные
// зависимости - его достаточно скопировать на нужный компьютер и запустить.
package main

// Импортируем пакеты, используемые в данном файле.
//
// Документацию по стандартным и сторонним пакетам легко найти по адресу
// https://godoc.org/<package_path>.
// Например,
//
//   * https://godoc.org/flag
//   * https://godoc.org/github.com/valyala/fasthttp
import (
	"flag"
	"github.com/valyala/fasthttp"
	"log"
)

// Объявляем глобальную переменную addr, куда будет записано значение параметра
// -addr при запуске программы.
//
// Например, параметр addr станет равным ":80" для следующей строки
// запуска:
//	./server -addr=:80
//
// Пропущенный IP в TCP адресе говорит о том, чтобы сервер "слушал"
// на всех доступных IP-адресах.
//
// flag.String указывает на то, что значение -addr - строка.
// flag.String принимает три аргумента:
//
//   * Название аргумента, который нужно распарсить. "addr" в данном случае.
//   * Значение аргумента по умолчанию. "127.0.0.1:8080" в данном случае.
//   * Описание аргумента, которое выводится при вызове программы
//     с параметром -help.
//
// flag.String возвращает указатель на строку, где хранится значение -addr.
var addr = flag.String("addr", "127.0.0.1:8080",
	"TCP address to listen to for incoming connections")

// main - функция, с которой начинается выполнение программы.
// Эта функцию должна находиться в package main.
func main() {
	// Парсим параметры, указанные в строке запуска программы.
	flag.Parse()

	// Конфигурируем http сервер.
	//
	// См. возможные параметры конфигурации
	// в https://godoc.org/github.com/valyala/fasthttp#Server
	s := fasthttp.Server{
		// Hanlder - функция-обработчик входящих http запросов.
		// См. код функции handlerDOU ниже.
		Handler: handlerDOU,
	}

	// Запускаем сервер.
	//
	// ListenAndServe принимает TCP адрес, где будет запущен сервер.
	// ListenAndServe возвращает результат только в двух случаях:
	//
	//   * Если во во время запуска сервера произошла ошибка.
	//     Например, указанный адрес уже занят другим сервером.
	//     Тогда соответствующая ошибка попадет в err.
	//   * Если сервер был остановлен. Тогда err будет равно nil.
	err := s.ListenAndServe(*addr)
	if err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

// handlerDOU обрабатывает входящие запросы.
func handlerDOU(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello, world!\n")
}