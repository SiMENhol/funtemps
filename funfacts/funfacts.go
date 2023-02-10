package funfacts

/*
*

	Implementer funfacts-funksjon:
	  GetFunFacts(about string) []string
	    hvor about kan ha en av tre testverdier, -
	      sun, luna eller terra

	Sett inn alle Funfucts i en struktur
	type FunFacts struct {
	    Sun []string
	    Luna []string
	    Terra []string
	}
*/
func GetFunFacts(about string) []string {
	//	funfacts := []FunFacts{
	//		FunFacts{
	//			Sun: []string{"sol"},
	//		},
	//	}
	//	return (funfacts.FunFacts).Sun
	var funfact FunFacts
	funfact.Sun = []string{"sol"}

	if about == "sun" {
		return funfact.Sun
	}
	return nil
}

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}
