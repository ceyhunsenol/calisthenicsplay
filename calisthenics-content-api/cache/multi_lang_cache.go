package cache

type MultiLangCache struct {
	Code    string
	langMap map[string]string
}

func NewMultiLangCache(code string) *MultiLangCache {
	return &MultiLangCache{
		Code:    code,
		langMap: make(map[string]string),
	}
}

func (mc *MultiLangCache) SetLangMap(val map[string]string) {
	mc.langMap = val
}

func (mc *MultiLangCache) SetByLang(langCode, value string) {
	mc.langMap[langCode] = value
}

func (mc *MultiLangCache) GetByLang(langCode string) string {
	value, exists := mc.langMap[langCode]
	if !exists {
		value, exists = mc.langMap["base"]
		if !exists {
			value = mc.Code
		}
	}
	return value
}
