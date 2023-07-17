// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.
package id

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tj/assert"
)

// Only the first occurrence of a double used ID gets in the list.
// Only the list generation out of sources ist tested here.
func TestRefreshIDListSingle1(t *testing.T) {
	text := `
	TRICE16_3( Id(12345), "hi %2d, %13u, %64b\n",1,2,3 );	
	TRICE16_3( Id(12345), "DIFFERENT! %2d, %13u, %64b\n",1,2,3 );	
`
	expJSON := `{
	"12345": {
		"Type": "TRICE16_3",
		"Strg": "hi %2d, %13u, %64b\\n"
	}
}`
	check(t, text, expJSON)
}

// Only the first occurrence of a double used ID gets in the list.
// Only the list generation out of sources ist tested here.
func TestRefreshIDListSingle0(t *testing.T) {
	text := `
	TRICE16_3( Id(12345), "hi %2d, %13u, %64b\n", 1, 2, 3 );
	TRICE16_3( Id(12345), "DIFFERENT! %2d, %13u, %64b\n", 1, 2, 3 );
	Trice16_1( iD(12344), "hi %2d\n", 1 );
	Trice16_1( iD(12344), "DIFFERENT! %2d\n", 2 );
`
	expJSON := `{
	"12344": {
		"Type": "Trice16_1",
		"Strg": "hi %2d\\n"
	},
	"12345": {
		"Type": "TRICE16_3",
		"Strg": "hi %2d, %13u, %64b\\n"
	}
}`
	check(t, text, expJSON)
}

func TestRefreshIDListSingle1WithNewline1(t *testing.T) {
	var text string
	fm := `hi %2d, %13u, %64b\n`
	text += `...   TRICE( Id(1234), "` + fm + `",1,2,3 ); ...
	`
	text += `...   TRICE16_3( Id(12345), "` + fm + `",1,2,3 ); ...
	`
	text += `...   TRICE16_3( Id(  123), "` + fm + `",1,2,3 ); ...
	`
	text += `...   Trice16_1( "` + fm + `", 3 );               ...
	`

	el := make(TriceIDLookUp)
	el[12345] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	el[1234] = TriceFmt{Type: "TRICE", Strg: fm}
	el[123] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	checkTil(t, text, el)
}

func TestRefreshIDListSingle1WithNewline0(t *testing.T) {
	var text string
	fm := `hi %2d, %13u, %64b\n`
	text += `...   TRICE( Id(1234), "` + fm + `",1,2,3 ); ...
	`
	text += `...   TRICE16_3( Id(12345), "` + fm + `",1,2,3 ); ...
	`
	text += `...   TRICE16_3( Id(  123), "` + fm + `",1,2,3 ); ...
	`
	text += `...   Trice16_1( iD(   13), "` + fm + `", 3 );               ...
	`
	el := make(TriceIDLookUp)
	el[12345] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	el[1234] = TriceFmt{Type: "TRICE", Strg: fm}
	el[123] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	el[13] = TriceFmt{Type: "Trice16_1", Strg: fm}
	checkTil(t, text, el)
}

func TestRefreshIDListSingle1WithoutNewline1(t *testing.T) {
	var text string
	fm := `hi %2d, %13u, %64b\n`
	text += `...   TRICE( Id(1234), "` + fm + `",1,2,3 ); ...`
	text += `...   TRICE16_3( Id(12345), "` + fm + `",1,2,3 ); ...`
	text += `...   TRICE16_3( Id(  123), "` + fm + `",1,2,3 ); ...`
	text += `...   Trice16_1(  "` + fm + `", 3 );               ...`

	el := make(TriceIDLookUp)
	el[12345] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	el[1234] = TriceFmt{Type: "TRICE", Strg: fm}
	el[123] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	checkTil(t, text, el)
}

func TestRefreshIDListSingle1WithoutNewline0(t *testing.T) {
	var text string
	fm := `hi %2d, %13u, %64b\n`
	text += `...   TRICE( Id(1234), "` + fm + `",1,2,3 ); ...`
	text += `...   TRICE16_3( Id(12345), "` + fm + `",1,2,3 ); ...`
	text += `...   TRICE16_3( Id(  123), "` + fm + `",1,2,3 ); ...`
	text += `...   Trice16_1( iD(   13), "` + fm + `", 3 );               ...`

	el := make(TriceIDLookUp)
	el[12345] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	el[1234] = TriceFmt{Type: "TRICE", Strg: fm}
	el[123] = TriceFmt{Type: "TRICE16_3", Strg: fm}
	el[13] = TriceFmt{Type: "Trice16_1", Strg: fm}
	checkTil(t, text, el)
}

func TestRefreshIDListSingle2(t *testing.T) {
	text := `
	TRICE16_3( Id(12345), "hi %2d, %13u, %64b\n",1,2,3 );
	trice16_3( iD(12345), "hi %2d, %13u, %64b\n",1,2,3 );
	Trice16_1( iD(12344), "ho %64b\n",1 );
`
	expJSON := `{
	"12344": {
		"Type": "Trice16_1",
		"Strg": "ho %64b\\n"
	},
	"12345": {
		"Type": "TRICE16_3",
		"Strg": "hi %2d, %13u, %64b\\n"
	}
}`
	check(t, text, expJSON)
}

func TestTriceFmtParse(t *testing.T) {
	for i := range tryOkSet {
		tf, ok := triceFmtParse(tryOkSet[i].nbTRICE)
		assert.True(t, ok == tryOkSet[i].ok)
		assert.True(t, tf == tryOkSet[i].tf, fmt.Sprint(tf, tryOkSet[i].tf))
	}
}

func TestInsertSharedIDs0WithParamCount(t *testing.T) {
	tt := testTable{
		{
			`TRICE(Id(0), "Go is fun");`,
			`TRICE(Id(10000), "Go is fun");`, true, true},
		{
			`trice8_1( Id(0), "hi %d", 5); // first id`,
			`trice8_1( Id(10001), "hi %d", 5); // first id`, true, true},
		{
			`trice8_1( Id(0), "hi %d", 5); // different format string needs a new id`,
			`trice8_1( Id(10002), "hi %d", 5); // different format string needs a new id`, true, true},
		{
			`TRICE8_1( Id(       0   ), "hi %d", 5); // different type case gets same id`,
			`TRICE8_1( Id(10003), "hi %d", 5); // different type case gets same id`, true, true},
		{
			`TRICE8_1( Id(0), "hi %d", 5); // same format string gets same id`,
			`TRICE8_1( Id(10004), "hi %d", 5); // same format string gets same id`, true, true},
		{
			`trice8_1(  Id( 0       ),  "hi %d", 5); // same format string gets same id`,
			`trice8_1(  Id(10005),  "hi %d", 5); // same format string gets same id`, true, true},
		{
			`Trice8_1( Id(0), "hi %d", 5); // first id`,
			`Trice8_1( Id(10006), "hi %d", 5); // first id`, true, true},
		{
			`Trice8_2( Id(0), "hi %d %u", 5, 7); // first id`,
			`Trice8_2( Id(10007), "hi %d %u", 5, 7); // first id`, true, true},
		{
			`Trice16_1( Id(0), "hi %d", 5); // first id`,
			`Trice16_1( Id(10008), "hi %d", 5); // first id`, true, true},
	}

	im := make(TriceIDLookUp)
	em := make(TriceIDLookUp)
	em[10000] = TriceFmt{Type: "TRICE", Strg: "Go is fun"}
	em[10001] = TriceFmt{Type: "trice8_1", Strg: "hi %d"}
	em[10002] = TriceFmt{Type: "trice8_1", Strg: "hi %d"}
	em[10003] = TriceFmt{Type: "TRICE8_1", Strg: "hi %d"}
	em[10004] = TriceFmt{Type: "TRICE8_1", Strg: "hi %d"}
	em[10005] = TriceFmt{Type: "trice8_1", Strg: "hi %d"}
	em[10006] = TriceFmt{Type: "Trice8_1", Strg: "hi %d"}
	em[10007] = TriceFmt{Type: "Trice8_2", Strg: "hi %d %u"}
	em[10008] = TriceFmt{Type: "Trice16_1", Strg: "hi %d"}
	checkList3(t, false, 10000, 10099, "upward", tt, false, im, em)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// todo

//  func TestInsertSharedIDsInvalid0(t *testing.T) {
//  	var listModified bool
//  	var fileModified bool
//  	tt := testTable{
//  		{
//  			`trice8_1( Id(0x0), "Hi %d", 5); // this gets not changed because no decimal number`,
//  			`trice8_1( Id(0x0), "Hi %d", 5); // this gets not changed because no decimal number`, fileModified, listModified},
//  		{
//  			`trice8_1( Id(-0), "hi %d", 5); // minus id's are not touched`,
//  			`trice8_1( Id(-0), "hi %d", 5); // minus id's are not touched`, fileModified, listModified},
//  	}
//  	eList := `map[]
//  `
//  	checkList(t, SharedIDs, 10000, 20000, "downward", tt, eList, true)
//  	checkList(t, SharedIDs, 10000, 20000, "downward", tt, eList, false)
//  }

//  func TestTriceIDParseOK(t *testing.T) {
//  	set := tryOkSet
//  	for i := range set {
//  		nbID, id, ok := triceIDParse(set[i].nbTRICE)
//  		assert.True(t, ok == set[i].ok)
//  		if ok {
//  			checkID(t, set, i, id)
//  			checkNbID(t, set, i, nbID)
//  		}
//  	}
//  }

//  func TestTriceParseOK(t *testing.T) {
//  	set := tryOkSet
//  	for i := range set {
//  		nbID, id, tf, ok := triceParse(set[i].nbTRICE)
//  		assert.True(t, ok == set[i].ok)
//  		if ok {
//  			checkID(t, set, i, id)
//  			checkNbID(t, set, i, nbID)
//  			assert.True(t, tf == tryOkSet[i].tf, fmt.Sprint(tf, tryOkSet[i].tf))
//  		}
//  	}
//  }

//  func TestTriceIDParseNotOK(t *testing.T) {
//  	var tryNotOkSetID = []idCheck{
//  		{`TRICE0(Id(0x5), "tt" )`, "Id(0x5)", 5, false, TriceFmt{"TRICE0", "tt"}},
//  		{`TRICE0(id(0x5 ), "tt" )`, "Id(0x59)", 0x59, false, TriceFmt{"TRICE0", "tt"}},
//  	}
//  	set := tryNotOkSetID
//  	for i := range set {
//  		nbID, id, ok := triceIDParse(set[i].nbTRICE)
//  		assert.True(t, ok == set[i].ok, fmt.Sprint(i))
//  		if ok {
//  			checkID(t, set, i, id)
//  			checkNbID(t, set, i, nbID)
//  		}
//  	}
//  }

// no more shared
func _TestInsertSharedIDs0ZeroParam1(t *testing.T) {
	tt := testTable{
		{`... TRICE0( "hi"); ...`, `... TRICE0( Id(   99), "hi"); ...`, true, true},
		{`... TRICE0( "hi"); ...`, `... TRICE0( Id(   98), "hi"); ...`, true, true},
	}
	eList := `map[98:{TRICE0 hi} 99:{TRICE0 hi}]
`
	checkList2(t, false, 10, 99, "downward", tt, true, "", eList)
}

// no more shared
func _TestInsertSharedIDs0ZeroParam2(t *testing.T) {
	tt := testTable{
		{`... TRICE( "hi"); ...`, `... TRICE( Id(   99), "hi"); ...`, true, true},
		{`... TRICE( "hi"); ...`, `... TRICE( Id(   98), "hi"); ...`, true, true},
	}
	eList := `map[98:{TRICE hi} 99:{TRICE hi}]
`
	checkList2(t, false, 10, 99, "downward", tt, true, "", eList)
}

// no more shared
func _TestInsertSharedIDs0ZeroParam3(t *testing.T) {
	tt := testTable{
		{`... TRICE( "hi %d", 7);       ...`, `... TRICE( Id(   99), "hi %d", 7);       ...`, true, true},
		{`... TRICE( "hi %u %b", 6, 6); ...`, `... TRICE( Id(   98), "hi %u %b", 6, 6); ...`, true, true},
	}
	eList := `map[98:{TRICE_2 hi %u %b} 99:{TRICE_1 hi %d}]
`
	checkList2(t, false, 10, 99, "downward", tt, false, "", eList)
}

// no more shared
func _TestInsertSharedIDs0ZeroParam3032(t *testing.T) {
	tt := testTable{
		{`... TRICE32( "hi %d", 7);       ...`, `... TRICE32( Id(   99), "hi %d", 7);       ...`, true, true},
		{`... TRICE32( "hi %u %b", 6, 6); ...`, `... TRICE32( Id(   98), "hi %u %b", 6, 6); ...`, true, true},
	}
	eList := `map[98:{TRICE32_2 hi %u %b} 99:{TRICE32_1 hi %d}]
`
	checkList2(t, false, 10, 99, "downward", tt, false, "", eList)
}

// no more shared
func _TestInsertSharedIDs0ZeroParam332(t *testing.T) {
	tt := testTable{
		{`... TRICE32( "hi %d", 7);       ...`, `... TRICE32_1( Id(   99), "hi %d", 7);       ...`, true, true},
		{`... TRICE32( "hi %u %b", 6, 6); ...`, `... TRICE32_2( Id(   98), "hi %u %b", 6, 6); ...`, true, true},
	}
	eList := `map[98:{TRICE32_2 hi %u %b} 99:{TRICE32_1 hi %d}]
`
	checkList2(t, false, 10, 99, "downward", tt, true, "", eList)
}

func _TestInsertSharedIDs0ZeroParam4noExtend(t *testing.T) {
	tt := testTable{
		{`... TRICE( "hi %d", 7);       ...`, `... TRICE( Id(   99), "hi %d", 7);       ...`, true, false},
		{`... TRICE( "hi %u %b", 6, 6); ...`, `... TRICE( Id(   98), "hi %u %b", 6, 6); ...`, true, false},
	}
	il := make(TriceIDLookUp)
	il[99] = TriceFmt{Type: "TRICE", Strg: "hi %d"}
	il[98] = TriceFmt{Type: "TRICE", Strg: "hi %u %b"}
	checkList3(t, true, 10, 99, "downward", tt, false, il, il)
}

func _TestInsertSharedIDs0ZeroParam4extend(t *testing.T) {
	tt := testTable{
		{`... TRICE( "hi %d", 7);       ...`, `... TRICE_1( Id(   99), "hi %d", 7);       ...`, true, false},
		{`... TRICE( "hi %u %b", 6, 6); ...`, `... TRICE_2( Id(   98), "hi %u %b", 6, 6); ...`, true, false},
	}
	il := make(TriceIDLookUp)
	il[99] = TriceFmt{Type: "TRICE_1", Strg: "hi %d"}
	il[98] = TriceFmt{Type: "TRICE_2", Strg: "hi %u %b"}
	checkList3(t, true, 10, 99, "downward", tt, true, il, il)
}

func _TestInsertSharedIDs0ZeroParam16noExtend(t *testing.T) {
	tt := testTable{
		{`... TRICE16( "hi %d", 7);       ...`, `... TRICE16( Id(   99), "hi %d", 7);       ...`, true, false},
		{`... TRICE16( "hi %u %b", 6, 6); ...`, `... TRICE16( Id(   98), "hi %u %b", 6, 6); ...`, true, false},
	}
	il := make(TriceIDLookUp)
	il[99] = TriceFmt{Type: "TRICE16", Strg: "hi %d"}
	il[98] = TriceFmt{Type: "TRICE16", Strg: "hi %u %b"}
	checkList3(t, true, 10, 99, "downward", tt, false, il, il)
}

func _TestInsertSharedIDs0ZeroParam16extend(t *testing.T) {
	tt := testTable{
		{`... TRICE16( "hi %d", 7);       ...`, `... TRICE16_1( Id(   99), "hi %d", 7);       ...`, true, false},
		{`... TRICE16( "hi %u %b", 6, 6); ...`, `... TRICE16_2( Id(   98), "hi %u %b", 6, 6); ...`, true, false},
	}
	il := make(TriceIDLookUp)
	il[99] = TriceFmt{Type: "TRICE16_1", Strg: "hi %d"}
	il[98] = TriceFmt{Type: "TRICE16_2", Strg: "hi %u %b"}
	checkList3(t, true, 10, 99, "downward", tt, true, il, il)
}

func _TestInsertSharedIDs0ZeroParam(t *testing.T) {
	im := make(TriceIDLookUp)
	im[99] = TriceFmt{Type: "TRICE0", Strg: "hi"}
	im[98] = TriceFmt{Type: "TRICE0", Strg: "hi"}

	// When at least 2 IDs in til with same information, it is not determined, which ID is used again.
	// Therefore, all cases must be checked.
	tt := testTable{ // case one
		{`... TRICE0 ( "hi"); ...`, `... TRICE0 ( Id(   99), "hi"); ...`, true, false},
		{`... TRICE0( "hi");  ...`, `... TRICE0( Id(   99), "hi");  ...`, true, false},
		{`... trice0 ( "hi"); ...`, `... trice0 ( Id(   99), "hi"); ...`, true, false},
		{`... trice0( "hi");  ...`, `... trice0( Id(   99), "hi");  ...`, true, false},
		{`... Trice0 ( "hi"); ...`, `... Trice0 ( Id(   99), "hi"); ...`, true, false},
		{`... Trice0( "hi");  ...`, `... Trice0( Id(   99), "hi");  ...`, true, false},
	}
	om := checkList4(t, true, 10, 99, "downward", tt, false, im)
	eq := reflect.DeepEqual(om, im)
	if eq {
		return
	}

	// next try
	tt = testTable{ // case two
		{`... TRICE0 ( "hi"); ...`, `... TRICE0 ( Id(   98), "hi"); ...`, true, false},
		{`... TRICE0( "hi");  ...`, `... TRICE0( Id(   98), "hi");  ...`, true, false},
		{`... trice0 ( "hi"); ...`, `... trice0 ( Id(   98), "hi"); ...`, true, false},
		{`... trice0( "hi");  ...`, `... trice0( Id(   98), "hi");  ...`, true, false},
		{`... Trice0 ( "hi"); ...`, `... Trice0 ( Id(   98), "hi"); ...`, true, false},
		{`... Trice0( "hi");  ...`, `... Trice0( Id(   98), "hi");  ...`, true, false},
	}
	om = checkList4(t, true, 10, 99, "downward", tt, false, im)
	assert.True(t, reflect.DeepEqual(om, im))
}

func _TestInsertSharedIDs1WithExtendN(t *testing.T) {
	tt := testTable{
		{`...  Trice8 ( "hi %03u", 5); ...`, `...  Trice8_1 ( Id(10000), "hi %03u", 5); ...`, true, false},
		{`...  TRICE8 ( "hi %03u", 5); ...`, `...  TRICE8_1 ( Id(10000), "hi %03u", 5); ...`, true, false},
		{`...  trice8 ( "hi %03u", 5); ...`, `...  trice8_1 ( Id(10000), "hi %03u", 5); ...`, true, false},

		{`... Trice16 ( "hi %03u", 5); ...`, `... Trice16_1 ( Id(10001), "hi %03u", 5); ...`, true, false},
		{`... TRICE16 ( "hi %03u", 5); ...`, `... TRICE16_1 ( Id(10001), "hi %03u", 5); ...`, true, false},
		{`... trice16 ( "hi %03u", 5); ...`, `... trice16_1 ( Id(10001), "hi %03u", 5); ...`, true, false},

		{`... Trice32 ( "hi %03u", 5); ...`, `... Trice32_1 ( Id(10002), "hi %03u", 5); ...`, true, false},
		{`... TRICE32 ( "hi %03u", 5); ...`, `... TRICE32_1 ( Id(10002), "hi %03u", 5); ...`, true, false},
		{`... trice32 ( "hi %03u", 5); ...`, `... trice32_1 ( Id(10002), "hi %03u", 5); ...`, true, false},

		{`... Trice64 ( "hi %03u", 5); ...`, `... Trice64_1 ( Id(10003), "hi %03u", 5); ...`, true, false},
		{`... TRICE64 ( "hi %03u", 5); ...`, `... TRICE64_1 ( Id(10003), "hi %03u", 5); ...`, true, false},
		{`... trice64 ( "hi %03u", 5); ...`, `... trice64_1 ( Id(10003), "hi %03u", 5); ...`, true, false},
	}
	im := make(TriceIDLookUp)
	im[10000] = TriceFmt{Type: "TRICE8_1", Strg: "hi %03u"}
	im[10001] = TriceFmt{Type: "TRICE16_1", Strg: "hi %03u"}
	im[10002] = TriceFmt{Type: "TRICE32_1", Strg: "hi %03u"}
	im[10003] = TriceFmt{Type: "TRICE64_1", Strg: "hi %03u"}
	em := im
	checkList3(t, true, 10, 99, "upward", tt, true, im, em)
}

// The trice map distinguishes between TRICE8 and TRICE8_2 for example, so even "sameID" is selected,
// there are 2 different IDs used when the format string is identical.
func _TestInsertSharedIDs2NoExtendN(t *testing.T) {
	tt := []struct {
		text, exp        string
		fileMod, listMod bool
	}{
		{`...  TRICE8 ( "hi %03u, %5x", 5, 7); ...`, `...  TRICE8 ( Id(10000), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`...  trice8 ( "hi %03u, %5x", 5, 7); ...`, `...  trice8 ( Id(10000), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`...  Trice8 ( "hi %03u, %5x", 5, 7); ...`, `...  Trice8 ( Id(10000), "hi %03u, %5x", 5, 7); ...`, true, false},

		{`... Trice16 ( "hi %03u, %5x", 5, 7); ...`, `... Trice16 ( Id(10001), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`... TRICE16 ( "hi %03u, %5x", 5, 7); ...`, `... TRICE16 ( Id(10001), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`... trice16 ( "hi %03u, %5x", 5, 7); ...`, `... trice16 ( Id(10001), "hi %03u, %5x", 5, 7); ...`, true, false},

		{`... Trice32 ( "hi %03u, %5x", 5, 7); ...`, `... Trice32 ( Id(10002), "hi %03u, %5x", 5, 7); ...`, true, false}, // does not exist but allowed matching
		{`... TRICE32 ( "hi %03u, %5x", 5, 7); ...`, `... TRICE32 ( Id(10002), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`... trice32 ( "hi %03u, %5x", 5, 7); ...`, `... trice32 ( Id(10002), "hi %03u, %5x", 5, 7); ...`, true, false},

		{`... Trice64 ( "hi %03u, %5x", 5, 7); ...`, `... Trice64 ( Id(10003), "hi %03u, %5x", 5, 7); ...`, true, false}, // does not exist but allowed matching
		{`... TRICE64 ( "hi %03u, %5x", 5, 7); ...`, `... TRICE64 ( Id(10003), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`... trice64 ( "hi %03u, %5x", 5, 7); ...`, `... trice64 ( Id(10003), "hi %03u, %5x", 5, 7); ...`, true, false},

		{`... Trice( "hi %03u, %5x", 5, 7); ...`, `... Trice( Id(10004), "hi %03u, %5x", 5, 7); ...`, true, false}, // does not exist but allowed matching
		{`... TRICE( "hi %03u, %5x", 5, 7); ...`, `... TRICE( Id(10004), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`... trice( "hi %03u, %5x", 5, 7); ...`, `... trice( Id(10004), "hi %03u, %5x", 5, 7); ...`, true, false},

		{`...  Trice8_2 ( "hi %03u, %5x", 5, 7); ...`, `...  Trice8_2 ( Id(10005), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`...  Trice8_2 ( "hi %03u, %5x", 5, 7); ...`, `...  Trice8_2 ( Id(10005), "hi %03u, %5x", 5, 7); ...`, true, false},
		{`...  Trice8_2 ( Id(10005), "hi %03u, %5x", 5, 7); ...`, `...  Trice8_2 ( Id(10005), "hi %03u, %5x", 5, 7); ...`, false, false},
	}

	im := make(TriceIDLookUp)
	im[10000] = TriceFmt{Type: "TRICE8", Strg: "hi %03u, %5x"}
	im[10001] = TriceFmt{Type: "TRICE16", Strg: "hi %03u, %5x"}
	im[10002] = TriceFmt{Type: "TRICE32", Strg: "hi %03u, %5x"}
	im[10003] = TriceFmt{Type: "TRICE64", Strg: "hi %03u, %5x"}
	im[10004] = TriceFmt{Type: "TRICE", Strg: "hi %03u, %5x"}
	im[10005] = TriceFmt{Type: "TRICE8_2", Strg: "hi %03u, %5x"}
	em := im
	checkList3(t, true, 10000, 10099, "upward", tt, false, im, em)
}

func _TestTrice0(t *testing.T) { // wip
	tt := testTable{
		{`TRICE( "Go is fun");`, `TRICE( Id(10000), "Go is fun");`, true, true},
		{`TRICE( "Go is fun");`, `TRICE( Id(10000), "Go is fun");`, true, false},
	}
	im := make(TriceIDLookUp)
	em := make(TriceIDLookUp)
	em[10000] = TriceFmt{Type: "TRICE", Strg: "Go is fun"}
	checkList3(t, true, 10000, 10099, "upward", tt, false, im, em)
}

func _TestTrice1(t *testing.T) { // wip
	tt := testTable{
		{`TRICE( "Go is fun");`, `TRICE( Id(10000), "Go is fun");`, true, true},
		{`TRICE( "Go is fun");`, `TRICE( Id(10001), "Go is fun");`, true, true},
	}
	im := make(TriceIDLookUp)
	em := make(TriceIDLookUp)
	em[10000] = TriceFmt{Type: "TRICE", Strg: "Go is fun"}
	em[10001] = TriceFmt{Type: "TRICE", Strg: "Go is fun"}
	om := checkList4(t, false, 10000, 10099, "upward", tt, false, im)
	assert.True(t, reflect.DeepEqual(om, em))
}
