package config

type Config struct {
	Colors        Colors        `yaml:"colors"`
	Date          Date          `yaml:"date"`
	BatteryCharge BatteryCharge `yaml:"BatteryCharge"`
}

type Colors struct {
	Primary string `yaml:"primary"`
	Alert   string `yaml:"alert"`
	Info    string `yaml:"info"`
}

type Date struct {
	Colorize bool       `yaml:"colorize"`
	Format   string     `yaml:"format"`
	JDate    string     `yaml:"jdate"`
	GDate    string     `yaml:"gdate"`
	Time     string     `yaml:"time"`
	Colors   DateColors `yaml:"colors"`
	Icons    DateIcons  `yaml:"icons"`
}

type DateColors struct {
	Saturday  string `yaml:"saturday"`
	Sunday    string `yaml:"sunday"`
	Monday    string `yaml:"monday"`
	Tuesday   string `yaml:"tuesday"`
	Wednesday string `yaml:"wednesday"`
	Thursday  string `yaml:"thursday"`
	Friday    string `yaml:"friday"`

	January   string `yaml:"january"`
	February  string `yaml:"february"`
	March     string `yaml:"march"`
	April     string `yaml:"april"`
	May       string `yaml:"may"`
	June      string `yaml:"june"`
	July      string `yaml:"july"`
	August    string `yaml:"august"`
	September string `yaml:"september"`
	October   string `yaml:"october"`
	November  string `yaml:"november"`
	December  string `yaml:"december"`

	Spring string `yaml:"spring"`
	Summer string `yaml:"summer"`
	Autumn string `yaml:"autumn"`
	Winter string `yaml:"winter"`
}

type DateIcons struct {
	Spring string `yaml:"spring"`
	Summer string `yaml:"summer"`
	Autumn string `yaml:"autumn"`
	Winter string `yaml:"winter"`

	Time      string `yaml:"time"`
	Gregorian string `yaml:"gregorian"`
	Jalali    string `yaml:"jalali"`
}

type BatteryCharge struct {
	Colorize       bool      `yaml:"colorize"`
	FullAt         int       `yaml:"full-at"`
	LowAt          int       `yaml:"low-at"`
	AcPath         string    `yaml:"ac-path"`
	CurBatLvlPath  string    `yaml:"cur-bat-lvl-path"`
	FullBatLvlPath string    `yaml:"full-bat-lvl-path"`
	LvlIconMap     []LvlIcon `yaml:"lvl-icon-map"`
}

type LvlIcon struct {
	Threshold int    `yaml:"threshold"`
	AcIcon    string `yaml:"ac-icon"`
	BatIcon   string `yaml:"bat-icon"`
}
