package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"time"
	"gopkg.in/mgo.v2"
)

type MoocHotArticle struct {
	Title  string
	Link string
}

func getMooc()  {

	session, err := mgo.Dial("113.209.119.170:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("MoocHotArticle").C("moochotarticle")

	for i:=1; i<3 ; i++  {
		url := "https://www.imooc.com/article/hot/" + strconv.Itoa(i)
		fmt.Println(url)
		doc, err := goquery.NewDocument( url )
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(doc)
		doc.Find(".title-detail").Each(func(i int, s *goquery.Selection){
			fmt.Println(s.Text())
			link, _ := s.Attr("href")
			fullLink := "https://www.imooc.com"+link
			fmt.Println( fullLink )

			error := c.Insert(&MoocHotArticle{ s.Text(), fullLink })
			if error != nil {
				log.Fatal(err)
			}
		})

		time.Sleep(time.Minute)
	}
}

func main(){
	getMooc()
}




