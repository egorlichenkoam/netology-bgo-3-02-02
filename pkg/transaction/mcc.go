package transaction

type Mcc string

func Mccs() map[Mcc]string {
	return map[Mcc]string{
		"5411": "Типа магазин",
		"0000": "Оплата услуг сверхсекретного агента",
		"5812": "Кто девушку платит тот ее и танцует",
		"5555": "Жижино три тотора",
		"666":  "ОплОчено",
	}
}
