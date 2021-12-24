package robot

type Metadata struct {
	Name       string
	Identifier string
	Detail     string
	Avatar     string
	Setting    []SettingItem
}

type SettingItemType string

const (
	SettingItemTypeString SettingItemType = "string"
	SettingItemTypeInt    SettingItemType = "int"
	SettingItemTypeFloat  SettingItemType = "float"
	SettingItemTypeBool   SettingItemType = "bool"
)

type SettingItem struct {
	Key      string          `json:"key"`
	Type     SettingItemType `json:"type"`
	Required bool            `json:"required"`
	Value    interface{}     `json:"value"`
}
