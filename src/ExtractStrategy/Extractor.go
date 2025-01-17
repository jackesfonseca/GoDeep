package extract

import (
	"../basicdata"
	"path/filepath"
	"fmt"
	"strings"
	"errors"
	"os"
	"log"
)

func (dr *DataReader) SetReadStrategy(rs readStrategy) {
	dr.Strategy = rs
}

func (dr *DataReader) ReadFolder(folderindex int,index int) int{
	
	var files []string
	var name string
	var first bool = true
	var i int
	nametemp := []string{"\"./","\""} //that is a tecnical detail, necessary to use the name that we will get as a path to read

	err := filepath.Walk((*dr).readOrigins[folderindex], visit(&files)) //get all files names at a folder

	if err != nil {
		panic(err)
	} else {
		if (*dr).Readinfo.SizeData == 0{
			(*dr).Readinfo.SizeData = len((*dr).readOrigins)*(len(files)-1)
			
			dr.Strategy.Allocate()
		}
	}

	for _, file := range files {

		if first {	//the first entry is the name of the folder, we don't want that entry
			i = index
			first = false
			continue
		}

		name = strings.Join(nametemp, file)

		if (*dr).Print {
			fmt.Println("geting file:     ", name)
		}

		(*dr).Strategy.ReadData(file,i)
		i++
	}
	return len(files)-1
}

func (dr *DataReader) getFolderName(index int){
	if len((*dr).split) == 0{
		(*dr).split = make([][]string,len((*dr).readOrigins))
	}
	(*dr).split[index] = append(strings.Split((*dr).readOrigins[index], "/"))
}

func (dr *DataReader) SetOrigins(origins []string,rs readStrategy) ([]bool,error){

	err := (*dr).verifyorigins(origins)
	if err != nil {
		return nil,err
	}
	
	(*dr).SetReadStrategy(rs)
	
	var originsIntegrity bool = true
	path := make([][]string,len(origins))
	statusorigins := make([]bool,len(origins))

	for i := 0; i < len(origins); i++ {
		
		path[i] = append(strings.Split(origins[i], "/"))
		statusorigins[i] = (*dr).verifycandidate(path[i])
		
		if originsIntegrity {
			originsIntegrity = statusorigins[i] 
		}
	}

	if originsIntegrity{
		(*dr).readOrigins = origins
		return statusorigins,nil
	}else{
		return statusorigins,errors.New("There was an error to set the origins, path provided is not valid")
	}
}

func (dr *DataReader) Read(	format bool, show bool, print bool) error{
	(*dr).Format = format
	(*dr).Show = show
	(*dr).Print = print
	
	if len((*dr).readOrigins) == 0{
		return errors.New("Origins were not provided, use ReadFloder or define the Origins")
	} else{
		for i := 0; i < len((*dr).readOrigins); i++ {
			(*dr).getFolderName(i)
			(*dr).setLabelbyPath(i)
			
			if i == 0{
				(*dr).Readinfo.Labelsize[i].Size_l = (*dr).ReadFolder(i,i)
				fmt.Println("at zero: ",(*dr).Readinfo.Labelsize[i].Size_l)
			} else{
				(*dr).Readinfo.Labelsize[i].Size_l = (*dr).ReadFolder(i,i*(*dr).Readinfo.Labelsize[i-1].Size_l) //temporary solution
				fmt.Println("at next: ",(*dr).Readinfo.Labelsize[i].Size_l)

			}
		}
		return nil
	}
}

func (dr *DataReader) setLabelbyPath(index int,meaningfulname ...int){
	if len((*dr).Readinfo.Labelsize) == 0{
		(*dr).Readinfo.Labelsize = make([]cartesian.Sizelabel,len((*dr).readOrigins))
	}
	if len(meaningfulname) == 0{
		(*dr).Readinfo.Labelsize[index].Label = (*dr).split[index][len((*dr).split[index])-1]
		fmt.Println("The label that was defined: ", (*dr).Readinfo.Labelsize[index].Label)
	} else{
		fmt.Println("REMEMBER TO IMPLEMENT THAT OPTION")
		(*dr).Readinfo.Labelsize[index].Label = (*dr).split[index][len((*dr).split[index])-1]
		fmt.Println("The label that was defined: ", (*dr).Readinfo.Labelsize[index].Label)
	}
}

func (dr DataReader) verifycandidate(candidate []string) bool{
	if candidate[0] == ".." || candidate[0] == "." { //"../src/imagehandler/Images/danger" or "./Images/grass_1.png"
		return true
	} else{
		return false
	}
}

func (dr *DataReader) verifyorigins(origins []string) error{
	if len(origins) == 0 {
		return errors.New("no Origins provided")
	} else{
		return nil
	}
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("INSIDE THAT PROBLEMATIC IF")
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func FolderLength(folder string) int {
	var files []string

	err := filepath.Walk(folder, visit(&files))

	if err != nil {
		panic(err)
	}
	return (len(files) - 1)
}