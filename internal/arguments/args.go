package arguments

import (
	"flag"
	"time"
)

var link string

/*
-b (--background) - работать в фоновом режиме
-d (--debug) - включить режим отладки
-v (--verbose) - выводить максимум информации о работе утилиты
-q (--quiet) - выводить минимум информации о работе
-i файл (--input-file) - прочитать URL из файла
--force-html - читать файл указанный в предыдущем параметре как html
-O файл (--output-document) - файл в который будут сохранены полученные данные
-с (--continue) - продолжить ранее прерванную загрузку
-S (--server-response) - вывести ответ сервера
--spider - проверить работоспособность URL
--limit-rate - ограничить скорость загрузки
-w (--wait) - интервал между запросами
-Q (--quota) - максимальный размер загрузки
-4 (--inet4only) - использовать протокол ipv4
-6 (--inet6only) - использовать протокол ipv6
-U (--user-agent)- строка USER AGENT отправляемая серверу
-r (--recursive)- рекурсивная работа утилиты
-l (--level) - глубина при рекурсивном сканировании
-k (--convert-links) - конвертировать ссылки в локальные при загрузке страниц
-m (--mirror) - скачать сайт на локальную машину
-p (--page-requisites) - во время загрузки сайта скачивать все необходимые ресурсы
*/

type Options struct {
	Outdir    string    // -P (--directory-prefix) - каталог, в который будут загружаться файлы
	Timeout   time.Time //-T время (--timeout) - таймаут подключения к серверу
	Tries     int       // -t (--tries) - количество попыток подключения к серверу
	Outlogs   string    // -o файл (--out-file) - указать лог файл
	Help      bool      // -h (--help) - вывести справку
	Version   bool      // -V (--version) - вывести версию программы
	Inputfile string    // -i файл (--input-file) - прочитать URL из файла
	Output    string
}

var (
	help    = flag.Bool("h", false, "help")
	version = flag.Bool("v", false, "version")
	output  = flag.String("o", "", "usage")
)

func Parse() (*Options, []string) {
	flag.Parse()
	return &Options{
		Help:    *help,
		Version: *version,
		Output:  *output,
	}, flag.CommandLine.Args()
}
