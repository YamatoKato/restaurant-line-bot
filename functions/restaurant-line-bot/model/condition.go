package model

type Condition struct {
	Name             string
	ImgURL           string
	Desc             string
	ConditionOptions []ConditionOption
}

func CreateCondition(name, imgURL, desc string, options ...ConditionOption) Condition {
	return Condition{
		Name:             name,
		ImgURL:           imgURL,
		Desc:             desc,
		ConditionOptions: options,
	}
}

type ConditionOption struct {
	Code       string
	Name       string
	ParamKey   string
	ParamValue string
}

func CreateConditionOption(name, paramKey string, paramValue string) ConditionOption {
	return ConditionOption{
		Name:       name,
		ParamKey:   paramKey,
		ParamValue: paramValue,
	}
}
