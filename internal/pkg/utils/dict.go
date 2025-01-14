package utils

type SystemDictData struct {
	DictType  string `json:"dictType"`
	DictLabel string `json:"dictLabel"`
	DictValue string `json:"dictValue"`
	DictSort  int    `json:"dictSort"`
}

// 默认字典类型
var (
	AssistLivePlatform = "sys_internal_assist_live_platform"
	internalDictType   = []string{AssistLivePlatform}
	m                  = make(map[string]([]SystemDictData), 1)
)

func init() {
	addDictData(AssistLivePlatform, []SystemDictData{
		{DictType: AssistLivePlatform, DictLabel: "抖音", DictValue: "douyin", DictSort: 1},
	})
}

func IsInternalDictData(dictType string) bool {
	return InSliceString(dictType, &internalDictType)
}

func GetDictDataByType(dictType string) *[]SystemDictData {
	data, ok := m[dictType]
	if !ok {
		return &[]SystemDictData{}
	}
	return &data
}

func addDictData(dictType string, dictData []SystemDictData) {
	data, ok := m[dictType]
	if !ok {
		data = []SystemDictData{}
	}
	m[dictType] = append(data, dictData...)
}
