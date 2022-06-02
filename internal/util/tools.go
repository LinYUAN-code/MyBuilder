package util


type KeyValue struct {
	Name string
	Value string
	IsStmt bool
}

func StructToJson(attri... KeyValue) string {
	ans := "{"
	for _,item := range attri {
		ans += "\t" + "\"" + item.Name + "\": " 
		if !item.IsStmt {
			ans += "\""
		}
		ans += item.Value
		if !item.IsStmt {
			ans += "\""
		}
		ans += "\n"
	}
	ans += "}"
	return ans
}

func KV(name string,value string,isStmt bool) KeyValue {
	return KeyValue{
		Name: name,
		Value: value,
		IsStmt: isStmt,
	}
}

func JsonArray(value... string) string {
	ans := "[ "
	for index,item := range value {
		if index == len(value)-1 {
			ans += item + " "
		} else {
			ans += item + ", "
		}
	}
	ans += "]"
	return ans
}