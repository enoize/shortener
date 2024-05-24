### Проверка в Windows

Скачать репозиторий

```
git clone https://github.com/enoize/shortener.git shortener
```

Запустить ssdb-server
```
cd shortener\ssdb
ssdb-server-1.9.4.exe ssdb.conf 
```

Перейти в папку проекта и запустить приложение
```
cd shortener
go run .
```

Открыть в браузере
```
http://127.0.0.1:8080/a?url=http%3A%2F%2Fgoogle.com%2F%3Fq%3Dgolang
```

Открыть в браузере
```
http://127.0.0.1:8080/s/cDQIQzgt
```
