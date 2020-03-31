package main

import (
 "fmt" 
 "path/filepath"
 "os"
 "strings" 
)

var paths = [...]string{"C:/prosjekt/largeScale" , "C:/prosjekt/smallscale"}


type project struct {
	projName string
	nugets string
	npms string	
}


var outfile = "audit.bat"

func main() {
	projMap := make(map[string]project)

	for path := range paths {
		fmt.Printf("looking through %s\n", paths[path])

		err := filepath.Walk(paths[path], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(info.Name(), ".csproj") {
				if val, inmap := projMap[path];  inmap {
					val.projName = info.Name()					
				} else {
					projMap[path] = project { projName: info.Name() }					
				}				
			} else if strings.ToLower(info.Name()) == "package.json" {
				if val, inmap := projMap[path];  inmap {
					val.npms = info.Name()				
				} else {
					projMap[path] = project { npms: info.Name() }										
				}				
			} else if strings.ToLower(info.Name()) == "packages.config" {
				if val, inmap := projMap[path];  inmap {
					val.nugets = info.Name()				
				} else {
					projMap[path] = project { nugets: info.Name() }										
				}				
			}

			return nil
		})

		fmt.Printf("Found %d projects ######################\n", len(projMap));

		//for _, p := range(projMap){
		for key, p := range(projMap){
			fmt.Printf("Project: %s, nuget: %s, npms: %s in path: %s\n", p.projName, p.nugets, p.npms, key)
		}

		
		if err != nil {
			fmt.Println("Error", err, paths[path])
		}		
	}

	fmt.Println("finished")		
}
