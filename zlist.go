package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/zlisthq/zlistutil"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	Num         int    = 10
	CacheExpire string = "1800"
)

var (
	conn, err = redis.Dial("tcp", os.Getenv("REDIS_PORT_6379_TCP_ADDR")+":"+os.Getenv("REDIS_PORT_6379_TCP_PORT"))
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

func getJSONStringCached(site string, url string, num int) string {
	if conn == nil {
		return getJSONString(site, url, num)
	}
	jsonString, err := redis.String(conn.Do("GET", url))
	if err != nil {
		jsonString = getJSONString(site, url, num)
		conn.Do("SETEX", url, CacheExpire, jsonString)
		log.Println("Cache: set " + url)
	}
	return jsonString
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
	str := getJSONStringCached(zlistutil.SITE_V2EX, url, Num)
	fmt.Fprint(w, str)
	return
}
func ZhihuDaily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_ZHIHUDAILY, zlistutil.DAILY_FETCH_NOW, Num)
	fmt.Fprint(w, str)
	return
}
func Next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_NEXT, zlistutil.NEXT, Num)
	fmt.Fprint(w, str)
	return
}
func ProductHunt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_PRODUCTHUNT, zlistutil.PRODUCTHUNT_TODAY, Num)
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
	str := getJSONStringCached(zlistutil.SITE_HACKERNEWS, url, Num)
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
	str := getJSONStringCached(zlistutil.SITE_JIANSHU, url, Num)
	fmt.Fprint(w, str)
	return
}
func Wanqu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_WANQU, zlistutil.WANQU, Num)
	fmt.Fprint(w, str)
	return
}
func PingWestNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_PINGWEST, zlistutil.PINGWEST_NEWS, Num)
	fmt.Fprint(w, str)
	return
}
func Solidot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_SOLIDOT, zlistutil.SOLIDOT, Num)
	fmt.Fprint(w, str)
	return
}
func Github(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_GITHUB, zlistutil.GITHUB, Num)
	fmt.Fprint(w, str)
	return
}
func DoubanMoment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_DOUBANMOMENT, zlistutil.DOUBAN_MOMENT, Num)
	fmt.Fprint(w, str)
	return
}
func IfanrSurvey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_IFANR, zlistutil.IFANR, Num)
	fmt.Fprint(w, str)
	return
}
func MindStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_MINDSTORE, zlistutil.MINDSTORE, Num)
	fmt.Fprint(w, str)
	return
}
func Kickstarter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_KICKSTARTER, zlistutil.KICKSTARTER, Num)
	fmt.Fprint(w, str)
	return
}
func Toutiao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := getJSONStringCached(zlistutil.SITE_TOUTIAO, zlistutil.TOUTIAO, Num)
	fmt.Fprint(w, str)
	return
}
func Refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	refreshCache(true)
	str := "{'code':'OK'}"
	fmt.Fprint(w, str)
	return
}
func refreshCache(flag bool) {
	urlSite := map[string]string{
		zlistutil.V2EX_BASE_URL + "hot.json":                        zlistutil.SITE_V2EX,
		zlistutil.V2EX_BASE_URL + "latest.json":                     zlistutil.SITE_V2EX,
		zlistutil.DAILY_FETCH_NOW:                                   zlistutil.SITE_ZHIHUDAILY,
		zlistutil.NEXT:                                              zlistutil.SITE_NEXT,
		zlistutil.PRODUCTHUNT_TODAY:                                 zlistutil.SITE_PRODUCTHUNT,
		zlistutil.HACKER_NEWS_BASE_API_URL + "/v0/topstories.json":  zlistutil.SITE_HACKERNEWS,
		zlistutil.HACKER_NEWS_BASE_API_URL + "/v0/newstories.json":  zlistutil.SITE_HACKERNEWS,
		zlistutil.HACKER_NEWS_BASE_API_URL + "/v0/askstories.json":  zlistutil.SITE_HACKERNEWS,
		zlistutil.HACKER_NEWS_BASE_API_URL + "/v0/showstories.json": zlistutil.SITE_HACKERNEWS,
		zlistutil.JIANSHU_BASE_URL + "/trending/now":                zlistutil.SITE_JIANSHU,
		zlistutil.JIANSHU_BASE_URL + "/trending/weekly":             zlistutil.SITE_JIANSHU,
		zlistutil.JIANSHU_BASE_URL + "/trending/monthly":            zlistutil.SITE_JIANSHU,
		zlistutil.WANQU:                                             zlistutil.SITE_WANQU,
		zlistutil.PINGWEST_NEWS:                                     zlistutil.SITE_PINGWEST,
		zlistutil.SOLIDOT:                                           zlistutil.SITE_SOLIDOT,
		zlistutil.GITHUB:                                            zlistutil.SITE_GITHUB,
		zlistutil.DOUBAN_MOMENT:                                     zlistutil.SITE_DOUBANMOMENT,
		zlistutil.IFANR:                                             zlistutil.SITE_IFANR,
		zlistutil.MINDSTORE:                                         zlistutil.SITE_MINDSTORE,
		zlistutil.KICKSTARTER:                                       zlistutil.SITE_KICKSTARTER,
		zlistutil.TOUTIAO:                                           zlistutil.SITE_TOUTIAO,
	}
	log.Println("start refresh...")
	log.Println(time.Now())
	log.Printf("Clean exist cache ? %t", flag)
	for url, site := range urlSite {
		if flag == true && conn != nil {
			conn.Do("DEL", url)
		}
		getJSONStringCached(site, url, Num)
	}
	log.Println("stop refresh...")
}
func jobRefreshCache() {
	refreshCache(false)
}
func main() {
	c := cron.New()
	c.AddFunc("0 */5 * * * ?", jobRefreshCache)
	c.Start()
	log.Println("REDIS HOST:" + os.Getenv("REDIS_PORT_6379_TCP_ADDR"))
	log.Println("REDIS PORT:" + os.Getenv("REDIS_PORT_6379_TCP_PORT"))
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/", Index)
	router.HandleFunc("/refresh", Refresh)
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
	router.HandleFunc("/toutiao/top", Toutiao)
	log.Println(http.ListenAndServe(":8080", router))
}
