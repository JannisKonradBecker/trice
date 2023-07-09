// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package id

// source tree management

import (
	"fmt"
	"io"

	"github.com/rokath/trice/pkg/ant"
	"github.com/rokath/trice/pkg/msg"
	"github.com/spf13/afero"
)

// idData holds the Id specific data.
type idData struct {
	idToTrice  TriceIDLookUp   // idToFmt is a trice ID lookup map and is generated from existing til.json file at the begin of SubCmdIdInsert. This map is only extended during SubCmdIdInsert and goes back into til.json afterwards.
	triceToId  triceFmtLookUp  // fmtToId is a trice fmt lookup map (reversed idToFmt for faster operation). Each fmt can have several trice IDs (slice). This map is only reduced during SubCmdIdInsert and goes _not_ back into til.json afterwards.
	idToLocRef TriceIDLookUpLI // idToLocRef is the trice ID location information as reference generated from li.json (if exists) at the begin of SubCmdIdInsert and is not modified at all. At the end of SubCmdIdInsert a new li.json is generated from itemToId.
	idToLocNew TriceIDLookUpLI // idToLocNew is the trice ID location information generated during insertTriceIDs. At the end of SubCmdIdInsert a new li.json is generated from idToLocNew.
	//itemToId   TriceItemLookUpID // itemToId is a trice item lookup ID map, extended from source tree during SubCmdIdInsert after each found and maybe modified trice item.
	//idToItem   TriceIDLookupItem // idToItem is a trice ID lookup item map (reversed itemToId for faster operation) and kept in sync with itemToId.
	idCount int       // idCount is the initial used ID count.
	IDSpace []TriceID // IDSpace contains unused IDs.
}

var (
	idd idData
)

func (p *idData) preProcessing(w io.Writer, fSys *afero.Afero) {

	// get state
	p.idToTrice = NewLut(w, fSys, FnJSON)
	p.triceToId = p.idToTrice.reverseS()
	p.idCount = len(p.idToTrice)
	p.idToLocRef = NewLutLI(w, fSys, LIFnJSON) // for reference lookup
	p.idToLocNew = make(TriceIDLookUpLI, 4000) // for new li.json

	// create IDSpace
	p.IDSpace = make([]TriceID, 0, Max-Min+1)
	for id := Min; id <= Max; id++ {
		_, usedFmt := p.idToTrice[id]
		_, usedLoc := p.idToLocRef[id]
		if !usedFmt && !usedLoc {
			p.IDSpace = append(p.IDSpace, id)
		}
		if usedFmt && !usedLoc {
			fmt.Fprintln(w, "ID", id, "only inside til.json")
		}
		if !usedFmt && usedLoc {
			fmt.Fprintln(w, "ID", id, "only inside li.json")
		}
	}

	// prepare
	//p.itemToId = make(TriceItemLookUpID, 4000)
	//p.idToItem = make(TriceIDLookupItem, 4000)
}

func (p *idData) postProcessing(w io.Writer, fSys *afero.Afero) {

	// finalize
	if Verbose {
		fmt.Fprintln(w, len(p.idToTrice), "ID's in List", FnJSON)
	}
	idsAdded := len(p.idToTrice) - p.idCount
	if idsAdded > 0 && !DryRun {
		if Verbose {
			fmt.Fprintln(w, idsAdded, "ID's were added to List", FnJSON)
		}
		msg.FatalOnErr(p.idToTrice.toFile(fSys, FnJSON))
	}
	if LIFnJSON == "off" || LIFnJSON == "none" {
		return
	}
	msg.FatalInfoOnErr(p.idToLocNew.toFile(fSys, LIFnJSON), "could not write LIFnJSON")
}

func cmdManageTriceIDs(w io.Writer, fSys *afero.Afero, action ant.Processing) error {

	// initialize
	a := new(ant.Admin)
	a.Action = action
	if len(Srcs) == 0 {
		a.Trees = append(Srcs, "./") // default value
	} else {
		a.Trees = Srcs
	}
	a.MatchingFileName = isSourceFile

	// process
	idd.preProcessing(w, fSys)
	err := a.Walk(w, fSys)
	idd.postProcessing(w, fSys)

	return err
}
