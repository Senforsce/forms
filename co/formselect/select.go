package formselect

type ComponentSelectConfig struct {
	InputType   string
	Placeholder string
	Name        string
	Id          string
	OptionList  []ComponentOptionConfig
}

type ComponentOptionConfig struct {
	Name  string
	Value string
}
