package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jellydator/ttlcache/v2"
	"net/http"
	"sync"
	"time"
	"weatherCase/models"
	"weatherCase/services"
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

type WeatherController struct {
	Service services.DefaultService
}

func (w WeatherController) GetWeather(c *fiber.Ctx) error {
	location := c.Params("location")
	var serviceResponse models.Response

	isExpire := Expired(location)
	if isExpire {
		serviceResponse = w.Service.WeatherInsert(location)
		cache.SetWithTTL(location, serviceResponse.Temperature, time.Duration(5*time.Second))

	} else {
		cached, _ := cache.Get(location)
		serviceResponse = models.Response{
			Location:    location,
			Temperature: fmt.Sprintf("%v", cached),
		}
	}
	return c.Status(http.StatusOK).JSON(serviceResponse)
}

func Expired(id string) bool {
	response, ttl, _ := cache.GetWithTTL(id)
	if response != nil || ttl != 0 {
		return false
	}
	return true
}

func testWaitGroup(w WeatherController, c *fiber.Ctx) {
	location := c.Params("location")
	var urls []string
	urls = append(urls, c.OriginalURL())
	jsonResponses := make(chan string)

	var wg sync.WaitGroup

	wg.Add(10)

	for i, url := range urls {
		tsleep := i
		go func(url string) {
			defer wg.Done()
			time.Sleep(time.Duration(tsleep) * time.Second * 5)
			fmt.Println("Sending " + url + ", at " + time.Now().String())
			res := w.Service.WeatherInsert(location)
			fmt.Println(res)
			t := time.Now()
			jsonResponses <- string("GOT id: " + url + ",  at: " + t.String())

		}(url)
	}

	go func() {
		for response := range jsonResponses {
			fmt.Println(response)
		}
	}()

	wg.Wait()
}
