package es

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/olivere/elastic.v5"
)

const (
	URL="http://localhost:9200"
)
var (
	client *elastic.Client
)

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(URL))
	if err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}
	info, code, err := client.Ping(URL).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}


func Create(Index,Type string,Id string,doc interface{}) (*elastic.IndexResponse,error){
	resp, err := client.Index().Index(Index).Type(Type).Id(Id).BodyJson(doc).Do(context.Background())
	return resp,err
}

func Get(Index,Type,Id string)(*elastic.GetResult, error) {
	result, err := client.Get().Index(Index).Type(Type).Id(Id).Do(context.Background())
	return result,err
}

func Delete(Index,Type,Id string)(*elastic.DeleteResponse, error) {
	res, err := client.Delete().Index(Index).Type(Type).Id(Id).Do(context.Background())
	return res,err
}