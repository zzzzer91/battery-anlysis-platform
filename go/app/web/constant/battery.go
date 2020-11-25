package constant

type dbTable struct {
	Name     string
	FieldSet map[string]struct{}
}

var BatteryNameToTable = map[string]dbTable{
	"宇通_4F37195C1A908CFBE0532932A8C0EECB": {
		Name: MongoCollectionYuTongVehicle,
		FieldSet: map[string]struct{}{
			"时间":      {},
			"总电压":     {},
			"总电流":     {},
			"车速":      {},
			"正向累计电量":  {},
			"反向累计电量":  {},
			"总里程":     {},
			"SOC":     {},
			"单体最高温度":  {},
			"单体最低温度":  {},
			"单体最高电压":  {},
			"单体最低电压":  {},
			"最高温度电池号": {},
			"最低温度电池号": {},
			"最高电压电池号": {},
			"最低电压电池号": {},
		},
	},
	"北汽_LNBSCU3HXJR884327": {
		Name: MongoCollectionBeiQiVehicle,
		FieldSet: map[string]struct{}{
			"时间":          {},
			"动力电池内部总电压V1": {},
			"动力电池充/放电电流":  {},
			"动力电池可用能量":    {},
			"动力电池可用容量":    {},
			"动力电池剩余电量SOC": {},
			"MSODO总里程":    {},
		},
	},
}
