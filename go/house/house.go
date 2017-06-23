package house

const testVersion = 1

var strReplace = map[string][2]string{
	"is":      {"malt", "lay"},
	"lay":     {"rat", "ate"},
	"ate":     {"cat", "killed"},
	"killed":  {"dog", "worried"},
	"worried": {"cow with the crumpled horn", "tossed"},
	"tossed":  {"maiden all forlorn", "milked"},
	"milked":  {"man all tattered and torn", "kissed"},
	"kissed":  {"priest all shaven and shorn", "married"},
	"married": {"rooster that crowed in the morn", "woke"},
	"woke":    {"the farmer sowing his corn", "kept"},
	"kept":    {"the horse and the hound and the horn", "belonged"},
}

var pre = "This"
var after = "that"

// Song : All house song
func Song() string {

	verb := "is"
	pharse := "house that Jack built"
	song := pre + " " + verb + " " + pharse

	for i := 0; i < 12; i++ {
		// later

	}
	return song
}

// Verse : Verse the song
func Verse() string {

	return ""
}
