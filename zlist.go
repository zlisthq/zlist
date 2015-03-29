package main

import (
	"encoding/json"
	"fmt"
	"github.com/WhiteWorld/zlistutil"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func perror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"css":    staticDirectory + "/css/",
		"js":     staticDirectory + "/js/",
		"images": staticDirectory + "/images/"}

	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix, http.FileServer(http.Dir(pathValue))))
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	perror(err)
	err = t.Execute(w, nil)
	perror(err)

}
func V2ex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	vars := mux.Vars(r)
	listName := vars["list_name"]
	var url string
	if listName == "hot" || listName == "latest" {
		url = zlistutil.V2EX_BASE_URL + listName + ".json"
	} else {
		http.NotFound(w, r)
		return
	}
	items = zlistutil.FetchV2ex(url, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func ZhihuDaily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchZhihuDaily(zlistutil.DAILY_FETCH_NOW, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func Next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchNEXT(zlistutil.NEXT, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func ProductHunt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchProductHunt(zlistutil.PRODUCTHUNT_TODAY, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func HackerNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	vars := mux.Vars(r)
	listName := vars["list_name"]
	var url string
	if listName == "topstories" || listName == "newstories" {
		url = zlistutil.HACKER_NEWS_BASE_API_URL + "/v0/" + listName + ".json"
	} else {
		http.NotFound(w, r)
		return
	}
	// fmt.Println(url)
	items = zlistutil.FetchHackerNews(url, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func Jianshu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	vars := mux.Vars(r)
	listName := vars["list_name"]
	var url string
	if listName == "now" || listName == "weekly" || listName == "monthly" {
		url = zlistutil.JIANSHU_BASE_URL + "/trending/" + listName
	} else {
		http.NotFound(w, r)
		return
	}
	items = zlistutil.FetchJianshu(url, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func Wanqu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchWanqu(zlistutil.WANQU, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func PingWestNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchPingWestNews(zlistutil.PINGWEST_NEWS, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func Solidot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchSolidot(zlistutil.SOLIDOT, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func Github(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchGitHub(zlistutil.GITHUB, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func DoubanMoment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchDoubanMoment(zlistutil.DOUBAN_MOMENT, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func IfanrSurvey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchIfanr(zlistutil.IFANR, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func MindStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []zlistutil.Item
	items = zlistutil.FetchMindStore(zlistutil.MINDSTORE, 10)
	json_items, err := json.Marshal(&items)
	perror(err)
	fmt.Fprint(w, string(json_items))
	return
}
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/", Index)
	router.HandleFunc("/producthunt/top", ProductHunt)
	router.HandleFunc("/jianshu/{list_name}", Jianshu)
	router.HandleFunc("/36kr/next", Next)
	router.HandleFunc("/hackernews/{list_name}", HackerNews)
	router.HandleFunc("/v2ex/{list_name}", V2ex)
	router.HandleFunc("/zhihudaily/latest", ZhihuDaily)
	router.HandleFunc("/wanqu/top", Wanqu)
	router.HandleFunc("/pingwest/news", PingWestNews)
	router.HandleFunc("/solidot/top", Solidot)
	router.HandleFunc("/github/top", Github)
	router.HandleFunc("/douban/moment", DoubanMoment)
	router.HandleFunc("/ifanr/survey", IfanrSurvey)
	router.HandleFunc("/mindstore/top", MindStore)
	log.Fatal(http.ListenAndServe(":8080", router))
}
