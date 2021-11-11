package common

type ModelProperties []map[string]interface{}

func (ma ModelProperties) Len() int {
	return len(ma)
}

func (ma ModelProperties) Less(i, j int) bool {
	return ma[i]["titledAttributeName"].(string) <= ma[j]["titledAttributeName"].(string)
}

func (ma ModelProperties) Swap(i, j int) {
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
