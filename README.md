# Spurrow Pipeline 🚀

Фреймворк для построения гибкой конвейерной обработки данных с поддержкой параллельного выполнения.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🌟 Особенности
- 🧩 Простое создание цепочек обработки данных
- ⚡ Параллельное выполнение независимых веток
- 🧠 Поддержка динамических входных данных
- 🚦 Контроль зависимостей между шагами

## 🚀 Быстрый старт

### Требования
- Go 1.20+

### Установка
```bash
go get github.com/bdrlv/spurrow
```

### Пример использования
```go
package main

import (
	"fmt"
	"github.com/bdrlv/spurrow"
)

func main() {
	spl := spurrow.NewPipeline()

	// Регистрация цепочек
	spl.RegisterChain("input", func(input ...interface{}) (interface{}, error) {
		return "Initial Data", nil
	})

	spl.RegisterChain("processor", func(input ...interface{}) (interface{}, error) {
		return fmt.Sprintf("Processed: %v", input[0]), nil
	}, "input")

	spl.RegisterChain("output", func(input ...interface{}) (interface{}, error) {
		return fmt.Sprintf("Final: %v", input[0]), nil
	}, "processor")


	// Выполнение
	result, _ := spl.Execute("output", nil)
	fmt.Println(result)
}
```

## 📚 Документация

### Основные компоненты

#### Pipeline
Центральный элемент системы (фасад), предоставляющий:
- `RegisterChain()` - регистрация шага в цепи выполнения
- `Execute()` - запуск выполнения пайплайна

#### Processor
Обработчик данных:
```go
type Processor struct {
	DoFunc func(...interface{}) (interface{}, error)
}
```

### Расширенные сценарии

#### Передача параметров
```go
result, err := spl.Execute("chain", map[string]interface{}{
	"input_step": customData,
})
```

#### Параллельная обработка
```go
spl.RegisterChain("parallel", func(inputs ...interface{}) (interface{}, error) {
	// Параллельная обработка входов
	results := processInParallel(inputs)
	return results, nil
}, "dep1", "dep2")
```

## 📜 Лицензия
MIT License © 2025 Aleksej Badryzlov