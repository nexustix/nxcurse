package nxcurse

import "github.com/nexustix/nxduck"

//GetMinecraftModSearchphrase generates a Searchphrase for a minecraft mod
func GetMinecraftModSearchphrase(modName string) string {
	return GenerateCurseSearchString(modName, "Minecraft", "Files")
}

//GetMinecraftModResults filters results for Mod urls
func GetMinecraftModResults(theResults []nxduck.SearchResult) []nxduck.SearchResult {
	return FilterCurseResults(theResults, "minecraft", "projects")
}
