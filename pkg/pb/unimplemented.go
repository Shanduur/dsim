package pb

type PBUnimplemented interface{}

func FPBUnimplemented(req PBUnimplemented) (res PBUnimplemented, err error) {
	return
}
