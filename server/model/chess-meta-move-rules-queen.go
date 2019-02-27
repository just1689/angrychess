package model

import "log"

func isMovePossibleQueen(player *Player, description *MoveDescription) (ok bool, wouldTake *Piece, msg string) {

	//For now moving boards is not possible
	if description.MovingBoards {
		log.Println("Can't move boards")
		ok = false
		return
	}

	//Can't land on your own piece
	if description.LandingOnPieceOwn {
		log.Println("Can't move onto own piece")
		ok = false
		return
	}

	//There is a whole set of rules for placing pieces
	if description.BeingPlaced {
		//TODO: other checks
		// - is empty tile
		// - not in check

		//Can't place on last two
		if description.LastTwoRows {
			log.Println("Can't place on last two enemy lines")
			ok = false
			return
		}

		//Can't place in check
		//if ??? {
		//	log.Println("Can't place in check")
		//	ok = false
		//	return
		//}

		ok = true
		return
	}

	queenMove := (description.XDiff == description.YDiff) || (description.XDiff == 0 && description.YDiff > 0) || (description.YDiff == 0 && description.XDiff > 0)

	if !queenMove {
		log.Println("Queens can only move horizontally, vertically or diagonally")
		ok = false
		return
	}

	if len(description.PiecesBetween) > 0 {
		log.Println("Queens can't move over pieces: ", len(description.PiecesBetween))
		ok = false
		return
	}

	if description.OtherBoard {
		log.Println("Can't move onto another board")
		ok = false
		return
	}

	ok = true
	wouldTake = description.LandingOnPiece

	return

}
