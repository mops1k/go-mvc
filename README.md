## go-mvc
Библиотека для создания MVC веб-приложений на Golang. Эволюция SimpleMVC приложения.
_Библиотека разрабатывается для личных нужд_

#### Быстрый старт
Создайте пакет с поддержкой модулей go.

- В `main.go` напишите следующее:
```go
package main

import (
    "github.com/mops1k/go-mvc"

    _ "<Your package>/config"
)

func main() {
    mvc.Run()
}
```

- В папке `config` создайте файл `app_config.yaml` следующего содержания:
```yaml
server:
    host: localhost
    port: 8082
    timeout:
        read:  10
        write: 15
        idle:  30

database:
  default:
    enabled: false
    type:    sqlite3 # sqlite3, mysql, mssql, postgres
    url:     database.db

translation:
    default: ru
    fallback: ru
``` 
_В данной папке все `yaml` и `yml` файлы собираются конфигурацией автоматически_

- Создайте папку `app/controller` и в ней файл `index.go` следующего содержания:
```go
package controller

import (
    "github.com/mops1k/go-mvc/http"
)

type IndexController struct {
    http.BaseController
    PathInfo interface{} `path:"/" name:"index" methods:"GET"`
}

func (i *IndexController) Action(c *http.Context) (string, error) {
    return i.RenderString("Hello, {{ name }}!", map[string]interface{}{"name": "World"}), nil
}
```

- В папке config создайте файл `controllers.go` следующего содержания:
```go
package config

import (
    "github.com/mops1k/go-mvc/http"

    "<Your package>/app/controller"
)

func init() {
    // Add your controllers here
    http.Controllers.Add(&controller.IndexController{})
}
```

- Выполните:
```bash
go mod vendor
go build .
```

#### [WIP]Документация
_Здесь пока пусто_
