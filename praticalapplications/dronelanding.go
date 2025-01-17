package main

import (
	"../src/imageprocessing"
	"../src/generalizecartesian"
	"../src/nonparametric"
	"../src/basicdata"
	"gocv.io/x/gocv"
	"fmt"
	//"math"
)

func main() {
	
	var dataset generalizecartesian.Labelfeatures
	
	/*size of train and know groups*/
	var size int
	var knowsize int
	var trainsize int
	
	/*normalize flags*/
	var normtype gocv.NormType = gocv.NormMinMax

	/*calc sizes*/
	size  = imageprocessing.FolderLength("../src/imageprocessing/Images/danger")
	trainsize = 25//int(size/2.5)
	knowsize = size - trainsize

	/*set labelsizes*/
	knowls := make([]cartesian.Sizelabel,3)
	trainls := make([]cartesian.Sizelabel,3)

	for i := 0; i < 3; i++ {
		knowls[i].Size_l  = knowsize
		trainls[i].Size_l = trainsize	
	}

	knowls[0].Label  = "danger"
	trainls[0].Label = "danger"

	knowls[1].Label  = "asphalt"
	trainls[1].Label = "asphalt"

	knowls[2].Label  = "grass"
	trainls[2].Label = "grass"

	/* Know images and features allocation*/
	knowImages 			:= make([]gocv.Mat,3*knowsize)	// 	images
	
	knowGLCMs 			:= make([]gocv.Mat,3*knowsize)	// 	GLCMs
	normalizedknow	 	:= make([]gocv.Mat,3*knowsize)	// 	normalizedGLCMs
	/*Know gclm and normalized glcm internal allocation*/
	for i := 0; i < 3*knowsize; i++ {
		knowGLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedknow[i]		= gocv.NewMat()
	}
	/*Know Features*/
	knowEnergys			:= make([]float64,3*knowsize)	// 	Energy
	knowCorrelations	:= make([]float64,3*knowsize)	// 	Correlation
	knowContrasts		:= make([]float64,3*knowsize)	// 	Contrast

	/* Train images and features allocation*/
	trainImages 		:= make([]gocv.Mat,3*trainsize)	// 	images
	
	trainGLCMs 			:= make([]gocv.Mat,3*trainsize)	// 	GLCMs
	normalizedtrain		:= make([]gocv.Mat,3*trainsize)	// 	normalizedGLCMs
	/*Train gclm and normalized glcm internal allocation*/
	for i := 0; i < 3*trainsize; i++ {
		trainGLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedtrain[i]		= gocv.NewMat()
	}
	/*Train Features*/
	trainEnergys		:= make([]float64,3*trainsize)	// 	Energy
	trainCorrelations	:= make([]float64,3*trainsize)	// 	Correlation
	trainContrasts		:= make([]float64,3*trainsize)	// 	Contrast	
	
	/*temporary set of images that will be used to read each folder*/
	auxImages 			:= make([]gocv.Mat,size)

	/*read and separe each group of images*/
	fmt.Println("Reading danger folder")
	imageprocessing.ReadFolder(auxImages,"../src/imageprocessing/Images/danger",false,true,false)
	
	for i := 0; i < size; i++ {
		if i < trainsize{
			trainImages[i] = auxImages[i]
		} else{
			knowImages[i-trainsize] = auxImages[i]
		}
	}
	
	fmt.Println("Reading asphalt folder")
	imageprocessing.ReadFolder(auxImages,"../src/imageprocessing/Images/asphalt",false,true,false)
	for i := 0; i < size; i++ {
		if i < trainsize{
			trainImages[i+trainsize] = auxImages[i]
		} else{
			knowImages[i+(knowsize-trainsize)] = auxImages[i]
		}
	}
	
	fmt.Println("Reading grass folder")
	imageprocessing.ReadFolder(auxImages,"../src/imageprocessing/Images/grass",false,true,false)
	for i := 0; i < size; i++ {
		if i < trainsize{
			trainImages[i+(2*trainsize)] = auxImages[i]
		} else{
			knowImages[i+((2*knowsize)-trainsize)] = auxImages[i]
		}
	}	

	/*compute GLCMs and them the normalized GLCM*/
	fmt.Println("Computing know GLCMs")
	imageprocessing.GroupGLCM(knowImages, &knowGLCMs, false, true)
	for i := 0; i < 3*knowsize; i++ {
		gocv.Normalize(knowGLCMs[i], &normalizedknow[i], 0.0, 255.0, normtype )		
	}

	fmt.Println("Computing train GLCMs")
	imageprocessing.GroupGLCM(trainImages, &trainGLCMs, false, true)
	for i := 0; i < 3*trainsize; i++ {
		gocv.Normalize(trainGLCMs[i], &normalizedtrain[i], 0.0, 255.0, normtype )

	}

	/*Extract the features*/
	fmt.Println("Computing know features")
	imageprocessing.GroupFeature(&normalizedknow,knowEnergys,imageprocessing.EnergyFeature, false)
	imageprocessing.GroupFeature(&normalizedknow,knowCorrelations,imageprocessing.CorrelationFeature, false)
	imageprocessing.GroupFeature(&normalizedknow,knowContrasts,imageprocessing.ContrastFeature, false)

	fmt.Println("Computing train features")
	imageprocessing.GroupFeature(&normalizedtrain,trainEnergys,imageprocessing.EnergyFeature, false)
	imageprocessing.GroupFeature(&normalizedtrain,trainCorrelations,imageprocessing.CorrelationFeature, false)
	imageprocessing.GroupFeature(&normalizedtrain,trainContrasts,imageprocessing.ContrastFeature, false)


	fmt.Println("Generalizing know data set")
	generalizecartesian.Generalize_for_nonparametric(&dataset, knowEnergys, knowCorrelations, knowContrasts,knowls,generalizecartesian.Knowflag,3*knowsize)
	
	fmt.Println("Generalizing train data set")
	generalizecartesian.Generalize_for_nonparametric(&dataset, trainEnergys, trainCorrelations, trainContrasts,trainls,generalizecartesian.Trainflag,3*trainsize)

	// fmt.Println("Computing centroid")
	// dataset.Centroid()

	// fmt.Println("Conputing radius")
	// dataset.Calcradius()

	// fmt.Println("Conputing CalcCenterdistance")
	// dataset.CalcCenterdistance()

	// fmt.Println("Filtering data set")
	// dataset.Filterdataset(dataset.MinCaoszoneRule)	

	fmt.Println("Calling KNN")
	nonparametric.KNN(&dataset,3)

	// fmt.Println("Calling Kmeans")
	// nonparametric.Kmeans(&dataset)

	dataset.Printresults()

	//dataset.GroupCenterdists()
}