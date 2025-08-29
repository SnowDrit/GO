// @title           MyServer API
// @version         1.0
// @description     Учебный проект: API с информацией обо мне и о причинах изучения Go.
// @host            localhost:8080
// @BasePath        /
package main // показываем что это исполняемый файл

import (
	_ "myserver/docs" // ⚠️ заменить myserver на имя твоего модуля (из go.mod)

	httpSwagger "github.com/swaggo/http-swagger" // импортируем свагер, чтобы открывался по ссылке
)
import (
	"encoding/json" // работа с JSON
	"fmt"           // для фораматированного вывода
	"log"           // логирование ошибок и
	"net/http"      // стандартная библиотека для веб-сервера
)

// Обертка для about-me при помощи extensions даем понять swagger в каком порядке должны выводится поля в документации
type About struct {
	Last_Name   string   `json:"last_name" extensions:"x-order=1"`
	First_Name  string   `json:"first_name" extensions:"x-order=2"`
	Middle_Name string   `json:"middle_name" extensions:"x-order=3"`
	University  string   `json:"university" extensions:"x-order=4"`
	Class       int      `json:"class" extensions:"x-order=5"`
	Hobbies     []string `json:"hobbies" extensions:"x-order=6"`
}

// Обертка для why-go
type Why struct {
	FromCPlus  string "json:from_c_plus"
	FromPython string "json:from_python"
	Universal  string "json:universal"
}

// @Summary      Информация обо мне
// @Description  Возвращает данные о студенте (JSON или HTML)
// @Tags         about-me
// @Accept       json
// @Produce      json
// @Success      200  {object}  About
// @Router       /api/about-me [get]
// функция для обработки запросов /api/about-me и /about-me
func about(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")  // проверяем заголовок Accept: JSON или HTML
	if accept == "application/json" { // если cmd, то идем по этому пути
		w.Header().Set("Content-Type", "application/json; charset=utf-8") // задаем параметры вывода
		data := About{ // задаем переменную в которую записываем все поля обертки
			Last_Name:   "Артюхин",
			First_Name:  "Матфей",
			Middle_Name: "Андреевич",
			University:  "ПСТГУ",
			Class:       3,
			Hobbies:     []string{"Наука", "Игры", "Спорт", "Программирование", "Фантастика"},
		}
		json.NewEncoder(w).Encode(data) // отдаем обертку запросившему
	} else { // иной случай если это не cmd
		w.Header().Set("Content-Type", "text/html; charset=utf-8")                                                                                                                          // задаем параметры вывода
		fmt.Fprintln(w, "ФИО: <b>Артюхин Матфей Андреевич</b><br>Университет: <b>ПСТГУ</b><br>Курс: <b>3</b><br>Увлечения: <b>Наука, игры, спорт, программирование, фантастика</b> и т.д.") // отдаем текст с обаботкой html
	}
}

// @Summary      Почему Go
// @Description  Мотивация изучения Go
// @Tags         why-go
// @Accept       json
// @Produce      json
// @Success      200  {object}  Why
// @Router       /api/why-go [get]
// фунция для обработки запросов /api/why-go и /why-go
func why(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")  // проверяем заголовок Accept: JSON или HTML
	if accept == "application/json" { // если cmd, то идем по этому пути
		w.Header().Set("Content-Type", "application/json; charset=utf-8") // задаем параметры вывода
		data := Why{ // задаем переменную в которую записываем все поля обертки
			FromCPlus:  "Go даёт почти скорость C++, но без сложностей с памятью и наследованием",
			FromPython: "Go проще, чем Python в плане развёртывания и быстрее работает, с нормальной многопоточностью",
			Universal:  "Go годится и для веба, и для системного программирования, и для микросервисов",
		}
		json.NewEncoder(w).Encode(data) // отдаем обертку запросившему
	} else { // иной случай если это не cmd
		w.Header().Set("Content-Type", "text/html; charset=utf-8")                                                                                                                                                                                                              // задаем параметры вывода
		fmt.Fprintln(w, "Go даёт почти скорость C++, но без сложностей с памятью и наследованием, Go проще, чем Python в плане развёртывания и быстрее работает, с нормальной многопоточностью, Go годится и для веба, и для системного программирования, и для микросервисов") // отдаем текст с обаботкой html
	}
}

// исполнительная функция
func main() {

	// для запросов api
	http.HandleFunc("/api/about-me", about)
	http.HandleFunc("/api/why-go", why)

	// для запросов через браузеры
	http.HandleFunc("/about-me", about)
	http.HandleFunc("/why-go", why)

	// подлючаем для свагера
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// запускаем сервер на порту 8080 и ловим логи крашей, вдруг что
	log.Fatal(http.ListenAndServe(":8080", nil))
}
