package graphlib


/*
	GraphLibrary for GeoGraphical Production and Supply Chain Management
*/

type GeoNode interface {
	getID() 		string
	get_DB_key()	string
}

type EdgeWeight interface {
	getCost(key int) float32
}

type GeoEdge interface {
	getNodeFrom() 	string 
	getNodeTo() 	string 
	getEdgeWeight() EdgeWeight
}

type GeoGraph interface {
	AddNode(GeoNode)
	AddEdge(GeoEdge)
	
	AllPairsShortestPath(sourceNode int, Distances *[]int)
	MaxCut(sourceNodes *[]int, destinationNodes *[]int)
}

// Get an Create a super Regional Node

