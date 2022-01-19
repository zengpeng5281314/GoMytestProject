package main

import "fmt"

func main() {
	//var countryCapitalMap map[string]string
	//countryCapitalMap := make(map[string]string)
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["北京"] = "北京"
	countryCapitalMap["湖北"] = "武汉"
	countryCapitalMap["河北"] = "石家庄"
	countryCapitalMap["湖南"] = "长沙"

	for s, s2 := range countryCapitalMap {
		fmt.Println("省分:" + s + "   省会：" + s2)
	}

	for s := range countryCapitalMap {
		fmt.Println("省分2:" + s + "   省会：" + countryCapitalMap[s])
	}

	/*查看元素在集合中是否存在 */
	capital1, ok2 := countryCapitalMap["湖北"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok2 {
		fmt.Println("湖北 的省会是", capital1)
	} else {
		fmt.Println("湖北 的省会不存在")
	}

	delete(countryCapitalMap, "北京")

	for s, s2 := range countryCapitalMap {
		fmt.Println("省分:" + s + "   省会：" + s2)
	}

	fmt.Printf("countryCapitalMap 长度  %d \n", len(countryCapitalMap))

}
