package identity

import (
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
)

// IsValid
//  Description: 是否有效身份证
//  Author: Kevin·CC
//  Param: idCard 身份证
//  return bool true=有效,false=无效
func IsValid(idCard string) bool {
	for _, length := range constant.Any {
		if len(idCard) == length.Value() {
			factory, err := New(length)
			if err != nil {
				return false
			}

			return factory.IsLocalValidIDCard(idCard)
		}
	}
	return false
}