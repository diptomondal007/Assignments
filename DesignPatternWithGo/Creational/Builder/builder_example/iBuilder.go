package builder_example

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &NormalBuilder{}
	}
	if builderType == "igloo"{
		return &IglooBuilder{}
	}
	return nil
}
