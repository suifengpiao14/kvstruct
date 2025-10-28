package kvstruct_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/kvstruct"
)

func TestKVSJson(t *testing.T) {
	t.Run("boolen", func(t *testing.T) {
		kv := kvstruct.KV{
			Type:  kvstruct.KV_TYPE_BOOLEAN,
			Key:   "doc.example.response.200.language",
			Value: `是`,
		}
		kvs := make(kvstruct.KVS, 0)
		kvs = append(kvs, kv)
		jsonStr, err := kvs.Json(true)
		require.NoError(t, err)
		fmt.Println(jsonStr)
	})

	t.Run("many", func(t *testing.T) {
		kv := kvstruct.KV{
			Key:   "doc.parameter.response.16.type",
			Value: "string",
		}
		kvs := make(kvstruct.KVS, 0)
		kvs = append(kvs, kv)
		kvs = append(kvs, kvstruct.KV{Key: "a", Value: "b"})
		jsonStr, err := kvs.Json(false)
		require.NoError(t, err)
		fmt.Println(jsonStr)
	})
}

func TestJsonToKVS(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		jsonstr := `{"ad":{"advertise":{"id":{"database":"ad","table":"advertise","name":"id","goType":"int","dbType":"int(11)","comment":"主键","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"false","size":"11"},"advertiser_id":{"database":"ad","table":"advertise","name":"advertiser_id","goType":"string","dbType":"varchar(32)","comment":"广告主","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"title":{"database":"ad","table":"advertise","name":"title","goType":"string","dbType":"varchar(32)","comment":"广告标题","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"begin_at":{"database":"ad","table":"advertise","name":"begin_at","goType":"string","dbType":"datetime","comment":"投放开��时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},"end_at":{"database":"ad","table":"advertise","name":"end_at","goType":"string","dbType":"datetime","comment":"投放结束时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},"summary":{"database":"ad","table":"advertise","name":"summary","goType":"string","dbType":"varchar(128)","comment":"广告素材-文字描述","nullable":"true","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"128"},"image":{"database":"ad","table":"advertise","name":"image","goType":"string","dbType":"varchar(256)","comment":"广告素材-图片地址","nullable":"true","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},"link":{"database":"ad","table":"advertise","name":"link","goType":"string","dbType":"varchar(512)","comment":"连接地址","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"512"},"remark":{"database":"ad","table":"advertise","name":"remark","goType":"string","dbType":"varchar(255)","comment":"备注","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"255"},"type":{"database":"ad","table":"advertise","name":"type","goType":"string","dbType":"enum('text','image','vido')","comment":"广告素材(类型),text-文字,image-图片,vido-视频","nullable":"false","enums":"text,image,vido","autoIncrement":"false","default":"text","onUpdate":"false","unsigned":"false","size":"0"},"value_obj":{"database":"ad","table":"advertise","name":"value_obj","goType":"string","dbType":"varchar(1024)","comment":"json扩展,广告的值属性对象","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"1024"},"created_at":{"database":"ad","table":"advertise","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"ad","table":"advertise","name":"updated_at","goType":"string","dbType":"datetime","comment":"修改时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},"deleted_at":{"database":"ad","table":"advertise","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"}},"window":{"id":{"database":"ad","table":"window","name":"id","goType":"int","dbType":"int(11) unsigned","comment":"主键","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"true","size":"11"},"code":{"database":"ad","table":"window","name":"code","goType":"string","dbType":"varchar(32)","comment":"位置编码","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"title":{"database":"ad","table":"window","name":"title","goType":"string","dbType":"varchar(32)","comment":"位置名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"remark":{"database":"ad","table":"window","name":"remark","goType":"string","dbType":"varchar(255)","comment":"位置描述","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"255"},"content_types":{"database":"ad","table":"window","name":"content_types","goType":"string","dbType":"varchar(50)","comment":"广告素材(类型),text-文字,image-图片,vido-视频,多个逗号分隔","nullable":"true","enums":"","autoIncrement":"false","default":"text","onUpdate":"false","unsigned":"false","size":"50"},"width":{"database":"ad","table":"window","name":"width","goType":"int","dbType":"smallint(6)","comment":"橱窗宽度","nullable":"true","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"6"},"high":{"database":"ad","table":"window","name":"high","goType":"int","dbType":"smallint(6)","comment":"橱窗高度","nullable":"true","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"6"},"created_at":{"database":"ad","table":"window","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"ad","table":"window","name":"updated_at","goType":"string","dbType":"datetime","comment":"修改时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},"deleted_at":{"database":"ad","table":"window","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"}},"window_advertise":{"id":{"database":"ad","table":"window_advertise","name":"id","goType":"int","dbType":"int(11)","comment":"主键","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"false","size":"11"},"code":{"database":"ad","table":"window_advertise","name":"code","goType":"string","dbType":"varchar(32)","comment":"橱窗编码","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"32"},"advertise_id":{"database":"ad","table":"window_advertise","name":"advertise_id","goType":"int","dbType":"int(11)","comment":"广告ID","nullable":"false","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"11"},"advertise_priority":{"database":"ad","table":"window_advertise","name":"advertise_priority","goType":"int","dbType":"int(11)","comment":"广告优先级(同一个橱窗有多个广告时,按照优先级展示)","nullable":"true","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"false","size":"11"},"created_at":{"database":"ad","table":"window_advertise","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"ad","table":"window_advertise","name":"updated_at","goType":"string","dbType":"datetime","comment":"修改时间","nullable":"true","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},"deleted_at":{"database":"ad","table":"window_advertise","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"}}}}`
		kvs := kvstruct.JsonToKVS(jsonstr, "root")
		fmt.Println(kvs)
	})
	t.Run("int", func(t *testing.T) {
		jsonstr := "2"
		kvs := kvstruct.JsonToKVS(jsonstr, "")
		fmt.Println(kvs)
	})

}
func TestFormatValue2String(t *testing.T) {
	jsonstr := `{"id":1,"list":["3",4,"aa"],"jsonKey":"{\"id\":4}"}`
	newJsonstr, err := kvstruct.FormatValue2String(jsonstr, "root")
	require.NoError(t, err)
	fmt.Println(newJsonstr)

}

func TestGetNextIndex(t *testing.T) {
	keySeparator := "."
	t.Run("end number", func(t *testing.T) {
		prefix := "doc.request.parameter."
		kvs := kvstruct.KVS{
			kvstruct.KV{Key: fmt.Sprintf("%s2", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 3, nextIndex)
	})

	t.Run("middle number", func(t *testing.T) {
		prefix := "doc.request.parameter."
		kvs := kvstruct.KVS{
			kvstruct.KV{Key: fmt.Sprintf("%s2.hahhah", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 3, nextIndex)
	})
	t.Run("star number", func(t *testing.T) {
		prefix := "d"
		kvs := kvstruct.KVS{
			kvstruct.KV{Key: fmt.Sprintf("%s2.hahhah", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 3, nextIndex)
	})
	t.Run("no number", func(t *testing.T) {
		prefix := "d"
		kvs := kvstruct.KVS{
			kvstruct.KV{Key: fmt.Sprintf("%s2eigj.hahhah", prefix)},
		}
		nextIndex := kvs.GetNextIndex(prefix, keySeparator)
		assert.Equal(t, 0, nextIndex)
	})

}

func TestOrder(t *testing.T) {
	kvs := kvstruct.KVS{
		{Key: "a", Value: "value_a"},
		{Key: "c", Value: "value_c"},
		{Key: "b", Value: "value_b"},
	}
	keyOrder := []string{"b", "a", "d"}
	expectedKeys := []string{"b", "a", "c"}
	orderedKVS := kvs.Order(keyOrder)
	ok := len(orderedKVS) == len(expectedKeys)
	if ok {
		for i, kv := range orderedKVS {
			key := expectedKeys[i]
			if key != kv.Key {
				ok = false
			}
		}
	}
	assert.Equal(t, true, ok)
}

func TestArrSlic(t *testing.T) {
	arr := []int{1, 2}
	subArr := arr[1+1:]
	fmt.Println(subArr)
}

func TestIndex(t *testing.T) {
	kvs := kvstruct.KVS{}
	kvs.Add(kvstruct.KV{
		Key:   "Dictionary.0.fullname",
		Value: "viedo.id",
	}, kvstruct.KV{
		Key:   "Dictionary.0.type",
		Value: "int",
	},
		kvstruct.KV{
			Key:   "Dictionary.1.fullname",
			Value: "viedo.name",
		},
		kvstruct.KV{
			Key:   "Dictionary.1.type",
			Value: "string",
		},
	)

	err := kvs.Index("Dictionary.{index}.fullname")
	require.NoError(t, err)
	fmt.Println(kvs.Json(false))
}

func TestIndex2(t *testing.T) {
	kvsJson := `[{"type":"","key":"Dictionary.0.fullname","value":"pagination.index"},{"type":"","key":"Dictionary.0.title","value":"页码"},{"type":"","key":"Dictionary.0.explain","value":""},{"type":"","key":"Dictionary.1.fullname","value":"pagination.size"},{"type":"","key":"Dictionary.1.title","value":"每页数量"},{"type":"","key":"Dictionary.1.explain","value":""},{"type":"","key":"Dictionary.2.fullname","value":"pagination.total"},{"type":"","key":"Dictionary.2.title","value":"总数"},{"type":"","key":"Dictionary.2.explain","value":""},{"type":"","key":"Dictionary.3.fullname","value":"limit.size"},{"type":"","key":"Dictionary.3.title","value":"单次记录数"},{"type":"","key":"Dictionary.3.explain","value":"sql limit size"},{"type":"","key":"Dictionary.4.fullname","value":"limit.offset"},{"type":"","key":"Dictionary.4.title","value":"查询记录偏移量"},{"type":"","key":"Dictionary.4.explain","value":"sql limit ,offset"},{"type":"","key":"Dictionary.5.fullname","value":"class.api.withSelf"},{"type":"","key":"Dictionary.5.title","value":"包含节点本身"},{"type":"","key":"Dictionary.5.explain","value":"获取子分类时，是否返回父节点(1-是,2-否)"},{"type":"","key":"Dictionary.6.fullname","value":"class.api.depth"},{"type":"","key":"Dictionary.6.title","value":"深度"},{"type":"","key":"Dictionary.6.explain","value":"分类是树型结构数据，深度表示数深度，如1表示直接子节点"},{"type":"","key":"Dictionary.7.fullname","value":"request.params.token"},{"type":"","key":"Dictionary.7.title","value":"jwt token"},{"type":"","key":"Dictionary.7.explain","value":"用户登录token"},{"type":"","key":"Dictionary.8.fullname","value":"content.position"},{"type":"","key":"Dictionary.8.title","value":"内容展示位置"},{"type":"","key":"Dictionary.8.explain","value":"前端页面元素位置格式:pageName.elementName[...elementName]"},{"type":"","key":"Dictionary.9.fullname","value":"content.format"},{"type":"","key":"Dictionary.9.title","value":"内容形式"},{"type":"","key":"Dictionary.9.explain","value":"一定的内容需要一定的形式承载，形式确定后，输出的内容格式是固定的"},{"type":"","key":"Dictionary.10.fullname","value":"nav.url"},{"type":"","key":"Dictionary.10.title","value":"导航跳转地址"},{"type":"","key":"Dictionary.10.explain","value":"导航跳转地址"},{"type":"","key":"Dictionary.11.fullname","value":"response.code"},{"type":"","key":"Dictionary.11.title","value":"业务状态码"},{"type":"","key":"Dictionary.11.explain","value":"业务状态码，0-正常，非0-失败"},{"type":"","key":"Dictionary.12.fullname","value":"response.message"},{"type":"","key":"Dictionary.12.title","value":"业务提示"},{"type":"","key":"Dictionary.12.explain","value":"业务提示，失败时，表达失败原因"},{"type":"","key":"Dictionary.13.fullname","value":"cognition.catalog.introduceObj.title"},{"type":"","key":"Dictionary.13.title","value":"标题"},{"type":"","key":"Dictionary.13.explain","value":"课程介绍项标题"},{"type":"","key":"Dictionary.14.fullname","value":"cognition.catalog.introduceObj.content"},{"type":"","key":"Dictionary.14.title","value":"内容"},{"type":"","key":"Dictionary.14.explain","value":"课程介绍项内容"},{"type":"","key":"Dictionary.15.fullname","value":"cognition.catalog.introduceObj.sort"},{"type":"","key":"Dictionary.15.title","value":"排序"},{"type":"","key":"Dictionary.15.explain","value":"课程介绍项排序,倒序"}]`
	kvs := kvstruct.KVS{}
	err := json.Unmarshal([]byte(kvsJson), &kvs)
	require.NoError(t, err)
	format := "Dictionary.{index}.fullname"
	err = kvs.Index(format)
	require.NoError(t, err)
	newKvsJsonb, err := kvs.Json(false)
	require.NoError(t, err)
	newKvsJson := string(newKvsJsonb)
	fmt.Println(newKvsJson)
	require.NotContains(t, newKvsJson, "size0")
}

func TestReplaceKey(t *testing.T) {
	kv := kvstruct.KV{
		Key: "data.data.0.fxy_spu_id",
	}
	repair := kvstruct.KeyPair{
		OldKeyRegexp: `data.data.(\d+).fxy_spu_id`,
		NewKeyRegexp: `_data.items.$1.xySpuId`,
	}
	err := kv.ReplaceKey(repair)
	require.NoError(t, err)
	require.Equal(t, "_data.items.0.xySpuId", kv.Key)

}
