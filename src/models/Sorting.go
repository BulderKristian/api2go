package models

type ModelAttributes []map[string]interface{}

func (ma ModelAttributes) Len() int {
	return len(ma)
}

func (ma ModelAttributes) Less(i, j int) bool {
	return ma[i]["titeledAttributeName"].(string) <= ma[j]["titeledAttributeName"].(string)
}

func (ma ModelAttributes) Swap(i, j int) {
	ma[i], ma[j] = ma[j], ma[i]
}

type ModelEnums []map[string]string

func (me ModelEnums) Len() int {
	return len(me)
}

func (me ModelEnums) Less(i, j int) bool {
	return me[i]["enumConstName"] <= me[j]["enumConstName"]
}

func (me ModelEnums) Swap(i, j int) {
	me[i], me[j] = me[j], me[i]
}
