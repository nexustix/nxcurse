package nxcurse

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	bp "github.com/nexustix/boilerplate"
	"github.com/nexustix/nxduck"
	//"github.com/nexustixOLD_1/nxprovider"
)

func GenerateCurseSearchString(modName, gameName, pagename string) string {
	//gameName := "Minecraft"

	//return " t:(- Mods - Projects - " + gameName + " CurseForge) " + modName + " site:curseforge.com"
	//return fmt.Sprintf("t:(%s - %s - Mods - Projects - %s CurseForge) site:curseforge.com", pagename, modName, gameName)
	return fmt.Sprintf("%s site:%s.curseforge.com", modName, gameName)
}

//FilterCurseResults filters curse results from search results
//TODO (old code)
// func FilterCurseResults(theResults []nxduck.SearchResult, prefix string, postfix string) []nxduck.SearchResult {
func FilterCurseResults(theResults []nxduck.SearchResult, prefix string, postfix string) []nxduck.SearchResult {
	var curseResults []nxduck.SearchResult

	//prefix := "minecraft"
	//postfix := "projects"

	for _, v := range theResults {
		//if strings.HasPrefix(v.URL, "http://"+prefix+".curseforge.com/"+postfix) {
		//fmt.Println(v.URL)
		if strings.Contains(v.URL, prefix+".curseforge.com/"+postfix) {
			//fmt.Println(v.URL)
			tmpURL := rationalizeURL(v.URL)

			var tmpSearchResult nxduck.SearchResult

			tmpSearchResult.Title = v.Title
			if !strings.HasSuffix(tmpURL, "/") {
				tmpURL = tmpURL + "/"
			}
			tmpSearchResult.URL = tmpURL

			curseResults = append(curseResults, tmpSearchResult)
			//curseURLS = append(curseURLS, v)
		}
		//curseURLS = append(curseURLS, v)
	}

	curseResults = eliminateDuplicates(curseResults)
	return curseResults
}

// GetCurseFilename gets filename from download url
//TODO (old code)
//XXX deprecated
func GetCurseFilename(infoURL, fileEnding string) string {
	//get webpage html
	//resp, err := http.Get(infoURL)
	//bp.FailError(err)
	//doc, err := goquery.NewDocumentFromResponse(resp)
	doc, err := goquery.NewDocument(infoURL)
	if err != nil {
		log.Fatal(err)
	}

	//find "links_main"
	sel := doc.Find(".details-info")
	for i := range sel.Nodes {
		node := sel.Eq(i)
		stuff := node.Find(".info-data")

		for n := range stuff.Nodes {
			theNode := stuff.Eq(n)

			if strings.Contains(theNode.Text(), fileEnding) {
				//fmt.Printf(">%s<\n", theNode.Text())
				/////safeFilename := strings.Replace(theNode.Text(), " ", "_", -1)
				return theNode.Text()
				/////return safeFilename
			}
		}
	}
	return ""
}

func GetCurseDownloads(curseDownloadURL, version string) []bp.Download {
	//fmt.Printf(">%s<\n", curseDownloadURL)

	baseURL := curseDownloadURL
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}

	//var downloadsFound []nxprovider.ModDownload
	//var preReleaseFound []nxprovider.ModDownload
	var downloadsFound []bp.Download
	var preReleaseFound []bp.Download

	var downloadURL string
	var downloadFileName string
	//var downloadSize int

	doc, err := goquery.NewDocument(curseDownloadURL)
	if err != nil {
		log.Fatal(err)
	}
	sel := doc.Find(".project-file-list-item")
	for i := range sel.Nodes {
		node := sel.Eq(i)

		versionNode := node.Find(".version-label")
		//fmt.Printf("version>%s<<\n", versionNode.Text())
		if !strings.Contains(versionNode.Text(), version) {
			//fmt.Printf("<!>version>%s<\n", versionNode.Text())
			continue
		}
		//fmt.Printf("<->version>%s<\n", versionNode.Text())

		urlNode := node.Find(".overflow-tip")
		tmpDownURL, _ := urlNode.Attr("href")
		tmpSegments := strings.SplitAfter(tmpDownURL, "/files/")
		downloadURL = baseURL + tmpSegments[len(tmpSegments)-1] + "/download"
		//fmt.Printf("<->url>%s<\n", tmpSegments[0])

		//HACK cheap-ish way to get filename without using too much traffic (is prone to errors but works most of the time)
		downloadFileName = urlNode.Text()

		if !strings.HasSuffix(downloadFileName, ".jar") && !strings.HasSuffix(downloadFileName, ".zip") {
			//FIXME use GetCurseFilename instead ?
			downloadFileName = downloadFileName + ".jar"

		}
		//fmt.Printf("<->file>%s<\n", urlNode.Text())

		kindNode := node.Find(".project-file-release-type")
		typeNode := kindNode.Find(".release-phase")
		kind, _ := typeNode.Attr("title")
		//fmt.Printf("<->kind>%s<\n", kind)

		tmpModDownload := bp.Download{}

		tmpModDownload.URL = downloadURL
		tmpModDownload.Filename = downloadFileName

		if kind == "Release" {
			downloadsFound = append(downloadsFound, tmpModDownload)
		} else {
			preReleaseFound = append(preReleaseFound, tmpModDownload)
		}

		//tmpDownURL, _ := urlNode.Text()
		//tmpSegments := strings.SplitAfter(tmpDownURL, "/files/")
		//downloadURL = baseURL + tmpSegments[len(tmpSegments)-1] + "/download"

		/*
			theNode := node.Find(".overflow-tip")
			tmpDownURL, _ := theNode.Attr("href")
			tmpSegments := strings.SplitAfter(tmpDownURL, "/files/")
			downloadURL = baseURL + tmpSegments[len(tmpSegments)-1] + "/download"
		*/

		/*
			theNode := node.Find(".project-file-name")
			tmpSpaceless := strings.TrimSpace(theNode.Text())
			tmpSegments := strings.SplitAfter(tmpSpaceless, ".jar")
			fmt.Printf(">%s<\n", tmpSpaceless)
			fmt.Printf(">>%s<<\n", tmpSegments[0])
			downloadFileName = tmpSegments[0]
		*/
	}
	//fmt.Println("end")

	if len(downloadsFound) < 1 && len(preReleaseFound) >= 1 {
		// if no releases are fond just add the newest alpha/beta
		downloadsFound = append(downloadsFound, preReleaseFound[0])
	}

	return downloadsFound
}
