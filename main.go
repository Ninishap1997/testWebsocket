package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// WebSocketHandler - обработчик вэбсокета
func WebSocketHandler(c *websocket.Conn) {
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)

	for {
		// Читаем сообщение от клиента
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			return
		}

		// Выводим сообщение в консоль сервера
		log.Printf("Получено сообщение: %s", msg)

		// Отправляем обратно то же сообщение
		if err := c.WriteMessage(msgType, msg); err != nil {
			log.Println("Ошибка отправки сообщения:", err)
			return
		}
	}
}

func main() {
	// Создаем новый экземпляр Fiber
	app := fiber.New()

	// Настраиваем обработчик вэбсокета
	app.Get("/ws", websocket.New(WebSocketHandler))

	// Слушаем порт 3000
	log.Fatal(app.Listen(":3000"))
}
