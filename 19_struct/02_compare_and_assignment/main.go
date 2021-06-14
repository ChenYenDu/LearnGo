package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	// song1 := song{title: "wonderwall", artist: "oasis"}
	// song2 := song{title: "super sonic", artist: "oasis"}

	// song2 = song1

	// fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	// if song1 == song2 {
	// 	fmt.Println("Songs are equal.")
	// } else {
	// 	fmt.Println("Songs are not equal.")
	// }

	rock := playlist{
		genre: "indie rock",
		songs: []song{
			{title: "wonderwall", artist: "oasis"},
			{title: "super sonic", artist: "oasis"},
		},
	}

	// clone := rock
	// if rock == clone {
	// 		... this doesn't work, because playlist contains a struct field -> which is not compariable
	// }

	song := rock.songs[0]       // this clones a new slice
	song.title = "live forever" // won't change the original value in rock songs

	fmt.Printf("%+v\n%+v\n", rock.songs[0], song)

	rock.songs[0].title = "live forever" // this change the original value in rock songs

	fmt.Printf("\n%-20s %-20s\n", "TITLE", "ARTIST")

	for _, s := range rock.songs {

		// s := rock.songs[i]
		s.title = "destory" // s is a clone of song, so the later loop still prints the original value
		fmt.Printf("%-20s %-20s\n", s.title, s.artist)

	}

	fmt.Printf("\n%-20s %-20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %-20s\n", s.title, s.artist)

	}

}
