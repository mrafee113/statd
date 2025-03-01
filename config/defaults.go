package config

var DefaultConfig = Config{
	Colors: Colors{
		Primary: "#987CD6",
		Alert:   "#A54242",
		Info:    "#adadad",
	},
	Date: Date{
		Colorize: true,
		Format:   "{{.Icons.Time}} {{.time}} {{.dayOfWeek}} {{.Icons.Jalali}} {{.jdate}} {{.Icons.Gregorian}} {{.gdate}} {{.seasonIcon}}",
		JDate:    "%Y %B %d",     // jalali date format
		GDate:    "%Y %b(%m) %d", // # gregorian date format
		Time:     "%H:%M:%S",
		Colors: DateColors{
			Saturday:  "#924CC4",
			Sunday:    "#EE82EE",
			Monday:    "#FF0000",
			Tuesday:   "#FFA500",
			Wednesday: "#FFFF00",
			Thursday:  "#008000",
			Friday:    "#0000FF",

			January:   "#7A2F75",
			February:  "#924CC4",
			March:     "#388C93",
			April:     "#389341",
			May:       "#87A043",
			June:      "#98A043",
			July:      "#BF852F",
			August:    "#BF692F",
			September: "#BF582F",
			October:   "#BA3B2A",
			November:  "#BA2A3D",
			December:  "#872424",

			Spring: "#389341",
			Summer: "#BF852F",
			Autumn: "#BA3B2A",
			Winter: "#924CC4",
		},
		Icons: DateIcons{
			Spring: "󰉊",
			Summer: "",
			Autumn: "󰲓",
			Winter: "",

			Time:      "󱑆",
			Gregorian: "",
			Jalali:    "",
		},
	},
	BatteryCharge: BatteryCharge{
		Colorize:       true,
		FullAt:         65,
		LowAt:          35,
		AcPath:         "/sys/class/power_supply/AC0/online",
		CurBatLvlPath:  "/sys/class/power_supply/BAT0/energy_now",
		FullBatLvlPath: "/sys/class/power_supply/BAT0/energy_full",
		LvlIconMap: []LvlIcon{
			{Threshold: 100, AcIcon: "󰂅", BatIcon: "󰁹"},
			{Threshold: 90, AcIcon: "󰂋", BatIcon: "󰂂"},
			{Threshold: 80, AcIcon: "󰂊", BatIcon: "󰂁"},
			{Threshold: 70, AcIcon: "󰢞", BatIcon: "󰂀"},
			{Threshold: 60, AcIcon: "󰂉", BatIcon: "󰁿"},
			{Threshold: 50, AcIcon: "󰢝", BatIcon: "󰁾"},
			{Threshold: 40, AcIcon: "󰂈", BatIcon: "󰁽"},
			{Threshold: 30, AcIcon: "󰂇", BatIcon: "󰁼"},
			{Threshold: 20, AcIcon: "󰂆", BatIcon: "󰁻"},
			{Threshold: 10, AcIcon: "󰢜", BatIcon: "󰁺"},
			{Threshold: 0, AcIcon: "󰢜", BatIcon: "󰂎"},
		},
	},
}
