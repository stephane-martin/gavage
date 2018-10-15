package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/inconshreveable/log15"
	"github.com/olivere/elastic"
	cli "gopkg.in/urfave/cli.v1"
)

var Version string

func main() {
	BuildCliApp().Run(os.Args)
}

func MainCommand(c *cli.Context) error {
	args := BuildArguments(c)
	loglevel := strings.TrimSpace(c.GlobalString("loglevel"))
	toSyslog := c.GlobalBool("syslog")

	err := args.Verify()
	if err != nil {
		return ExitCliWithError("Arguments validation failed", err)
	}

	if len(loglevel) == 0 {
		loglevel = "info"
	}
	logger := getLogger(loglevel, toSyslog)
	logger.Debug(
		"Arguments",
		"eshost", args.EsHost,
		"esport", args.EsPort,
	)
	clt, err := BuildElasticClient(args, logger)
	if err != nil {
		return ExitCliWithError("Error bulding Elasticsearch client", err)
	}
	resp, err := clt.ClusterHealth().Do(context.Background())
	if err != nil {
		return ExitCliWithError("Error checking ES health", err)
	}
	logger.Info("Elasticsearch cluster health", "status", resp.Status)
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	clients := make(map[int]*elastic.Client)
	for _, month := range months {
		clt, err := BuildElasticClient(args, logger)
		if err != nil {
			return ExitCliWithError("Error bulding Elasticsearch client", err)
		}
		clients[month] = clt
	}
	var wg sync.WaitGroup
	for _, month := range months {
		m := month
		wg.Add(1)
		go func() {
			upload(clients[m], m, args, logger)
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}

func upload(clt *elastic.Client, month int, args Arguments, logger log15.Logger) error {
	nbLogLines := args.Size * 10000000 / 12
	indexName := fmt.Sprintf("%s-%04d-%02d", args.EsPrefix, args.Year, month)

	_, err := clt.CreateIndex(indexName).BodyString(GetEsOpts(indexName, 16, 0)).Do(context.Background())
	if err != nil {
		logger.Warn("Error creating index", "name", indexName, "error", err)
	}

	gen := NewMonthGenerator(args.Year, month, nbLogLines)

	for {
		res := gen.Next()
		if len(res) == 0 {
			break
		}
		bulk := clt.Bulk()
		for _, line := range res {
			bulk = bulk.Add(elastic.NewBulkIndexRequest().Index(indexName).Type(indexName).Doc(&line))
		}
		resp, err := bulk.Do(context.Background())
		if err != nil {
			return err
		}
		failed := resp.Failed()
		if len(failed) > 0 {
			logger.Warn("Some documents failed to be uploaded", "nb", len(failed))
		} else {
			logger.Debug("Progress", "month", month, "percent", gen.Percent())
		}

	}
	return nil
}
