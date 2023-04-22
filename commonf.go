package string

import (
	"strings"

	"github.com/shopspring/decimal"
)

// 现金每三位逗号显示
func PriceCommaDisplay(price float64) string {
	// 最终结果字符串
	var comma_display_str string = ""

	// 判断正负符号
	pm := ""
	if price < 0 {
		pm = "-"
		price = 0 - price
	}

	// 浮点型转换为字符串 - 保证精度
	price_str := decimal.NewFromFloat(price).String()

	// 分割整数小数部分
	decimal_arr := strings.Split(price_str, ".")
	integer_part_str := decimal_arr[0]
	fractional_part_str := decimal_arr[1]

	// 添加逗号逻辑
	integer_part_str_len := len(integer_part_str)
	first_index := integer_part_str_len%3 - 1
	for i, ch := range integer_part_str {
		comma_display_str += string(ch)
		// 所有 3*x + first_index 为索引的位置添加逗号
		if i%3 == first_index && i != integer_part_str_len-1 {
			comma_display_str += ","
		}
	}

	return pm + comma_display_str + "." + fractional_part_str
}

// 现金每三位逗号隐藏
func PriceCommaHide(price_str string) string {
	// 逗号分割，再空字符连接
	price_str_arr := strings.Split(price_str, ",")
	price_origin := strings.Join(price_str_arr, "")
	return price_origin
}
