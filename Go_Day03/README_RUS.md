# День 03 — Go Boot Camp

## Вкусные открытия

💡 [Нажмите здесь](https://new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624), чтобы **оставить отзыв о проекте**. Это анонимно и поможет нашей команде улучшить ваш образовательный опыт. Мы рекомендуем заполнять опрос сразу после завершения проекта.

## Содержание

1. [Глава I](#глава-i) \
    1.1. [Общие правила](#общие-правила)
2. [Глава II](#глава-ii) \
    2.1. [Правила дня](#правила-дня)
3. [Глава III](#глава-iii) \
    3.1. [Введение](#введение)
4. [Глава IV](#глава-iv) \
    4.1. [Упражнение 00: Загрузка данных](#упр00)
5. [Глава V](#глава-v) \
    5.1. [Упражнение 01: Простой интерфейс](#упр01)
6. [Глава VI](#глава-vi) \
    6.1. [Упражнение 02: Полноценный API](#упр02)
7. [Глава VII](#глава-vii) \
    7.1. [Упражнение 03: Ближайшие рестораны](#упр03)
8. [Глава VIII](#глава-viii) \
    8.1. [Упражнение 04: JWT](#упр04)

<h2 id="глава-i">Глава I</h2>
<h2 id="общие-правила">Общие правила</h2>

- Ваши программы не должны завершаться неожиданно (выдавать ошибку при корректном вводе данных). Если это произойдёт, проект будет считаться нефункциональным и получит 0 баллов при оценке.
- Мы рекомендуем создавать тестовые программы для вашего проекта. Хотя они не обязательно должны быть отправлены на проверку и не будут оцениваться, они позволят вам легко протестировать свою работу и работу ваших коллег. Эти тесты особенно пригодятся на защите. Вы можете использовать свои тесты и/или тесты коллег, которых вы оцениваете, во время защиты.
- Отправьте свою работу в назначенный git-репозиторий. Будет оцениваться только работа, находящаяся в репозитории.
- Если ваш код использует внешние зависимости, он должен управляться с помощью [Go Modules](https://go.dev/blog/using-go-modules).

<h2 id="глава-ii">Глава II</h2>
<h2 id="правила-дня">Правила дня</h2>

- Вы должны отправить только файлы с расширением `*.go`, а также `go.mod` и `go.sum` (если используются внешние зависимости).
- Ваш код для этой задачи должен собираться с помощью команды `go build`.
- Все входные данные (например, `page`, `lat`, `long`) должны быть тщательно проверены и не должны приводить к HTTP 500. Допустимы только HTTP 400/401 с содержательным сообщением об ошибке, как указано в Упражнении 02.

<h2 id="глава-iii">Глава III</h2>
<h2 id="введение">Введение</h2>

Люди любят приложения с рекомендациями. Они помогают меньше задумываться о том, что купить, куда пойти и что поесть.

К тому же у большинства людей есть телефон с функцией геолокации. Сколько раз вы пытались найти рестораны в вашем районе для ужина?

Давайте подумаем, как работают такие сервисы, и создадим свой, совсем простой, ладно?

<h2 id="глава-iv">Глава IV</h2>
<h3 id="упр00">Упражнение 00: Загрузка данных</h3>

На рынке существует множество разных баз данных. Но поскольку мы пытаемся предоставить возможность поиска, давайте использовать [Elasticsearch](https://www.elastic.co/downloads/elasticsearch).

Elasticsearch — это полнотекстовый поисковый движок, построенный на основе [Lucene](https://en.wikipedia.org/wiki/Apache_Lucene). Он предоставляет HTTP API, которое мы будем использовать в этой задаче.

Предоставленный набор данных о ресторанах (взятый с портала открытых данных) содержит более 13 тысяч ресторанов в Москве, Россия (вы можете создать аналогичный набор данных для любой другой локации). Каждая запись содержит:

- ID
- Название
- Адрес
- Телефон
- Долгота
- Широта

Перед загрузкой всех записей в базу данных давайте создадим индекс и сопоставление (явно указав типы данных). Без этого Elasticsearch попытается угадать типы полей на основе предоставленных документов, что иногда приводит к ошибкам, особенно для геоточек.

Вот несколько ссылок для начала работы:
- https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html
- https://www.elastic.co/guide/en/elasticsearch/reference/8.4/geo-point.html

Запустите базу данных командой `~$ /path/to/elasticsearch/dir/bin/elasticsearch` и начнем экспериментировать.

Для простоты используем "places" как название индекса и "place" как название записи. Вы можете создать индекс с помощью cURL следующим образом:

```bash
~$ curl -XPUT "http://localhost:9200/places"
```

Но в этой задаче вы должны использовать Elasticsearch-библиотеки для Go для выполнения этой задачи. Следующее, что нужно сделать, — это предоставить сопоставление типов для наших данных. С помощью cURL это выглядело бы так:

```
~$ curl -XPUT http://localhost:9200/places/place/_mapping?include_type_name=true -H "Content-Type: application/json" -d @"schema.json"
```

где `schema.json` выглядит следующим образом:

```
{
  "properties": {
    "name": {
        "type":  "text"
    },
    "address": {
        "type":  "text"
    },
    "phone": {
        "type":  "text"
    },
    "location": {
      "type": "geo_point"
    }
  }
}
```

Again, assuming the cURL commands are just a reference for self-testing, this action should be performed by the Go program you write.

Now you have a dataset to upload. You should use the [Bulk API](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/docs-bulk.html) to do this. All existing Elasticsearch bindings provide wrappers for this, for example [here's a good example](https://github.com/elastic/go-elasticsearch/blob/master/_examples/bulk/indexer.go) for an official client<!--- (note that you need to use client v7 for ES version 7.9, not v8)-->. There are also a number of third-party clients, choose which you prefer.

Чтобы проверить себя, вы можете использовать cURL. Так:

```
~$ curl -s -XGET "http://localhost:9200/places"
```

должен дать вам что-то вроде этого:

```
{
  "places": {
    "aliases": {},
    "mappings": {
      "properties": {
        "address": {
          "type": "text"
        },
        "id": {
          "type": "long"
        },
        "location": {
          "type": "geo_point"
        },
        "name": {
          "type": "text"
        },
        "phone": {
          "type": "text"
        }
      }
    },
    "settings": {
      "index": {
        "creation_date": "1601810777906",
        "number_of_shards": "1",
        "number_of_replicas": "1",
        "uuid": "4JKa9fgISd6-N130rpNYtQ",
        "version": {
          "created": "7090299"
        },
        "provided_name": "places"
      }
    }
  }
}
```

и запрос записи по ее идентификатору будет выглядеть так:

```
~$ curl -s -XGET "http://localhost:9200/places/_doc/1"
```

```
{
  "_index": "places",
  "_type": "place",
  "_id": "1",
  "_version": 1,
  "_seq_no": 0,
  "_primary_term": 1,
  "found": true,
  "_source": {
    "id": 1,
    "name": "SMETANA",
    "address": "gorod Moskva, ulitsa Egora Abakumova, dom 9",
    "phone": "(499) 183-14-10",
    "location": {
      "lat": 55.879001531303366,
      "lon": 37.71456500043604
    }
  }
}
```

Обратите внимание, что запись с ID=1 может отличаться от той, что в наборе данных, если вы решили использовать горутины для ускорения процесса (хотя это не обязательное требование для этой задачи).

<h2 id="глава-v" >Глава V</h2>
<h3 id="упр01">Упражнение 01: Самый простой интерфейс</h3>

Теперь давайте создадим HTML интерфейс для нашей базы данных. Ничего сложного, нужно всего лишь отобразить страницу со списком имен, адресов и телефонов, чтобы пользователи могли увидеть это в браузере.

Вы должны абстрагировать вашу базу данных за интерфейсом. Для того, чтобы просто вернуть список записей и иметь возможность [пагинировать](https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html) через них, этого интерфейса будет достаточно:

```
type Store interface {
    // возвращает список элементов, общее количество записей и (или) ошибку в случае её наличия
    GetPlaces(limit int, offset int) ([]types.Place, int, error)
}
```

В пакете `main` не должно быть импортов, связанных с Elasticsearch, так как все, что касается работы с базой данных, должно находиться в пакете `db` вашего проекта, и вы должны использовать только этот интерфейс для взаимодействия с ним.

Ваше HTTP-приложение должно работать на порту 8888, отвечать списком ресторанов и обеспечивать простую пагинацию. Таким образом, если вы запрашиваете "http://127.0.0.1:8888/?page=2" (обратите внимание на параметр GET `page`), вы должны получить страницу, похожую на эту:

```
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Places</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
<h5>Total: 13649</h5>
<ul>
    <li>
        <div>Sushi Wok</div>
        <div>gorod Moskva, prospekt Andropova, dom 30</div>
        <div>(499) 754-44-44</div>
    </li>
    <li>
        <div>Ryba i mjaso na ugljah</div>
        <div>gorod Moskva, prospekt Andropova, dom 35A</div>
        <div>(499) 612-82-69</div>
    </li>
    <li>
        <div>Hleb nasuschnyj</div>
        <div>gorod Moskva, ulitsa Arbat, dom 6/2</div>
        <div>(495) 984-91-82</div>
    </li>
    <li>
        <div>TAJJ MAHAL</div>
        <div>gorod Moskva, ulitsa Arbat, dom 6/2</div>
        <div>(495) 107-91-06</div>
    </li>
    <li>
        <div>Balalaechnaja</div>
        <div>gorod Moskva, ulitsa Arbat, dom 23, stroenie 1</div>
        <div>(905) 752-88-62</div>
    </li>
    <li>
        <div>IL Pizzaiolo</div>
        <div>gorod Moskva, ulitsa Arbat, dom 31</div>
        <div>(495) 933-28-34</div>
    </li>
    <li>
        <div>Bufet pri Astrahanskih banjah</div>
        <div>gorod Moskva, Astrahanskij pereulok, dom 5/9</div>
        <div>(495) 344-11-68</div>
    </li>
    <li>
        <div>MU-MU</div>
        <div>gorod Moskva, Baumanskaja ulitsa, dom 35/1</div>
        <div>(499) 261-33-58</div>
    </li>
    <li>
        <div>Bek tu Blek</div>
        <div>gorod Moskva, Tatarskaja ulitsa, dom 14</div>
        <div>(495) 916-90-55</div>
    </li>
    <li>
        <div>Glav Pirog</div>
        <div>gorod Moskva, Begovaja ulitsa, dom 17, korpus 1</div>
        <div>(926) 554-54-08</div>
    </li>
</ul>
<a href="/?page=1">Previous</a>
<a href="/?page=3">Next</a>
<a href="/?page=1364">Last</a>
</body>
</html>
```

Ссылка "Previous" должна исчезать на первой странице, а ссылка "Next" должна исчезать на последней странице.

ВАЖНОЕ ЗАМЕЧАНИЕ: Вы можете заметить, что по умолчанию Elasticsearch не позволяет работать с пагинацией для более чем 10000 записей. Существует два способа обойти это ограничение — либо использовать Scroll API (см. ссылку выше на пагинацию), либо просто увеличить лимит в настройках индекса для этой задачи. Последний способ приемлем для этой задачи, но не рекомендуется для использования в продакшн. Запрос, который поможет вам установить это, приведен ниже:

```
~$ curl -XPUT -H "Content-Type: application/json" "http://localhost:9200/places/_settings" -d '
{
  "index" : {
    "max_result_window" : 20000
  }
}'
```

Также, если параметр 'page' задан с неправильным значением (вне диапазона [0..last_page] или не является числом), ваша страница должна возвращать ошибку HTTP 400 и обычный текст с описанием ошибки:

```
Неверное значение 'page': 'foo'.
```

<h2 id="глава-vi" >Глава VI</h2>
<h3 id="упр02">Упражнение 02: Правильный API</h3>

В современном мире большинство приложений предпочитают API вместо обычного HTML. В этом упражнении все, что вам нужно сделать — это реализовать другой обработчик, который будет отвечать с заголовком `Content-Type: application/json` и JSON-версией того же, что в Ex01 (пример для http://127.0.0.1:8888/api/places?page=3):

```
{
  "name": "Places",
  "total": 13649,
  "places": [
    {
      "id": 65,
      "name": "AZERBAJDZhAN",
      "address": "gorod Moskva, ulitsa Dem'jana Bednogo, dom 4",
      "phone": "(495) 946-34-30",
      "location": {
        "lat": 55.769830485601204,
        "lon": 37.486914061171504
      }
    },
    {
      "id": 69,
      "name": "Vojazh",
      "address": "gorod Moskva, Beskudnikovskij bul'var, dom 57, korpus 1",
      "phone": "(499) 485-20-00",
      "location": {
        "lat": 55.872553383512496,
        "lon": 37.538326789741
      }
    },
    {
      "id": 70,
      "name": "GBOU Shkola № 1411 (267)",
      "address": "gorod Moskva, ulitsa Bestuzhevyh, dom 23",
      "phone": "(499) 404-15-09",
      "location": {
        "lat": 55.87213179130298,
        "lon": 37.609625999999984
      }
    },
    {
      "id": 71,
      "name": "Zhigulevskoe",
      "address": "gorod Moskva, Bibirevskaja ulitsa, dom 7, korpus 1",
      "phone": "(964) 565-61-28",
      "location": {
        "lat": 55.88024342230735,
        "lon": 37.59308635976602
      }
    },
    {
      "id": 75,
      "name": "Hinkal'naja",
      "address": "gorod Moskva, ulitsa Marshala Birjuzova, dom 16",
      "phone": "(499) 728-47-01",
      "location": {
        "lat": 55.79476126986192,
        "lon": 37.491709793339744
      }
    },
    {
      "id": 76,
      "name": "ShAURMA ZhI",
      "address": "gorod Moskva, ulitsa Marshala Birjuzova, dom 19",
      "phone": "(903) 018-74-64",
      "location": {
        "lat": 55.794378830665885,
        "lon": 37.49112002224252
      }
    },
    {
      "id": 80,
      "name": "Bufet Shkola № 554",
      "address": "gorod Moskva, Bolotnikovskaja ulitsa, dom 47, korpus 1",
      "phone": "(929) 623-03-21",
      "location": {
        "lat": 55.66186417434049,
        "lon": 37.58323602169326
      }
    },
    {
      "id": 83,
      "name": "Kafe",
      "address": "gorod Moskva, 1-j Botkinskij proezd, dom 2/6",
      "phone": "(495) 945-22-34",
      "location": {
        "lat": 55.781141341601696,
        "lon": 37.55643137063551
      }
    },
    {
      "id": 84,
      "name": "STARYJ BATUM'",
      "address": "gorod Moskva, ulitsa Akademika Bochvara, dom 7, korpus 1",
      "phone": "(495) 942-44-85",
      "location": {
        "lat": 55.8060307318284,
        "lon": 37.461669109923506
      }
    },
    {
      "id": 89,
      "name": "Cheburechnaja SSSR",
      "address": "gorod Moskva, Bol'shaja Bronnaja ulitsa, dom 27/4",
      "phone": "(495) 694-54-76",
      "location": {
        "lat": 55.764134959774346,
        "lon": 37.60256453956346
      }
    }
  ],
  "prev_page": 2,
  "next_page": 4,
  "last_page": 1364
}
```

Также, если параметр 'page' задан с неправильным значением (вне диапазона [0..last_page] или не является числом), ваш API должен ответить с соответствующей ошибкой HTTP 400 и аналогичным JSON:

```
{
    "error": "Invalid 'page' value: 'foo'"
}
```

<h2 id="глава-vii" >Глава VII</h2>
<h3 id="упр03">Упражнение 03: Ближайшие рестораны</h3>

Теперь давайте реализуем нашу основную функцию — поиск *трех* ближайших ресторанов! Для этого вам нужно настроить сортировку для вашего запроса:

```
"sort": [
    {
      "_geo_distance": {
        "location": {
          "lat": 55.674,
          "lon": 37.666
        },
        "order": "asc",
        "unit": "km",
        "mode": "min",
        "distance_type": "arc",
        "ignore_unmapped": true
      }
    }
]
```

где "lat" и "lon" — это ваши текущие координаты. Например, для URL http://127.0.0.1:8888/api/recommend?lat=55.674&lon=37.666 ваше приложение должно вернуть JSON, подобный этому:

```
{
  "name": "Recommendation",
  "places": [
    {
      "id": 30,
      "name": "Ryba i mjaso na ugljah",
      "address": "gorod Moskva, prospekt Andropova, dom 35A",
      "phone": "(499) 612-82-69",
      "location": {
        "lat": 55.67396575768212,
        "lon": 37.66626689310591
      }
    },
    {
      "id": 3348,
      "name": "Pizzamento",
      "address": "gorod Moskva, prospekt Andropova, dom 37",
      "phone": "(499) 612-33-88",
      "location": {
        "lat": 55.673075576456,
        "lon": 37.664533747576
      }
    },
    {
      "id": 3347,
      "name": "KOFEJNJa «KAPUChINOFF»",
      "address": "gorod Moskva, prospekt Andropova, dom 37",
      "phone": "(499) 612-33-88",
      "location": {
        "lat": 55.672865251005106,
        "lon": 37.6645689561318
      }
    }
  ]
}
```

<h2 id="глава-viii" >Глава VIII</h2>
<h3 id="упр04">Упражнение 04: JWT</h3>

Последнее (но не менее важное), что нам нужно сделать, — это предоставить простую форму аутентификации. На данный момент один из самых популярных способов реализации этого для API — использовать [JWT](https://jwt.io/introduction/). К счастью, Go имеет довольно хорошие инструменты для работы с этим.

Во-первых, вам нужно реализовать конечную точку API http://127.0.0.1:8888/api/get_token, основная цель которой — генерировать токен и возвращать его вот так (это пример, ваш токен, вероятно, будет другим):

```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjAxOTc1ODI5LCJuYW1lIjoiTmlrb2xheSJ9.FqsRe0t9YhvEC3hK1pCWumGvrJgz9k9WvhJgO8HsIa8"
}
```

Не забудьте установить заголовок 'Content-Type: application/json'.

Во-вторых, вам нужно защитить конечную точку `/api/recommend` с помощью JWT-мидлваре, которая проверяет действительность этого токена.

Так что по умолчанию, когда этот API запрашивается из браузера, он должен вернуть ошибку HTTP 401, но будет работать, если клиент отправит заголовок `Authorization: Bearer <token>` (вы можете проверить это с помощью cURL или Postman).

Это самый простой способ предоставить аутентификацию, подробности можно будет рассмотреть позже.

