package model

import (
	"log"
)

func IsMovePossible(game *Game, player *Player, piece *Piece, description *MoveDescription) (ok bool, taken *Piece, msg string) {
	if piece.Identity == identityPawn {
		return isMovePossiblePawn(game, player, piece, description)
	}
	ok = false
	msg = "Unknown identity. Cannot check isMovePossible"
	log.Println(msg)
	return
}

func isMovePossiblePawn(game *Game, player *Player, piece *Piece, description *MoveDescription) (ok bool, wouldTake *Piece, msg string) {

	//For now moving boards is not possible
	if description.MovingBoards {
		log.Println("Can't move boards")
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

	if description.XDiff > 1 {
		log.Println("X diff can only be 1 or 0 ")
		ok = false
		return
	}

	if description.YDiff > 2 {
		log.Println("Must only move 0 or 1 blocks")
		ok = false
		return
	}

	if description.YDiff == 0 && description.XDiff == 1 {
		log.Println("Can't just move one tile left or right.")
		ok = false
		return
	}

	if description.YDiff == 2 && description.XDiff == 1 {
		log.Println("Can't just move 2F and one tile left or right.")
		ok = false
		return
	}

	if player.shouldGoDown() != description.Down {
		log.Println("Expected goingDown ", description.Down, " equal to should go down ", player.shouldGoDown())
		ok = false
		return
	}

	if description.YDiff == 2 && !description.PawnOnSpawn {
		log.Println("Can't move two when not on start row")
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
