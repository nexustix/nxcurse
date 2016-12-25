package nxcurse

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/nexustix/nxduck"
)

//DeobfuscateURL tries to find the destination URL the given URL leads to
//TODO (old code)
func deobfuscateURL(theURL string) string {

	resp, err := http.Get(theURL)
	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}

	defer resp.Body.Close()
	// Request in the Response is the last URL the client tried to access.
	finalURL := resp.Request.URL.String()

	fmt.Printf("info;The URL you ended up at is: %v\n", finalURL)

	//TODO does not make sense for general use but is fine in context
	finalURL = strings.Replace(finalURL, "?cookieTest=1", "", -1)

	//fmt.Printf("info;Guessing files at: %v\n", finalURL+"/files/")
	//curseURLS = append(curseURLS, finalURL+"/files/")

	return finalURL
}

//RationalizeURL tries to trim unrelevant parts from the URL to return a "clean" URL
//TODO (old code)
func rationalizeURL(theURL string) string {
	tmpURL := ""

	theURLParts := strings.Split(theURL, "/")

	for k, v := range theURLParts {
		tmpURL = tmpURL + v + "/"
		if k == 4 {
			break
		}
	}

	argumentizedParts := strings.Split(tmpURL, "?")

	return argumentizedParts[0]
}

// http:
// /
// /minecraft.curseforge.com
// /projects
// /mcheli-minecraft-helicopter-mod?gameCategorySlug=mc-mods&projectID=224304

//TODO (old code)
func isCurseURL(theURL, prefix, postfix string) bool {

	if strings.HasPrefix(theURL, "http://"+prefix+".curseforge.com/"+postfix) {
		return true
	}
	return false
}

//TODO (old code)
func isURLInResult(theSlice []nxduck.SearchResult, theURL string) bool {
	for _, v := range theSlice {
		if v.URL == theURL {
			return true
		}
	}
	return false
}

//TODO (old code)
func eliminateDuplicates(theSlice []nxduck.SearchResult) []nxduck.SearchResult {
	var tmpSlice []nxduck.SearchResult

	for _, v := range theSlice {
		if !isURLInResult(tmpSlice, v.URL) {
			tmpSlice = append(tmpSlice, v)
		}
	}
	return tmpSlice
}
