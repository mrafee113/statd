package scripts

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"statd/config"
	"statd/pkg/utils"
)

func BatteryCharge() (string, error) {
	if config.Cfg == nil {
		return "", fmt.Errorf("UninitializedConfigError")
	}

	var (
		cfg = config.Cfg

		ac             int    = 0
		batteryLevel   uint64 = 0
		batteryMax     uint64 = 0
		batteryPercent int    = 0

		labelPrefix = '-'
		label       string
		labelColor  string
		icon        string
	)

	data, err := os.ReadFile(config.Cfg.BatteryCharge.AcPath)
	if err != nil {
		return "", err
	}
	switch strings.TrimSpace(string(data)) {
	case "1":
		ac = 1
		labelPrefix = '+'
	case "0":
		ac = 0
	default:
		return "", fmt.Errorf("ValueError: the value found in %s regarding ac was %s.",
			config.Cfg.BatteryCharge.AcPath, string(data))
	}

	data, err = os.ReadFile(config.Cfg.BatteryCharge.CurBatLvlPath)
	if err != nil {
		return "", err
	}
	batteryLevel, err = strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
	if err != nil {
		return "", err
	}

	data, err = os.ReadFile(config.Cfg.BatteryCharge.FullBatLvlPath)
	if err != nil {
		return "", err
	}
	batteryMax, err = strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
	if err != nil {
		return "", err
	}

	if batteryMax != 0 {
		batteryPercent = int(100 * batteryLevel / batteryMax)
	}

	label = fmt.Sprintf("%d%%", batteryPercent)
	if cfg.BatteryCharge.LowAt < batteryPercent && batteryPercent < cfg.BatteryCharge.FullAt {
		labelColor = config.Cfg.Colors.Info
	} else {
		labelColor = config.Cfg.Colors.Alert
	}
	label = utils.Fontify(cfg.BatteryCharge.Colorize, "1",
		utils.Colorize(cfg.BatteryCharge.Colorize, labelColor,
			fmt.Sprintf("%c%s", labelPrefix, label)))
	for _, lvlIcon := range config.Cfg.BatteryCharge.LvlIconMap {
		if batteryPercent >= lvlIcon.Threshold {
			if ac == 1 {
				icon = lvlIcon.AcIcon
			} else {
				icon = lvlIcon.BatIcon
			}
			break
		}
	}
	icon = utils.Colorize(true, cfg.Colors.Primary, icon)
	if cfg.BatteryCharge.Colorize {
		return fmt.Sprintf("%s %s", icon, label), nil
	} else {
		return label, nil
	}
}
