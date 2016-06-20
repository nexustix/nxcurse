package nxcurse

import (
	"log"
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
