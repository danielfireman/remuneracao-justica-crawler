package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var (
	pairsFlag = flag.String("p", "", "")
	out       = flag.String("out", "", "")
)

const (
	basePath        = "http://www.cnj.jus.br"
	remuneracaoPath = "http://www.cnj.jus.br/transparencia/remuneracao-dos-magistrados/remuneracao-"
)

func main() {
	flag.Parse()

	pairs := strings.Split(*pairsFlag, ",")
	var wg sync.WaitGroup
	for _, p := range pairs {
		remuneracaoURL := fmt.Sprintf("%s%s", remuneracaoPath, p)
		doc, err := goquery.NewDocument(remuneracaoURL)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("td").Each(func(index int, item *goquery.Selection) {
			linkTag := item.Find("a")
			link, _ := linkTag.Attr("href")
			if strings.HasSuffix(link, "xls") || strings.HasSuffix(link, "xlsx") {
				wg.Add(1)
				go func() {
					defer wg.Done()
					dLink := fmt.Sprintf("%s%s", basePath, link)
					resp, err := http.Get(dLink)
					if err != nil {
						log.Fatal(err)
					}
					defer resp.Body.Close()
					b, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Fatal(err)
					}
					outPath := filepath.Join(*out, fmt.Sprintf("%s-%s", p, filepath.Base(dLink)))
					f, err := os.Create(outPath)
					defer f.Close()
					if err != nil {
						log.Fatal(err)
					}
					w := bufio.NewWriter(f)
					w.Write(b)
					w.Flush()
					fmt.Printf("%s downloaded to %s\n", dLink, outPath)
				}()
			}
		})
	}
	wg.Wait()
}
