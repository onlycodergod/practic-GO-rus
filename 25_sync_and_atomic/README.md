**Задание:**
- Перепишите класс Cache из предыдущего задания, используя sync.RWMutex.
Замените вашу функцию main() и добавьте константу из кода, который предложен далее:

```go
const (
	k1   = "key1"
	step = 7 
)
```
- Модифицируйте код, использующий класс Cache, из прошлого задания так, чтобы все горутины завершали свое выполнение после отработки (используйте для этого sync.WaitGroup).