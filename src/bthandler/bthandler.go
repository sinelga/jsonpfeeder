package bthandler

import (
	//	"clean_pathinfo"
	//	"createfirstgz"
	//	"createpage"
	"domains"
	"encoding/json"
	"findfreeparagraph"
	"fmt"
	"keywords_and_phrases"
	"log/syslog"
	"net/http"
	"somekeywords"
	"somephrases"
	"strconv"
	"strings"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string, startparameters []string, callback string, quant string) {

	//	pathinfoclean := clean_pathinfo.CleanPath(golog, pathinfo)

	golog.Info(site + " ")
	var jsonBytes []byte

	if strings.HasPrefix(pathinfo, "/freeparagraph") {

		paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, "test.com", "google", startparameters)

		//	fmt.Println(paragraph)

		if bparagraph, err := json.Marshal(paragraph); err != nil {

			golog.Err(err.Error())

		} else {

			resp.Header().Add("Content-type", "application/javascript")
			

			if callback != "" {

				jsonBytes = []byte(fmt.Sprintf("%s(%s)", callback, bparagraph))

			} else {

				resp.Header().Add("Access-Control-Allow-Origin", "*")
				jsonBytes = []byte(fmt.Sprintf("%s", bparagraph))

			}

			resp.Write(jsonBytes)

		}

	}

	if strings.HasPrefix(pathinfo, "/keywords?") {

		golog.Info("pathinfo  " + pathinfo + " " + quant)

		keywords, phrases := keywords_and_phrases.GetAll(golog, locale, themes, startparameters)

		quantint, _ := strconv.Atoi(quant)

		somekeywordsres := somekeywords.GetSome(golog, keywords, quantint)
		somephrasesres := somephrases.GetSome(golog, phrases, quantint)

		var keyword_phrasearr []domains.Keyword_phrase

		if len(somekeywordsres) <= len(somephrasesres) {

			for i, keyword := range somekeywordsres {

				keyword_phrase := domains.Keyword_phrase{keyword, somephrasesres[i]}
				keyword_phrasearr = append(keyword_phrasearr, keyword_phrase)

			}

		} else {

			for i, phrase := range somephrasesres {

				keyword_phrase := domains.Keyword_phrase{somekeywordsres[i], phrase}
				keyword_phrasearr = append(keyword_phrasearr, keyword_phrase)

			}

		}

		if bkeyword_phrasearr, err := json.Marshal(keyword_phrasearr); err != nil {

			golog.Err(err.Error())

		} else {

			resp.Header().Add("Content-type", "application/javascript")

			if callback != "" {

				jsonBytes = []byte(fmt.Sprintf("%s(%s)", callback, bkeyword_phrasearr))

			} else {

				resp.Header().Add("Access-Control-Allow-Origin", "*")
				jsonBytes = []byte(fmt.Sprintf("%s", bkeyword_phrasearr))

			}
			resp.Write(jsonBytes)
		}

	}

}
