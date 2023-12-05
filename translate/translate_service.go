package translate

import (
	"handy-translate/config"
	"handy-translate/translate/baidu"
	"handy-translate/translate/caiyun"
	"handy-translate/translate/deepl"
	"handy-translate/translate/youdao"
	"sync"
)

type Translate interface {
	GetName() string
	PostQuery(query, sourceLang, targetLang string) ([]string, error)
}

func GetTransalteWay(way string) Translate {
	var t Translate
	switch way {
	case youdao.Way:
		t = &youdao.Youdao{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	case caiyun.Way:
		t = &caiyun.Caiyun{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	case baidu.Way:
		t = &baidu.Baidu{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	case deepl.Way:
		t = &deepl.DeepL{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	}

	return t
}

var queryText string

var lk sync.RWMutex

// SetQueryText
func SetQueryText(value string) {
	lk.Lock()
	queryText = value
	lk.Unlock()
}

// GetQueryText
func GetQueryText() string {
	lk.RLock()
	defer lk.RUnlock()
	return queryText
}
