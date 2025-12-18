package kvstruct

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type KVType string

const (
	KV_TYPE_STRING  = KVType("string")
	KV_TYPE_INT     = KVType("int")
	KV_TYPE_BOOLEAN = KVType("boolean")
	KV_TYPE_FLOAT   = KVType("float")
	KV_TYPE_JSON    = KVType("json")
)

type KV struct {
	Type  KVType `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func toJsonString(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	s := string(b)
	return s
}

func (kv KV) String() string {
	return toJsonString(kv)
}

func (kv *KV) ReplaceKey(keyPairs ...KeyPair) (err error) {
	newKey, err := KeyPairs(keyPairs).ReplaceKey(kv.Key)
	if err != nil {
		return err
	}
	kv.Key = newKey
	return nil
}

type KVS []KV

func (kvs KVS) String() string {
	return toJsonString(kvs)
}

func (kvs KVS) Json(WithType bool) (jsonStr string, err error) {
	for _, kv := range kvs {
		// 任何情况,都处理特殊处理json和boolean 类型
		if kv.Type == KV_TYPE_JSON {
			jsonStr, err = sjson.SetRaw(jsonStr, kv.Key, kv.Value)
			if err != nil {
				return "", err
			}
			continue
		}

		strValue := kv.Value
		if kv.Type == KV_TYPE_BOOLEAN {
			switch strings.ToLower(strValue) {
			case "是", "对", "1", "yes":
				strValue = "true"
			case "否", "错", "0", "no":
				strValue = "false"
			}
		}
		if !WithType {
			jsonStr, err = sjson.Set(jsonStr, kv.Key, strValue)
			if err != nil {
				return "", err
			}
			continue
		}

		var value interface{}
		value = strValue
		switch kv.Type {
		case KV_TYPE_BOOLEAN:
			value, err = strconv.ParseBool(strValue)
			if err != nil {
				return "", err
			}
		case KV_TYPE_INT:
			value, err = strconv.Atoi(strValue)
			if err != nil {
				return "", err
			}
		case KV_TYPE_FLOAT:
			value, err = strconv.ParseFloat(strValue, 64)
			if err != nil {
				return "", err
			}
		}
		jsonStr, err = sjson.Set(jsonStr, kv.Key, value)
		if err != nil {
			return "", err
		}
	}
	return jsonStr, nil
}

// GetIndex 获取相同前置后带数字名称的最大数字,加1后作为新元素下标,返回
func (kvs KVS) GetNextIndex(keyPrefix string, keySeparator string) (maxIndex int) {
	maxIndex = -1
	for _, kv := range kvs {
		if strings.HasPrefix(kv.Key, keyPrefix) && len(kv.Key) > len(keyPrefix) {
			numberStr := kv.Key[len(keyPrefix):]
			numberStr = strings.Trim(numberStr, keySeparator)
			dotIndex := strings.Index(numberStr, keySeparator)
			if dotIndex > -1 {
				numberStr = numberStr[:dotIndex]
			}
			if index, err := strconv.Atoi(numberStr); err == nil && maxIndex < index {
				maxIndex = index
			}

		}
	}
	maxIndex++ // 增加1,作为新元素下标
	return maxIndex
}

func (kvs KVS) Exists(key string) (exists bool) {
	for _, kv := range kvs {
		if key == kv.Key {
			return true
		}
	}
	return false
}

func (kvs KVS) GetFirstByKey(key string) (kv KV, index int) {
	index = -1
	for i, tplKV := range kvs {
		if key == tplKV.Key {
			return tplKV, i
		}
	}
	return kv, index
}

func (kvs KVS) GetByIndex(index int) (kv KV, exists bool) {
	if index > len(kvs)-1 || index < 0 {
		return kv, false
	}
	kv = kvs[index]
	return kv, true
}

// Order 对kv 集合排序
func (kvs KVS) Order(keyOrder []string) (orderedKVS KVS) {
	orderedKVS = make(KVS, 0)
	orderIndex := make([]int, 0)
	// 确定顺序
	for _, key := range keyOrder {
		kv, index := kvs.GetFirstByKey(key)
		if index < 0 {
			continue
		}
		orderIndex = append(orderIndex, index)
		orderedKVS = append(orderedKVS, kv)
	}

	if len(orderIndex) == len(kvs) {
		return orderedKVS
	}
	//复制剩余kv
	for i, kv := range kvs {
		notExists := true
		for _, index := range orderIndex {
			if i == index {
				notExists = false
				break
			}
		}
		if notExists {
			orderedKVS = append(orderedKVS, kv)
		}
	}

	return orderedKVS
}

func (kvs KVS) Map() (m map[string]string) {
	m = make(map[string]string, 0)
	for _, kv := range kvs {
		m[kv.Key] = kv.Value
	}
	return m
}

// Add 新增,不排除重复
func (kvs *KVS) Add(addkvs ...KV) {
	*kvs = append(*kvs, addkvs...)
}

// AddIgnore 引用解析到的kv，批量添加
func (kvs *KVS) AddIgnore(addkvs ...KV) {
	for _, addKv := range addkvs {
		exists := false
		for _, existsKv := range *kvs {
			if existsKv.Key == addKv.Key {
				exists = true
				break
			}
		}
		if !exists {
			*kvs = append(*kvs, addKv)
		}
	}
}

// AddReplace 模板解析后获取的kv，批量新增/替换
func (kvs *KVS) AddReplace(replacekvs ...KV) {
	for _, replaceKv := range replacekvs {
		exists := false
		for i, existsKv := range *kvs {
			if existsKv.Key == replaceKv.Key {
				(*kvs)[i] = replaceKv
				exists = true
				break
			}
		}
		if !exists {
			*kvs = append(*kvs, replaceKv)
		}
	}
}

// Pop 弹出key对应的元素
func (kvs *KVS) Pop(key string) (targetKv *KV, ok bool) {
	var index int
	for i, kv := range *kvs {
		if kv.Key == key {
			targetKv = &kv
			index = i
			ok = true
			break
		}
	}
	if !ok {
		return nil, false
	}
	newKVs := (*kvs)[:index]
	newKVs = append(newKVs, (*kvs)[index+1:]...)
	*kvs = newKVs
	return targetKv, ok
}

// AppendRow 在二维数组内增加行,prefix 会自动补齐后缀.
func (kvs *KVS) AppendRows(rows KVS, prefix string) {
	prefix = fmt.Sprintf("%s.", strings.TrimRight(prefix, ".")) // 确保以.结尾
	existsKvs := kvs.FillterByPrefix(prefix)
	if len(existsKvs) == 0 {
		kvs.AddReplace(rows...)
		return
	}
	maxIndex := -1
	for _, kv := range existsKvs {
		number := strings.TrimPrefix(kv.Key, prefix)
		dotIndex := strings.Index(number, ".")
		if dotIndex > -1 {
			number = number[:dotIndex]
		}
		intNumber, _ := strconv.Atoi(number) // 转不了int，直接认为0开始，所以err忽略
		if maxIndex < intNumber {
			maxIndex = intNumber
		}
	}
	//生成新的下标
	newIndex := maxIndex + 1
	for _, kv := range rows {
		ext := ""
		number := strings.Trim(strings.TrimPrefix(kv.Key, prefix), ".")
		dotIndex := strings.Index(number, ".")
		if dotIndex > -1 {
			ext = number[dotIndex:] // 先保留ext
			number = number[:dotIndex]
		}
		intNumber, _ := strconv.Atoi(number) // 转不了int，直接认为0开始，所以err忽略
		newNumber := newIndex + intNumber
		kv.Key = fmt.Sprintf("%s%d%s", prefix, newNumber, ext)
	}
}

// ReplacePrefix 引用解析获得的新数据，需要批量替换id前缀
func (kvs *KVS) ReplacePrefix(old, new string) {
	old = fmt.Sprintf("%s.", strings.TrimSuffix(old, ".")) // 确保以.结尾,避免 Dictionary.1  可以匹配 Dictionary.1.fullname,Dictionary.11.fullname
	new = fmt.Sprintf("%s.", strings.TrimSuffix(new, "."))
	for i, kv := range *kvs {
		if strings.HasPrefix(kv.Key, old) {
			kv.Key = fmt.Sprintf("%s%s", new, kv.Key[len(old):])
			(*kvs)[i] = kv
		}
	}
}

// FillterByPrefix 引用解析获得的新数据，获取指定前缀kv
func (kvs *KVS) FillterByPrefix(prefix string) (newKVs KVS) {
	newKVs = KVS{}
	for _, kv := range *kvs {
		if strings.HasPrefix(kv.Key, prefix) {
			newKVs = append(newKVs, kv)
		}
	}
	return newKVs
}

func (kvs *KVS) Fillter(fn func(kv KV) bool) (newKVs KVS) {
	newKVs = KVS{}
	for _, kv := range *kvs {
		if fn(kv) {
			newKVs = append(newKVs, kv)
		}
	}
	*kvs = newKVs
	return newKVs
}

func (kvs KVS) Walk(fn func(kvIn KV) (kvOut KV, err error)) error {
	for i := range kvs {
		kv, err := fn(kvs[i])
		if err != nil {
			return err
		}
		kvs[i] = kv
	}
	return nil
}

// ReplaceKey 替换key值(支持正则表达式)，但不改变value值,可用于转换数据格式，比如新旧接口返回字段不一致，可以将旧接口数据转换为新接口数据格式
type KeyPair struct {
	OldKeyRegexp string
	NewKeyRegexp string
	regexp       *regexp.Regexp
}

func (kp *KeyPair) compile() (err error) {
	kp.regexp, err = regexp.Compile(kp.OldKeyRegexp)
	return err
}

func (kp *KeyPair) ReplaceKey(oldKey string) (newStr string, matched bool, err error) {
	if kp.regexp == nil {
		err = kp.compile()
		if err != nil {
			return oldKey, false, err
		}
	}
	if kp.regexp.MatchString(oldKey) {
		newStr = kp.regexp.ReplaceAllString(oldKey, kp.NewKeyRegexp)
		return newStr, true, nil
	}
	return oldKey, false, nil
}

type KeyPairs []KeyPair

func (kps KeyPairs) ReplaceKey(oldKey string) (newStr string, err error) {
	for i := range kps {
		newStr, matched, err := kps[i].ReplaceKey(oldKey)
		if err != nil {
			return "", err
		}
		if matched {
			return newStr, nil
		}
	}
	return oldKey, nil
}

func (kvs KVS) ReplaceKey(keyPairs ...KeyPair) (err error) {
	kvs.Walk(func(kv KV) (KV, error) {
		err = kv.ReplaceKey(keyPairs...)
		if err != nil {
			return kv, err
		}
		return kv, nil
	})
	return nil
}

const (
	KVS_INDEX_PLACEHOLDLER = ".{index}." // 数组占位符
)

// Index 使用指定的可以模板值作为新的可以 format 格式为 xxx.{index}.yyy,  将xxx{index} 相同的部分作为一行数据,这行数据中yyy 的值将成为新的key（模拟二维数据依据指定列值转成map格式）,注意.{index}.为固定格式
func (kvs *KVS) Index(format string) (err error) {
	arr := strings.SplitN(format, KVS_INDEX_PLACEHOLDLER, 2)
	if len(arr) != 2 {
		err = errors.Errorf("required format:  xxx.{index}.yyy ,got %s", format)
		return err
	}
	prefix, suffix := arr[0], arr[1]
	patten := fmt.Sprintf("%s.%%d.%s", prefix, suffix)
	km := map[string]string{}
	for _, kv := range *kvs {
		var index int
		_, err = fmt.Sscanf(kv.Key, patten, &index)
		if err == nil {
			k := strings.TrimSuffix(kv.Key, "."+suffix)
			km[k] = fmt.Sprintf("%s.%s", prefix, strings.TrimLeft(kv.Value, "."))
		}
	}
	if len(km) == 0 {
		err = errors.Errorf("not match any key, format:%s", format)
		return err
	}
	for k, v := range km {
		kvs.ReplacePrefix(k, v)
	}
	return nil
}

// JsonToKVS 将json 转换为key->value 对,key 的规则为github.com/tidwall/gjson 的path
func JsonToKVS(jsonStr string, namespace string) (kvs KVS) {
	kvs = make(KVS, 0)
	paths := make([]string, 0)
	if !gjson.Valid(jsonStr) {
		err := errors.Errorf("invalid json:\n%s\n", jsonStr)
		panic(err)
	}
	result := gjson.Parse(jsonStr)
	allResult := getAllJsonResult(result)
	for _, result := range allResult {
		subPath := result.Path(jsonStr)
		paths = append(paths, subPath)
	}
	namespace = strings.TrimLeft(namespace, ".")
	if namespace != "" {
		namespace = fmt.Sprintf("%s.", namespace)
	}
	for _, path := range paths {
		value := result.Get(path).String()
		path := strings.TrimSuffix(path, "@this")
		key := fmt.Sprintf("%s%s", namespace, strings.Trim(path, "."))
		key = strings.Trim(key, ".")
		kv := KV{
			Key:   key,
			Value: value,
		}
		kvs = append(kvs, kv)
	}
	return kvs
}

func getAllJsonResult(result gjson.Result) (allResult []gjson.Result) {
	allResult = make([]gjson.Result, 0)
	result.ForEach(func(key, value gjson.Result) bool {
		if !value.IsArray() && !value.IsObject() {
			allResult = append(allResult, value)
		} else {
			subAllResult := getAllJsonResult(value)
			allResult = append(allResult, subAllResult...)
		}
		return true
	})
	return
}

// IsJsonStr 判断是否为json字符串
func IsJsonStr(str string) (yes bool) {
	str = strings.TrimSpace(str)
	yes = len(str) > 0 && (str[0] == '{' || str[0] == '[') && gjson.Valid(str)
	return yes

}

// FormatValue2String 将json 中所有的value转换成字符串
func FormatValue2String(jsonstr string, prefix string) (strJosn string, err error) {
	kvs := JsonToKVS(jsonstr, prefix)
	strJosn, err = kvs.Json(false)
	if err != nil {
		return "", err
	}
	return strJosn, nil
}
