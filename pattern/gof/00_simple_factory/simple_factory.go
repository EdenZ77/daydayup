package simple_factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (p jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (p yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
