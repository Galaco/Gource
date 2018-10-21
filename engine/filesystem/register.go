package filesystem

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/lib/vpk"
	"github.com/galaco/KeyValues"
	"path/filepath"
	"regexp"
	"strings"
)

// Read game resource data paths from gameinfo.txt
// All games should ship with a gameinfo.txt, but it isn't actually mandatory
func RegisterGameResourcePaths(basePath string, gameInfo *keyvalues.KeyValue) {
	searchPaths := gameInfo.FindByKey("GameInfo").FindByKey("FileSystem").FindByKey("SearchPaths")
	basePath,_ = filepath.Abs(basePath)
	basePath = strings.Replace(basePath, "\\", "/", -1)

	for _, searchPath := range *searchPaths.GetAllValues() {
		kv := searchPath.(keyvalues.KeyValue)
		path := (*kv.GetAllValues())[0].(string)

		// Current directory
		gameInfoPathRegex := regexp.MustCompile(`(?i)\|gameinfo_path\|`)
		if gameInfoPathRegex.MatchString(path) {
			path = gameInfoPathRegex.ReplaceAllString(path, basePath+"/")
		}

		// Executable directory
		allSourceEnginePathsRegex := regexp.MustCompile(`(?i)\|all_source_engine_paths\|`)
		if allSourceEnginePathsRegex.MatchString(path) {
			path = allSourceEnginePathsRegex.ReplaceAllString(path, basePath+"/../")
		}
		if strings.Contains(strings.ToLower(*kv.GetKey()), "mod") && !strings.HasPrefix(path, basePath) {
			path = basePath + "/../" + path
		}

		// Strip vpk extension, then load it
		if strings.HasSuffix(path, ".vpk") {
			path = strings.Replace(path, ".vpk", "", 1)
			vpkHandle, err := vpk.OpenVPK(path)
			if err != nil {
				debug.Log(err)
				continue
			}
			AddVpk(vpkHandle)
			debug.Log("Registered vpk: " + path)
		} else {
			// wildcard suffixes not useful
			if strings.HasSuffix(path, "/*") {
				path = strings.Replace(path, "/*", "",  -1)
			}
			AddSearchDirectory(path)
			debug.Log("Registered path: " + path)
		}

	}
}