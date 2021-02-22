package nonparametric

import (
	//"gocv.io/x/gocv"
	"fmt"
	"math"
	"sort"

)

type labeldist struct{

	dist []float64

	learnedlabel string
}

type features struct {

	features 	[3]float64 
	
	label 		string
}


type labelfeatures struct {

	study 		[]features

	know 		[]features
	
	distance 	[]labeldist
}

/**
 * [func description]
 * @param  {[type]} lf *labelfeatures) calcdistance( [description]
 * @return {[type]}    [description]
 */

type ByDistance []labeldist

func (d Bydist) Len() int { return len(d) }
func (d Bydist) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d Bydist) Less(i, j int) bool { return d.dist[i] < d.dist[j] }

func (lf *labelfeatures) calcdistance() { 

	var sum float64 = 0.0

	for i := 0; i < len((*lf).study); i++ {
		for j := 0; j < len((*lf).know); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*lf).know[j].features[f] - (*lf).study[j].features[f],2))
			}			
			(*lf).distance[i].dist[j] = math.Sqrt(sum)
			(*lf).distance[i].learnedlabel[j] = (*lf).know[j].label
		}
	}

}

func KNN(LabelFeatures *labelfeatures, result *map[string]int,k int){
	
	(*LabelFeatures).calcdistance()

	for i := 0; i < len((*LabelFeatures).know); i++ {
		(*result)[(*LabelFeatures).distance[i].label] = 0
	}

	for i := 0; i < len((*LabelFeatures).distance); i++ {
		sort.Sort(Bydist((*LabelFeatures).distance[i]))
		for j := 0; j < k; j++ {
			(*result)[(*LabelFeatures).distance[j].label]++
		}
		
	}


}