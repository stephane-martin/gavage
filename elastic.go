package main

import (
	"net/http"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/olivere/elastic"
)

func BuildElasticClient(args Arguments, logger log15.Logger) (*elastic.Client, error) {
	off := time.Duration(-1)
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}
	esClient, err := elastic.NewClient(
		elastic.SetURL(args.GetEsURL()),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHealthcheckInterval(off),
		elastic.SetHealthcheckTimeout(off),
		elastic.SetHealthcheckTimeoutStartup(off),
		elastic.SetSnifferInterval(off),
		elastic.SetSnifferTimeout(off),
		elastic.SetSnifferTimeoutStartup(off),
		elastic.SetHttpClient(httpClient),
		elastic.SetErrorLog(AdaptElasticErrorLogger(logger)),
		elastic.SetInfoLog(AdaptElasticInfoLogger(logger)),
	)
	if err != nil {
		return nil, err
	}
	return esClient, nil
}
