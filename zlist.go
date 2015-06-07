package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zlisthq/zlistutil"
	"html/template"
	"log"
	"net/http"
)
const NUM int = 10
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

func getJSONString(site string, url string, num int) string {
	var items []zlistutil.Item
	items = zlistutil.GetItem(site, url, num)
	json_items, err := json.Marshal(&items)
	perror(err)
	return string(json_items)
}

func V2ex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	listName := vars["list_name"]
	var url string
	if listName == "hot" || listName == "latest" {
		url = zlistutil.V2EX_BASE_URL + listName + ".json"
	} else {
		http.NotFound(w, r)
		return
	}
	str := getJSONString(zlistutil.SITE_V2EX, url, NUM)
	fmt.Fprint(w, str)
	return
}
func ZhihuDaily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_ZHIHUDAILY, zlistutil.DAILY_FETCH_NOW, NUM)
	fmt.Fprint(w, str)
	return
}
func Next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_NEXT, zlistutil.NEXT, NUM)
	fmt.Fprint(w, str)
	return
}
func ProductHunt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_PRODUCTHUNT, zlistutil.PRODUCTHUNT_TODAY, NUM)
	fmt.Fprint(w, str)
	return
}
func HackerNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	listName := vars["list_name"]
	var url string
	if listName == "topstories" || listName == "newstories" || listName == "askstories" || listName == "showstories" {
		url = zlistutil.HACKER_NEWS_BASE_API_URL + "/v0/" + listName + ".json"
	} else {
		http.NotFound(w, r)
		return
	}
	// fmt.Println(url)
	str := getJSONString(zlistutil.SITE_HACKERNEWS, url, NUM)
	fmt.Fprint(w, str)
	return
}
func Jianshu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	listName := vars["list_name"]
	var url string
	if listName == "now" || listName == "weekly" || listName == "monthly" {
		url = zlistutil.JIANSHU_BASE_URL + "/trending/" + listName
	} else {
		http.NotFound(w, r)
		return
	}
	str := getJSONString(zlistutil.SITE_JIANSHU, url, NUM)
	fmt.Fprint(w, str)
	return
}
func Wanqu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_WANQU, zlistutil.WANQU, NUM)
	fmt.Fprint(w, str)
	return
}
func PingWestNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_PINGWEST, zlistutil.PINGWEST_NEWS, NUM)
	fmt.Fprint(w, str)
	return
}
func Solidot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_SOLIDOT, zlistutil.SOLIDOT, NUM)
	fmt.Fprint(w, str)
	return
}
func Github(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_GITHUB, zlistutil.GITHUB, NUM)
	fmt.Fprint(w, str)
	return
}
func DoubanMoment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_DOUBANMOMENT, zlistutil.DOUBAN_MOMENT, NUM)
	fmt.Fprint(w, str)
	return
}
func IfanrSurvey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_IFANR, zlistutil.IFANR, NUM)
	fmt.Fprint(w, str)
	return
}
func MindStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_MINDSTORE, zlistutil.MINDSTORE, NUM)
	fmt.Fprint(w, str)
	return
}
func Kickstarter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONString(zlistutil.SITE_KICKSTARTER, zlistutil.KICKSTARTER, NUM)
	fmt.Fprint(w, str)
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
	router.HandleFunc("/kickstarter/latest", Kickstarter)
	log.Fatal(http.ListenAndServe(":8080", router))
}
