package component

import (
	"encoding/json"
	"fmt"
	"github.com/strebul/gogy/model"
	"github.com/strebul/gogy/model/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Host     string
	Login    string
	Password string
}

func (c *Client) FindLogs(query model.Request) []log.Log {
	aliases := strings.Join(c.getAliases(), ",")

	url := fmt.Sprintf("https://%s:%s@%s/%s/_search", c.Login, c.Password, c.Host, aliases)

	request := c.buildRequest(query)

	resp, err := http.Post(url, "", strings.NewReader(request))
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	var response model.Response

	errJson := json.Unmarshal(bytes, &response)

	if errJson != nil {
		panic(errJson)
	}

	var list []log.Log

	for _, hit := range response.Hits.Hit {
		t, e := time.Parse(time.RFC3339, hit.Source["@timestamp"].(string))

		if e != nil {
			panic(e)
		}

		level := log.LogLevel{}
		level.SetFromString(hit.Source["log-level"].(string))

		message := ""
		if v := hit.Source["message"]; v != nil {
			message = v.(string)
		}
		host := ""
		if v := hit.Source["host"]; v != nil {
			host = v.(string)
		}
		scriptId := ""
		if v := hit.Source["script-id"]; v != nil {
			scriptId = v.(string)
		}
		sessionId := ""
		if v := hit.Source["sessionId"]; v != nil {
			sessionId = v.(string)
		}

		list = append(list, log.Log{
			hit.Id,
			level,
			message,
			t,
			host,
			scriptId,
			sessionId,
			hit.Source,
		})
	}

	return list
}

func (c *Client) getAliases() []string {
	url := fmt.Sprintf("https://%s:%s@%s/_aliases", c.Login, c.Password, c.Host)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	var response map[string]interface{}

	errJson := json.Unmarshal(bytes, &response)
	if errJson != nil {
		panic(errJson)
	}

	var aliases []string
	for key, _ := range response {
		aliases = append(aliases, key)
	}

	return aliases
}

func (c *Client) buildRequest(q model.Request) string {
	request := `{
        "query":{
          "filtered":{
             "query":{
                "bool":{
                   "should":[
                      {
                         "query_string":{
                            "query":"%s"
                         }
                      }
                   ]
                }
             },
             "filter":{
                "bool":{
                   "must":[
                      {
                         "range":{
                            "@timestamp":{
                               "from": %d,
                               "to": %d
                            }
                         }
                      }
                   ]
                }
             }
          }
        },
        "size":%d,
        "sort":[
          {
             "@timestamp":{
                "order":"desc",
                "ignore_unmapped":true
             }
          },
          {
             "@timestamp":{
                "order":"desc",
                "ignore_unmapped":true
             }
          }
        ]
    }`

	timeStart := q.TimeStart.Unix() * 1000
	timeEnd := q.TimeEnd.Unix() * 1000

	request = fmt.Sprintf(
		request,
		strings.Replace(q.Query, "\"", "\\\"", -1),
		timeStart,
		timeEnd,
		q.Size,
	)

	return request
}
