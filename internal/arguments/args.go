package arguments

import (
	"flag"
)

var link string

/*
-V (--version) - вывести версию программы
-h (--help) - вывести справку
-b (--background) - работать в фоновом режиме
-o файл (--out-file) - указать лог файл
-d (--debug) - включить режим отладки
-v (--verbose) - выводить максимум информации о работе утилиты
-q (--quiet) - выводить минимум информации о работе
-i файл (--input-file) - прочитать URL из файла
--force-html - читать файл указанный в предыдущем параметре как html
-t (--tries) - количество попыток подключения к серверу
-O файл (--output-document) - файл в который будут сохранены полученные данные
-с (--continue) - продолжить ранее прерванную загрузку
-S (--server-response) - вывести ответ сервера
--spider - проверить работоспособность URL
-T время (--timeout) - таймаут подключения к серверу
--limit-rate - ограничить скорость загрузки
-w (--wait) - интервал между запросами
-Q (--quota) - максимальный размер загрузки
-4 (--inet4only) - использовать протокол ipv4
-6 (--inet6only) - использовать протокол ipv6
-U (--user-agent)- строка USER AGENT отправляемая серверу
-r (--recursive)- рекурсивная работа утилиты
-l (--level) - глубина при рекурсивном сканировании
-k (--convert-links) - конвертировать ссылки в локальные при загрузке страниц
-P (--directory-prefix) - каталог, в который будут загружаться файлы
-m (--mirror) - скачать сайт на локальную машину
-p (--page-requisites) - во время загрузки сайта скачивать все необходимые ресурсы
*/

type Options struct {
	Help    bool
	Version bool
	Output  string
}

var (
	help    = flag.Bool("h", false, "help")
	version = flag.Bool("v", false, "version")
	output  = flag.String("o", "", "usage")
)

func Parse() (*Options, string) {
	flag.Parse()
	args := flag.CommandLine.Args()
	if len(args) > 0 {
		link = args[0]
	}
	return &Options{
		Help:    *help,
		Version: *version,
		Output:  *output,
	}, link
}
