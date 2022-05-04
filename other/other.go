package other

type privateStruct struct {
	PublicField string
}

type privateInterface interface {
	Get() privateStruct
}

type getter struct {
}

func (getter) Get( ) privateStruct {
	return privateStruct{ PublicField: "hello world"}
}

func GetGetter() privateInterface {
	return getter{}
}
