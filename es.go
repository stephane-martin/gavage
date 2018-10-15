package main

import "encoding/json"

func GetEsOpts(typename string, shards int, replicas int) string {
	b, _ := json.Marshal(newEsOpts(typename, shards, replicas))
	return string(b)
}

func newEsOpts(typename string, shards int, replicas int) map[string]interface{} {
	return map[string]interface{}{
		"settings": newSettings(shards, replicas),
		"mappings": newMappings(typename),
	}
}

type esSettings struct {
	Shards   int    `json:"number_of_shards"`
	Replicas int    `json:"number_of_replicas"`
	Wait     string `json:"index.write.wait_for_active_shards"`
}

func newSettings(s int, r int) esSettings {
	return esSettings{
		Shards:   s,
		Replicas: r,
		Wait:     "all",
	}
}

func newMappings(typename string) map[string]interface{} {
	return map[string]interface{}{
		typename: esType{
			Properties: newMessageFields(),
		},
	}
}

type esType struct {
	Properties esFields `json:"properties"`
}

type esFields map[string]anyEsField

func newMessageFields() esFields {
	return map[string]anyEsField{
		"cs-bytes":          newLongField(),
		"sc-bytes":          newLongField(),
		"cs-categories":     newKeyword(),
		"c-ip":              newIPField(),
		"cs-uri-scheme":     newKeyword(),
		"cs-host":           newKeyword(),
		"tld":               newKeyword(),
		"topdomain":         newKeyword(),
		"cs-method":         newKeyword(),
		"cs-uri-extension":  newKeyword(),
		"cs-uri-path":       newKeyword(),
		"cs-uri-port":       newLongField(),
		"cs-uri-query":      newKeyword(),
		"cs-h-user-agent":   newKeyword(),
		"rs-h-content-type": newKeyword(),
		"time-taken":        newDoubleField(),
		"cs-username":       newKeyword(),
		"s-action":          newKeyword(),
		"sc-filter-result":  newKeyword(),
		"x-exception-id":    newKeyword(),
		"sc-status":         newLongField(),
		"spamitude":         newDoubleField(),
		"work":              newBoolField(),
		"@timestamp":        newDatetimeField(),
		"dayofweek":         newLongField(),
		"date":              newDateField(),
		"time":              newKeyword(),
		"year":              newLongField(),
		"yearmonth":         newKeyword(),
	}
}

type anyEsField interface{}

type doubleEsField struct {
	Typ   string `json:"type"`
	Store bool   `json:"store"`
}

func newDoubleField() doubleEsField {
	return doubleEsField{
		Typ: "double",
	}
}

type boolEsField struct {
	Typ   string `json:"type"`
	Store bool   `json:"store"`
}

func newBoolField() boolEsField {
	return boolEsField{
		Typ: "boolean",
	}
}

type ipEsField struct {
	Typ   string `json:"type"`
	Store bool   `json:"store"`
}

func newIPField() ipEsField {
	return ipEsField{
		Typ: "ip",
	}
}

type strMultiEsField struct {
	Typ    string     `json:"type"`
	Fields rawEsField `json:"fields,omitempty"`
	Copy   string     `json:"copy_to,omitempty"`
}

type rawEsField struct {
	Raw rawRawEsField `json:"raw"`
}

type rawRawEsField struct {
	Typ string `json:"type"`
}

func newMulti() strMultiEsField {
	return strMultiEsField{
		Typ:  "text",
		Copy: "fulltext",
		Fields: rawEsField{
			Raw: rawRawEsField{
				Typ: "keyword",
			},
		},
	}
}

type strEsField struct {
	Typ  string `json:"type"`
	Copy string `json:"copy_to,omitempty"`
}

func newKeyword() strEsField {
	return strEsField{
		Typ: "keyword",
	}
}

func newTextField(copyfull bool) strEsField {
	return strEsField{
		Typ: "text",
	}
}

type dateEsField struct {
	Typ    string `json:"type"`
	Format string `json:"format"`
}

func newDateField() dateEsField {
	return dateEsField{
		Typ:    "date",
		Format: "strict_date",
	}
}

type datetimeEsField struct {
	Typ    string `json:"type"`
	Format string `json:"format"`
}

func newDatetimeField() datetimeEsField {
	return datetimeEsField{
		Typ:    "date",
		Format: "strict_date_time_no_millis||strict_date_time",
	}
}

type longEsField struct {
	Typ string `json:"type"`
}

func newLongField() longEsField {
	return longEsField{
		Typ: "long",
	}
}
