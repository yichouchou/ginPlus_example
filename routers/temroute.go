package routers

import (
	"reflect"

	"github.com/yichouchou/ginPlus/annotation"
	"github.com/yichouchou/ginPlus/utils"

	bind "github.com/yichouchou/ginPlus/bind"
)

func init() {
	annotation.SetVersion(1637040167)

	QEPDHMPLUY := new(string)

	ZWRSZZAVWH := new(string)

	EDILPDWJHV := new(int)

	QRCUHTHFZD := new(bind.ReqTest)

	RJEWUSGMXE := new(bind.ReqTest)

	TWYPJDMMEJ := new(string)

	YYXKIBTHNA := new(string)

	PAGPNUIZSS := new(int)

	RWUXYAHLZN := new(int)

	HRXATBWGNO := new([]bind.ReqTest)

	YGJLTOFLRM := new([]*bind.ReqTest)

	AZZBNQQRJC := new(bind.ReqTest)

	BVGPCEUEBQ := new(string)

	CWHVYAVHJV := new(int)

	PAWHLPFHLJ := new(string)

	FZUICEJIYJ := new(int)

	BQNTUSYANZ := new(string)

	DMAMECUTVF := new(error)

	SPLMPWFGLZ := new(bind.ReqTest)

	DEQYWBJNXP := new(error)

	TDQQQELAFW := new(int)

	AKFIMOYJCJ := new(error)

	AKNIKSEUSQ := new(int)

	MNAVPJTEVU := new(error)

	SUJPGSFDCO := new(int)

	LVODAAWUKQ := new(error)

	PQMTYSUWDD := new(int)

	SIVHHIONEY := new(error)

	ZXMUHQVKIF := new(int)

	RWJFLJWWOX := new(error)

	JOWMNGSYGW := new(bool)

	ERWETHLSJL := new(bool)

	annotation.AddGenOne("Hello.Hi1", utils.GenComment{
		RouterPath: "/block1",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*QEPDHMPLUY),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(*ZWRSZZAVWH),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*EDILPDWJHV),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(*QRCUHTHFZD),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "hi",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi1",
				ParmType: reflect.TypeOf(*BQNTUSYANZ),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "errHi1",
				ParmType: reflect.TypeOf(*DMAMECUTVF),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi2", utils.GenComment{
		RouterPath: "/block2",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "hiValue",
				ParmType: reflect.TypeOf(*RJEWUSGMXE),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "hi",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi2",
				ParmType: reflect.TypeOf(*SPLMPWFGLZ),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},

			{
				ParmName: "errHi2",
				ParmType: reflect.TypeOf(*DEQYWBJNXP),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi3", utils.GenComment{
		RouterPath: "/block3",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*TWYPJDMMEJ),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "password",
				ParmType: reflect.TypeOf(*YYXKIBTHNA),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*PAGPNUIZSS),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "year",
				ParmType: reflect.TypeOf(*RWUXYAHLZN),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "commentHi3",
				ParmType: reflect.TypeOf(*TDQQQELAFW),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi3",
				ParmType: reflect.TypeOf(*AKFIMOYJCJ),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi4", utils.GenComment{
		RouterPath: "/block4",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "List",
				ParmType: reflect.TypeOf(*HRXATBWGNO),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*AKNIKSEUSQ),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(*MNAVPJTEVU),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi5", utils.GenComment{
		RouterPath: "/block5",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "reqList",
				ParmType: reflect.TypeOf(new(bind.ReqTest)),
				IsMust:   false,
				ParmKind: reflect.Ptr,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*SUJPGSFDCO),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(*LVODAAWUKQ),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi6", utils.GenComment{
		RouterPath: "/block6",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "List",
				ParmType: reflect.TypeOf(*YGJLTOFLRM),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*PQMTYSUWDD),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi4",
				ParmType: reflect.TypeOf(*SIVHHIONEY),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("Hello.Hi7", utils.GenComment{
		RouterPath: "/block7",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "reqList",
				ParmType: reflect.TypeOf(*AZZBNQQRJC),
				IsMust:   false,
				ParmKind: reflect.Struct,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "index",
				ParmType: reflect.TypeOf(*ZXMUHQVKIF),
				IsMust:   false,
				ParmKind: reflect.Int,
			},

			{
				ParmName: "errHi5",
				ParmType: reflect.TypeOf(*RWJFLJWWOX),
				IsMust:   false,
				ParmKind: reflect.Interface,
			},
		},
	})
	annotation.AddGenOne("UserRest.LogOutUser", utils.GenComment{
		RouterPath: "/LogOutUser",
		Note:       "",
		Methods:    []string{"GET"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*BVGPCEUEBQ),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*CWHVYAVHJV),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*JOWMNGSYGW),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
	annotation.AddGenOne("UserRest.RegistUser", utils.GenComment{
		RouterPath: "/RegistUser",
		Note:       "",
		Methods:    []string{"POST"},
		Parms: []*utils.Parm{

			{
				ParmName: "name",
				ParmType: reflect.TypeOf(*PAWHLPFHLJ),
				IsMust:   false,
				ParmKind: reflect.String,
			},

			{
				ParmName: "age",
				ParmType: reflect.TypeOf(*FZUICEJIYJ),
				IsMust:   false,
				ParmKind: reflect.Int,
			},
		},
		Result: []*utils.Parm{

			{
				ParmName: "success",
				ParmType: reflect.TypeOf(*ERWETHLSJL),
				IsMust:   false,
				ParmKind: reflect.Slice,
			},
		},
	})
}
