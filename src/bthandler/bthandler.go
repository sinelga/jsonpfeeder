package bthandler

import (
	//	"clean_pathinfo"
	//	"createfirstgz"
	//	"createpage"
	"log/syslog"
	"net/http"
	"encoding/json"
	"findfreeparagraph"
	"fmt"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string, startparameters []string, callback string) {

	//	pathinfoclean := clean_pathinfo.CleanPath(golog, pathinfo)

	golog.Info(site + " ")

	paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, "test.com", "google", startparameters)

	//	fmt.Println(paragraph)

	if bparagraph, err := json.Marshal(paragraph); err != nil {

		golog.Err(err.Error())

	} else {

		resp.Header().Add("Content-type", "application/javascript")

		//		fmt.Sprintf(resp, "%s(%s)", "callback", bparagraph)
		jsonBytes := []byte(fmt.Sprintf("%s(%s)", callback, bparagraph))
		resp.Write(jsonBytes)

	}

	//	if blocksite {
	//
	//		golog.Info("BTrequestHandler:will block-> " + site+pathinfo)
	//
	//	}

	//	var bytepage []byte
	//	if strings.HasSuffix(pathinfoclean, ".html")  {
	//
	//		bytepage = createpage.CreateHtmlPage(golog, locale, themes,site, bot, startparameters,blocksite,variant)
	//
	//		resp.Write(bytepage)
	//
	//	} else if (strings.HasSuffix(pathinfoclean, "sitemap.xml") || strings.HasSuffix(pathinfoclean, "index.xml"))  {
	//
	//
	//		bytepage =sitemaphandler.Create(golog, locale, themes,site,startparameters)
	//		resp.Header().Add("Content-type","text/xml")
	//		resp.Write(bytepage)
	//
	//	} else if strings.HasSuffix(pathinfoclean, "robots.txt") {
	//
	//
	//		bytepage =robots_txt.Create(golog, locale, themes,site)
	//		resp.Header().Add("Content-type","text/plain")
	//		resp.Write(bytepage)
	//
	//
	//	} else {
	//
	//		resp.WriteHeader(404)
	//
	//	}

	//	if strings.HasSuffix(pathinfoclean, ".html") {
	//
	//		go createfirstgz.Creategzhtml(golog, locale, themes, site, pathinfoclean, bytepage)
	//
	//	}

}
