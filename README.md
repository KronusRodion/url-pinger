# URL Pinger - Go HTTP Status Checker

## 📌 Описание
Утилита для массовой проверки доступности URL-адресов из файла с сохранением результатов. Поддерживает параллельные запросы и обработку ошибок.

## 🚀 Быстрый старт

### 1. Подготовка файла с URL
Создайте в директории проекта файл `url.txt` и добавьте адреса (по одному на строку):

<pre>
https://example.com
https://google.com
https://invalid-site.test
</pre>

### 2. Запуск программы
Выполните в терминале из директории проекта:
```bash
go run main.go
```

Или с указанием своего файла:

```bash
go run main.go -file path/to/your_urls.txt
```
### 3.Результаты
Программа создаст файл result.txt с отчетом вида:

<pre>
Error while ping next urls:
URL: https://invalid-site.test, err: Get "https://invalid-site.test": dial tcp: lookup invalid-site.test: no such host

Successfully pinged: 
URL: https://example.com, status: 200 OK
URL: https://google.com, status: 200 OK
</pre>

