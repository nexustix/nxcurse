package nxcurse

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//GetDependencies gets dependencies of a given mod
//TODO (old code)
func GetDependencies(ModPostfix string) []string {

	baseURL := "http://minecraft.curseforge.com/projects/"
	postURL := "/relations/dependencies?filter-related-dependencies=3"

	theURL := baseURL + ModPostfix + postURL

	var tmpFound string
	var tmpSegments []string
	var foundDependencies []string

	//fmt.Println(theURL)

	doc, err := goquery.NewDocument(theURL)
	if err != nil {
		log.Fatal(err)
	}

	//"name-wrapper overflow-tip"

	sel := doc.Find(".project-list-item")
	//fmt.Printf(">start\n")
	for i := range sel.Nodes {
		//fmt.Printf(">ding\n")
		/*
			    node := sel.Eq(i)
					stuff := node.Find(".name-wrapper overflow-tip")
					dataNode := stuff.Find("a")
					tmpURL, _ := dataNode.Attr("href")
		*/
		node := sel.Eq(i)
		dataNode := node.Find("a")
		tmpURL, _ := dataNode.Attr("href")
		//fmt.Printf(">%s<\n", tmpURL)
		//fmt.Printf(">%s<\n", strings.TrimSpace(dataNode.Text()))
		tmpSegments = strings.Split(tmpURL, "/")
		tmpFound = tmpSegments[len(tmpSegments)-1]
		/*
			fmt.Printf("<-> dependency found: %s\n", tmpFound)
		*/
		foundDependencies = append(foundDependencies, tmpFound)
		//fmt.Printf(">%s<\n", stuff.Text())
	}
	//fmt.Printf(">done\n")
	return foundDependencies

	/*
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
	*/
}

//HACK could cause unreasonable traffic (relative to information retrieved) if used too often
//only retrieving filename from a querry is a bit wasteful (get filename during download instead if you want to save traffic)

//GetFilenameFromDownloadURL retrieves filename from a curse URL (like https://minecraft.curseforge.com/projects/vending-machines-revamped/files/2266299/download")
func GetFilenameFromDownloadURL(downURL string) string {

	var filename string

	tr := &http.Transport{
		//TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://minecraft.curseforge.com/projects/vending-machines-revamped/files/2266299/download")
	//resp, err := client.Get("http://addons.cursecdn.com/files/2266/299/VendingMachinesRevamped_Alpha0.0.1_1.8.jar")
	//resp, err := client.Get("https://addons-origin.cursecdn.com/files/2307/310/immcraft-1.9.4-1.1.7.jar") // does not work
	resp, err := client.Get(downURL)
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
	//bp.FailError(err)

	//fmt.Printf("%#v\n", resp.Request.Header.Get("Referer"))

	referer := resp.Request.Header.Get("Referer")

	segments := strings.Split(referer, "/")

	if len(segments) >= 1 {
		filename = segments[len(segments)-1]
	}

	return filename

}
