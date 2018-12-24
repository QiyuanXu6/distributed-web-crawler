package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"web-crawler/engine"
	"web-crawler/model"
)


const jsonRe = `window\.__INITIAL_STATE__=(.+);\(function\(\){var s;`
func ParseProfile(content []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(jsonRe)

	profile := model.Profile{}
	submatch := re.FindSubmatch(content)

	if submatch != nil {
		var dat map[string]interface{}
		if err := json.Unmarshal(submatch[1], &dat); err != nil {
			panic(err)
		}

		objectInfo := dat["objectInfo"].(map[string]interface{})
		profile.Age = fmt.Sprintf("%.0f", objectInfo["age"].(float64))
		profile.AvatarUrl = objectInfo["avatarURL"].(string)
		profile.BasicInfo = fmt.Sprintf("%v", objectInfo["basicInfo"])
		profile.DetailInfo = fmt.Sprintf("%v", objectInfo["detailInfo"])
		profile.Education = objectInfo["educationString"].(string)
		profile.Gender = objectInfo["genderString"].(string)
		profile.Height = objectInfo["heightString"].(string)
		profile.Marriage = objectInfo["marriageString"].(string)
		//profile.Name = dat["nickname"].(string)
		profile.Name = name
		profile.Salary = objectInfo["salaryString"].(string)

	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	fmt.Println("value %v\n", result.Items[0])
	return result
}