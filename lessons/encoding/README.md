## tags.go : 
+ Форматы XML, JSON, GOB
+ +как добавлять поддержку сериализации пользовательских типов данных и \
работать с динамическими спецификациями
+ Пакет easyjson
+ YAML, TOML, MessagePack и Protocol Buffers.

<b>Итак тэги: </b>
+ Теги разделяют пробелом;
+ имя тега и его значения разделяют двоеточием без пробелов;
+ значения тега разделяют запятой и заключают в кавычки;
+ имя и значения тега пишут в формате snake_case или camelCase.

```go
type User struct {
	ID        string `json:"id" format:"uuid"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at" format:"unixtime,seconds"`
}
```
+ __json__ — имя поля JSON-представления объекта (используется JSON-сериализаторами);
+ __format__ — дополнительные данные о формате поля (может использоваться другими библиотеками).