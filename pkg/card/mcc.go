package card

type Mcc string

func Mccs() map[Mcc]string {
	return map[Mcc]string{
		"5411": "Типа магазин",
		"0000": "ОПлата услуг сверхсекретного агента",
		"5812": "Кто девушку платит тот ее и танцует",
		"5555": "Жижино три тотора",
	}
}
