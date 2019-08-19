package esClient

import (
	"gin-practice/pkg/setting"
	"github.com/olivere/elastic"
)

var esConn *elastic.Client

func ConnectElasticSearchClient() (client *elastic.Client, err error) {
	esUrl := setting.Ini.Section("elasticsearch").Key("HTTP_PORT").MustString("http://127.0.0.1:9200")
	esConn, err = elastic.NewClient(elastic.SetURL("http://" + esUrl))
	return esConn, err
}