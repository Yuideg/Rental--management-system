package handler

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/stretchr/testify/assert"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/news/news_repository"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/news/news_services"
)

func TestGetNews(t *testing.T) {
	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewNewsService(EventsRepo)
	handler := NewAdminNewsHandlerApi(EventsServ)
	router := httprouter.New()
	router.GET("/hotel/news", handler.GetNews)

	req, _ := http.NewRequest("GET", "/hotel/news", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	if rr.Code != 200 {
		t.Errorf("Response code is %v", rr.Code)
	}
	//var post=entity.News{
	//	Id:          1,
	//	Header:      "Mock newss 01",
	//	Description: "two newss",
	//	Image:       "tutu.png",
	//}
	//
	//json.Unmarshal(rr.Body.Bytes(), &post)
	//if post.Id != 1 {
	//	t.Errorf("Cannot retrieve JSON News")
	//}


}
func TestGetNewsById(t *testing.T) {
	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewNewsService(EventsRepo)
	handler := NewAdminNewsHandlerApi(EventsServ)
	router := httprouter.New()
	router.GET("/hotel/news/:id", handler.GetNewsById)

	req, _ := http.NewRequest("GET", "/hotel/news/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	if rr.Code != 200 {
		t.Errorf("Response code is %v", rr.Code)
	}
	//var post=entity.News{
	//	Id:          1,
	//	Header:      "Mock newss 01",
	//	Description: "two newss",
	//	Image:       "tutu.png",
	//}
	//
	//json.Unmarshal(rr.Body.Bytes(), &post)
	//if post.Id != 1 {
	//	t.Errorf("Cannot retrieve JSON News")
	//}
}

func TestPostNews(t *testing.T) {

	var jsonStr = []byte(`{"id": 74,
        "header": "my news",
        "description": "this is news",
        "image": "hh.jpg"}`)
	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewNewsService(EventsRepo)
	handler := NewAdminNewsHandlerApi(EventsServ)

	req, err := http.NewRequest("POST", "/hotel/news", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/hotel/news/", handler.PostNews)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusPermanentRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPermanentRedirect)
	}

}
func TestDeleteNews(t *testing.T) {
	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewNewsService(EventsRepo)
	handler := NewAdminNewsHandlerApi(EventsServ)
	router := httprouter.New()
	router.DELETE("/hotel/news/:id", handler.DeleteNews)

	req, _ := http.NewRequest("DELETE", "/hotel/news/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code it is %v",rr.Code)
	}


}